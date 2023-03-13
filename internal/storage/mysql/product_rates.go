package mysql

import (
	"github.com/venomuz/kegel-backend/internal/models"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProductRatesRepo struct {
	db *gorm.DB
}

func NewProductRatesRepo(db *gorm.DB) *ProductRatesRepo {
	return &ProductRatesRepo{
		db: db,
	}
}

func (p *ProductRatesRepo) Create(ctx context.Context, productRate *models.ProductRates) error {
	err := p.db.WithContext(ctx).Select(
		"order_id",
		"product_id",
		"account_id",
		"account_firstname",
		"rate",
		"description",
		"created_at",
	).Create(productRate).Error

	return err
}

func (p *ProductRatesRepo) GetAll(ctx context.Context) ([]models.ProductRates, error) {
	var productRates []models.ProductRates

	err := p.db.WithContext(ctx).Order("id desc").Find(&productRates).Error

	return productRates, err
}

func (p *ProductRatesRepo) GetAllOrderByCreatedAtByProductID(ctx context.Context, productID string) ([]models.ProductRates, error) {
	var productRates []models.ProductRates

	err := p.db.WithContext(ctx).Find(&productRates, "product_id = ?", productID).Order("created_at").Error

	return productRates, err
}

func (p *ProductRatesRepo) GetByID(ctx context.Context, ID uint64) (models.ProductRates, error) {
	var productRates models.ProductRates

	err := p.db.WithContext(ctx).First(&productRates, "id = ?", ID).Error

	return productRates, err
}
