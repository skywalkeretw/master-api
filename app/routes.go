package app

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/skywalkeretw/master-api/app/v1"
	_ "github.com/skywalkeretw/master-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// AttatchRoutes configures all route handlers
func AttatchRoutes(router *gin.Engine) {

	// Healthcheck routes
	// router.GET("/healthcheck", healthcheck)
	// router.GET("/readiness", readiness)
	// router.GET("/liveness", liveness)

	//

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Configured routes grouped in diffrent versions for future development
	v1.ConfigureRoutes(router.Group("/api/v1"))

}
