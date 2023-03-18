package mysql

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
)

type Accounts interface {
	Create(ctx context.Context, account *models.Accounts) error
	Update(ctx context.Context, account *models.Accounts) error
	GetAll(ctx context.Context) ([]models.Accounts, error)
	GetByID(ctx context.Context, ID uint32) (models.Accounts, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (models.Accounts, error)
	DeleteByID(ctx context.Context, ID uint32) error
}

type Groups interface {
	Create(ctx context.Context, group *models.Groups) error
	Update(ctx context.Context, group *models.Groups) error
	GetAll(ctx context.Context) ([]models.Groups, error)
	GetAllByFilter(ctx context.Context, input models.GetGroupsByFilterInput) ([]models.Groups, error)
	GetByID(ctx context.Context, ID string) (models.Groups, error)
	GetByUrl(ctx context.Context, url string) (models.Groups, error)
	GetCountNameUz(ctx context.Context, nameUz string, id ...string) (int64, error)
	DeleteByID(ctx context.Context, ID string) error
}

type OrderProducts interface {
	Create(ctx context.Context, order *models.OrderProducts) error
	Update(ctx context.Context, order *models.OrderProducts) error
	GetAll(ctx context.Context) ([]models.OrderProducts, error)
	GetAllByOrderID(ctx context.Context, orderID uint64) ([]models.OrderProducts, error)
	GetAllPriceFullSumByOrderID(ctx context.Context, orderID uint64) (float64, error)
	GetByID(ctx context.Context, ID uint64) (models.OrderProducts, error)
	DeleteByID(ctx context.Context, ID uint64) error
}

type Orders interface {
	Create(ctx context.Context, order *models.Orders) error
	Update(ctx context.Context, order *models.Orders) error
	UpdateFullSumByID(ctx context.Context, ID uint64, fullSum float64) error
	GetAll(ctx context.Context) ([]models.Orders, error)
	GetByID(ctx context.Context, ID uint64) (models.Orders, error)
	DeleteByID(ctx context.Context, ID uint64) error
}

type ProductImages interface {
	Create(ctx context.Context, account *models.ProductImages) error
	Update(ctx context.Context, account *models.ProductImages) error
	GetAll(ctx context.Context) ([]models.ProductImages, error)
	GetAllByProductID(ctx context.Context, productID string) ([]models.ProductImages, error)
	GetByID(ctx context.Context, ID uint32) (models.ProductImages, error)
	DeleteByID(ctx context.Context, ID uint32) error
	DeleteByIDs(ctx context.Context, IDs []uint32) error
}

type ProductRates interface {
	Create(ctx context.Context, productRate *models.ProductRates) error
	GetAll(ctx context.Context) ([]models.ProductRates, error)
	GetAllOrderByCreatedAtByProductID(ctx context.Context, productID string) ([]models.ProductRates, error)
	GetByID(ctx context.Context, ID uint64) (models.ProductRates, error)
}

type Products interface {
	Create(ctx context.Context, product *models.Products) error
	Update(ctx context.Context, product *models.Products) error
	GetAll(ctx context.Context) ([]models.Products, error)
	GetAllByFilterAndGroupID(ctx context.Context, groupID string) ([]models.Products, error)
	GetAllByFilter(ctx context.Context, input models.GetProductsByFilterInput) ([]models.Products, error)
	GetAllByIDs(ctx context.Context, input models.GetProductsByIDsInput) ([]models.Products, error)
	GetByID(ctx context.Context, ID string) (models.Products, error)
	GetByUrl(ctx context.Context, url string) (models.Products, error)
	GetCountNameUz(ctx context.Context, nameUz string, id ...string) (int64, error)
	GetCount(ctx context.Context) (int64, error)
	DeleteByID(ctx context.Context, ID string) error
}

type Settings interface {
	GetByID(ctx context.Context, ID uint32) (models.Settings, error)
	GetAll(ctx context.Context) ([]models.Settings, error)
	GetByKey(ctx context.Context, key string) (models.Settings, error)
}

type Users interface {
	GetByUsername(ctx context.Context, username string) (models.Users, error)
}

type Repositories struct {
	Accounts      Accounts
	Groups        Groups
	OrderProducts OrderProducts
	Orders        Orders
	ProductImages ProductImages
	ProductRates  ProductRates
	Products      Products
	Settings      Settings
	Users         Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Accounts:      NewAccountsRepo(db),
		Groups:        NewGroupsRepo(db),
		OrderProducts: NewOrderProductsRepo(db),
		Orders:        NewOrdersRepo(db),
		ProductImages: NewProductImagesRepo(db),
		ProductRates:  NewProductRatesRepo(db),
		Products:      NewProductsRepo(db),
		Settings:      NewSettingsRepo(db),
		Users:         NewUsersRepo(db),
	}
}
