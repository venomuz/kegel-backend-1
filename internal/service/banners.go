package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"time"
)

type BannersService struct {
	bannersRepo  mysql.Banners
	filesService Files
}

func NewBannersService(bannersRepo mysql.Banners, filesService Files) *BannersService {
	return &BannersService{bannersRepo: bannersRepo, filesService: filesService}
}

func (s *BannersService) Create(ctx context.Context, input models.CreateBannerInput) (models.Banners, error) {
	banner := models.Banners{
		Name:      input.Name,
		Position:  input.Position,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if input.FileImage != nil {
		name, err := s.filesService.Save(ctx, models.File{File: input.FileImage, Path: models.FilePathBanners})
		if err != nil {
			return models.Banners{}, err
		}
		banner.Image = name
	}

	err := s.bannersRepo.Create(ctx, &banner)

	return banner, err
}

func (s *BannersService) Update(ctx context.Context, input models.UpdateBannerInput) (models.Banners, error) {
	bannerInfo, err := s.bannersRepo.GetByID(ctx, input.ID)
	if err != nil {
		return models.Banners{}, err
	}

	banner := models.Banners{
		ID:        input.ID,
		Name:      input.Name,
		Position:  input.Position,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if input.FileImage != nil && input.FileImage.Size != 0 {

		name, err := s.filesService.Save(ctx, models.File{File: input.FileImage, Path: models.FilePathBanners})
		if err != nil {
			return models.Banners{}, err
		}

		banner.Image = name

		if bannerInfo.Image != "" {
			_ = s.filesService.DeleteByName(ctx, models.FilePathBanners, banner.Image)
		}
	}

	err = s.bannersRepo.Update(ctx, &banner)

	return banner, err
}

func (s *BannersService) GetByID(ctx context.Context, ID uint32) (models.Banners, error) {
	return s.bannersRepo.GetByID(ctx, ID)
}

func (s *BannersService) GetAll(ctx context.Context) ([]models.Banners, error) {
	return s.bannersRepo.GetAll(ctx)
}

func (s *BannersService) GetByKey(ctx context.Context, key string) (models.Banners, error) {
	return s.bannersRepo.GetByKey(ctx, key)
}

func (s *BannersService) DeleteByID(ctx context.Context, ID uint32) error {
	return s.bannersRepo.DeleteByID(ctx, ID)
}
