package models

type Users struct {
	ID        uint32 `json:"id"  gorm:"type:int unsigned not null;primaryKey"`
	Username  string `json:"username"  gorm:"type:varchar(150);default:null"`
	Password  string `json:"-"  gorm:"type:varchar(150);default:null"`
	Email     string `json:"email"  gorm:"type:varchar(150);default:null"`
	LastVisit string `json:"lastVisit"  gorm:"type:datetime;default:null"`
	CreatedAt string `json:"createdAt"  gorm:"type:datetime;default:null"`
	UpdatedAt string `json:"updatedAt"  gorm:"type:datetime;default:null"`
	DeletedAt string `json:"deletedAt"  gorm:"type:datetime;default:null"`
}

type LoginUserInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
