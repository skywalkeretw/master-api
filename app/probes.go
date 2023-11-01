package app

import "github.com/gin-gonic/gin"

// @Summary Healthcheck
// @Description Perform a healthcheck
// @ID healthcheck
// @Produce json
// @Success 200 {object} string "OK"
// @Router /healthcheck [get]
func healthcheck(c *gin.Context) {
	c.JSON(200, "OK")
}

// @Summary Readiness
// @Description Check if the service is ready
// @ID readiness
// @Produce json
// @Success 200 {object} string "Ready"
// @Router /readiness [get]
func readiness(c *gin.Context) {
	c.JSON(200, "Ready")
}

// @Summary Liveness
// @Description Check if the service is live
// @ID liveness
// @Produce json
// @Success 200 {object} string "Live"
// @Router /liveness [get]
func liveness(c *gin.Context) {
	c.JSON(200, "Live")
}
