package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	mysqlrepo "github.com/venomuz/kegel-backend/internal/storage/mysql"
	"time"
)

type OrderProductsService struct {
	orderProductsRepo mysqlrepo.OrderProducts
	productsService   Products
}

func NewOrderProductsService(orderProductsRepo mysqlrepo.OrderProducts, productsService Products) *OrderProductsService {
	return &OrderProductsService{
		orderProductsRepo: orderProductsRepo,
		productsService:   productsService,
	}
}

func (o *OrderProductsService) Create(ctx context.Context, input models.CreateOrderProductsInput) (models.OrderProducts, error) {
	product, err := o.productsService.GetByID(ctx, input.ProductID)
	if err != nil {
		return models.OrderProducts{}, err
	}

	orderProduct := models.OrderProducts{
		OrderID:     input.OrderID,
		ProductID:   input.ProductID,
		ProductName: product.NameUz,
		Price:       product.Price,
		Amount:      input.Amount,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	err = o.orderProductsRepo.Create(ctx, &orderProduct)
	if err != nil {
		return models.OrderProducts{}, err
	}

	//fullSum, err := o.orderProductsRepo.GetAllPriceFullSumByOrderID(ctx, input.OrderID)
	//if err != nil {
	//	return models.OrderProducts{}, err
	//}
	//
	////err = o.ordersService.IncreaseFullSum(ctx, input.OrderID, fullSum)
	////if err != nil {
	////	return models.OrderProducts{}, err
	////}

	return orderProduct, err
}

func (o *OrderProductsService) Update(ctx context.Context, input models.UpdateOrderProductsInput) (models.OrderProducts, error) {
	product, err := o.productsService.GetByID(ctx, input.ProductID)
	if err != nil {
		return models.OrderProducts{}, err
	}

	orderProduct := models.OrderProducts{
		ID:          input.ID,
		OrderID:     input.OrderID,
		ProductID:   input.ProductID,
		ProductName: product.NameUz,
		Price:       product.Price,
		Amount:      input.Amount,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	err = o.orderProductsRepo.Update(ctx, &orderProduct)
	if err != nil {
		return models.OrderProducts{}, err
	}
	//
	//fullSum, err := o.orderProductsRepo.GetAllPriceFullSumByOrderID(ctx, input.OrderID)
	//
	//err = o.ordersService.IncreaseFullSum(ctx, input.OrderID, fullSum)
	//if err != nil {
	//	return models.OrderProducts{}, err
	//}

	return orderProduct, err
}

func (o *OrderProductsService) GetAllByOrderID(ctx context.Context, orderID uint64) ([]models.OrderProducts, error) {
	return o.orderProductsRepo.GetAllByOrderID(ctx, orderID)
}

func (o *OrderProductsService) GetByID(ctx context.Context, ID uint64) (models.OrderProducts, error) {
	return o.orderProductsRepo.GetByID(ctx, ID)
}

func (o *OrderProductsService) GetPriceSumByOrderID(ctx context.Context, orderID uint64) (float64, error) {
	return o.orderProductsRepo.GetAllPriceFullSumByOrderID(ctx, orderID)
}

func (o *OrderProductsService) DeleteByID(ctx context.Context, ID uint64) error {
	return o.orderProductsRepo.DeleteByID(ctx, ID)
}
