package models

import "mime/multipart"

type Banners struct {
	ID        uint32 `json:"id" gorm:"type:int unsigned not null;primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(255) not null"`
	Image     string `json:"image" gorm:"type:varchar(255);default:null"`
	Position  *int32 `json:"position" gorm:"type:int;default:null"`
	CreatedAt string `json:"createdAt" gorm:"type:datetime;default:null"`
	UpdatedAt string `json:"updatedAt" gorm:"type:datetime;default:null"`
	DeletedAt string `json:"deletedAt" gorm:"type:datetime;default:null"`
}

type CreateBannerInput struct {
	FileImage *multipart.FileHeader `form:"fileImage" swaggerignore:"true" binding:"required"`
	Position  *int32                `form:"position"`
	Name      string                `form:"name" binding:"required"`
}

type UpdateBannerInput struct {
	ID        uint32                `form:"-" json:"-"`
	FileImage *multipart.FileHeader `form:"fileImage" swaggerignore:"true"`
	Position  *int32                `form:"position"`
	Name      string                `form:"name" binding:"required"`
}
