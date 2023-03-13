package models

import "mime/multipart"

type Groups struct {
	ID             string `json:"id" gorm:"type:varchar(50) not null;primaryKey" extensions:"x-order=1"`
	Url            string `json:"url" gorm:"type:varchar(255) not null;unique" extensions:"x-order=2"`
	NameUz         string `json:"nameUz" gorm:"type:varchar(255) not null" extensions:"x-order=3"`
	NameRu         string `json:"nameRu"  gorm:"type:varchar(255);default:null" extensions:"x-order=3"`
	NameEn         string `json:"nameEn" gorm:"type:varchar(255);default:null" extensions:"x-order=4"`
	DescriptionUz  string `json:"descriptionUz" gorm:"type:text not null" extensions:"x-order=5"`
	DescriptionRu  string `json:"descriptionRu" gorm:"type:text;default:null" extensions:"x-order=6"`
	DescriptionEn  string `json:"descriptionEn" gorm:"type:text;default:null" extensions:"x-order=7"`
	Position       *int32 `json:"position" gorm:"type:int;default:null" extensions:"x-order=8"`
	ParentGroup    string `json:"parentGroup" gorm:"type:varchar(50);default:null" extensions:"x-order=9"`
	Image          string `json:"image" gorm:"type:varchar(255);default:null" extensions:"x-order=10"`
	SeoDescription string `json:"seoDescription" gorm:"type:text;default:null" extensions:"x-order=11"`
	SeoKeywords    string `json:"seoKeywords" gorm:"type:text;default:null" extensions:"x-order=12"`
	SeoText        string `json:"seoText" gorm:"type:text;default:null" extensions:"x-order=13"`
	SeoTitle       string `json:"seoTitle" gorm:"type:text;default:null" extensions:"x-order=14"`
	Enabled        *bool  `json:"enabled" gorm:"type:bool;default:true" extensions:"x-order=15"`
	CreatedAt      string `json:"createdAt" gorm:"type:datetime;default:null" extensions:"x-order=16"`
	UpdatedAt      string `json:"updatedAt" gorm:"type:datetime;default:null" extensions:"x-order=17"`
	DeletedAt      string `json:"deletedAt" gorm:"type:datetime;default:null" extensions:"x-order=18"`
}

type CreateGroupInput struct {
	Url            string                `form:"url" extensions:"x-order=1"`
	NameUz         string                `form:"nameUz" binding:"required" extensions:"x-order=2"`
	NameRu         string                `form:"nameRu" extensions:"x-order=3"`
	NameEn         string                `form:"nameEn" extensions:"x-order=3"`
	DescriptionUz  string                `form:"descriptionUz" binding:"required" extensions:"x-order=4"`
	DescriptionRu  string                `form:"descriptionRu" extensions:"x-order=5"`
	DescriptionEn  string                `form:"descriptionEn" extensions:"x-order=6"`
	Position       *int32                `form:"position" extensions:"x-order=7"`
	ParentGroup    string                `form:"parentGroup" extensions:"x-order=8"`
	FileImage      *multipart.FileHeader `form:"fileImage" swaggerignore:"true" extensions:"x-order=9"`
	SeoDescription string                `form:"seoDescription" extensions:"x-order=10"`
	SeoKeywords    string                `form:"seoKeywords" extensions:"x-order=11"`
	SeoText        string                `form:"seoText" extensions:"x-order=12"`
	SeoTitle       string                `form:"seoTitle" extensions:"x-order=13"`
	Enabled        *bool                 `form:"enabled" extensions:"x-order=14"`
}

type UpdateGroupInput struct {
	ID             string                `json:"-" form:"-" extensions:"x-order=1"`
	Url            string                `form:"url" extensions:"x-order=2"`
	NameUz         string                `form:"nameUz" binding:"required" extensions:"x-order=3"`
	NameRu         string                `form:"nameRu" extensions:"x-order=3"`
	NameEn         string                `form:"nameEn" extensions:"x-order=4"`
	DescriptionUz  string                `form:"descriptionUz" binding:"required" extensions:"x-order=5"`
	DescriptionRu  string                `form:"descriptionRu" extensions:"x-order=6"`
	DescriptionEn  string                `form:"descriptionEn" extensions:"x-order=7"`
	Position       *int32                `form:"position" extensions:"x-order=8"`
	ParentGroup    string                `form:"parentGroup" extensions:"x-order=9"`
	FileImage      *multipart.FileHeader `form:"fileImage" swaggerignore:"true" extensions:"x-order=10"`
	SeoDescription string                `form:"seoDescription" extensions:"x-order=11"`
	SeoKeywords    string                `form:"seoKeywords" extensions:"x-order=12"`
	SeoText        string                `form:"seoText" extensions:"x-order=13"`
	SeoTitle       string                `form:"seoTitle" extensions:"x-order=14"`
	Enabled        *bool                 `form:"enabled" extensions:"x-order=15"`
}

type GetGroupsByFilterInput struct {
	All      bool   `form:"all" extensions:"x-order=1"`
	ParentId string `form:"parentId" extensions:"x-order=2"`
	NameUz   string `form:"nameUz" extensions:"x-order=3"`
	Url      string `form:"url" extensions:"x-order=3"`
	Enabled  *bool  `form:"enabled" extensions:"x-order=4"`
}

type GroupsWithChild struct {
	Groups `extensions:"x-order=1"`
	Child  []Groups `json:"child" extensions:"x-order=2"`
}
