package function

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/skywalkeretw/master-api/app/utils"
	"github.com/skywalkeretw/master-api/app/v1/kubernetes"
)

// @Summary Create Function Data
// @Description Represents data for creating a function
type CreateFunctionHandlerData struct {
	Name            string                   `json:"name" binding:"required"`
	Description     string                   `json:"description" binding:"required"`
	Language        string                   `json:"language" binding:"required"`
	SourceCode      string                   `json:"sourcecode" binding:"required"`
	InputParameters string                   `json:"inputparameters" binding:"required"`
	ReturnValue     string                   `json:"returnvalue" binding:"required"`
	FunctionModes   kubernetes.FunctionModes `json:"functionmodes" binding:"required"`
}

// @Summary Create a new Function
// @Description Create and deploy a new function
// @Tags Function
// @Accept json
// @Produce json
// @Success 200 {array} v1.Pod
// @Router /api/v1/function [post]
func CreateFunctionHandler(ctx *gin.Context) {
	var data CreateFunctionHandlerData
	// Bind the JSON data from the request body to the updateOptions struct
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	if !utils.ValidateAllowedLanguages(data.Language) {
		ctx.AbortWithError(http.StatusForbidden, fmt.Errorf("unsupported language: %s", data.Language))
	}

	go CreateFunction(data)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Your Function is beeing Processed"})
}

// @Summary Get Function List
// @Description Returns a list of all functions in Kubernetes Cluster
// @Tags Function
// @Produce json
// @Success 200 {object} []FunctionsData
// @Router /api/v1/function [get]
func GetFunctionsHandler(ctx *gin.Context) {
	functions, err := GetFunctions()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to retrive functions"))
	}
	ctx.JSON(http.StatusOK, functions)
}

type GenerateAdapterCodeData struct {
	Function string `json:"function" binding:"required"`
	Language string `json:"language" binding:"required"`
	Mode     string `json:"mode" binding:"required"`
}

func GenerateAdapterCodeHandler(ctx *gin.Context) {
	var data GenerateAdapterCodeData
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	clientDataZipPath, err := GenerateAdapterCode(data.Function, data.Mode, data.Language)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	clientDataZip, err := os.ReadFile(clientDataZipPath)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Failed to read zip file"))
		return
	}
	ctx.Data(http.StatusOK, "application/zip", clientDataZip)
}
