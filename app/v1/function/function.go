package function

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	asyncapi "github.com/skywalkeretw/master-api/app/v1/asyncAPI"
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

	var data map[string]string
	err = json.Unmarshal(decodedInputParametersBytes, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	numEntries := len(data)
	i := 0
	var callStr string
	for key, value := range data {
		i++
		if functionData.Language == "golang" {
			callStr = fmt.Sprintf("%srb[\"%s\"].(%s)", callStr, key, value)
		} else {
			callStr = fmt.Sprintf("%srb[\"%s\"]", callStr, key)
		}
		if i < numEntries {
			callStr = fmt.Sprintf("%s, ", callStr)
		}

	}
	buildDeployData.FuncInput = callStr
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

	asyncAPISpecData := asyncapi.AsyncAPISpecData{
		Name:            functionData.Name,
		Description:     functionData.Description,
		InputParameters: inputParameters,
		ReturnValue:     returnValue,
	}
	// Create AsyncAPI file
	buildDeployData.AsyncAPIJSON, err = asyncapi.CreateAsyncAPISpec(asyncAPISpecData)
	if err != nil {
		fmt.Println("Error creating AsyncAPI Specification: ", err.Error())
	}

	// Continue with the next command or operation

	fmt.Println("created Build Deploy struct that will be used to create the function")
	fmt.Println(buildDeployData)
	rabbitmq.RPCclient(buildDeployData)

}
