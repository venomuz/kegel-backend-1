package models

type OrderProducts struct {
	ID          uint64  `json:"id" gorm:"type:bigint not null auto_increment;primaryKey" extensions:"x-order=1"`
	OrderID     uint64  `json:"orderId" gorm:"type:bigint not null" extensions:"x-order=2"`
	ProductID   string  `json:"productId" gorm:"type:varchar(40)" extensions:"x-order=3"`
	ProductName string  `json:"productName" gorm:"type:varchar(255)" extensions:"x-order=4"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2) not null" extensions:"x-order=5"`
	Amount      float64 `json:"amount" gorm:"type:decimal(10,2) not null" extensions:"x-order=6"`
	CreatedAt   string  `json:"createdAt" gorm:"type:datetime;default:null" extensions:"x-order=7"`
	UpdatedAt   string  `json:"updatedAt" gorm:"type:datetime;default:null" extensions:"x-order=8"`
	DeletedAt   string  `json:"deletedAt" gorm:"type:datetime;default:null" extensions:"x-order=9"`
}

type CreateOrderProductsInput struct {
	OrderID   uint64  `json:"-"`
	ProductID string  `json:"productId"  binding:"required" extensions:"x-order=1"`
	Amount    float64 `json:"amount" binding:"required" extensions:"x-order=2"`
}

type UpdateOrderProductsInput struct {
	ID        uint64  `json:"-"`
	OrderID   uint64  `json:"orderId" binding:"required" extensions:"x-order=1"`
	ProductID string  `json:"productId"  binding:"required" extensions:"x-order=2"`
	Amount    float64 `json:"amount" binding:"required" extensions:"x-order=3"`
}
