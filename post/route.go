package post

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/ujwaldhakal/go-blog-example/post/service"
)


func RegisterPostAuth(route *gin.Engine) {
	v1 := route.Group("/v1/posts")
	v1.POST("/", post_service.Create)
}
