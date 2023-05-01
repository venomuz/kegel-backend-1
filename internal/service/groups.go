package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"github.com/venomuz/kegel-backend/pkg/humanizer"
	"time"
)

type GroupsService struct {
	groupsRepo   mysql.Groups
	filesService Files
	humanizerUrl humanizer.Url
}

func NewGroupsService(groupsRepo mysql.Groups, filesService Files, humanizerUrl humanizer.Url) *GroupsService {
	return &GroupsService{
		groupsRepo:   groupsRepo,
		filesService: filesService,
		humanizerUrl: humanizerUrl,
	}
}

func (g *GroupsService) Create(ctx context.Context, input models.CreateGroupInput) (models.Groups, error) {

	group := models.Groups{
		ID:             uuid.New().String(),
		Url:            input.Url,
		NameRu:         input.NameRu,
		NameUz:         input.NameUz,
		NameEn:         input.NameEn,
		DescriptionRu:  input.DescriptionRu,
		DescriptionUz:  input.DescriptionUz,
		DescriptionEn:  input.DescriptionEn,
		Position:       input.Position,
		ParentGroup:    input.ParentGroup,
		SeoDescription: input.SeoDescription,
		SeoKeywords:    input.SeoKeywords,
		SeoText:        input.SeoText,
		SeoTitle:       input.SeoTitle,
		Enabled:        input.Enabled,
		CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
	}
	if input.FileImage != nil {
		name, err := g.filesService.Save(ctx, models.File{File: input.FileImage, Path: models.FilePathGroups})
		if err != nil {
			return models.Groups{}, err
		}
		group.Image = name
	}

	if input.ParentGroup == "" {
		group.ParentGroup = "0"
	}

	if input.Url == "" {
		count, _ := g.groupsRepo.GetCountNameUz(ctx, input.NameUz)

		url := g.humanizerUrl.Generate(input.NameUz)
		if count != 0 {
			url = url + fmt.Sprintf("-%d", count+1)
		}

		group.Url = url
	}

	err := g.groupsRepo.Create(ctx, &group)

	return group, err
}

func (g *GroupsService) Update(ctx context.Context, input models.UpdateGroupInput) (models.Groups, error) {

	groupInfo, err := g.groupsRepo.GetByID(ctx, input.ID)
	if err != nil {
		return models.Groups{}, err
	}

	group := models.Groups{
		ID:            input.ID,
		Url:           input.Url,
		NameRu:        input.NameRu,
		NameUz:        input.NameUz,
		NameEn:        input.NameEn,
		DescriptionRu: input.DescriptionRu,
		DescriptionUz: input.DescriptionUz,
		DescriptionEn: input.DescriptionEn,
		Position:      input.Position,
		ParentGroup:   input.ParentGroup,
		Enabled:       input.Enabled,
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	if input.FileImage != nil && input.FileImage.Size != 0 {

		name, err := g.filesService.Save(ctx, models.File{File: input.FileImage, Path: models.FilePathGroups})
		if err != nil {
			return models.Groups{}, err
		}

		group.Image = name

		if groupInfo.Image != "" {
			_ = g.filesService.DeleteByName(ctx, models.FilePathGroups, groupInfo.Image)
		}
	}

	err = g.groupsRepo.Update(ctx, &group)

	return group, err
}

func (g *GroupsService) GetAll(ctx context.Context) ([]models.Groups, error) {
	return g.groupsRepo.GetAll(ctx)
}

func (g *GroupsService) GetAllByFilterWithChild(ctx context.Context, input models.GetGroupsByFilterInput) ([]models.GroupsWithChild, error) {

	var groupsWithChild []models.GroupsWithChild

	if input.All {
		groups, err := g.groupsRepo.GetAll(ctx)

		for _, group := range groups {
			groupWithChild := models.GroupsWithChild{
				Groups: group,
			}
			for _, check := range groups {

				if check.ParentGroup == "0" {
					continue
				}
				if group.ID == check.ParentGroup {
					groupWithChild.Child = append(groupWithChild.Child, check)
				}
			}

			if groupWithChild.Child == nil {
				groupWithChild.Child = []models.Groups{}
			}

			groupsWithChild = append(groupsWithChild, groupWithChild)
		}
		return groupsWithChild, err
	}

	if input.ParentID == "" {
		input.ParentID = "0"
	}

	groupsFilter, err := g.groupsRepo.GetAllByFilter(ctx, input)

	groupsAll, err := g.groupsRepo.GetAll(ctx)

	for _, group := range groupsFilter {
		groupWithChild := models.GroupsWithChild{
			Groups: group,
		}
		for _, check := range groupsAll {

			if check.ParentGroup == "0" {
				continue
			}
			if group.ID == check.ParentGroup {
				groupWithChild.Child = append(groupWithChild.Child, check)
			}
		}

		if groupWithChild.Child == nil {
			groupWithChild.Child = []models.Groups{}
		}

		groupsWithChild = append(groupsWithChild, groupWithChild)
	}

	if groupsWithChild == nil {
		groupsWithChild = []models.GroupsWithChild{}
	}

	return groupsWithChild, err
}

func (g *GroupsService) GetByID(ctx context.Context, ID string) (models.Groups, error) {
	return g.groupsRepo.GetByID(ctx, ID)
}

func (g *GroupsService) GetByUrl(ctx context.Context, url string) (models.Groups, error) {
	return g.groupsRepo.GetByUrl(ctx, url)
}

func (g *GroupsService) DeleteByID(ctx context.Context, ID string) error {
	group, err := g.groupsRepo.GetByID(ctx, ID)
	if err != nil {
		return err
	}

	if group.Image != "" {
		if err = g.filesService.DeleteByName(ctx, models.FilePathGroups, group.Image); err != nil {
			return err
		}

	}

	if group.ParentGroup == "0" {
		err = g.groupsRepo.DeleteChildByParentID(ctx, ID)
	}

	return g.groupsRepo.DeleteByID(ctx, ID)
}
