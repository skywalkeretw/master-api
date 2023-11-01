package v1

import "github.com/gin-gonic/gin"

func ConfigureRoutes(v1 *gin.RouterGroup) {
	v1.Group("/G1").GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})

}
