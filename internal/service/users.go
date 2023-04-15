package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"github.com/venomuz/kegel-backend/pkg/hash"
)

type UsersService struct {
	usersRepo    mysql.Users
	hashPassword hash.PasswordHasher
}

func NewUsersService(usersRepo mysql.Users, hashPassword hash.PasswordHasher) *UsersService {
	return &UsersService{
		usersRepo:    usersRepo,
		hashPassword: hashPassword,
	}
}

func (u *UsersService) Login(ctx context.Context, input models.LoginUserInput) (models.Users, error) {

	user, err := u.usersRepo.GetByUsername(ctx, input.Username)
	if err != nil {
		return models.Users{}, err
	}

	err = u.hashPassword.CheckString(user.Password, input.Password)
	if err != nil {
		return models.Users{}, err
	}

	return user, err
}
