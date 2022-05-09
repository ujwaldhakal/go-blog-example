package main

import (
	"github.com/ujwaldhakal/go-blog-example/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=pgsql user=postgres password=postgres dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	db.AutoMigrate(&user.User{})
}