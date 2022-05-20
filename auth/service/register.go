package auth

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/ujwaldhakal/go-blog-example/auth/repository"
	"github.com/ujwaldhakal/go-blog-example/common"
)

type RegisterRequest struct {
	Email                string `json:"email"  binding:"required"`
	Password             string `json:"password"  binding:"required"`
	ConfirmationPassword string `json:"confirmation_password"  binding:"required"`
}

var confirmPasswordsDoNotMatch = "Passwords & Confirm Passwords donot match"
var emailAlreadyExists = "The email you are trying to register already exists"

// @BasePath /v1
// Login
// @Summary Registers a user when provided with details
// @Schemes
// @Description Registers when you provide details
// @Accept json
// @Produce json
// @Param Body body RegisterRequest true "Parameters should not be empty"
// @Success 200 {object} Response
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Router /register [post]
func Register(c *gin.Context) {

	var registerRequest RegisterRequest

	if err := c.BindJSON(&registerRequest); err != nil {
		common.RespondBadRequest(c, err.Error())
		return
	}

	email := registerRequest.Email
	password := registerRequest.Password
	confirmPassword := registerRequest.ConfirmationPassword

	if password != confirmPassword {
		common.RespondBadRequest(c, confirmPasswordsDoNotMatch)
		return
	}

	if !repository.IsUniqueEmail(email) {
		common.RespondBadRequest(c, emailAlreadyExists)
		return
	}

	repository.Register(email, password)
	common.RespondCreated(c, common.Response{Message: "Created successfully"})

}
