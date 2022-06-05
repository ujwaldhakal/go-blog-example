package post

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	auth_middleware "github.com/ujwaldhakal/go-blog-example/auth/middleware"
	"github.com/ujwaldhakal/go-blog-example/post/service"
)

func RegisterPostAuth(route *gin.Engine) {
	v1 := route.Group("/v1/posts")
	v1.Use(auth_middleware.AuthJwtHeaderToken)
	{
		v1.POST("/", post_service.Create)
	}
}
