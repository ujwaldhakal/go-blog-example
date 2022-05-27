package auth_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	auth "github.com/ujwaldhakal/go-blog-example/auth/service"
	"net/http"
)

func AuthJwtHeaderToken(c *gin.Context) {

	authHeader := c.Request.Header.Get("Authorization")
	jwtToken := authHeader[len("Bearer "):]

	fmt.Println("tok", jwtToken)
	status, _ := auth.ValidateJwtToken(jwtToken)

	if !status {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	// Continue down the chain to handler etc
	c.Next()
}
