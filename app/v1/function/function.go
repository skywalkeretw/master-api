package function

import (
	"encoding/base64"
	"fmt"

	openapi "github.com/skywalkeretw/master-api/app/v1/openAPI"
	"github.com/skywalkeretw/master-api/app/v1/rabbitmq"
)

func CreateFunction(functionData CreateFunctionHandlerData) {
	fmt.Println("function data", functionData)
	fmt.Println("Creating Function")
	var err error

	buildDeployData := rabbitmq.FunctionBuildDeployData{
		Name:       functionData.Name,
		Language:   functionData.Language,
		SourceCode: functionData.SourceCode,
	}

	decodedInputParametersBytes, err := base64.StdEncoding.DecodeString(functionData.InputParameters)
	if err != nil {
		fmt.Println("Error decoding  input parameters: ", err)
	}
	inputParameters := string(decodedInputParametersBytes)

	decodedReturnValueBytes, err := base64.StdEncoding.DecodeString(functionData.ReturnValue)
	if err != nil {
		fmt.Println("Error decoding  return value: ", err)
	}
	returnValue := string(decodedReturnValueBytes)

	openAPISpecData := openapi.OpenAPISpecData{
		Name:            functionData.Name,
		Description:     functionData.Description,
		InputParameters: inputParameters,
		ReturnValue:     returnValue,
	}
	// Create OpenAPI file
	buildDeployData.OpenAPIJSON, err = openapi.CreateOpenAPISpec(openAPISpecData)
	if err != nil {
		fmt.Println("Error creating OpenAPI Specification: ", err.Error())
	}

	// asyncAPISpecData := asyncapi.AsyncAPISpecData{
	// 	Name:            functionData.Name,
	// 	Description:     functionData.Description,
	// 	InputParameters: inputParameters,
	// 	ReturnValue:     returnValue,
	// }
	// // Create AsyncAPI file
	// buildDeployData.AsyncAPIJSON, err = asyncapi.CreateAsyncAPISpec(asyncAPISpecData)
	// if err != nil {
	// 	fmt.Println("Error creating AsyncAPI Specification: ", err.Error())
	// }

	// Continue with the next command or operation

	fmt.Println("created Build Deploy struct that will be used to create the function")
	fmt.Println(buildDeployData)
	rabbitmq.RPCclient(buildDeployData)

}
