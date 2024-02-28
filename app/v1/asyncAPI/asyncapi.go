package asyncapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/skywalkeretw/master-api/app/utils"
)

// AsyncAPI represents the AsyncAPI specification.
type AsyncAPI struct {
	Asyncapi           string               `json:"asyncapi,omitempty"`
	Id                 string               `json:"id,omitempty"` //"urn:example:com:smartylighting:streetlights:server"
	Info               AsyncAPIInfo         `json:"info,omitempty"`
	Servers            map[string]Server    `json:"servers,omitempty"`
	DefaultContentType string               `json:"defaultContentType,omitempty"`
	Channels           map[string]Channel   `json:"channels,omitempty"`
	Components         Components           `json:"components,omitempty"`
	Operations         map[string]Operation `json:"operations,omitempty"`
	// ExternalDocs       ExternalDocs           `json:"externalDocs,omitempty"`
	Tags       []Tag                  `json:"tags,omitempty"`
	Extensions map[string]interface{} `json:"x-extensions,omitempty"`
}

// AsyncAPIInfo represents the metadata information in the AsyncAPI specification.
type AsyncAPIInfo struct {
	Title          string  `json:"title,omitempty"`
	Version        string  `json:"version,omitempty"`
	Description    string  `json:"description,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
}

// Server represents a server in the AsyncAPI specification.
type Server struct {
	Host        string                    `json:"host,omitempty"`
	Protocol    string                    `json:"protocol,omitempty"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

// ServerVariable represents a variable for a server in the AsyncAPI specification.
type ServerVariable struct {
	Default     string   `json:"default"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}

// Channel represents a channel in the AsyncAPI specification.
type Channel struct {
	Address     string         `json:"address,omitempty"`
	Message     map[string]Ref `json:"messages,omitempty"`
	Title       string         `json:"title,omitempty"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
	Servers     []Ref          `json:"servers,omitempty"`
	Parameters  map[string]Ref `json:"parameters,omitempty"`
	Tags        []Tag          `json:"tags,omitempty"`
	// ExternalDocs ExternalDocs    `json:"externalDocs,omitempty"`
	Bindings ChannelBindings `json:"bindings,omitempty"`
}

type OperationBindings struct {
	AMQP AMQPBinding `json:"amqp,omitempty"`
}

type ChannelBindings struct {
	AMQP AMQPBinding `json:"amqp,omitempty"`
}

type AMQPBinding struct {
	Is             string `json:"is,omitempty"` //"queue" or "exchange"
	Queue          Queue  `json:"queue,omitempty"`
	BindingVersion string `json:"bindingVersion,omitempty"`
}

type Queue struct {
	Name       string `json:"name,omitempty"`
	Durable    bool   `json:"durable"`
	Exclusive  bool   `json:"exclusive"`
	AutoDelete bool   `json:"autoDelete"`
	VHost      string `json:"vhost,omitempty"`
}

// AsyncAPIComponents represents the components section in the AsyncAPI specification.
type Components struct {
	Schemas    map[string]SchemaObject `json:"schemas,omitempty"`
	Servers    map[string]Server       `json:"servers,omitempty"`
	Channels   map[string]Channel      `json:"channels,omitempty"`
	Operations map[string]Operation    `json:"operations,omitempty"`
	Messages   map[string]Message      `json:"messages,omitempty"`
	// Add other AsyncAPIComponents fields as needed
}

type SchemaObject struct {
	SchemaFormat string      `json:"schemaFormat,omitempty"`
	Schema       interface{} `json:"schema,omitempty"`
}

// Tag represents a tag in the AsyncAPI specification.
type Tag struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// ExternalDocs ExternalDocs `json:"externalDocs,omitempty"`
}

// Operation represents an operation in the AsyncAPI specification.
type Operation struct {
	Action      string             `json:"action,omitempty"` // "send" | "receive"
	Channel     Ref                `json:"channel,omitempty"`
	Title       string             `json:"title,omitempty"`
	Summary     string             `json:"summary,omitempty"`
	Description string             `json:"description,omitempty"`
	Bindings    *OperationBindings `json:"bindings,omitempty"`
	Message     map[string]Ref     `json:"messages,omitempty"`
	Reply       *Reply             `json:"reply,omitempty"`
}

type Reply struct {
	Address Address `json:"addressee,omitempty"`
	Channel Ref     `json:"channel,omitempty"`
}

type Address struct {
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"` // $message.header#/replyTo
}

type Message struct {
	Headers       map[string]interface{} `json:"headers,omitempty"`
	Payload       Payload                `json:"payload,omitempty"`
	CorrelationId *CorrelationId         `json:"correlationId,omitempty"`
	ContentType   string                 `json:"contentType,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Summary       string                 `json:"summary,omitempty"`
}

type Payload struct {
	Type       string              `json:"type,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
}

type Property struct {
	Type string `json:"type,omitempty"`
}

type CorrelationId struct {
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"` //  $message.header#/correlationId
}

// ExternalDocs represents external documentation in the AsyncAPI specification.
type ExternalDocs struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
	Summary     string `json:"summary,omitempty"`
}

// Contact represents contact information in the AsyncAPI specification.
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// License represents license information in the AsyncAPI specification.
type License struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type AsyncAPISpecData struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	InputParameters string `json:"inputparameters" binding:"required"`
	ReturnValue     string `json:"returnvalue" binding:"required"`
}

type Ref struct {
	Ref string `json:"$ref,omitempty"`
}

// getStringFromInterface checks if the interface contains a string of a valid OpenAPI type and returns it
func isValidAsyncAPIType(strValue string) (string, error) {

	if utils.IsJSONObject(strValue) {
		strValue = "object"
	}
	// Check if the string is a valid OpenAPI type
	openAPITypes := map[string]bool{
		"string":  true,
		"number":  true,
		"integer": true,
		"boolean": true,
		"array":   true,
		"object":  true,
	}

	_, valid := openAPITypes[strValue]
	if !valid {
		return "", fmt.Errorf("'%s' is not a valid OpenAPI type", strValue)
	}

	return strValue, nil
}

func generateProperties(dataString string) map[string]Property {
	properties := map[string]Property{}

	if utils.IsJSONObject(dataString) {
		dataMap, err := utils.JsonToMap(dataString)
		if err != nil {
			fmt.Println("error converting Json string to Map", err.Error())
			return properties // Return empty properties on error
		}
		for key, inputType := range dataMap {
			inputTypeStr := fmt.Sprintf("%v", inputType)
			strType, err := isValidAsyncAPIType(inputTypeStr)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			if strType != "object" {
				properties[key] = Property{Type: strType}
			} else {
				// Recursively handle object type
			}
		}
	}
	return properties
}

func CreateAsyncAPISpec(functionData AsyncAPISpecData) (string, error) {

	inputPayload := Payload{}
	if functionData.InputParameters != "" {
		if utils.IsJSONObject(functionData.InputParameters) {

			inputPayload = Payload{
				Type:       "object",
				Properties: generateProperties(functionData.InputParameters),
			}
		} else {
			return "", fmt.Errorf("unsopported format create a correct json file")
		}
	}

	returnPayload := Payload{}
	if utils.IsJSONObject(functionData.ReturnValue) {
		// If the return value is a JSON object
		returnPayload = Payload{
			Type:       "object",
			Properties: generateProperties(functionData.ReturnValue),
		}

	} else {
		returnPayload = Payload{
			Type: strings.ReplaceAll(strings.ReplaceAll(functionData.ReturnValue, "'", ""), `"`, ""),
		}
		// // If the return value is a string, number, or bool
		// responses = map[string]Response{
		// 	"200": {
		// 		Description: functionData.Description,
		// 		Content: map[string]MediaType{
		// 			"application/json": {Schema: Schema{Type: functionData.ReturnValue}}, // Example for a string response, modify as needed
		// 			// For number: Schema{Type: "number"}
		// 			// For bool: Schema{Type: "boolean"}
		// 		},
		// 	},
		// }
	}

	serverName := "cluster-local"
	serverRef := fmt.Sprintf("#/servers/%s", serverName)
	serverHost := fmt.Sprintf("%s:%d", utils.GetEnvSting("RABBITMQ_HOST", "localhost"), utils.GetEnvInt("RABBITMQ_PORT", 5672))

	inputChannelName := fmt.Sprintf("input-func-%s", functionData.Name)
	inputChannelRef := fmt.Sprintf("#/channels/%s", inputChannelName)
	inputMessageName := inputChannelName
	inputMessagRef := fmt.Sprintf("#/components/messages/%s", inputMessageName)

	returnChanelName := fmt.Sprintf("return-func-%s", functionData.Name)
	returnChanelRef := fmt.Sprintf("#/channels/%s", returnChanelName)
	returnMessageName := returnChanelName
	returnMessageRef := fmt.Sprintf("#/components/messages/%s", returnMessageName)

	receiveOperation := fmt.Sprintf("receive-%s", functionData.Name)
	sendOperation := fmt.Sprintf("send-%s", functionData.Name)

	asyncAPISpec := AsyncAPI{
		Asyncapi: "2.6.0",
		Info: AsyncAPIInfo{
			Title:       functionData.Name,
			Version:     "1.0.0",
			Description: functionData.Description,
			Contact: Contact{
				Name: "Developer",
			},
			License: License{
				Name: "MIT",
				URL:  "https://opensource.org/license/mit",
			},
		},
		DefaultContentType: "application/json",
		Servers: map[string]Server{
			serverName: {
				Host:     serverHost,
				Protocol: "amqp",
			},
		},
		Channels: map[string]Channel{
			inputChannelName: {
				Address:     "input",
				Title:       "Input",
				Description: "The channel for sending input parameters to a serverless function.",
				Message:     map[string]Ref{inputMessageName: {Ref: inputMessagRef}},
				// Parameters:  map[string]Ref{"": {Ref: ""}},
				Servers: []Ref{{Ref: serverRef}},
				Bindings: ChannelBindings{
					AMQP: AMQPBinding{
						Is: "queue",
						Queue: Queue{
							Name:       inputChannelName,
							Durable:    true,
							Exclusive:  false,
							AutoDelete: false,
						},
						BindingVersion: "0.3.0",
					},
				},
			},
			returnChanelName: {
				Address:     fmt.Sprintf("%s.%s", functionData.Name, "return"),
				Title:       "Return",
				Description: "The channel for sending the return data to the client from a serverless function.",
				Message:     map[string]Ref{"returnData": {Ref: returnMessageRef}},
				// Parameters:  map[string]Ref{"": {Ref: ""}},
				Servers: []Ref{{Ref: serverRef}},
				Bindings: ChannelBindings{
					AMQP: AMQPBinding{
						Is: "queue",
						Queue: Queue{
							Name:       returnChanelName,
							Durable:    true,
							Exclusive:  false,
							AutoDelete: false,
						},
						BindingVersion: "0.3.0",
					},
				},
			},
		},
		Operations: map[string]Operation{
			receiveOperation: {
				Action:  "receive",
				Channel: Ref{Ref: inputChannelRef},
				// Bindings: OperationBindings{},
				// Reply:    Reply{},
			},
			sendOperation: {
				Action:  "send",
				Channel: Ref{Ref: returnChanelRef},
				// Bindings: OperationBindings{},
				// Reply:    Reply{},
			},
		},
		Components: Components{
			Messages: map[string]Message{
				inputMessageName: {
					Name:    inputMessageName,
					Payload: inputPayload,
					// CorrelationId: CorrelationId{},
				},
				returnMessageName: {
					Name:    returnMessageName,
					Payload: returnPayload,
					// CorrelationId: CorrelationId{},
				},
			},
		},
	}

	asyncAPISpecBytes, err := json.Marshal(asyncAPISpec)
	if err != nil {
		return "", fmt.Errorf("error marshalling AsyncAPI Specification JSON: %v", err.Error())
	}
	return string(base64.StdEncoding.EncodeToString(asyncAPISpecBytes)), nil
}

func GenerateClient(asyncapiSpecPath, name, language string) (string, error) {
	// utils.Contains(clients, language)
	clientCodeTmpDirPath, err := utils.GenerateTempFolder()
	if err != nil {
		return "", err
	}
	var generator string
	switch language {
	case "go", "golang":
		generator = "@asyncapi/go-watermill-template"
	case "python":
		generator = "@asyncapi/python-paho-template"
	case "javascript":
		generator = "@asyncapi/nodejs-template"
	}
	// npm install -g @asyncapi/cli
	// asyncapi generate fromTemplate https://bit.ly/asyncapi @asyncapi/nodejs-template  -o example
	_, err = utils.RunShellCommand("asyncapi", "generate", "fromTemplate", asyncapiSpecPath, generator, "-o", clientCodeTmpDirPath)
	if err != nil {
		return "", err
	}

	zipPath := fmt.Sprintf("/generate/%s-%s.zip", name, language)
	err = utils.ZipFolder(clientCodeTmpDirPath, zipPath)
	if err != nil {
		return "", err
	}
	// err = utils.DeleteFolder(clientCodeTmpDirPath)
	// if err != nil {
	// 	fmt.Println("failed to delete folder")
	// 	// return "", err
	// }
	// err = os.Remove(swaggerSpecPath)
	// if err != nil {
	// 	fmt.Println("failed to delete swagger json")
	// 	// return "", err
	// }
	return zipPath, nil
}
