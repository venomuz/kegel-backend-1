package models

type Settings struct {
	ID        uint32 `json:"id" gorm:"type:int unsigned not null;primaryKey"`
	Title     string `json:"title" gorm:"not null;type:varchar(255) not null"`
	Key       string `json:"key" gorm:"not null;type:varchar(50) not null;unique"`
	Value     string `json:"value" gorm:"not null;type:varchar(255) not null"`
	CreatedAt string `json:"created_at" gorm:"type:datetime;default:null" example:"2022-01-15T11:27:04+05:00"`
	UpdatedAt string `json:"updated_at" gorm:"type:datetime;default:null" example:"2022-01-15T11:27:04+05:00"`
	DeletedAt string `json:"deleted_at" gorm:"type:datetime;default:null" example:"2022-01-15T11:27:04+05:00"`
}

type CreateSettingInput struct {
	Title string `json:"title" binding:"required"`
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type UpdateSettingInput struct {
	ID    uint32 `json:"-"`
	Title string `json:"title" binding:"required"`
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}
