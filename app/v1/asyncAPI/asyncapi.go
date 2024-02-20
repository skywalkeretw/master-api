package asyncapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/skywalkeretw/master-api/app/utils"
)

// AsyncAPI represents the AsyncAPI specification.
type AsyncAPI struct {
	Asyncapi           string                 `json:"asyncapi"`
	Id                 string                 `json:"id,omitempty"` //"urn:example:com:smartylighting:streetlights:server"
	Info               AsyncAPIInfo           `json:"info,omitempty"`
	Servers            map[string]Server      `json:"servers,omitempty"`
	DefaultContentType string                 `json:"defaultContentType,omitempty"`
	Channels           map[string]Channel     `json:"channels"`
	Components         Components             `json:"components,omitempty"`
	Operation          map[string]Operation   `json:"operation,omitempty"`
	ExternalDocs       ExternalDocs           `json:"externalDocs,omitempty"`
	Tags               []Tag                  `json:"tags,omitempty"`
	Extensions         map[string]interface{} `json:"x-extensions,omitempty"`
}

// AsyncAPIInfo represents the metadata information in the AsyncAPI specification.
type AsyncAPIInfo struct {
	Title          string  `json:"title"`
	Version        string  `json:"version"`
	Description    string  `json:"description,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
}

// Server represents a server in the AsyncAPI specification.
type Server struct {
	Host        string                    `json:"host"`
	Protocol    string                    `json:"protocol"`
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
	Address      string          `json:"address,omitempty"`
	Message      map[string]Ref  `json:"messages,omitempty"`
	Title        string          `json:"title,omitempty"`
	Summary      string          `json:"summary,omitempty"`
	Description  string          `json:"description,omitempty"`
	Servers      []Ref           `json:"servers,omitempty"`
	Parameters   map[string]Ref  `json:"parameters,omitempty"`
	Tags         []Tag           `json:"tags,omitempty"`
	ExternalDocs ExternalDocs    `json:"externalDocs,omitempty"`
	Bindings     ChannelBindings `json:"bindings,omitempty"`
}

type OperationBindings struct {
	AMQP AMQPBinding `json:"amqp,omitempty"`
}

type ChannelBindings struct {
	AMQP AMQPBinding `json:"amqp,omitempty"`
}

type AMQPBinding struct {
	Is    string `json:"is,omitempty"`
	Queue Queue  `json:"queue,omitempty"`
}

type Queue struct {
	Name       string `json:"name,omitempty"`
	Durable    bool   `json:"durable,omitempty"`
	Exclusive  bool   `json:"exclusive,omitempty"`
	AutoDelete bool   `json:"auto_delete,omitempty"`
	VHost      string `json:"vHost,omitempty"`
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
	Name         string       `json:"name,omitempty"`
	Description  string       `json:"description,omitempty"`
	ExternalDocs ExternalDocs `json:"externalDocs,omitempty"`
}

// Operation represents an operation in the AsyncAPI specification.
type Operation struct {
	Action      string            `json:"action,omitempty"` // "send" | "receive"
	Channel     Ref               `json:"channel,omitempty"`
	Title       string            `json:"title,omitempty"`
	Summary     string            `json:"summary,omitempty"`
	Description string            `json:"description,omitempty"`
	Bindings    OperationBindings `json:"bindings,omitempty"`
	Message     map[string]Ref    `json:"messages,omitempty"`
	Reply       Reply             `json:"reply,omitempty"`
}

type Reply struct {
	Address Address        `json:"addressee,omitempty"`
	Channel Ref            `json:"channel,omitempty"`
	Message map[string]Ref `json:"messages,omitempty"`
}

type Address struct {
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"` // $message.header#/replyTo
}

type Message struct {
	Headers       map[string]interface{} `json:"headers,omitempty"`
	Payload       Payload                `json:"payload,omitempty"`
	CorrelationId CorrelationId          `json:"correlationId,omitempty"`
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
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

type AsyncAPISpecData struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	InputParameters string `json:"inputparameters" binding:"required"`
	ReturnValue     string `json:"returnvalue" binding:"required"`
}

type Ref struct {
	Ref string `json:"$ref"`
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

	asyncAPISpec := AsyncAPI{
		Asyncapi: "3.0.0",
		Info: AsyncAPIInfo{
			Title:       functionData.Name,
			Version:     "1.0.0",
			Description: functionData.Description,
		},
		DefaultContentType: "application/json",
		Servers: map[string]Server{
			"local": {
				Host:     fmt.Sprintf("%s:%d", utils.GetEnvSting("RABBITMQ_HOST", "localhost"), utils.GetEnvInt("RABBITMQ_PORT", 5672)),
				Protocol: "amqp",
			},
		},
		Channels: map[string]Channel{
			fmt.Sprintf("input-func-%s", functionData.Name): {
				Address:     "input",
				Title:       "Input",
				Description: "The channel for sending input parameters to a serverless function.",
				Message:     map[string]Ref{"inputParameter": {Ref: "#/components/messages/InputParameter"}},
				Parameters:  map[string]Ref{"": {Ref: ""}},
				Servers:     []Ref{{Ref: "#/servers/local"}},
				Bindings: ChannelBindings{
					AMQP: AMQPBinding{
						Is: "queue",
						Queue: Queue{
							Exclusive: false,
						},
					},
				},
			},
			fmt.Sprintf("return-func-%s", functionData.Name): {
				Address:     fmt.Sprintf("%s.%s", functionData.Name, "return"),
				Title:       "Return",
				Description: "The channel for sending the return data to the client from a serverless function.",
				Message:     map[string]Ref{"returnData": {Ref: "#/components/messages/ReturnData"}},
				Parameters:  map[string]Ref{"": {Ref: ""}},
				Servers:     []Ref{{Ref: "#/servers/local"}},
				Bindings: ChannelBindings{
					AMQP: AMQPBinding{
						Is: "queue",
						Queue: Queue{
							Exclusive: false,
						},
					},
				},
			},
		},
		Operation: map[string]Operation{
			fmt.Sprintf("receive_%s", functionData.Name): {
				Action:  "receive",
				Channel: Ref{Ref: fmt.Sprintf("#/channels/input-func-%s", functionData.Name)},
			},
			fmt.Sprintf("send_%s", functionData.Name): {
				Action:  "send",
				Channel: Ref{Ref: fmt.Sprintf("#/channels/return-func-%s", functionData.Name)},
			},
		},
		Components: Components{
			Messages: map[string]Message{
				"parameterMessage": {
					Name:    "input",
					Payload: inputPayload,
				},
				"returnMessage": {
					Name:    "returnMessage",
					Payload: returnPayload,
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
