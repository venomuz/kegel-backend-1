package service

import (
	"context"
	"github.com/venomuz/kegel-backend/config"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"github.com/venomuz/kegel-backend/internal/storage/rdb"
	"github.com/venomuz/kegel-backend/pkg/gen"
	"github.com/venomuz/kegel-backend/pkg/hash"
	"github.com/venomuz/kegel-backend/pkg/humanizer"
	"github.com/venomuz/kegel-backend/pkg/logger"
)

type Accounts interface {
	Create(ctx context.Context, input models.RegistrationAccountInput) (models.Accounts, error)
	Update(ctx context.Context, input models.UpdateAccountInput) (models.Accounts, error)
	GetAll(ctx context.Context) ([]models.Accounts, error)
	GetByID(ctx context.Context, ID uint32) (models.Accounts, error)
	Login(ctx context.Context, input models.LoginAccountInput) (models.Accounts, error)
	Registration(ctx context.Context, input models.RegistrationAccountInput) (models.Accounts, error)
	SendVerificationCode(ctx context.Context, input models.AccountSendVerificationInput) error
	DeleteByID(ctx context.Context, ID uint32) error
}

type Banners interface {
	Create(ctx context.Context, input models.CreateBannerInput) (models.Banners, error)
	Update(ctx context.Context, input models.UpdateBannerInput) (models.Banners, error)
	GetByID(ctx context.Context, ID uint32) (models.Banners, error)
	GetAll(ctx context.Context) ([]models.Banners, error)
	GetByKey(ctx context.Context, key string) (models.Banners, error)
	DeleteByID(ctx context.Context, ID uint32) error
}

type Files interface {
	Save(ctx context.Context, file models.File) (string, error)
	DeleteByName(ctx context.Context, path, filename string) error
}

type Groups interface {
	Create(ctx context.Context, input models.CreateGroupInput) (models.Groups, error)
	Update(ctx context.Context, input models.UpdateGroupInput) (models.Groups, error)
	GetAll(ctx context.Context) ([]models.Groups, error)
	GetAllByFilterWithChild(ctx context.Context, input models.GetGroupsByFilterInput) ([]models.GroupsWithChild, error)
	GetByID(ctx context.Context, ID string) (models.Groups, error)
	GetByUrl(ctx context.Context, url string) (models.Groups, error)
	DeleteByID(ctx context.Context, ID string) error
}

type OrderProducts interface {
	Create(ctx context.Context, input models.CreateOrderProductsInput) (models.OrderProducts, error)
	Update(ctx context.Context, input models.UpdateOrderProductsInput) (models.OrderProducts, error)
	GetAllByOrderID(ctx context.Context, orderID uint64) ([]models.OrderProducts, error)
	GetByID(ctx context.Context, ID uint64) (models.OrderProducts, error)
	GetPriceSumByOrderID(ctx context.Context, orderID uint64) (float64, error)
	DeleteByID(ctx context.Context, ID uint64) error
}

type Orders interface {
	Create(ctx context.Context, input models.CreateOrderWithProductsInput) (models.OrderWithOrderProducts, error)
	Update(ctx context.Context, input models.UpdateOrderInput) (models.Orders, error)
	GetAll(ctx context.Context) ([]models.Orders, error)
	GetAllByFilter(ctx context.Context, input models.GetOrdersByFilterInput) ([]models.Orders, error)
	GetByID(ctx context.Context, ID uint64) (models.OrderWithOrderProducts, error)
	IncreaseFullSum(ctx context.Context, ID uint64, fullSum float64) error
	DeleteByID(ctx context.Context, ID uint64) error
}

type ProductImages interface {
	Create(ctx context.Context, input models.CreateProductImageInput) (models.ProductImages, error)
	Update(ctx context.Context, input models.UpdateProductImageInput) (models.ProductImages, error)
	GetAll(ctx context.Context) ([]models.ProductImages, error)
	GetAllByProductID(ctx context.Context, productID string) ([]models.ProductImages, error)
	GetByID(ctx context.Context, ID uint32) (models.ProductImages, error)
	DeleteAllByProductID(ctx context.Context, productID string) error
	DeleteByID(ctx context.Context, ID uint32) error
}

type ProductRates interface {
	Create(ctx context.Context, input models.CreateProductRateInput) (models.ProductRates, error)
	GetAll(ctx context.Context) ([]models.ProductRates, error)
	GetAllOrderByCreatedAtByProductID(ctx context.Context, productID string) ([]models.ProductRates, error)
	GetByID(ctx context.Context, ID uint64) (models.ProductRates, error)
}

type Products interface {
	Create(ctx context.Context, input models.CreateProductInput) (models.Products, error)
	Update(ctx context.Context, input models.UpdateProductInput) (models.Products, error)
	GetAll(ctx context.Context) ([]models.Products, error)
	// GetAllWithImagesByFilter GetAllByGroupUrlWithImages(ctx context.Context, url string) ([]models.ProductWithImages, error)
	GetAllWithImagesByFilter(ctx context.Context, input models.GetProductsByFilterInput) (models.ProductsWithImagesAndPagination, error)
	GetAllByIDs(ctx context.Context, input models.GetProductsByIDsInput) ([]models.ProductWithImages, error)
	GetByID(ctx context.Context, ID string) (models.Products, error)
	GetByIDWithImagesAndRates(ctx context.Context, ID string) (models.ProductWithImagesAndRates, error)
	GetByUrlWithImagesAndRates(ctx context.Context, url string) (models.ProductWithImagesAndRates, error)
	DeleteByID(ctx context.Context, ID string) error
}

type Settings interface {
	Create(ctx context.Context, input models.CreateSettingInput) (models.Settings, error)
	Update(ctx context.Context, input models.UpdateSettingInput) (models.Settings, error)
	GetByID(ctx context.Context, ID uint32) (models.Settings, error)
	GetAll(ctx context.Context) ([]models.Settings, error)
	GetByKey(ctx context.Context, key string) (models.Settings, error)
	DeleteByID(ctx context.Context, ID uint32) error
}

type Sms interface {
	SendVerificationCode(ctx context.Context, phone string) (string, error)
}

type Users interface {
	Login(ctx context.Context, input models.LoginUserInput) (models.Users, error)
}

type Deps struct {
	MysqlRepos   *mysql.Repositories
	RdbRepos     rdb.Repository
	Log          logger.Logger
	Cfg          config.Config
	HumanizerUrl humanizer.Url
	Hasher       hash.PasswordHasher
	Generator    gen.Generator
}

type Services struct {
	Accounts      Accounts
	Banners       Banners
	Files         Files
	Groups        Groups
	OrderProducts OrderProducts
	Orders        Orders
	ProductImages ProductImages
	ProductRates  ProductRates
	Products      Products
	Settings      Settings
	Sms           Sms
	Users         Users
}

func NewServices(deps Deps) *Services {
	smsService := NewSmsService(deps.MysqlRepos.Settings, deps.Generator, deps.Log)
	accountsService := NewAccountsService(deps.MysqlRepos.Accounts, deps.RdbRepos, smsService, deps.Hasher)
	filesService := NewFilesService(deps.Cfg)
	bannersService := NewBannersService(deps.MysqlRepos.Banners, filesService)
	groupsService := NewGroupsService(deps.MysqlRepos.Groups, filesService, deps.HumanizerUrl, deps.Log)
	productImagesService := NewProductImagesService(deps.MysqlRepos.ProductImages, filesService, deps.Log)
	settingsService := NewSettingsService(deps.MysqlRepos.Settings)
	productRatesService := NewProductRatesService(deps.MysqlRepos.ProductRates, accountsService)
	productsService := NewProductsService(deps.MysqlRepos.Products, filesService, groupsService, productImagesService, productRatesService, settingsService, deps.HumanizerUrl)
	orderProductsService := NewOrderProductsService(deps.MysqlRepos.OrderProducts, productsService)
	ordersService := NewOrdersService(deps.MysqlRepos.Orders, orderProductsService)
	usersService := NewUsersService(deps.MysqlRepos.Users, deps.Hasher)
	return &Services{
		Accounts:      accountsService,
		Banners:       bannersService,
		Files:         filesService,
		Groups:        groupsService,
		OrderProducts: orderProductsService,
		Orders:        ordersService,
		ProductImages: productImagesService,
		ProductRates:  productRatesService,
		Products:      productsService,
		Settings:      settingsService,
		Sms:           smsService,
		Users:         usersService,
	}
}
