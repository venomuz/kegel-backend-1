package database

import (
	"fmt"
	"github.com/venomuz/kegel-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewClient(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.MysqlHost,
		cfg.MysqlUser,
		cfg.MysqlPassword,
		cfg.MysqlDatabase,
		cfg.MysqlPort,
	)
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	return db, err
}
