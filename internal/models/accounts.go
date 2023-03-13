package models

type Accounts struct {
	ID          uint32  `json:"id" gorm:"type:int unsigned not null;primaryKey" extensions:"x-order=1"`
	RegionID    uint32  `json:"regionId"  gorm:"type:int;default:1" extensions:"x-order=2"`
	ChatID      *uint32 `json:"chatId" gorm:"type:int;default:null" extensions:"x-order=3"`
	System      string  `json:"system" gorm:"type:varchar(15);default:'WebSite'" extensions:"x-order=3"`
	FirstName   string  `json:"firstName" gorm:"type:varchar(150);default:null" extensions:"x-order=4"`
	LastName    string  `json:"lastName" gorm:"type:varchar(150);default:null" extensions:"x-order=5"`
	Birthday    *string `json:"birthday"  gorm:"type:date;default:null" extensions:"x-order=6"`
	PhoneNumber string  `json:"phoneNumber" gorm:"type:varchar(20) not null;index;unique" extensions:"x-order=7"`
	Password    string  `json:"password" gorm:"type:varchar(150);default:null" extensions:"x-order=8"`
	LastVisit   string  `json:"lastVisit" gorm:"type:datetime;default:null" extensions:"x-order=9"`
	Language    string  `json:"language" gorm:"type:varchar(5);default:'ru'" extensions:"x-order=10"`
	AuthKey     string  `json:"authKey" gorm:"type:varchar(255);default:null" extensions:"x-order=11"`
	Blocked     *int8   `json:"blocked" gorm:"type:smallint;default:0;index" extensions:"x-order=12"`
	CreatedAt   string  `json:"createdAt" gorm:"type:datetime;default:null" extensions:"x-order=13"`
	UpdatedAt   string  `json:"updatedAt" gorm:"type:datetime;default:null" extensions:"x-order=14"`
	DeletedAt   string  `json:"deletedAt" gorm:"type:datetime;default:null" extensions:"x-order=15"`
}

type RegistrationAccountInput struct {
	RegionID         uint32  `json:"regionId" extensions:"x-order=1"`
	ChatID           *uint32 `json:"chatId" extensions:"x-order=2"`
	System           string  `json:"system" extensions:"x-order=3"`
	FirstName        string  `json:"firstName" binding:"min=2,required" extensions:"x-order=3"`
	LastName         string  `json:"lastName" binding:"min=2,required" extensions:"x-order=4"`
	Birthday         *string `json:"birthday" example:"2006-11-22" time_format:"2006-01-02" extensions:"x-order=5"`
	PhoneNumber      string  `json:"phoneNumber" binding:"required" example:"998901234323" extensions:"x-order=6"`
	Password         string  `json:"password" binding:"required" extensions:"x-order=7"`
	Language         string  `json:"language" binding:"max=2" example:"uz" extensions:"x-order=8"`
	VerificationCode string  `json:"verificationCode" biding:"max=5,min=5,required" extensions:"x-order=9"`
}

type UpdateAccountInput struct {
	ID        uint32  `json:"-" extensions:"x-order=1"`
	RegionID  uint32  `json:"regionId" extensions:"x-order=2"`
	ChatID    *uint32 `json:"chatId" extensions:"x-order=3"`
	System    string  `json:"system" extensions:"x-order=3"`
	FirstName string  `json:"firstName" binding:"min=2,required" extensions:"x-order=4"`
	LastName  string  `json:"lastName" extensions:"x-order=5"`
	Birthday  *string `json:"birthday" example:"2006-11-22" time_format:"2006-01-02" extensions:"x-order=6"`
	Password  string  `json:"password" extensions:"x-order=7"`
	Language  string  `json:"language"  binding:"max=2" example:"uz" extensions:"x-order=8"`
	Blocked   *int8   `json:"blocked" extensions:"x-order=9"`
}

type AccountSendVerificationInput struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" example:"998901234323"`
}

type LoginAccountInput struct {
	PhoneNumber string `form:"phoneNumber" binding:"required" example:"998901234323" extensions:"x-order=1"`
	Password    string `form:"password" extensions:"x-order=2"`
}
