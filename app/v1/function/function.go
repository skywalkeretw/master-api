package function

import (
	"fmt"
	"sync"

	"github.com/skywalkeretw/master-api/app/utils"
	openapi "github.com/skywalkeretw/master-api/app/v1/openAPI"
)

func CreateFunction(functionData CreateFunctionHandlerData) {
	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	tempDirPath, err := utils.GenerateTempFolder()
	if err != nil {
		// handle error
	}
	fmt.Println(tempDirPath)
	// Increment the WaitGroup counter for each goroutine
	wg.Add(2)

	OpenAPISpecData := openapi.OpenAPISpecData{
		Name:            functionData.Name,
		Description:     functionData.Description,
		InputParameters: functionData.InputParameters,
		ReturnValue:     functionData.ReturnValue,
	}
	// Create OpenAPI file
	go func() {
		defer wg.Done()
		openapi.CreateOpenAPISpec(OpenAPISpecData)

		//firstFunction()
	}()

	// Create AsyncAPI file
	go func() {
		defer wg.Done()
		//secondFunction()
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Continue with the next command or operation
	fmt.Println("Both functions have completed.")
}
