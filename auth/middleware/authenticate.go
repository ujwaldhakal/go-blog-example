package auth_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	auth "github.com/ujwaldhakal/go-blog-example/auth/service"
	"net/http"
)

func AuthJwtHeaderToken(c *gin.Context) {

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader != "" {
		jwtToken := authHeader[len("Bearer "):]

		fmt.Println("tok", jwtToken)
		user, err := auth.ValidateJwtToken(jwtToken)
		c.Set("user", user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}
	}
	// Continue down the chain to handler etc

	c.Next()
}
