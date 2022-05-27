package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=127.0.0.1 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	return db
}
