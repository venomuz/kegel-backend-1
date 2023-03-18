package models

type Products struct {
	ID             string  `json:"id" gorm:"type:varchar(50) not null;primaryKey" extensions:"x-order=1"`
	Url            string  `json:"url"  gorm:"type:varchar(255) not null;unique" extensions:"x-order=2"`
	NameUz         string  `json:"nameUz"  gorm:"type:varchar(255) not null;" extensions:"x-order=3"`
	NameRu         string  `json:"nameRu"  gorm:"type:varchar(255);default:null" extensions:"x-order=4"`
	NameEn         string  `json:"nameEn"  gorm:"type:varchar(255);default:null" extensions:"x-order=5"`
	DescriptionUz  string  `json:"descriptionUz" gorm:"type:text not null;" extensions:"x-order=6"`
	DescriptionRu  string  `json:"descriptionRu" gorm:"type:text;default:null" extensions:"x-order=7"`
	DescriptionEn  string  `json:"descriptionEn" gorm:"type:text;default:null" extensions:"x-order=8"`
	Position       *int32  `json:"position" gorm:"type:int;default:null" extensions:"x-order=9"`
	GroupID        string  `json:"groupId" gorm:"type:varchar(50) not null;" extensions:"x-order=10"`
	Price          float64 `json:"price" gorm:"type:decimal(11,2) not null" extensions:"x-order=11"`
	Rate           int8    `json:"rate" gorm:"-:migration"`
	SeoDescription string  `json:"seoDescription" gorm:"type:text;default:null" extensions:"x-order=12"`
	SeoKeywords    string  `json:"seoKeywords" gorm:"type:text;default:null" extensions:"x-order=13"`
	SeoText        string  `json:"seoText" gorm:"type:text;default:null" extensions:"x-order=14"`
	SeoTitle       string  `json:"seoTitle" gorm:"type:text;default:null" extensions:"x-order=15"`
	Enabled        *bool   `json:"enabled" gorm:"type:bool;default:true" extensions:"x-order=16"`
	CreatedAt      string  `json:"createdAt" gorm:"type:datetime;default:null" extensions:"x-order=17"`
	UpdatedAt      string  `json:"updatedAt" gorm:"type:datetime;default:null" extensions:"x-order=18"`
	DeletedAt      string  `json:"deletedAt" gorm:"type:datetime;default:null" extensions:"x-order=19"`
}

type ProductWithImages struct {
	Products
	Images []ProductImages `json:"images" extensions:"x-order=21"`
}

type ProductWithImagesAndRates struct {
	Products
	Images []ProductImages        `json:"images" extensions:"x-order=21"`
	Rates  []ProductRatesResponse `json:"rates" extensions:"x-order=22"`
}

type CreateProductInput struct {
	Url            string  `json:"url" extensions:"x-order=1"`
	NameUz         string  `json:"nameUz" binding:"required" extensions:"x-order=2"`
	NameRu         string  `json:"nameRu" extensions:"x-order=3"`
	NameEn         string  `json:"nameEn" extensions:"x-order=4"`
	DescriptionUz  string  `json:"descriptionUz" binding:"required" extensions:"x-order=5"`
	DescriptionRu  string  `json:"descriptionRu" extensions:"x-order=6"`
	DescriptionEn  string  `json:"descriptionEn" extensions:"x-order=7"`
	Position       *int32  `json:"position" extensions:"x-order=8"`
	GroupID        string  `json:"groupId" binding:"required" extensions:"x-order=9"`
	Price          float64 `json:"price"  extensions:"x-order=10"`
	SeoDescription string  `json:"seoDescription" extensions:"x-order=11"`
	SeoKeywords    string  `json:"seoKeywords" extensions:"x-order=12"`
	SeoText        string  `json:"seoText" extensions:"x-order=13"`
	SeoTitle       string  `json:"seoTitle" extensions:"x-order=14"`
	Enabled        *bool   `json:"enabled" extensions:"x-order=15"`
}

type UpdateProductInput struct {
	ID             string  `json:"-" extensions:"x-order=1"`
	Url            string  `json:"url" extensions:"x-order=2"`
	NameUz         string  `json:"nameUz" binding:"required" extensions:"x-order=3"`
	NameRu         string  `json:"nameRu" extensions:"x-order=4"`
	NameEn         string  `json:"nameEn" extensions:"x-order=5"`
	DescriptionUz  string  `json:"descriptionUz" binding:"required" extensions:"x-order=6"`
	DescriptionRu  string  `json:"descriptionRu" extensions:"x-order=7"`
	DescriptionEn  string  `json:"descriptionEn" extensions:"x-order=8"`
	Position       *int32  `json:"position" extensions:"x-order=9"`
	GroupID        string  `json:"groupId" binding:"required" extensions:"x-order=10"`
	Price          float64 `json:"price"  extensions:"x-order=11"`
	SeoDescription string  `json:"seoDescription" extensions:"x-order=12"`
	SeoKeywords    string  `json:"seoKeywords" extensions:"x-order=13"`
	SeoText        string  `json:"seoText" extensions:"x-order=14"`
	SeoTitle       string  `json:"seoTitle" extensions:"x-order=15"`
	Enabled        *bool   `json:"enabled" extensions:"x-order=16"`
}

type GetProductsByFilterInput struct {
	NameUz        string   `form:"nameUz" extensions:"x-order=1"`
	Url           string   `form:"url" extensions:"x-order=2"`
	Price         *float64 `form:"price" extensions:"x-order=3"`
	DescriptionUz string   `form:"descriptionUz" extensions:"x-order=4"`
	GroupID       string   `form:"-" json:"-" extensions:"x-order=5"`
	GroupUrl      string   `form:"groupUrl" extensions:"x-order=6"`
	Enabled       *bool    `form:"enabled" extensions:"x-order=7"`
	SortBy        string   `form:"sortBy" enums:"name_uz,price,created_at" extensions:"x-order=8"`
	Desc          bool     `form:"desc" extensions:"x-order=9"`
	All           bool     `form:"all" extensions:"x-order=10"`
	Page          int      `form:"page" extensions:"x-order=11"`
	PageSize      int      `form:"pageSize" extensions:"x-order=12"`
}

type GetProductsByIDsInput struct {
	IDs []string `json:"ids"`
}

type ProductsWithImagesAndPagination struct {
	Products  []ProductWithImages `json:"products" extensions:"x-order=1"`
	Page      int                 `json:"page" extensions:"x-order=2"`
	PageSize  int                 `json:"pageSize" extensions:"x-order=3"`
	PageCount int                 `json:"pageCount" extensions:"x-order=4"`
	Count     int64               `json:"count" extensions:"x-order=5"`
}
