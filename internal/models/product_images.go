package models

import "mime/multipart"

type ProductImages struct {
	ID        uint32 `json:"id" gorm:"type:int unsigned not null;primaryKey"`
	ProductID string `json:"productId" gorm:"type:varchar(50) not null;index"`
	Image     string `json:"image" gorm:"type:varchar(50) not null"`
	Position  *int32 `json:"position" gorm:"type:int;default:null"`
	CreatedAt string `json:"createdAt" gorm:"type:datetime;default:null"`
	UpdatedAt string `json:"updatedAt" gorm:"type:datetime;default:null"`
	DeletedAt string `json:"deletedAt" gorm:"type:datetime;default:null"`
}

type CreateProductImageInput struct {
	ProductID string                `form:"productId" json:"productId" binding:"required"`
	Position  *int32                `form:"position"`
	FileImage *multipart.FileHeader `form:"fileImage" json:"fileImage" binding:"required" swaggerignore:"true"`
}

type UpdateProductImageInput struct {
	ID        uint32                `form:"-" json:"-"`
	ProductID string                `form:"productId" json:"productId" binding:"required"`
	Position  *int32                `form:"position"`
	FileImage *multipart.FileHeader `form:"fileImage" json:"fileImage"  binding:"required" swaggerignore:"true"`
}
