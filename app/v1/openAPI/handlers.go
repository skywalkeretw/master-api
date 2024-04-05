package openapi

import (
	"fmt"
	"net/http"
	"os"

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

	oasFilePath := utils.TransformTitle2FilenamePath("generate", "swaggerjson", oas.Info.Title)
	if err := utils.CreateJSONFile(oasFilePath, oas); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	serverZipPath, err := GenerateServerStub(oasFilePath, "go-server")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generete Server Stub" + err.Error()})
		return
	}
	ctx.File(serverZipPath)
	if err != nil {
		fmt.Println("failed to delete swagger json")
		// return "", err
	}
}

func GenerateClientHandler(ctx *gin.Context) {
	var oas OpenAPI
	if err := ctx.BindJSON(&oas); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oasFilePath := utils.TransformTitle2FilenamePath("generate", "swaggerjson", oas.Info.Title)
	utils.CreateJSONFile(oasFilePath, oas)

	clientZipPath, err := GenerateClient(oasFilePath, "name", "javascript")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generete Server Stub"})
		return
	}
	ctx.File(clientZipPath)
	err = os.Remove(clientZipPath)
	if err != nil {
		fmt.Println("failed to delete swagger json")
		// return "", err
	}
}
