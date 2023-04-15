package models

type ProductRates struct {
	ID               uint64 `json:"id"  gorm:"type:int unsigned not null;primaryKey"`
	OrderID          uint64 `json:"orderId" gorm:"type:bigint not null"`
	ProductID        string `json:"productId" gorm:"type:varchar(50) not null"`
	AccountID        uint32 `json:"accountId" gorm:"type:int unsigned not null"`
	AccountFirstname string `json:"accountFirstname" gorm:"type:varchar(150);default:'Anonymous'"`
	Rate             int8   `json:"rate" gorm:"type:int8;default:5"`
	Description      string `json:"description" gorm:"type:text"`
	CreatedAt        string `json:"createdAt" gorm:"type:datetime;default:null"`
}

type CreateProductRateInput struct {
	OrderID   uint64 `json:"orderId"`
	ProductID string `json:"productId"`
	AccountID uint32 `json:"-"`
	//AccountFirstname string `json:"-"`
	Noname      bool   `json:"noname"`
	Rate        int8   `json:"rate"`
	Description string `json:"description"`
}

type ProductRatesResponse struct {
	AccountFirstname string `json:"accountFirstname"`
	Description      string `json:"description"`
	Rate             int8   `json:"rate"`
	CreatedAt        string `json:"createdAt"`
}
