package models

type OrderProducts struct {
	ID          uint64  `json:"id" gorm:"type:bigint not null auto_increment;primaryKey"`
	OrderID     uint64  `json:"orderId" gorm:"type:bigint not null"`
	ProductID   string  `json:"productId" gorm:"type:varchar(40)"`
	ProductName string  `json:"productName" gorm:"type:varchar(255)"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2) not null"`
	Amount      float64 `json:"amount" gorm:"type:decimal(10,2) not null"`
	CreatedAt   string  `json:"createdAt" gorm:"type:datetime;default:null"`
	UpdatedAt   string  `json:"updatedAt" gorm:"type:datetime;default:null"`
	DeletedAt   string  `json:"deletedAt" gorm:"type:datetime;default:null"`
}

type CreateOrderProductsInput struct {
	OrderID   uint64  `json:"-"`
	ProductID string  `json:"productId"  binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

type UpdateOrderProductsInput struct {
	ID        uint64  `json:"-"`
	OrderID   uint64  `json:"orderId" binding:"required"`
	ProductID string  `json:"productId"  binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}
