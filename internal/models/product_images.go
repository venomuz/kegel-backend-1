package models

import "mime/multipart"

type ProductImages struct {
	ID        uint32 `json:"id" gorm:"type:int unsigned not null;primaryKey" extensions:"x-order=1"`
	ProductID string `json:"productId" gorm:"type:varchar(50) not null;index" extensions:"x-order=2"`
	Image     string `json:"image" gorm:"type:varchar(50) not null" extensions:"x-order=3"`
	Position  *int32 `json:"position" gorm:"type:int;default:null" extensions:"x-order=4"`
	CreatedAt string `json:"createdAt" gorm:"type:datetime;default:null" extensions:"x-order=5"`
	UpdatedAt string `json:"updatedAt" gorm:"type:datetime;default:null" extensions:"x-order=6"`
	DeletedAt string `json:"deletedAt" gorm:"type:datetime;default:null" extensions:"x-order=7"`
}

type CreateProductImageInput struct {
	ProductID string                `form:"productId" json:"productId" binding:"required" extensions:"x-order=1"`
	Position  *int32                `form:"position" extensions:"x-order=2"`
	FileImage *multipart.FileHeader `form:"fileImage" json:"fileImage" binding:"required" swaggerignore:"true" extensions:"x-order=3"`
}

type UpdateProductImageInput struct {
	ID        uint32                `form:"-" json:"-"`
	ProductID string                `form:"productId" json:"productId" binding:"required" extensions:"x-order=1"`
	Position  *int32                `form:"position" extensions:"x-order=2"`
	FileImage *multipart.FileHeader `form:"fileImage" json:"fileImage"  binding:"required" swaggerignore:"true" extensions:"x-order=3"`
}
