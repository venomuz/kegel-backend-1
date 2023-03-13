package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"github.com/venomuz/kegel-backend/pkg/humanizer"
	"math"
	"strconv"
	"time"
)

type ProductsService struct {
	productsRepo         mysql.Products
	filesService         Files
	groupsService        Groups
	productImagesService ProductImages
	productRatesService  ProductRates
	settingsService      Settings
	humanizerUrl         humanizer.Url
}

func NewProductsService(productsRepo mysql.Products, filesService Files, groupsService Groups, productImagesService ProductImages, productRatesService ProductRates, settingsService Settings, humanizerUrl humanizer.Url) *ProductsService {
	return &ProductsService{
		productsRepo:         productsRepo,
		filesService:         filesService,
		groupsService:        groupsService,
		productImagesService: productImagesService,
		productRatesService:  productRatesService,
		settingsService:      settingsService,
		humanizerUrl:         humanizerUrl,
	}
}

func (p *ProductsService) Create(ctx context.Context, input models.CreateProductInput) (models.Products, error) {

	product := models.Products{
		ID:             uuid.New().String(),
		Url:            input.Url,
		NameUz:         input.NameUz,
		NameRu:         input.NameRu,
		NameEn:         input.NameEn,
		DescriptionUz:  input.DescriptionUz,
		DescriptionRu:  input.DescriptionRu,
		DescriptionEn:  input.DescriptionEn,
		Position:       input.Position,
		GroupID:        input.GroupID,
		Price:          input.Price,
		SeoDescription: input.SeoDescription,
		SeoKeywords:    input.SeoKeywords,
		SeoText:        input.SeoText,
		SeoTitle:       input.SeoTitle,
		Enabled:        input.Enabled,
		CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
	}

	if input.Url == "" {
		count, _ := p.productsRepo.GetCountNameUz(ctx, input.NameUz)

		url := p.humanizerUrl.Generate(input.NameUz)
		if count != 0 {
			url = url + fmt.Sprintf("-%d", count+1)
		}
		product.Url = url
	}

	err := p.productsRepo.Create(ctx, &product)

	return product, err

}

func (p *ProductsService) Update(ctx context.Context, input models.UpdateProductInput) (models.Products, error) {
	_, err := p.productsRepo.GetByID(ctx, input.ID)
	if err != nil {
		return models.Products{}, err
	}

	product := models.Products{
		ID:             input.ID,
		Url:            input.Url,
		NameUz:         input.NameUz,
		NameRu:         input.NameRu,
		NameEn:         input.NameEn,
		DescriptionUz:  input.DescriptionUz,
		DescriptionRu:  input.DescriptionRu,
		DescriptionEn:  input.DescriptionEn,
		Position:       input.Position,
		GroupID:        input.GroupID,
		Price:          input.Price,
		SeoDescription: input.SeoDescription,
		SeoKeywords:    input.SeoKeywords,
		SeoText:        input.SeoText,
		SeoTitle:       input.SeoTitle,
		Enabled:        input.Enabled,
		UpdatedAt:      time.Now().Format("2006-01-02 15:04:05"),
	}

	err = p.productsRepo.Update(ctx, &product)

	return product, err
}

func (p *ProductsService) GetAll(ctx context.Context) ([]models.Products, error) {
	return p.productsRepo.GetAll(ctx)
}

//func (p *ProductsService) GetAllByGroupUrlWithImages(ctx context.Context, url string) ([]models.ProductWithImages, error) {
//	var productWithImages []models.ProductWithImages
//
//	group, err := p.groupsService.GetByUrlWithImagesAndRates(ctx, url)
//	if err != nil {
//		return nil, err
//	}
//
//	products, err := p.productsRepo.GetAllByFilterAndGroupID(ctx, group.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	images, err := p.productImagesService.GetAll(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, product := range products {
//		productsWithImage := models.ProductWithImages{
//			Products: product,
//		}
//
//		for _, image := range images {
//			if product.ID == image.ProductID {
//				productsWithImage.Images = append(productsWithImage.Images, image)
//			}
//		}
//
//		if productsWithImage.Images == nil {
//			productsWithImage.Images = []models.ProductImages{}
//		}
//
//		productWithImages = append(productWithImages, productsWithImage)
//	}
//
//	return productWithImages, err
//}

func (p *ProductsService) GetAllWithImagesByFilter(ctx context.Context, input models.GetProductsByFilterInput) (models.ProductsWithImagesAndPagination, error) {
	var productsWithImagesRes models.ProductsWithImagesAndPagination

	images, err := p.productImagesService.GetAll(ctx)
	if err != nil {
		return models.ProductsWithImagesAndPagination{}, err
	}

	setting, err := p.settingsService.GetByKey(ctx, "dollar")

	dollar, err := strconv.ParseFloat(setting.Value, 64)
	if err != nil {
		return models.ProductsWithImagesAndPagination{}, err
	}

	if input.All {
		products, err := p.productsRepo.GetAll(ctx)
		if err != nil {
			return models.ProductsWithImagesAndPagination{}, err
		}

		for _, product := range products {
			product.Price = 1000 * math.Round(product.Price/1000*dollar)

			productsWithImage := models.ProductWithImages{
				Products: product,
			}

			for _, image := range images {
				if product.ID == image.ProductID {
					productsWithImage.Images = append(productsWithImage.Images, image)
				}
			}

			if productsWithImage.Images == nil {
				productsWithImage.Images = []models.ProductImages{}
			}

			productsWithImagesRes.Products = append(productsWithImagesRes.Products, productsWithImage)
			productsWithImagesRes.Count = int64(len(productsWithImagesRes.Products))
		}
		return productsWithImagesRes, err
	}

	if input.Page == 0 {
		input.Page = 1
	}

	if input.PageSize == 0 {
		input.PageSize = 10
	}
	if input.GroupUrl != "" {

		group, err := p.groupsService.GetByUrl(ctx, input.GroupUrl)
		if err != nil {
			return models.ProductsWithImagesAndPagination{}, err
		}

		input.GroupID = group.ID
	}

	products, err := p.productsRepo.GetAllByFilter(ctx, input)
	for _, product := range products {
		product.Price = 1000 * math.Round(product.Price/1000*dollar)

		productsWithImage := models.ProductWithImages{
			Products: product,
		}

		for _, image := range images {
			if product.ID == image.ProductID {
				productsWithImage.Images = append(productsWithImage.Images, image)
			}
		}

		if productsWithImage.Images == nil {
			productsWithImage.Images = []models.ProductImages{}
		}

		productsWithImagesRes.Products = append(productsWithImagesRes.Products, productsWithImage)
	}

	count, err := p.productsRepo.GetCount(ctx)

	productsWithImagesRes.Count = count

	productsWithImagesRes.Page = input.Page

	productsWithImagesRes.PageSize = input.PageSize

	productsWithImagesRes.PageCount = int(math.Ceil(float64(count) / float64(input.PageSize)))

	if productsWithImagesRes.Products == nil {
		productsWithImagesRes.Products = []models.ProductWithImages{}
	}

	return productsWithImagesRes, err
}

func (p *ProductsService) GetByID(ctx context.Context, ID string) (models.Products, error) {
	return p.productsRepo.GetByID(ctx, ID)
}

func (p *ProductsService) GetByIDWithImagesAndRates(ctx context.Context, ID string) (models.ProductWithImagesAndRates, error) {
	var productWithImagesAndRates models.ProductWithImagesAndRates

	setting, err := p.settingsService.GetByKey(ctx, "dollar")

	dollar, err := strconv.ParseFloat(setting.Value, 64)
	if err != nil {
		return models.ProductWithImagesAndRates{}, err
	}

	product, err := p.productsRepo.GetByID(ctx, ID)
	if err != nil {
		return models.ProductWithImagesAndRates{}, err
	}

	images, err := p.productImagesService.GetAllByProductID(ctx, product.ID)
	if err != nil {
		return models.ProductWithImagesAndRates{}, err
	}

	rates, err := p.productRatesService.GetAllOrderByCreatedAtByProductID(ctx, product.ID)
	if err != nil {
		return models.ProductWithImagesAndRates{}, err
	}

	for _, rate := range rates {
		res := models.ProductRatesResponse{
			AccountFirstname: rate.AccountFirstname,
			Description:      rate.Description,
			Rate:             rate.Rate,
			CreatedAt:        rate.CreatedAt,
		}
		productWithImagesAndRates.Rates = append(productWithImagesAndRates.Rates, res)
	}

	if images == nil {
		images = []models.ProductImages{}
	}

	if productWithImagesAndRates.Rates == nil {
		productWithImagesAndRates.Rates = []models.ProductRatesResponse{}
	}

	product.Price = 1000 * math.Round(product.Price/1000*dollar)

	productWithImagesAndRates.Products = product

	productWithImagesAndRates.Images = images

	return productWithImagesAndRates, err
}

func (p *ProductsService) GetByUrlWithImagesAndRates(ctx context.Context, url string) (models.ProductWithImagesAndRates, error) {
	var productWithImagesAndRates models.ProductWithImagesAndRates

	setting, err := p.settingsService.GetByKey(ctx, "dollar")

	dollar, err := strconv.ParseFloat(setting.Value, 64)
	if err != nil {
		return models.ProductWithImagesAndRates{}, err
	}

	product, err := p.productsRepo.GetByUrl(ctx, url)
	if err != nil {
		return models.ProductWithImagesAndRates{}, err
	}

	images, err := p.productImagesService.GetAllByProductID(ctx, product.ID)
	if err != nil {
		return models.ProductWithImagesAndRates{}, err
	}

	rates, err := p.productRatesService.GetAllOrderByCreatedAtByProductID(ctx, product.ID)
	if err != nil {
		return models.ProductWithImagesAndRates{}, err
	}

	for _, rate := range rates {
		res := models.ProductRatesResponse{
			AccountFirstname: rate.AccountFirstname,
			Description:      rate.Description,
			Rate:             rate.Rate,
			CreatedAt:        rate.CreatedAt,
		}
		productWithImagesAndRates.Rates = append(productWithImagesAndRates.Rates, res)
	}

	if images == nil {
		images = []models.ProductImages{}
	}

	product.Price = 1000 * math.Round(product.Price/1000*dollar)

	productWithImagesAndRates.Products = product

	productWithImagesAndRates.Images = images

	return productWithImagesAndRates, err
}

func (p *ProductsService) DeleteByID(ctx context.Context, ID string) error {

	err := p.productImagesService.DeleteAllByProductID(ctx, ID)
	if err != nil {
		return err
	}

	return p.productsRepo.DeleteByID(ctx, ID)
}
