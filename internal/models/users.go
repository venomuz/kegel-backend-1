package models

type Users struct {
	ID        uint32 `json:"id"  gorm:"type:int unsigned not null;primaryKey" extensions:"x-order=1"`
	Username  string `json:"username"  gorm:"type:varchar(150);default:null" extensions:"x-order=2"`
	Password  string `json:"-"  gorm:"type:varchar(150);default:null" extensions:"x-order=3"`
	Email     string `json:"email"  gorm:"type:varchar(150);default:null" extensions:"x-order=4"`
	LastVisit string `json:"lastVisit"  gorm:"type:datetime;default:null" extensions:"x-order=5"`
	CreatedAt string `json:"createdAt"  gorm:"type:datetime;default:null" extensions:"x-order=6"`
	UpdatedAt string `json:"updatedAt"  gorm:"type:datetime;default:null" extensions:"x-order=7"`
	DeletedAt string `json:"deletedAt"  gorm:"type:datetime;default:null" extensions:"x-order=8"`
}

type LoginUserInput struct {
	Username string `form:"username" binding:"required" extensions:"x-order=1"`
	Password string `form:"password" binding:"required" extensions:"x-order=2"`
}
