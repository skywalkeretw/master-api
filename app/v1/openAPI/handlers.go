package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Run the Swagger Codegen Help
// @Description Run the swagger codegen command to check it is working
// @Tags Deployments
// @Accept json
// @Produce json
// @Success 200 string
// @Router /api/v1/openapi/codegen/help [get]
func SwaggerCodegenHelpHandler(ctx *gin.Context) {
	GetSwaggerCodegenHelp()
	ctx.JSON(http.StatusOK, gin.H{"ok": "cool"})
}
