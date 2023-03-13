package mysql

import (
	"context"
	"fmt"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductsRepo struct {
	db *gorm.DB
}

func NewProductsRepo(db *gorm.DB) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

func (p *ProductsRepo) Create(ctx context.Context, product *models.Products) error {
	err := p.db.WithContext(ctx).Select(
		"id",
		"url",
		"name_ru",
		"name_uz",
		"name_en",
		"description_ru",
		"description_uz",
		"description_en",
		"position",
		"group_id",
		"parent_product",
		"image",
		"price",
		"seo_description",
		"seo_keywords",
		"seo_text",
		"seo_title",
		"enabled",
		"created_at",
	).Create(product).Error

	return err
}

func (p *ProductsRepo) Update(ctx context.Context, product *models.Products) error {
	columns := map[string]interface{}{
		"url":             product.Url,
		"name_ru":         product.NameRu,
		"name_uz":         product.NameUz,
		"name_en":         product.NameEn,
		"description_ru":  product.DescriptionRu,
		"description_uz":  product.DescriptionUz,
		"description_en":  product.DescriptionEn,
		"position":        product.Position,
		"group_id":        product.GroupID,
		"price":           product.Price,
		"seo_description": product.SeoDescription,
		"seo_keywords":    product.SeoKeywords,
		"seo_text":        product.SeoText,
		"seo_title":       product.SeoTitle,
		"enabled":         product.Enabled,
		"updated_at":      product.UpdatedAt,
	}

	err := p.db.Clauses(clause.Returning{}).WithContext(ctx).Model(product).Updates(columns).Error

	return err
}

func (p *ProductsRepo) GetAll(ctx context.Context) ([]models.Products, error) {
	var products []models.Products

	err := p.db.WithContext(ctx).Order("id desc").Find(&products, "deleted_at IS NULL").Error

	return products, err
}

func (p *ProductsRepo) GetAllByFilterAndGroupID(ctx context.Context, groupID string) ([]models.Products, error) {
	var products []models.Products

	err := p.db.WithContext(ctx).Order("id desc").Find(&products, "deleted_at IS NULL AND group_id = ?", groupID).Error

	return products, err
}

func (p *ProductsRepo) GetAllByFilter(ctx context.Context, input models.GetProductsByFilterInput) ([]models.Products, error) {
	var (
		products []models.Products
		sub      = p.db.WithContext(ctx).Model(models.ProductRates{}).Select("ROUND(AVG(rate))").Where("product_id = p.id")
		query    = p.db.WithContext(ctx).Table("products AS p").Select("p.*, (?) AS rate", sub)
	)

	if input.NameUz != "" {
		query = query.Where("name_uz LIKE ?", fmt.Sprintf("%%%s%%", input.NameUz))
	}

	if input.Url != "" {
		query = query.Where("url LIKE ?", fmt.Sprintf("%%%s%%", input.Url))
	}

	if input.Price != nil {
		query = query.Where("price LIKE ?", fmt.Sprintf("%%%f%%", *input.Price))
	}

	if input.DescriptionUz != "" {
		query = query.Where("description_uz LIKE ?", fmt.Sprintf("%%%s%%", input.DescriptionUz))
	}

	if input.GroupID != "" {
		query = query.Where("group_id = ?", input.GroupID)
	}

	if input.Enabled != nil {
		query = query.Where("enabled = ?", *input.Enabled)
	}

	offset := (input.Page - 1) * input.PageSize

	if input.SortBy == "name_uz" || input.SortBy == "price" || input.SortBy == "created_at" {
		if input.Desc {
			query = query.Order(fmt.Sprintf("%s DESC", input.SortBy))
		} else {
			query = query.Order(input.SortBy)
		}
	} else {
		query = query.Order("-position DESC")
	}

	query = query.Limit(input.PageSize).Offset(offset)

	err := query.Debug().Find(&products).Error

	return products, err
}

func (p *ProductsRepo) GetByID(ctx context.Context, ID string) (models.Products, error) {
	var product models.Products

	err := p.db.WithContext(ctx).First(&product, "deleted_at IS NULL AND id = ?", ID).Error

	return product, err
}

func (p *ProductsRepo) GetByUrl(ctx context.Context, url string) (models.Products, error) {
	var product models.Products

	err := p.db.WithContext(ctx).First(&product, "deleted_at IS NULL AND url = ?", url).Error

	return product, err
}

func (p *ProductsRepo) GetCountNameUz(ctx context.Context, nameUz string, ID ...string) (int64, error) {
	var (
		count int64
	)

	query := p.db.WithContext(ctx).Model(models.Products{}).Select("name_uz").Where("name_uz = ?", nameUz)

	if len(ID) == 1 {
		query = query.Where("id != ?", ID[0])
	}

	err := query.Count(&count).Error

	return count, err
}

func (p *ProductsRepo) GetCount(ctx context.Context) (int64, error) {
	var count int64

	err := p.db.WithContext(ctx).Model(models.Products{}).Count(&count).Error

	return count, err
}

func (p *ProductsRepo) DeleteByID(ctx context.Context, ID string) error {

	err := p.db.WithContext(ctx).Delete(models.Products{}, "id = ?", ID).Error

	return err
}
