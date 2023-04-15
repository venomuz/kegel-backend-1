package models

import "mime/multipart"

type Groups struct {
	ID             string `json:"id" gorm:"type:varchar(50) not null;primaryKey"`
	Url            string `json:"url" gorm:"type:varchar(255) not null;unique"`
	NameUz         string `json:"nameUz" gorm:"type:varchar(255) not null"`
	NameRu         string `json:"nameRu"  gorm:"type:varchar(255);default:null"`
	NameEn         string `json:"nameEn" gorm:"type:varchar(255);default:null"`
	DescriptionUz  string `json:"descriptionUz" gorm:"type:text not null"`
	DescriptionRu  string `json:"descriptionRu" gorm:"type:text;default:null"`
	DescriptionEn  string `json:"descriptionEn" gorm:"type:text;default:null"`
	Position       *int32 `json:"position" gorm:"type:int;default:null"`
	ParentGroup    string `json:"parentGroup" gorm:"type:varchar(50);default:null"`
	Image          string `json:"image" gorm:"type:varchar(255);default:null"`
	SeoDescription string `json:"seoDescription" gorm:"type:text;default:null"`
	SeoKeywords    string `json:"seoKeywords" gorm:"type:text;default:null"`
	SeoText        string `json:"seoText" gorm:"type:text;default:null"`
	SeoTitle       string `json:"seoTitle" gorm:"type:text;default:null"`
	Enabled        *bool  `json:"enabled" gorm:"type:bool;default:true"`
	CreatedAt      string `json:"createdAt" gorm:"type:datetime;default:null"`
	UpdatedAt      string `json:"updatedAt" gorm:"type:datetime;default:null"`
	DeletedAt      string `json:"deletedAt" gorm:"type:datetime;default:null"`
}

type CreateGroupInput struct {
	Url            string                `form:"url"`
	NameUz         string                `form:"nameUz" binding:"required"`
	NameRu         string                `form:"nameRu"`
	NameEn         string                `form:"nameEn"`
	DescriptionUz  string                `form:"descriptionUz" binding:"required"`
	DescriptionRu  string                `form:"descriptionRu"`
	DescriptionEn  string                `form:"descriptionEn"`
	Position       *int32                `form:"position"`
	ParentGroup    string                `form:"parentGroup"`
	FileImage      *multipart.FileHeader `form:"fileImage" swaggerignore:"true"`
	SeoDescription string                `form:"seoDescription"`
	SeoKeywords    string                `form:"seoKeywords"`
	SeoText        string                `form:"seoText"`
	SeoTitle       string                `form:"seoTitle"`
	Enabled        *bool                 `form:"enabled"`
}

type UpdateGroupInput struct {
	ID             string                `json:"-" form:"-"`
	Url            string                `form:"url"`
	NameUz         string                `form:"nameUz" binding:"required"`
	NameRu         string                `form:"nameRu"`
	NameEn         string                `form:"nameEn"`
	DescriptionUz  string                `form:"descriptionUz" binding:"required"`
	DescriptionRu  string                `form:"descriptionRu"`
	DescriptionEn  string                `form:"descriptionEn"`
	Position       *int32                `form:"position"`
	ParentGroup    string                `form:"parentGroup"`
	FileImage      *multipart.FileHeader `form:"fileImage" swaggerignore:"true"`
	SeoDescription string                `form:"seoDescription"`
	SeoKeywords    string                `form:"seoKeywords"`
	SeoText        string                `form:"seoText"`
	SeoTitle       string                `form:"seoTitle"`
	Enabled        *bool                 `form:"enabled"`
}

type GetGroupsByFilterInput struct {
	All      bool   `form:"all"`
	ParentId string `form:"parentId"`
	NameUz   string `form:"nameUz"`
	Url      string `form:"url"`
	Enabled  *bool  `form:"enabled"`
}

type GroupsWithChild struct {
	Groups `extensions:"x-order=1"`
	Child  []Groups `json:"child"`
}
