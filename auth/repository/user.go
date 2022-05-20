package repository

import (
	db "github.com/ujwaldhakal/go-blog-example/db"
	"github.com/ujwaldhakal/go-blog-example/user"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(username string, password string) bool {
	con := db.GetConnection()
	var users []user.User
	con.Where(&user.User{Email: username}).Find(&users)
	if len(users) == 0 {
		return false
	}
	existingPw := users[0].Password
	err := bcrypt.CompareHashAndPassword([]byte(existingPw), []byte(password))

	return err == nil
}

func IsUniqueEmail(email string) bool {
	con := db.GetConnection()
	var users []user.User
	con.Where(&user.User{Email: email}).Find(&users)

	return len(users) == 0
}

func Register(username string, password string) error {
	con := db.GetConnection()

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	result := con.Create(&user.User{Email: username, Password: hashedPassword})

	return result.Error
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
