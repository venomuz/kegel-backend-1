package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
)

type SettingsService struct {
	settingsRepo mysql.Settings
}

func NewSettingsService(settingsRepo mysql.Settings) *SettingsService {
	return &SettingsService{settingsRepo: settingsRepo}
}

func (s *SettingsService) Create(ctx context.Context, input models.CreateSettingInput) (models.Settings, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingsService) Update(ctx context.Context, input models.UpdateSettingInput) (models.Settings, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingsService) GetByID(ctx context.Context, ID uint32) (models.Settings, error) {
	return s.settingsRepo.GetByID(ctx, ID)
}

func (s *SettingsService) GetAll(ctx context.Context) ([]models.Settings, error) {
	return s.settingsRepo.GetAll(ctx)
}

func (s *SettingsService) GetByKey(ctx context.Context, key string) (models.Settings, error) {
	return s.settingsRepo.GetByKey(ctx, key)
}

func (s *SettingsService) DeleteByID(ctx context.Context, ID uint32) error {
	//TODO implement me
	panic("implement me")
}
