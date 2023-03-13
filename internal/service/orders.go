package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"time"
)

type OrdersService struct {
	ordersRepo           mysql.Orders
	orderProductsService OrderProducts
}

func NewOrdersService(ordersRepo mysql.Orders, orderProductsService OrderProducts) *OrdersService {
	return &OrdersService{
		ordersRepo:           ordersRepo,
		orderProductsService: orderProductsService,
	}
}

func (o *OrdersService) Create(ctx context.Context, input models.CreateOrderWithProductsInput) (models.OrderWithOrderProducts, error) {
	var orderWithOrderProducts models.OrderWithOrderProducts

	if input.OrderProducts == nil {
		return models.OrderWithOrderProducts{}, models.ErrNotFoundOrderProducts
	}

	order := models.Orders{
		System:        input.System,
		AccountID:     input.AccountID,
		ChatID:        input.ChatID,
		RegionID:      input.RegionID,
		CustomerName:  input.CustomerName,
		CustomerPhone: input.CustomerPhone,
		City:          input.City,
		District:      input.District,
		Street:        input.Street,
		Home:          input.Home,
		Apartment:     input.Apartment,
		Comment:       input.Comment,
		PaymentType:   input.PaymentType,
		DeliveryPrice: input.DeliveryPrice,
		FullSum:       input.FullSum,
		Status:        input.Status,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	err := o.ordersRepo.Create(ctx, &order)
	if err != nil {
		return models.OrderWithOrderProducts{}, err
	}

	for _, orderProduct := range input.OrderProducts {
		orderProduct.OrderID = order.ID
		resOrderProduct, err := o.orderProductsService.Create(ctx, orderProduct)
		if err != nil {
			return models.OrderWithOrderProducts{}, err
		}
		orderWithOrderProducts.OrderProducts = append(orderWithOrderProducts.OrderProducts, resOrderProduct)
	}

	return orderWithOrderProducts, err
}

func (o *OrdersService) Update(ctx context.Context, input models.UpdateOrderInput) (models.Orders, error) {
	order := models.Orders{
		ID:            input.ID,
		System:        input.System,
		AccountID:     input.AccountID,
		ChatID:        input.ChatID,
		RegionID:      input.RegionID,
		CustomerName:  input.CustomerName,
		CustomerPhone: input.CustomerPhone,
		City:          input.City,
		District:      input.District,
		Street:        input.Street,
		Home:          input.Home,
		Apartment:     input.Apartment,
		Comment:       input.Comment,
		PaymentType:   input.PaymentType,
		DeliveryPrice: input.DeliveryPrice,
		FullSum:       input.FullSum,
		Status:        input.Status,
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	err := o.ordersRepo.Update(ctx, &order)

	return order, err
}

func (o *OrdersService) GetAll(ctx context.Context) ([]models.Orders, error) {
	return o.ordersRepo.GetAll(ctx)
}

func (o *OrdersService) GetAllByFilter(ctx context.Context, input models.GetOrdersByFilterInput) ([]models.Orders, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrdersService) GetByID(ctx context.Context, ID uint64) (models.OrderWithOrderProducts, error) {
	var orderWithOrderProducts models.OrderWithOrderProducts

	order, err := o.ordersRepo.GetByID(ctx, ID)
	if err != nil {
		return models.OrderWithOrderProducts{}, err
	}

	orderProducts, err := o.orderProductsService.GetAllByOrderID(ctx, ID)
	if err != nil {
		return models.OrderWithOrderProducts{}, nil
	}

	if orderProducts == nil {
		orderWithOrderProducts.OrderProducts = []models.OrderProducts{}
	}

	orderWithOrderProducts.Orders = order
	orderWithOrderProducts.OrderProducts = orderProducts

	return orderWithOrderProducts, err
}

func (o *OrdersService) IncreaseFullSum(ctx context.Context, ID uint64, fullSum float64) error {
	order, err := o.ordersRepo.GetByID(ctx, ID)
	if err != nil {
		return err
	}

	err = o.ordersRepo.UpdateFullSumByID(ctx, ID, order.FullSum+fullSum)

	return err
}

func (o *OrdersService) DeleteByID(ctx context.Context, ID uint64) error {
	return o.ordersRepo.DeleteByID(ctx, ID)
}
