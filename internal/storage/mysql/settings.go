package mysql

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SettingsRepo struct {
	db *gorm.DB
}

func NewSettingsRepo(db *gorm.DB) *SettingsRepo {
	return &SettingsRepo{
		db: db,
	}
}

func (s *SettingsRepo) Create(ctx context.Context, setting *models.Settings) error {

	err := s.db.WithContext(ctx).Create(setting).Error

	return err
}

func (s *SettingsRepo) Update(ctx context.Context, setting *models.Settings) error {
	columns := map[string]interface{}{
		"title":      setting.Title,
		"key":        setting.Key,
		"value":      setting.Value,
		"updated_at": setting.UpdatedAt,
	}

	err := s.db.Clauses(clause.Returning{}).WithContext(ctx).Model(setting).Updates(columns).Error

	return err
}

func (s *SettingsRepo) GetByID(ctx context.Context, ID uint32) (models.Settings, error) {
	var setting models.Settings

	err := s.db.WithContext(ctx).Where("deleted_at IS NULL").First(&setting, "ID = ?", ID).Error

	return setting, err
}

func (s *SettingsRepo) GetAll(ctx context.Context) ([]models.Settings, error) {
	var settings []models.Settings

	err := s.db.WithContext(ctx).Where("deleted_at IS NULL").Order("id desc").Find(&settings).Error

	return settings, err
}

func (s *SettingsRepo) GetByKey(ctx context.Context, key string) (models.Settings, error) {
	var setting models.Settings

	err := s.db.WithContext(ctx).Where("deleted_at IS NULL AND `key` = ?", key).First(&setting).Error

	return setting, err
}

func (s *SettingsRepo) DeleteByID(ctx context.Context, ID uint32) error {

	err := s.db.WithContext(ctx).Delete(models.Settings{}, "id = ?", ID).Error

	return err
}
