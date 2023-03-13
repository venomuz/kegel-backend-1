package mysql

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (u UsersRepo) GetByUsername(ctx context.Context, username string) (models.Users, error) {
	var user models.Users

	err := u.db.WithContext(ctx).First(&user, "deleted_at IS NULL AND username = ?", username).Error

	return user, err
}
