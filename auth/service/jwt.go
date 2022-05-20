package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = "secret"

func GenerateJwtToken(email string) (string, error) {

	var mySigningKey = []byte(jwtSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong while generating token: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}
