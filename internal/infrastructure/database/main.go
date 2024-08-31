package database

import (
	"loan_apps/config/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabaseConnection create a new database connection
func NewDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(mysql.Open(setting.DatabaseDNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
