package app

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()
	AttatchMiddleware(r)
	AttatchRoutes(r)
	r.Run(":8080")
}
