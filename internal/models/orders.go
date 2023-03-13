package models

type Orders struct {
	ID            uint64   `json:"id" gorm:"type:bigint not null auto_increment;primaryKey" extensions:"x-order=1"`
	System        string   `json:"system"  gorm:"type:varchar(20);default:'Web'" extensions:"x-order=2"`
	AccountID     uint32   `json:"accountID" gorm:"type:int not null" extensions:"x-order=3"`
	ChatID        *int64   `json:"chatID" gorm:"type:bigint;default:null" extensions:"x-order=3"`
	RegionID      uint32   `json:"regionID" gorm:"type:int;default:null" extensions:"x-order=4"`
	CustomerName  *string  `json:"customerName" gorm:"type:varchar(50);default:null" extensions:"x-order=5"`
	CustomerPhone string   `json:"customerPhone" gorm:"type:varchar(30) not null" extensions:"x-order=6"`
	City          string   `json:"city" gorm:"type:varchar(255);default:null" extensions:"x-order=7"`
	District      string   `json:"district" gorm:"type:varchar(255);default:null" extensions:"x-order=8"`
	Street        string   `json:"street" gorm:"type:varchar(255);default:null" extensions:"x-order=9"`
	Home          string   `json:"home" gorm:"type:varchar(255);default:null" extensions:"x-order=10"`
	Apartment     string   `json:"apartment" gorm:"type:varchar(255);default:null" extensions:"x-order=11"`
	Comment       string   `json:"comment" gorm:"type:text;default:null" extensions:"x-order=12"`
	PaymentType   string   `json:"paymentType" gorm:"type:varchar(255);default:null" extensions:"x-order=13"`
	DeliveryPrice *float64 `json:"deliveryPrice" gorm:"type:decimal(10,2);default:null" extensions:"x-order=14"`
	FullSum       float64  `json:"fullSum" gorm:"type:decimal(10,2) not null" extensions:"x-order=15"`
	Status        int      `json:"status" gorm:"type:int;default:0" extensions:"x-order=16"`
	CreatedAt     string   `json:"createdAt" gorm:"type:datetime;default:null" extensions:"x-order=17"`
	UpdatedAt     string   `json:"updatedAt" gorm:"type:datetime;default:null" extensions:"x-order=18"`
	DeletedAt     string   `json:"deletedAt" gorm:"type:datetime;default:null" extensions:"x-order=19"`
}

type CreateOrderWithProductsInput struct {
	System        string                     `json:"system" extensions:"x-order=1"`
	AccountID     uint32                     `json:"accountID" extensions:"x-order=2"`
	ChatID        *int64                     `json:"chatID" extensions:"x-order=3"`
	RegionID      uint32                     `json:"regionID" extensions:"x-order=4"`
	CustomerName  *string                    `json:"customerName" extensions:"x-order=5"`
	CustomerPhone string                     `json:"customerPhone" extensions:"x-order=6"`
	City          string                     `json:"city" extensions:"x-order=7"`
	District      string                     `json:"district" extensions:"x-order=8"`
	Street        string                     `json:"street" extensions:"x-order=9"`
	Home          string                     `json:"home" extensions:"x-order=10"`
	Apartment     string                     `json:"apartment" extensions:"x-order=11"`
	Comment       string                     `json:"comment" extensions:"x-order=12"`
	PaymentType   string                     `json:"paymentType" extensions:"x-order=13"`
	DeliveryPrice *float64                   `json:"deliveryPrice" extensions:"x-order=14"`
	FullSum       float64                    `json:"fullSum" extensions:"x-order=15"`
	Status        int                        `json:"status" extensions:"x-order=16"`
	OrderProducts []CreateOrderProductsInput `json:"orderProducts" extensions:"x-order=17"`
}

type UpdateOrderInput struct {
	ID            uint64   `json:"-"`
	System        string   `json:"system" extensions:"x-order=1"`
	AccountID     uint32   `json:"accountID" extensions:"x-order=2"`
	ChatID        *int64   `json:"chatID" extensions:"x-order=3"`
	RegionID      uint32   `json:"regionID" extensions:"x-order=4"`
	CustomerName  *string  `json:"customerName" extensions:"x-order=5"`
	CustomerPhone string   `json:"customerPhone" extensions:"x-order=6"`
	City          string   `json:"city" extensions:"x-order=7"`
	District      string   `json:"district" extensions:"x-order=8"`
	Street        string   `json:"street" extensions:"x-order=9"`
	Home          string   `json:"home" extensions:"x-order=10"`
	Apartment     string   `json:"apartment" extensions:"x-order=11"`
	Comment       string   `json:"comment" extensions:"x-order=12"`
	PaymentType   string   `json:"paymentType" extensions:"x-order=13"`
	DeliveryPrice *float64 `json:"deliveryPrice" extensions:"x-order=14"`
	FullSum       float64  `json:"fullSum" extensions:"x-order=15"`
	Status        int      `json:"status" extensions:"x-order=16"`
}

type OrderWithOrderProducts struct {
	Orders        `extensions:"x-order=1"`
	OrderProducts []OrderProducts `json:"orderProducts" extensions:"x-order=2"`
}

type GetOrdersByFilterInput struct {
}
