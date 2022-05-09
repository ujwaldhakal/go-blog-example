package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	auth "github.com/ujwaldhakal/go-blog-example/auth/service"
	"net/http"
)

func register(c *gin.Context) {
	fmt.Println("register called")
	c.String(http.StatusOK,"registered")
}

func login(c *gin.Context) {
	fmt.Println("login called")
	c.String(http.StatusOK,"Login")
}

func forgotPassword(c *gin.Context) {
	fmt.Println("forgotpassword called")
	c.String(http.StatusOK,"pw reset")
}

func RegisterAuthRoutes(route *gin.Engine)  {
	v1 := route.Group("/v1")
	v1.GET("/register", register)
	v1.POST("/login", auth.Login)
	v1.POST("/forgot-password", forgotPassword)
}
