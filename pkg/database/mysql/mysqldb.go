package mysqldb

import (
	"fmt"
	"github.com/venomuz/kegel-backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewClient(cfg config.Config) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.MysqlUser,
		cfg.MysqlPassword,
		cfg.MysqlHost,
		cfg.MysqlPort,
		cfg.MysqlDatabase,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.MysqlDatabase)

	return db, err
}
