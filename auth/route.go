package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	auth "github.com/ujwaldhakal/go-blog-example/auth/service"
	"net/http"
)

func forgotPassword(c *gin.Context) {
	fmt.Println("forgotpassword called")
	c.String(http.StatusOK, "pw reset")
}

func RegisterAuthRoutes(route *gin.Engine) {
	v1 := route.Group("/v1")
	v1.POST("/register", auth.Register)
	v1.POST("/login", auth.Login)
	v1.POST("/forgot-password", forgotPassword)
}
