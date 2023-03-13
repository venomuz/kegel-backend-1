package service

import (
	"context"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"time"
)

type ProductImagesService struct {
	productImagesRepo mysql.ProductImages
	filesService      Files
	log               logger.Logger
}

func NewProductImagesService(productImagesRepo mysql.ProductImages, filesService Files, log logger.Logger) *ProductImagesService {
	return &ProductImagesService{
		productImagesRepo: productImagesRepo,
		filesService:      filesService,
		log:               log,
	}
}

func (p *ProductImagesService) Create(ctx context.Context, input models.CreateProductImageInput) (models.ProductImages, error) {
	productImage := models.ProductImages{
		ProductID: input.ProductID,
		Position:  input.Position,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if input.FileImage != nil {
		name, err := p.filesService.Save(ctx, models.File{File: input.FileImage, Path: models.FilePathProducts})
		if err != nil {
			return models.ProductImages{}, err
		}
		productImage.Image = name
	}

	err := p.productImagesRepo.Create(ctx, &productImage)

	return productImage, err
}

func (p *ProductImagesService) Update(ctx context.Context, input models.UpdateProductImageInput) (models.ProductImages, error) {

	productImageInfo, err := p.productImagesRepo.GetByID(ctx, input.ID)
	if err != nil {
		return models.ProductImages{}, err
	}

	productImage := models.ProductImages{
		ID:        input.ID,
		ProductID: input.ProductID,
		Position:  input.Position,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if input.FileImage != nil {

		name, err := p.filesService.Save(ctx, models.File{File: input.FileImage, Path: models.FilePathProducts})
		if err != nil {
			return models.ProductImages{}, err
		}

		productImage.Image = name

		if productImageInfo.Image != "" {
			_ = p.filesService.DeleteByName(ctx, models.FilePathGroups, productImageInfo.Image)
		}
	}

	err = p.productImagesRepo.Update(ctx, &productImage)

	return productImage, err
}

func (p *ProductImagesService) GetAll(ctx context.Context) ([]models.ProductImages, error) {
	return p.productImagesRepo.GetAll(ctx)
}

func (p *ProductImagesService) GetAllByProductID(ctx context.Context, productID string) ([]models.ProductImages, error) {
	return p.productImagesRepo.GetAllByProductID(ctx, productID)
}

func (p *ProductImagesService) GetByID(ctx context.Context, ID uint32) (models.ProductImages, error) {
	return p.productImagesRepo.GetByID(ctx, ID)
}

func (p *ProductImagesService) DeleteAllByProductID(ctx context.Context, productID string) error {
	var IDs []uint32

	productImages, err := p.productImagesRepo.GetAllByProductID(ctx, productID)
	if productImages != nil {
		for _, productImage := range productImages {
			err = p.filesService.DeleteByName(ctx, models.FilePathProducts, productImage.Image)
			if err != nil {
				p.log.Error("error while delete image from products", logger.Error(err))
			}
			IDs = append(IDs, productImage.ID)
		}
		err = p.productImagesRepo.DeleteByIDs(ctx, IDs)
	}

	return err
}

func (p *ProductImagesService) DeleteByID(ctx context.Context, ID uint32) error {
	productImage, err := p.productImagesRepo.GetByID(ctx, ID)
	if err != nil {
		err = p.filesService.DeleteByName(ctx, models.FilePathProducts, productImage.Image)
		if err != nil {
			p.log.Error("error while delete image from products", logger.Error(err))
		}
	}

	return p.productImagesRepo.DeleteByID(ctx, ID)
}
