package models

type Settings struct {
	ID        uint32 `json:"id" gorm:"type:int unsigned not null;primaryKey" extensions:"x-order=1"`
	Title     string `json:"title" gorm:"not null;type:varchar(255) not null" extensions:"x-order=2"`
	Key       string `json:"key" gorm:"not null;type:varchar(50) not null;unique" extensions:"x-order=3"`
	Value     string `json:"value" gorm:"not null;type:varchar(255) not null" extensions:"x-order=4"`
	CreatedAt string `json:"created_at" gorm:"type:datetime;default:null" example:"2022-01-15T11:27:04+05:00" extensions:"x-order=5"`
	UpdatedAt string `json:"updated_at" gorm:"type:datetime;default:null" example:"2022-01-15T11:27:04+05:00" extensions:"x-order=6"`
	DeletedAt string `json:"deleted_at" gorm:"type:datetime;default:null" example:"2022-01-15T11:27:04+05:00" extensions:"x-order=7"`
}

type CreateSettingInput struct {
	Title     string `json:"title" extensions:"x-order=1"`
	Key       string `json:"key" extensions:"x-order=2"`
	Value     string `json:"value" extensions:"x-order=3"`
	CreatedAt string `json:"created_at" extensions:"x-order=4"`
	UpdatedAt string `json:"updated_at" extensions:"x-order=5"`
	DeletedAt string `json:"deleted_at" extensions:"x-order=6"`
}

type UpdateSettingInput struct {
	ID        uint32 `json:"id" extensions:"x-order=1"`
	Title     string `json:"title" extensions:"x-order=2"`
	Key       string `json:"key" extensions:"x-order=3"`
	Value     string `json:"value" extensions:"x-order=4"`
	CreatedAt string `json:"created_at" extensions:"x-order=5"`
	UpdatedAt string `json:"updated_at" extensions:"x-order=6"`
	DeletedAt string `json:"deleted_at" extensions:"x-order=7"`
}
