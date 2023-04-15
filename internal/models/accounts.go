package models

type Accounts struct {
	ID          uint32  `json:"id" gorm:"type:int unsigned not null;primaryKey"`
	RegionID    uint32  `json:"regionId"  gorm:"type:int;default:1"`
	ChatID      *uint32 `json:"chatId" gorm:"type:int;default:null"`
	System      string  `json:"system" gorm:"type:varchar(15);default:'WebSite'"`
	FirstName   string  `json:"firstName" gorm:"type:varchar(150);default:null"`
	LastName    string  `json:"lastName" gorm:"type:varchar(150);default:null"`
	Birthday    *string `json:"birthday"  gorm:"type:date;default:null"`
	PhoneNumber string  `json:"phoneNumber" gorm:"type:varchar(20) not null;index;unique"`
	Password    string  `json:"password" gorm:"type:varchar(150);default:null"`
	LastVisit   string  `json:"lastVisit" gorm:"type:datetime;default:null"`
	Language    string  `json:"language" gorm:"type:varchar(5);default:'ru'"`
	AuthKey     string  `json:"authKey" gorm:"type:varchar(255);default:null"`
	Blocked     *int8   `json:"blocked" gorm:"type:smallint;default:0;index"`
	CreatedAt   string  `json:"createdAt" gorm:"type:datetime;default:null"`
	UpdatedAt   string  `json:"updatedAt" gorm:"type:datetime;default:null"`
	DeletedAt   string  `json:"deletedAt" gorm:"type:datetime;default:null"`
}

type RegistrationAccountInput struct {
	PhoneNumber      string `json:"phoneNumber" binding:"required" example:"998901234323"`
	Password         string `json:"password" binding:"required"`
	VerificationCode string `json:"verificationCode" biding:"max=5,min=5,required"`
}

type UpdateAccountInput struct {
	ID        uint32  `json:"-"`
	RegionID  uint32  `json:"regionId"`
	ChatID    *uint32 `json:"chatId"`
	System    string  `json:"system"`
	FirstName string  `json:"firstName" binding:"min=2,required"`
	LastName  string  `json:"lastName"`
	Birthday  *string `json:"birthday" example:"2006-11-22" time_format:"2006-01-02"`
	Password  string  `json:"password"`
	Language  string  `json:"language"  binding:"max=2" example:"uz"`
	Blocked   *int8   `json:"blocked"`
}

type AccountSendVerificationInput struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" example:"998901234323"`
}

type LoginAccountInput struct {
	PhoneNumber string `form:"phoneNumber" binding:"required" example:"998901234323"`
	Password    string `form:"password"`
}
