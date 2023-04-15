package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"github.com/venomuz/kegel-backend/internal/storage/rdb"
	"github.com/venomuz/kegel-backend/pkg/hash"
	"regexp"
	"time"
)

type AccountsService struct {
	accountsRepo mysql.Accounts
	rdbRepos     rdb.Repository
	smsService   Sms
	hash         hash.PasswordHasher
}

func NewAccountsService(accountsRepo mysql.Accounts, rdbRepos rdb.Repository, smsService Sms, hash hash.PasswordHasher) *AccountsService {
	return &AccountsService{
		accountsRepo: accountsRepo,
		rdbRepos:     rdbRepos,
		smsService:   smsService,
		hash:         hash,
	}
}

func (a *AccountsService) Create(ctx context.Context, input models.RegistrationAccountInput) (models.Accounts, error) {

	account := models.Accounts{
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	err := a.accountsRepo.Create(ctx, &account)

	return account, err
}

func (a *AccountsService) Update(ctx context.Context, input models.UpdateAccountInput) (models.Accounts, error) {
	account := models.Accounts{
		ID:        input.ID,
		RegionID:  input.RegionID,
		ChatID:    input.ChatID,
		System:    input.System,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Birthday:  input.Birthday,
		Password:  input.Password,
		Language:  input.Language,
		Blocked:   input.Blocked,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := a.accountsRepo.Update(ctx, &account)

	return account, err
}

func (a *AccountsService) GetAll(ctx context.Context) ([]models.Accounts, error) {
	return a.accountsRepo.GetAll(ctx)
}

func (a *AccountsService) GetByID(ctx context.Context, ID uint32) (models.Accounts, error) {
	return a.accountsRepo.GetByID(ctx, ID)
}

func (a *AccountsService) Login(ctx context.Context, input models.LoginAccountInput) (models.Accounts, error) {
	account, err := a.accountsRepo.GetByPhoneNumber(ctx, input.PhoneNumber)
	if err != nil {
		return models.Accounts{}, err
	}

	err = a.hash.CheckString(account.Password, input.Password)
	if err != nil {
		return models.Accounts{}, err
	}

	return account, err
}

func (a *AccountsService) Registration(ctx context.Context, input models.RegistrationAccountInput) (models.Accounts, error) {

	code, err := a.rdbRepos.Get(ctx, input.PhoneNumber)
	if err != nil {
		return models.Accounts{}, models.ErrNotSendVerification
	}

	if input.VerificationCode != code {
		return models.Accounts{}, models.ErrVerificationCodeWrong
	}

	hashed, err := a.hash.String(input.Password)
	if err != nil {
		return models.Accounts{}, err
	}

	account := models.Accounts{
		PhoneNumber: input.PhoneNumber,
		Password:    hashed,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	err = a.accountsRepo.Create(ctx, &account)

	if err == nil {
		err = a.rdbRepos.Del(ctx, input.PhoneNumber)
	}

	return account, err
}

func (a *AccountsService) SendVerificationCode(ctx context.Context, input models.AccountSendVerificationInput) error {

	match, _ := regexp.MatchString("998[0-9]{3}[0-9]{6}$", input.PhoneNumber)
	if !match {
		return models.ErrPhoneNumber
	}

	randNum, err := a.smsService.SendVerificationCode(ctx, input.PhoneNumber)
	if err != nil {
		return err
	}

	err = a.rdbRepos.SetEX(ctx, input.PhoneNumber, randNum, time.Minute*3)
	if err != nil {
		return err
	}

	return err
}

func (a *AccountsService) DeleteByID(ctx context.Context, ID uint32) error {
	//TODO implement me
	panic("implement me")
}
