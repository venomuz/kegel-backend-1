package mysql

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type OrdersRepo struct {
	db *gorm.DB
}

func NewOrdersRepo(db *gorm.DB) *OrdersRepo {
	return &OrdersRepo{
		db: db,
	}
}

func (o *OrdersRepo) Create(ctx context.Context, order *models.Orders) error {
	err := o.db.WithContext(ctx).Select(
		"system",
		"account_id",
		"chat_id",
		"region_id",
		"customer_name",
		"customer_phone",
		"city",
		"district",
		"street",
		"home",
		"apartment",
		"comment",
		"payment_type",
		"delivery_price",
		"full_sum",
		"status",
		"created_at",
	).Create(order).Error

	return err
}

func (o *OrdersRepo) Update(ctx context.Context, order *models.Orders) error {
	columns := map[string]interface{}{
		"system":         order.System,
		"account_id":     order.AccountID,
		"chat_id":        order.ChatID,
		"region_id":      order.RegionID,
		"customer_name":  order.CustomerName,
		"customer_phone": order.CustomerPhone,
		"city":           order.City,
		"district":       order.District,
		"street":         order.Street,
		"home":           order.Home,
		"apartment":      order.Apartment,
		"comment":        order.Comment,
		"payment_type":   order.PaymentType,
		"delivery_price": order.DeliveryPrice,
		"full_sum":       order.FullSum,
		"status":         order.Status,
		"updated_at":     order.UpdatedAt,
	}

	err := o.db.Clauses(clause.Returning{}).WithContext(ctx).Model(order).Updates(columns).Error

	return err
}

func (o *OrdersRepo) UpdateFullSumByID(ctx context.Context, ID uint64, fullSum float64) error {
	err := o.db.WithContext(ctx).Table("orders").Where("id = ?", ID).Updates(map[string]interface{}{
		"full_sum": fullSum,
	}).Error

	return err
}

func (o *OrdersRepo) GetAll(ctx context.Context) ([]models.Orders, error) {
	var orders []models.Orders

	err := o.db.WithContext(ctx).Order("id desc").Find(&orders, "deleted_at IS NULL").Error

	return orders, err
}

func (o *OrdersRepo) GetByID(ctx context.Context, ID uint64) (models.Orders, error) {
	var order models.Orders

	err := o.db.WithContext(ctx).First(&order, "id = ?", ID).Error

	return order, err
}

func (o *OrdersRepo) DeleteByID(ctx context.Context, ID uint64) error {
	err := o.db.WithContext(ctx).Updates(&models.Orders{ID: ID, DeletedAt: time.Now().Format("2006-01-02 15:04:05")}).Error
	return err
}
