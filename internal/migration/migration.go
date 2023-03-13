package migration

import (
	"github.com/venomuz/kegel-backend/internal/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {

	err := db.AutoMigrate(
		&models.Accounts{},
		&models.Groups{},
		&models.OrderProducts{},
		&models.Orders{},
		&models.ProductImages{},
		&models.ProductRates{},
		&models.Products{},
		&models.Settings{},
		&models.Users{},
	)
	return err
}
