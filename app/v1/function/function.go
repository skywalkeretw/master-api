package function

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/skywalkeretw/master-api/app/utils"
	asyncapi "github.com/skywalkeretw/master-api/app/v1/asyncAPI"
	"github.com/skywalkeretw/master-api/app/v1/kubernetes"
	openapi "github.com/skywalkeretw/master-api/app/v1/openAPI"
	"github.com/skywalkeretw/master-api/app/v1/rabbitmq"
)

func CreateFunction(functionData CreateFunctionHandlerData) {
	fmt.Println("function data", functionData)
	fmt.Println("Creating Function")
	var err error

	buildDeployData := rabbitmq.FunctionBuildDeployData{
		Name:          functionData.Name,
		Language:      functionData.Language,
		Description:   functionData.Description,
		SourceCode:    functionData.SourceCode,
		FunctionModes: functionData.FunctionModes,
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

type FunctionsData struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Tags        map[string]string        `json:"tags"`
	Modes       kubernetes.FunctionModes `json:"modes"`
}

func GetFunctions() ([]FunctionsData, error) {
	var functions []FunctionsData
	deployments, err := kubernetes.GetKubernetesDeployments("functions")
	if err != nil {
		return functions, fmt.Errorf("cannot get deployments from Kubernetes API %v", err.Error())
	}

	for _, deployment := range deployments {
		newFunctionData := FunctionsData{
			Name: deployment.Name,
		}
		for _, container := range deployment.Spec.Template.Spec.Containers {
			for _, envVar := range container.Env {
				switch envVar.Name {
				case "DESCRIPTION":
					newFunctionData.Description = envVar.Value
				case "TAGS":
					tagsList := strings.Split(envVar.Value, ":")
					tags := make(map[string]string)

					// Iterate over pairs and split each pair by "=" to get key-value
					for _, pair := range tagsList {
						kv := strings.Split(pair, "=")
						if len(kv) == 2 {
							tags[kv[0]] = kv[1]
						}
					}
					newFunctionData.Tags = tags
				case "HTTPSYNC":
					newFunctionData.Modes.HTTPSync = utils.StringToBool(envVar.Value)
				case "HTTPASYNC":
					newFunctionData.Modes.HTTPAsync = utils.StringToBool(envVar.Value)
				case "MESSAGINGSYNC":
					newFunctionData.Modes.MessagingSync = utils.StringToBool(envVar.Value)
				case "MESSAGINGASYNC":
					newFunctionData.Modes.MessagingAsync = utils.StringToBool(envVar.Value)
				}
			}
		}
		functions = append(functions, newFunctionData)
	}

	return functions, nil
}
