package models

type Orders struct {
	ID            uint64   `json:"id" gorm:"type:bigint not null auto_increment;primaryKey"`
	System        string   `json:"system"  gorm:"type:varchar(20);default:'Web'"`
	AccountID     uint32   `json:"accountId" gorm:"type:int not null"`
	ChatID        *int64   `json:"chatId" gorm:"type:bigint;default:null"`
	RegionID      uint32   `json:"regionId" gorm:"type:int;default:null"`
	CustomerName  *string  `json:"customerName" gorm:"type:varchar(50);default:null"`
	CustomerPhone string   `json:"customerPhone" gorm:"type:varchar(30) not null"`
	City          string   `json:"city" gorm:"type:varchar(255);default:null"`
	District      string   `json:"district" gorm:"type:varchar(255);default:null"`
	Street        string   `json:"street" gorm:"type:varchar(255);default:null"`
	Home          string   `json:"home" gorm:"type:varchar(255);default:null"`
	Apartment     string   `json:"apartment" gorm:"type:varchar(255);default:null"`
	Comment       string   `json:"comment" gorm:"type:text;default:null"`
	PaymentType   string   `json:"paymentType" gorm:"type:varchar(255);default:null"`
	DeliveryPrice *float64 `json:"deliveryPrice" gorm:"type:decimal(10,2);default:null"`
	FullSum       float64  `json:"fullSum" gorm:"type:decimal(10,2) not null"`
	Status        int      `json:"status" gorm:"type:int;default:0"`
	CreatedAt     string   `json:"createdAt" gorm:"type:datetime;default:null"`
	UpdatedAt     string   `json:"updatedAt" gorm:"type:datetime;default:null"`
	DeletedAt     string   `json:"deletedAt" gorm:"type:datetime;default:null"`
}

type CreateOrderWithProductsInput struct {
	System        string                     `json:"system"`
	AccountID     uint32                     `json:"accountId" binding:"required"`
	ChatID        *int64                     `json:"chatId"`
	RegionID      uint32                     `json:"regionId"`
	CustomerName  *string                    `json:"customerName"`
	CustomerPhone string                     `json:"customerPhone"  binding:"required"`
	City          string                     `json:"city"`
	District      string                     `json:"district"`
	Street        string                     `json:"street"`
	Home          string                     `json:"home"`
	Apartment     string                     `json:"apartment"`
	Comment       string                     `json:"comment"`
	PaymentType   string                     `json:"paymentType"`
	DeliveryPrice *float64                   `json:"deliveryPrice"`
	FullSum       float64                    `json:"fullSum"`
	Status        int                        `json:"status"`
	OrderProducts []CreateOrderProductsInput `json:"orderProducts"`
}

type UpdateOrderInput struct {
	ID            uint64   `json:"-"`
	System        string   `json:"system"`
	AccountID     uint32   `json:"accountID"`
	ChatID        *int64   `json:"chatID"`
	RegionID      uint32   `json:"regionID"`
	CustomerName  *string  `json:"customerName"`
	CustomerPhone string   `json:"customerPhone"`
	City          string   `json:"city"`
	District      string   `json:"district"`
	Street        string   `json:"street"`
	Home          string   `json:"home"`
	Apartment     string   `json:"apartment"`
	Comment       string   `json:"comment"`
	PaymentType   string   `json:"paymentType"`
	DeliveryPrice *float64 `json:"deliveryPrice"`
	FullSum       float64  `json:"fullSum"`
	Status        int      `json:"status"`
}

type OrderWithOrderProducts struct {
	Orders        `extensions:"x-order=1"`
	OrderProducts []OrderProducts `json:"orderProducts"`
}

type GetOrdersByFilterInput struct {
}
