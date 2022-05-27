package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ujwaldhakal/go-blog-example/auth/repository"
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

func ValidateJwtToken(tokenString string) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		user := repository.FindUserByEmail(email)
		fmt.Println(user)
		if user {
			return true, nil
		}
	}
	return false, err
}
