package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"time"
)

type SettingsService struct {
	settingsRepo mysql.Settings
}

func NewSettingsService(settingsRepo mysql.Settings) *SettingsService {
	return &SettingsService{settingsRepo: settingsRepo}
}

func (s *SettingsService) Create(ctx context.Context, input models.CreateSettingInput) (models.Settings, error) {
	setting := models.Settings{
		Title:     input.Title,
		Key:       input.Key,
		Value:     input.Value,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := s.settingsRepo.Create(ctx, &setting)

	return setting, err
}

func (s *SettingsService) Update(ctx context.Context, input models.UpdateSettingInput) (models.Settings, error) {
	setting := models.Settings{
		ID:        input.ID,
		Title:     input.Title,
		Key:       input.Key,
		Value:     input.Value,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := s.settingsRepo.Update(ctx, &setting)

	return setting, err
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
