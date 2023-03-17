package mysql

import (
	"context"
	"fmt"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type GroupsRepo struct {
	db *gorm.DB
}

func NewGroupsRepo(db *gorm.DB) *GroupsRepo {
	return &GroupsRepo{
		db: db,
	}
}

func (g *GroupsRepo) Create(ctx context.Context, group *models.Groups) error {
	err := g.db.WithContext(ctx).Select(
		"id",
		"url",
		"name_ru",
		"name_uz",
		"name_en",
		"description_ru",
		"description_uz",
		"description_en",
		"position",
		"parent_group",
		"image",
		"seo_description",
		"seo_keywords",
		"seo_text",
		"seo_title",
		"enabled",
		"created_at",
	).Create(group).Error

	return err
}

func (g *GroupsRepo) Update(ctx context.Context, group *models.Groups) error {
	columns := map[string]interface{}{
		"url":             group.Url,
		"name_ru":         group.NameRu,
		"name_uz":         group.NameUz,
		"name_en":         group.NameEn,
		"description_ru":  group.DescriptionRu,
		"description_uz":  group.DescriptionUz,
		"description_en":  group.DescriptionEn,
		"position":        group.Position,
		"parent_group":    group.ParentGroup,
		"image":           group.Image,
		"seo_description": group.SeoDescription,
		"seo_keywords":    group.SeoKeywords,
		"seo_text":        group.SeoText,
		"seo_title":       group.SeoTitle,
		"enabled":         group.Enabled,
		"updated_at":      group.UpdatedAt,
	}

	err := g.db.Clauses(clause.Returning{}).WithContext(ctx).Model(group).Updates(columns).Error
	return err
}

func (g *GroupsRepo) GetAll(ctx context.Context) ([]models.Groups, error) {
	var groups []models.Groups

	err := g.db.WithContext(ctx).Order("-position DESC").Find(&groups, "deleted_at IS NULL").Error

	return groups, err
}

func (g *GroupsRepo) GetAllByFilter(ctx context.Context, input models.GetGroupsByFilterInput) ([]models.Groups, error) {
	var (
		groups []models.Groups
		query  = g.db.WithContext(ctx).Model(models.Groups{})
	)

	if input.ParentId != "" {
		query = query.Where("parent_group = ?", input.ParentId)
	}

	if input.NameUz != "" {
		query = query.Where("name_uz LIKE ?", fmt.Sprintf("%%%s%%", input.NameUz))
	}

	if input.Url != "" {
		query = query.Where("url LIKE ?", fmt.Sprintf("%%%s%%", input.Url))
	}

	if input.Enabled != nil {
		query = query.Where("enabled = ?", *input.Enabled)
	}

	query = query.Where("deleted_at IS NULL")

	query = query.Order("-position DESC")

	err := query.Debug().Find(&groups).Error

	return groups, err
}

func (g *GroupsRepo) GetByID(ctx context.Context, ID string) (models.Groups, error) {
	var group models.Groups

	err := g.db.WithContext(ctx).First(&group, "deleted_at IS NULL AND id = ?", ID).Error

	return group, err
}

func (g *GroupsRepo) GetByUrl(ctx context.Context, url string) (models.Groups, error) {
	var group models.Groups

	err := g.db.WithContext(ctx).First(&group, "url = ?", url).Error

	return group, err
}

func (g *GroupsRepo) GetCountNameUz(ctx context.Context, nameUz string, ID ...string) (int64, error) {
	var (
		count int64
	)

	query := g.db.WithContext(ctx).Model(models.Groups{}).Select("name_uz").Where("name_uz = ?", nameUz)

	if len(ID) == 1 {
		query = query.Where("id != ?", ID[0])
	}

	err := query.Count(&count).Error

	return count, err
}

func (g *GroupsRepo) DeleteByID(ctx context.Context, ID string) error {
	err := g.db.WithContext(ctx).Updates(&models.Groups{ID: ID, DeletedAt: time.Now().Format("2006-01-02 15:04:05")}).Error
	return err
}
