package mysql

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BannersRepo struct {
	db *gorm.DB
}

func NewBannersRepo(db *gorm.DB) *BannersRepo {
	return &BannersRepo{
		db: db,
	}
}

func (s *BannersRepo) Create(ctx context.Context, banner *models.Banners) error {

	err := s.db.WithContext(ctx).Create(banner).Error

	return err
}

func (s *BannersRepo) Update(ctx context.Context, banner *models.Banners) error {
	columns := map[string]interface{}{
		"name":       banner.Name,
		"position":   banner.Position,
		"updated_at": banner.UpdatedAt,
	}
	if banner.Image != "" {
		columns["image"] = banner.Image
	}

	err := s.db.Clauses(clause.Returning{}).WithContext(ctx).Model(banner).Updates(columns).Error

	return err
}

func (s *BannersRepo) GetByID(ctx context.Context, ID uint32) (models.Banners, error) {
	var banner models.Banners

	err := s.db.WithContext(ctx).Where("deleted_at IS NULL").First(&banner, "ID = ?", ID).Error

	return banner, err
}

func (s *BannersRepo) GetAll(ctx context.Context) ([]models.Banners, error) {
	var banners []models.Banners

	err := s.db.WithContext(ctx).Where("deleted_at IS NULL").Order("id desc").Find(&banners).Error

	return banners, err
}

func (s *BannersRepo) GetByKey(ctx context.Context, key string) (models.Banners, error) {
	var banner models.Banners

	err := s.db.WithContext(ctx).Where("deleted_at IS NULL AND `key` = ?", key).First(&banner).Error

	return banner, err
}

func (s *BannersRepo) DeleteByID(ctx context.Context, ID uint32) error {

	err := s.db.WithContext(ctx).Delete(models.Banners{}, "id = ?", ID).Error

	return err
}
