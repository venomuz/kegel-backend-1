package models

type ProductRates struct {
	ID               uint64 `json:"id"  gorm:"type:int unsigned not null;primaryKey" extensions:"x-order=1"`
	OrderID          uint64 `json:"orderId" gorm:"type:bigint not null" extensions:"x-order=2"`
	ProductID        string `json:"productId" gorm:"type:varchar(50) not null" extensions:"x-order=3"`
	AccountID        uint32 `json:"accountId" gorm:"type:int unsigned not null" extensions:"x-order=4"`
	AccountFirstname string `json:"accountFirstname" gorm:"type:varchar(150);default:'Anonymous'" extensions:"x-order=5"`
	Rate             int8   `json:"rate" gorm:"type:int8;default:5" extensions:"x-order=6"`
	Description      string `json:"description" gorm:"type:text" extensions:"x-order=7"`
	CreatedAt        string `json:"createdAt" gorm:"type:datetime;default:null" extensions:"x-order=8"`
}

type CreateProductRateInput struct {
	OrderID   uint64 `json:"orderId" extensions:"x-order=1"`
	ProductID string `json:"productId" extensions:"x-order=2"`
	AccountID uint32 `json:"-" extensions:"x-order=3"`
	//AccountFirstname string `json:"-" extensions:"x-order=4"`
	Noname      bool   `json:"noname" extensions:"x-order=4"`
	Rate        int8   `json:"rate" extensions:"x-order=5"`
	Description string `json:"description" extensions:"x-order=6"`
}

type ProductRatesResponse struct {
	AccountFirstname string `json:"accountFirstname"`
	Description      string `json:"description"`
	Rate             int8   `json:"rate"`
	CreatedAt        string `json:"createdAt"`
}
