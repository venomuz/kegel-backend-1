package mysql

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
	"time"
)

type OrderProductsRepo struct {
	db *gorm.DB
}

func NewOrderProductsRepo(db *gorm.DB) *OrderProductsRepo {
	return &OrderProductsRepo{
		db: db,
	}
}

func (o *OrderProductsRepo) Create(ctx context.Context, orderProduct *models.OrderProducts) error {
	err := o.db.WithContext(ctx).Select(
		"order_id",
		"product_id",
		"product_name",
		"price",
		"amount",
		"created_at",
	).Create(orderProduct).Error
	return err
}

func (o *OrderProductsRepo) Update(ctx context.Context, orderProduct *models.OrderProducts) error {
	columns := map[string]interface{}{
		"order_id":     orderProduct.OrderID,
		"product_id":   orderProduct.ProductID,
		"product_name": orderProduct.ProductName,
		"price":        orderProduct.Price,
		"amount":       orderProduct.Amount,
		"updated_at":   orderProduct.UpdatedAt,
	}

	err := o.db.WithContext(ctx).Model(orderProduct).Updates(columns).Error

	return err
}

func (o *OrderProductsRepo) GetByID(ctx context.Context, ID uint64) (models.OrderProducts, error) {
	var orderProduct models.OrderProducts

	err := o.db.WithContext(ctx).Where("deleted_at IS NULL").First(&orderProduct, "id = ?", ID).Error

	return orderProduct, err
}

func (o *OrderProductsRepo) GetAll(ctx context.Context) ([]models.OrderProducts, error) {
	var orderProducts []models.OrderProducts

	err := o.db.WithContext(ctx).Where("deleted_at IS NULL").Order("id desc").Find(&orderProducts).Error

	return orderProducts, err
}

func (o *OrderProductsRepo) GetAllByOrderID(ctx context.Context, orderID uint64) ([]models.OrderProducts, error) {
	var orderProducts []models.OrderProducts

	err := o.db.WithContext(ctx).Find(&orderProducts, "order_id = ?", orderID).Error

	return orderProducts, err
}

func (o *OrderProductsRepo) GetAllPriceFullSumByOrderID(ctx context.Context, orderID uint64) (float64, error) {
	var fullSum float64

	err := o.db.Select("SUM(price * amount)").Model(models.OrderProducts{}).Where("order_id = ?", orderID).Scan(&fullSum).Error

	return fullSum, err
}

func (o *OrderProductsRepo) DeleteByID(ctx context.Context, ID uint64) error {
	err := o.db.WithContext(ctx).Updates(models.OrderProducts{ID: ID, DeletedAt: time.Now().String()}).Error

	return err
}
