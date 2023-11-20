package openapi

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/skywalkeretw/master-api/app/utils"
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

func GenerateServerStubHandler(ctx *gin.Context) {
	var oas OpenAPI
	if err := ctx.BindJSON(&oas); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oasFilePath := filepath.Join("dir", "subdir", utils.TransformTitle2Filename(oas.Info.Title))
	utils.CreateJSONFile(oasFilePath, oas)

	serverZipPath, err := GenerateServerStub(oasFilePath, ":todo add language")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generete Server Stub"})
		return
	}
	ctx.File(serverZipPath)
}

func GenerateClientHandler(ctx *gin.Context) {
	GenerateClient()
	ctx.File()
}
