package app

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/skywalkeretw/master-api/app/v1"
)

// AttatchRoutes configures all route handlers
func AttatchRoutes(router *gin.Engine) {

	// Healthcheck routes
	router.GET("/healthcheck")
	router.GET("/readiness")
	router.GET("/liveness")

	//
	v1.ConfigureRoutes(router.Group("/api/v1"))

}
