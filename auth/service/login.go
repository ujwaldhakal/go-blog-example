package auth

import (
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/ujwaldhakal/go-blog-example/auth/repository"
	"net/http"
)

type LoginRequest struct {
	Email string `json:"email"  binding:"required"`

	Password string `json:"password"  binding:"required"`
}

type Response struct {
	status  string `json:"status" format:"string"`
	message string `json:"message" format:"string"`
	code    int    `json:"code" format:"int"`
	data    interface{}
}

func respond(response *Response) gin.H {

	g := gin.H{"code": response.code}
	g["message"] = response.message
	g["status"] = response.status
	g["data"] = response.data

	return g
}

// @BasePath /v1
// Login
// @Summary Authenticates when provided with login details
// @Schemes
// @Description Logins when you provide details
// @Accept json
// @Produce json
// @Param Body body LoginRequest true "Parameters should not be empty"
// @Success 200 {object} Response
// @Router /login [post]
func Login(c *gin.Context) {
	var requestBody LoginRequest

	fmt.Println("here man", c.Request.Body)
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userName := requestBody.Email
	password := requestBody.Password

	isAuthenticated := repository.Authenticate(userName, password)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Sorry username is incorrect"})
		return
	}

	token, _ := GenerateJwtToken(userName)
	dataMap := make(map[string]string)
	dataMap["token"] = token
	c.JSON(200, respond(&Response{status: "success", code: http.StatusOK, message: "success", data: dataMap}))
}
