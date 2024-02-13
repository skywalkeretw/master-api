package function

import (
	"encoding/json"
	"fmt"

	asyncapi "github.com/skywalkeretw/master-api/app/v1/asyncAPI"
	openapi "github.com/skywalkeretw/master-api/app/v1/openAPI"
	"github.com/skywalkeretw/master-api/app/v1/rabbitmq"
)

func CreateFunction(functionData CreateFunctionHandlerData) {
	var err error
	buildDeployData := rabbitmq.FunctionBuildDeployData{
		Name:       functionData.Name,
		Language:   functionData.Language,
		SourceCode: functionData.SourceCode,
	}
	// Increment the WaitGroup counter for each goroutine

	openAPISpecData := openapi.OpenAPISpecData{
		Name:            functionData.Name,
		Description:     functionData.Description,
		InputParameters: functionData.InputParameters,
		ReturnValue:     functionData.ReturnValue,
	}
	// Create OpenAPI file
	openAPISpec, err := openapi.CreateOpenAPISpec(openAPISpecData)
	if err != nil {
		fmt.Println("Error creating OpenAPI Specification: ", err.Error())
	}

	asyncAPISpecData := asyncapi.AsyncAPISpecData{
		Name:            functionData.Name,
		Description:     functionData.Description,
		InputParameters: functionData.InputParameters,
		ReturnValue:     functionData.ReturnValue,
	}
	// Create AsyncAPI file
	asyncAPISpec, err := asyncapi.CreateAsyncAPISpec(asyncAPISpecData)
	if err != nil {
		fmt.Println("Error creating AsyncAPI Specification: ", err.Error())
	}

	openAPISpecBytes, err := json.Marshal(openAPISpec)
	if err != nil {
		fmt.Println("Error marshalling OpenAPI Specification JSON: ", err.Error())
	}
	buildDeployData.OpenAPIJSON = string(openAPISpecBytes)

	asyncAPISpecBytes, err := json.Marshal(asyncAPISpec)
	if err != nil {
		fmt.Println("Error marshalling AsyncAPI Specification JSON: ", err.Error())
	}
	buildDeployData.AsyncAPIJSON = string(asyncAPISpecBytes)
	// Continue with the next command or operation
	fmt.Println("Both functions have completed.")
	fmt.Println(buildDeployData)
	rabbitmq.RPCclient(buildDeployData)

}
