package models

type Products struct {
	ID             string  `json:"id" gorm:"type:varchar(50) not null;primaryKey"`
	Url            string  `json:"url"  gorm:"type:varchar(255) not null;unique"`
	NameUz         string  `json:"nameUz"  gorm:"type:varchar(255) not null;"`
	NameRu         string  `json:"nameRu"  gorm:"type:varchar(255);default:null"`
	NameEn         string  `json:"nameEn"  gorm:"type:varchar(255);default:null"`
	DescriptionUz  string  `json:"descriptionUz" gorm:"type:text not null"`
	DescriptionRu  string  `json:"descriptionRu" gorm:"type:text;default:null"`
	DescriptionEn  string  `json:"descriptionEn" gorm:"type:text;default:null"`
	Position       *int32  `json:"position" gorm:"type:int;default:null"`
	GroupID        string  `json:"groupId" gorm:"type:varchar(50) not null;"`
	Price          float64 `json:"price" gorm:"type:decimal(11,2) not null"`
	Rate           int8    `json:"rate" gorm:"-:migration"`
	SeoDescription string  `json:"seoDescription" gorm:"type:text;default:null"`
	SeoKeywords    string  `json:"seoKeywords" gorm:"type:text;default:null"`
	SeoText        string  `json:"seoText" gorm:"type:text;default:null"`
	SeoTitle       string  `json:"seoTitle" gorm:"type:text;default:null"`
	Enabled        *bool   `json:"enabled" gorm:"type:bool;default:true"`
	CreatedAt      string  `json:"createdAt" gorm:"type:datetime;default:null"`
	UpdatedAt      string  `json:"updatedAt" gorm:"type:datetime;default:null"`
	DeletedAt      string  `json:"deletedAt" gorm:"type:datetime;default:null"`
}

type ProductWithImages struct {
	Products
	Images []ProductImages `json:"images"`
}

type ProductWithImagesAndRates struct {
	Products
	Images []ProductImages        `json:"images"`
	Rates  []ProductRatesResponse `json:"rates"`
}

type CreateProductInput struct {
	Url            string  `json:"url"`
	NameUz         string  `json:"nameUz" binding:"required"`
	NameRu         string  `json:"nameRu"`
	NameEn         string  `json:"nameEn"`
	DescriptionUz  string  `json:"descriptionUz" binding:"required"`
	DescriptionRu  string  `json:"descriptionRu"`
	DescriptionEn  string  `json:"descriptionEn"`
	Position       *int32  `json:"position"`
	GroupID        string  `json:"groupId" binding:"required"`
	Price          float64 `json:"price" `
	SeoDescription string  `json:"seoDescription"`
	SeoKeywords    string  `json:"seoKeywords"`
	SeoText        string  `json:"seoText"`
	SeoTitle       string  `json:"seoTitle"`
	Enabled        *bool   `json:"enabled"`
}

type UpdateProductInput struct {
	ID             string  `json:"-"`
	Url            string  `json:"url"`
	NameUz         string  `json:"nameUz" binding:"required"`
	NameRu         string  `json:"nameRu"`
	NameEn         string  `json:"nameEn"`
	DescriptionUz  string  `json:"descriptionUz" binding:"required"`
	DescriptionRu  string  `json:"descriptionRu"`
	DescriptionEn  string  `json:"descriptionEn"`
	Position       *int32  `json:"position"`
	GroupID        string  `json:"groupId" binding:"required"`
	Price          float64 `json:"price"`
	SeoDescription string  `json:"seoDescription"`
	SeoKeywords    string  `json:"seoKeywords"`
	SeoText        string  `json:"seoText"`
	SeoTitle       string  `json:"seoTitle"`
	Enabled        *bool   `json:"enabled"`
}

type GetProductsByFilterInput struct {
	NameUz        string   `form:"nameUz"`
	Url           string   `form:"url"`
	Price         *float64 `form:"price"`
	DescriptionUz string   `form:"descriptionUz"`
	GroupID       string   `form:"-" json:"-"`
	GroupUrl      string   `form:"groupUrl"`
	Enabled       *bool    `form:"enabled"`
	Dollar        bool     `form:"dollar"`
	SortBy        string   `form:"sortBy" enums:"name_uz,price,created_at"`
	Desc          bool     `form:"desc"`
	All           bool     `form:"all"`
	Page          int      `form:"page"`
	PageSize      int      `form:"pageSize"`
}

type GetProductsByIDsInput struct {
	IDs []string `json:"ids"`
}

type ProductsWithImagesAndPagination struct {
	Products  []ProductWithImages `json:"products"`
	Page      int                 `json:"page"`
	PageSize  int                 `json:"pageSize"`
	PageCount int                 `json:"pageCount"`
	Count     int64               `json:"count"`
}
