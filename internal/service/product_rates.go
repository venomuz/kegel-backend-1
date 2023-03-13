package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"time"
)

type ProductRatesService struct {
	productRatesRepo mysql.ProductRates
	accountsService  Accounts
}

func NewProductRatesService(productRatesRepo mysql.ProductRates, accountsService Accounts) *ProductRatesService {
	return &ProductRatesService{
		productRatesRepo: productRatesRepo,
		accountsService:  accountsService,
	}
}

func (p *ProductRatesService) Create(ctx context.Context, input models.CreateProductRateInput) (models.ProductRates, error) {
	productRate := models.ProductRates{
		OrderID:     input.OrderID,
		ProductID:   input.ProductID,
		AccountID:   input.AccountID,
		Rate:        input.Rate,
		Description: input.Description,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if !input.Noname {

		account, err := p.accountsService.GetByID(ctx, input.AccountID)
		if err != nil {
			return models.ProductRates{}, err
		}

		productRate.AccountFirstname = account.FirstName
	}

	err := p.productRatesRepo.Create(ctx, &productRate)

	return productRate, err
}

func (p *ProductRatesService) GetAll(ctx context.Context) ([]models.ProductRates, error) {
	return p.productRatesRepo.GetAll(ctx)
}

func (p *ProductRatesService) GetAllOrderByCreatedAtByProductID(ctx context.Context, productID string) ([]models.ProductRates, error) {
	return p.productRatesRepo.GetAllOrderByCreatedAtByProductID(ctx, productID)
}

func (p *ProductRatesService) GetByID(ctx context.Context, ID uint64) (models.ProductRates, error) {
	return p.productRatesRepo.GetByID(ctx, ID)
}
