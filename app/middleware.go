package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AttatchMiddleware(router *gin.Engine) {
	router.Use(cors.Default())
}
