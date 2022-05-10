package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ujwaldhakal/go-blog-example/auth/repository"
	"net/http"
)
type LoginRequest struct {
 	Username string
	Password string
}

type Response struct {
	status string
	data any.Any
}

func Login(c *gin.Context)  {
	fmt.Println("got it")
	var requestBody LoginRequest

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	 userName := requestBody.Username
	 password := requestBody.Password


	 fmt.Println("is logged in",repository.Authenticate(userName,password))

	c.JSON(200,gin.H{"code": http.StatusOK,"message": "Successfully authenticated"})
}
