package repository

import (
	db "github.com/ujwaldhakal/go-blog-example/db"
	"github.com/ujwaldhakal/go-blog-example/user"
)

func Authenticate(username string, password string) bool {
	con := db.GetConnection()
	var users []user.User
	con.Where(&user.User{Email: username, Password: password}).Find(&users)

	return len(users) > 0
}


func Register(username string, password string) bool {
	con := db.GetConnection()
	var users []user.User
	con.Where(&user.User{Email: username, Password: password}).Find(&users)

	return len(users) > 0
}