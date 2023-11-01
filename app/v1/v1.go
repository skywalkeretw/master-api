package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/skywalkeretw/master-api/app/v1/kubernets"
)

func ConfigureRoutes(v1 *gin.RouterGroup) {
	v1.Group("/G1").GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})
	v1.GET("/pods", kubernets.PodHandler)

	v1.GET("/app", kubernets.GetAllAppsHandler)
	v1.GET("/app/:name", kubernets.GetAppHandler)
	v1.DELETE("/app/:name", kubernets.DeleteAppHandler)
	v1.PATCH("/app", kubernets.UpdateAppHandler)

	// v1.POST("/generateadapter")
	// v1.POST("/deployapp")

}
