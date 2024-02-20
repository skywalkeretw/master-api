package function

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Function Data
// @Description Represents data for creating a function
type CreateFunctionHandlerData struct {
	Name            string        `json:"name" binding:"required"`
	Description     string        `json:"description" binding:"required"`
	Language        string        `json:"language" binding:"required"`
	SourceCode      string        `json:"sourcecode" binding:"required"`
	InputParameters string        `json:"inputparameters" binding:"required"`
	ReturnValue     string        `json:"returnvalue" binding:"required"`
	FunctionModes   FunctionModes `json:"functionmodes" binding:"required"`
}

type FunctionModes struct {
	HTTPSync       bool `json:"httpsync"`
	HTTPAsync      bool `json:"httpasync"`
	MessagingSync  bool `json:"messagingsync"`
	MessagingAsync bool `json:"messagingasync"`
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
	if !validateAllowedLanguages(data.Language) {
		ctx.AbortWithError(http.StatusForbidden, fmt.Errorf("unsupported language: %s", data.Language))
	}

	go CreateFunction(data)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Your Function is beeing Processed"})
}

// Custom validation function for allowed languages
func validateAllowedLanguages(language string) bool {
	allowedLanguages := []string{"golang", "python", "javascript"}
	for _, allowed := range allowedLanguages {
		if language == allowed {
			return true
		}
	}
	return false
}
