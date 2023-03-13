package mysql

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountsRepo struct {
	db *gorm.DB
}

func NewAccountsRepo(db *gorm.DB) *AccountsRepo {
	return &AccountsRepo{
		db: db,
	}
}

func (a *AccountsRepo) Create(ctx context.Context, account *models.Accounts) error {
	err := a.db.Debug().WithContext(ctx).Select(
		"region_id",
		"chat_id",
		"system",
		"first_name",
		"last_name",
		"birthday",
		"phone_number",
		"password",
		"language",
		"created_at",
	).Create(account).Error

	return err
}

func (a *AccountsRepo) Update(ctx context.Context, account *models.Accounts) error {

	columns := map[string]interface{}{
		"region_id":    account.RegionID,
		"chat_id":      account.ChatID,
		"system":       account.System,
		"first_name":   account.FirstName,
		"last_name":    account.LastName,
		"birthday":     account.Birthday,
		"phone_number": account.PhoneNumber,
		"language":     account.Language,
		"blocked":      account.Blocked,
		"updated_at":   account.UpdatedAt,
	}

	if account.Password != "" {
		columns["password"] = account.Password
	}

	err := a.db.Clauses(clause.Returning{}).WithContext(ctx).Model(&account).Updates(columns).Error

	return err
}

func (a *AccountsRepo) GetAll(ctx context.Context) ([]models.Accounts, error) {
	var accounts []models.Accounts

	err := a.db.WithContext(ctx).Debug().Order("id desc").Find(&accounts, "deleted_at IS NULL").Error

	return accounts, err
}

func (a *AccountsRepo) GetByID(ctx context.Context, ID uint32) (models.Accounts, error) {
	var account models.Accounts

	err := a.db.WithContext(ctx).First(&account, "deleted_at IS NULL AND id = ?", ID).Error

	return account, err
}

func (a *AccountsRepo) GetByPhoneNumber(ctx context.Context, phoneNumber string) (models.Accounts, error) {
	var user models.Accounts

	err := a.db.WithContext(ctx).First(&user, "deleted_at IS NULL AND phone_number = ?", phoneNumber).Error

	return user, err
}

func (a *AccountsRepo) DeleteByID(ctx context.Context, ID uint32) error {
	//TODO implement me
	panic("implement me")
}
