package function

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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
	deployments, err := kubernetes.GetKubernetesDeployments("default")
	if err != nil {
		return functions, fmt.Errorf("cannot get deployments from Kubernetes API %v", err.Error())
	}

	for _, deployment := range deployments {
		value, ok := deployment.Spec.Template.ObjectMeta.Labels["type"]
		if ok && value == "function" {

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
	}
	if len(functions) == 0 {
		return functions, errors.New("no functions found in kubernetes api")
	}

	return functions, nil
}

func GenerateAdapterCode(functionName, functionMode, language string) (string, error) {
	fmt.Println(functionName, functionMode, language)
	switch functionMode {
	case "httpsync":
		swaggerSpecPath, err := getSpec(functionName, "openapi")
		if err != nil {
			return "", err
		}
		return openapi.GenerateClient(swaggerSpecPath, functionName, language)
	case "httpasync":
		swaggerSpecPath, err := getSpec(functionName, "openapi")
		if err != nil {
			return "", err
		}
		return openapi.GenerateClient(swaggerSpecPath, functionName, language)
	case "messagingsync":
		asyncapiSpecPath, err := getSpec(functionName, "asyncapi")
		if err != nil {
			return "", err
		}
		return asyncapi.GenerateClient(asyncapiSpecPath, functionName, language)

	case "messagingasync":
		asyncapiSpecPath, err := getSpec(functionName, "asyncapi")
		if err != nil {
			return "", err
		}
		return asyncapi.GenerateClient(asyncapiSpecPath, functionName, language)

	}

	return "", fmt.Errorf("wrong function mode")

}

func getSpec(function, mode string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s.default:8080/%s", function, mode))
	if err != nil {
		return "", fmt.Errorf("failed to call function: %v", err.Error())
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read function body: %v", err.Error())
	}
	// Define a map to store the JSON data

	var jsonString string
	if err := json.Unmarshal(body, &jsonString); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON string: %v", err.Error())
	}

	var jsonData map[string]interface{}

	// Unmarshal JSON into the map
	if err := json.Unmarshal([]byte(jsonString), &jsonData); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %v", err)

	}
	fmt.Println("JSON Data: ", jsonData)

	specJson, err := json.Marshal(jsonData)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)

	}

	file := fmt.Sprintf("/specs/%s-openapi.json", function)
	err = utils.WriteFile(file, specJson)
	if err != nil {
		return "", fmt.Errorf("failed to write to file: %v", err.Error())
	}

	return file, nil
}
