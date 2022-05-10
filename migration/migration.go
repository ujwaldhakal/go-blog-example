package main

import (
	"github.com/ujwaldhakal/go-blog-example/db"
	"github.com/ujwaldhakal/go-blog-example/user"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db := db.GetConnection()

	db.AutoMigrate(&user.User{})
}