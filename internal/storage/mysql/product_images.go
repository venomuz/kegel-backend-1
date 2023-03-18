package mysql

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ProductImagesRepo struct {
	db *gorm.DB
}

func NewProductImagesRepo(db *gorm.DB) *ProductImagesRepo {
	return &ProductImagesRepo{
		db: db,
	}
}

func (p *ProductImagesRepo) Create(ctx context.Context, productImage *models.ProductImages) error {
	err := p.db.WithContext(ctx).Select(
		"product_id",
		"image",
		"position",
		"created_at",
	).Create(productImage).Error

	return err
}

func (p *ProductImagesRepo) Update(ctx context.Context, productImage *models.ProductImages) error {
	columns := map[string]interface{}{
		"product_id": productImage.ProductID,
		"position":   productImage.Position,
		"updated_at": productImage.UpdatedAt,
	}

	if productImage.Image != "" {
		columns["image"] = productImage.Image
	}

	err := p.db.Clauses(clause.Returning{}).WithContext(ctx).Model(productImage).Updates(columns).Error

	return err
}

func (p *ProductImagesRepo) GetAll(ctx context.Context) ([]models.ProductImages, error) {
	var productImages []models.ProductImages

	err := p.db.WithContext(ctx).Order("product_id").Find(&productImages, "deleted_at IS NULL").Error

	return productImages, err
}

func (p *ProductImagesRepo) GetByID(ctx context.Context, ID uint32) (models.ProductImages, error) {
	var productImage models.ProductImages

	err := p.db.WithContext(ctx).First(&productImage, "deleted_at IS NULL AND id = ?", ID).Error

	return productImage, err
}

func (p *ProductImagesRepo) GetAllByProductID(ctx context.Context, productID string) ([]models.ProductImages, error) {
	var productImages []models.ProductImages

	err := p.db.WithContext(ctx).Order("-position DESC").Find(&productImages, "deleted_at IS NULL AND product_id = ?", productID).Error

	return productImages, err
}

func (p *ProductImagesRepo) DeleteByID(ctx context.Context, ID uint32) error {

	err := p.db.WithContext(ctx).Updates(&models.ProductImages{ID: ID, DeletedAt: time.Now().Format("2006-01-02 15:04:05")}).Error

	return err
}

func (p *ProductImagesRepo) DeleteByIDs(ctx context.Context, IDs []uint32) error {

	err := p.db.WithContext(ctx).Delete(models.ProductImages{}, "id IN ?", IDs).Error

	return err
}
