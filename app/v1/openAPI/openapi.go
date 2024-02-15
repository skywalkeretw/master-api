package openapi

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/skywalkeretw/master-api/app/utils"
)

// OpenAPI represents the OpenAPI specification.
type OpenAPI struct {
	Openapi      string                `json:"openapi"`
	Info         Info                  `json:"info"`
	Servers      []Server              `json:"servers,omitempty"`
	Paths        map[string]Path       `json:"paths"`
	Components   Components            `json:"components,omitempty"`
	Security     []map[string][]string `json:"security,omitempty"`
	Tags         []Tag                 `json:"tags,omitempty"`
	ExternalDocs ExternalDocs          `json:"externalDocs,omitempty"`
}

// Info represents the metadata information in the OpenAPI specification.
type Info struct {
	Title          string  `json:"title"`
	Version        string  `json:"version"`
	Description    string  `json:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
}

// Server represents a server in the OpenAPI specification.
type Server struct {
	URL         string                    `json:"url"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

// ServerVariable represents a variable for a server in the OpenAPI specification.
type ServerVariable struct {
	Default     string   `json:"default"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}

// Path represents a path in the OpenAPI specification.
type Path struct {
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Get         *Operation `json:"get,omitempty"`
	Put         *Operation `json:"put,omitempty"`
	Patch       *Operation `json:"patch,omitempty"`
	Post        *Operation `json:"post,omitempty"`
	Delete      *Operation `json:"delete,omitempty"`
	// Add other HTTP methods and their corresponding Operation as needed
	Parameters []Parameter `json:"parameters,omitempty"`
}

// Operation represents an operation in the OpenAPI specification.
type Operation struct {
	Tags        []string              `json:"tags,omitempty"`
	Summary     string                `json:"summary,omitempty"`
	Description string                `json:"description,omitempty"`
	OperationID string                `json:"operationId,omitempty"`
	Deprecated  bool                  `json:"deprecated"`
	Responses   map[string]Response   `json:"responses"`
	RequestBody *RequestBody          `json:"requestBody,omitempty"`
	Callbacks   map[string]Callback   `json:"callbacks,omitempty"`
	Security    []map[string][]string `json:"security,omitempty"`
	Servers     []Server              `json:"servers,omitempty"`
}

// Response represents a response in the OpenAPI specification.
type Response struct {
	Description string               `json:"description"`
	Schema      Schema               `json:"schema,omitempty"`
	Headers     map[string]Header    `json:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
}

// MediaType represents a media type in the OpenAPI specification.
type MediaType struct {
	Schema   interface{}            `json:"schema,omitempty"`
	Example  interface{}            `json:"example,omitempty"`
	Examples map[string]interface{} `json:"examples,omitempty"`
}

// Parameter represents a parameter in the OpenAPI specification.
type Parameter struct {
	Name        string      `json:"name"`
	In          string      `json:"in"`
	Description string      `json:"description,omitempty"`
	Required    bool        `json:"required,omitempty"`
	Schema      interface{} `json:"schema,omitempty"`
	// Add other Parameter fields as needed
}

// Components represents the components in the OpenAPI specification.
type Components struct {
	Schemas         map[string]Schema         `json:"schemas,omitempty"`
	Responses       map[string]Response       `json:"responses,omitempty"`
	Parameters      map[string]Parameter      `json:"parameters,omitempty"`
	Examples        map[string]Example        `json:"examples,omitempty"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies,omitempty"`
	Headers         map[string]Header         `json:"headers,omitempty"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
	// Add other Components fields as needed
}

// Tag represents a tag in the OpenAPI specification.
type Tag struct {
	Name         string       `json:"name"`
	Description  string       `json:"description,omitempty"`
	ExternalDocs ExternalDocs `json:"externalDocs,omitempty"`
}

// ExternalDocs represents external documentation in the OpenAPI specification.
type ExternalDocs struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
}

// Contact represents contact information in the OpenAPI specification.
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// License represents license information in the OpenAPI specification.
type License struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

func GetSwaggerCodegenHelp() {
	cmd := exec.Command("swagger-codegen", "-h")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

type Schema struct {
	Type                 string            `json:"type,omitempty"`
	Format               string            `json:"format,omitempty"`
	Title                string            `json:"title,omitempty"`
	Description          string            `json:"description,omitempty"`
	Default              interface{}       `json:"default,omitempty"`
	Enum                 []interface{}     `json:"enum,omitempty"`
	Items                *Schema           `json:"items,omitempty"`
	Properties           map[string]Schema `json:"properties,omitempty"`
	AdditionalProperties *Schema           `json:"additionalProperties,omitempty"`
	Required             []string          `json:"required,omitempty"`
	Nullable             bool              `json:"nullable,omitempty"`
	ReadOnly             bool              `json:"readOnly,omitempty"`
	WriteOnly            bool              `json:"writeOnly,omitempty"`
	Example              interface{}       `json:"example,omitempty"`
	Deprecated           bool              `json:"deprecated,omitempty"`
	// Add other Schema fields as needed
}

// Example represents an example in the OpenAPI specification.
type Example struct {
	// Add Example fields as needed
}

// Header represents a header in the OpenAPI specification.
type Header struct {
	// Add Header fields as needed
}

// RequestBody represents a request body in the OpenAPI specification.
type RequestBody struct {
	Description string               `json:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
	Required    bool                 `json:"required,omitempty"`
}

// Callback represents a callback in the OpenAPI specification.
type Callback struct {
	// Add Callback fields as needed
}

// SecurityScheme represents a security scheme in the OpenAPI specification.
type SecurityScheme struct {
	// Add SecurityScheme fields as needed
}

type OpenAPISpecData struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	InputParameters string `json:"inputparameters" binding:"required"`
	ReturnValue     string `json:"returnvalue" binding:"required"`
}

func CreateOpenAPISpec(functionData OpenAPISpecData) (string, error) {

	requestBody := RequestBody{}
	if functionData.InputParameters != "" {
		if utils.IsJSONObject(functionData.InputParameters) {
			requestBody = RequestBody{
				Description: functionData.Description,
				Content: map[string]MediaType{
					"application/json": {Schema: generateSchema(functionData.InputParameters)},
				},
				Required: true,
			}
		} else {
			return "", fmt.Errorf("unsopported format create a correct json file")
		}
	}

	var responses map[string]Response
	if utils.IsJSONObject(functionData.ReturnValue) {
		// If the return value is a JSON object
		responses = map[string]Response{
			"200": {
				Description: functionData.Description,
				Content: map[string]MediaType{
					"application/json": {Schema: generateSchema(functionData.ReturnValue)},
				},
			},
		}
	} else {
		// If the return value is a string, number, or bool
		responses = map[string]Response{
			"200": {
				Description: functionData.Description,
				Content: map[string]MediaType{
					"application/json": {Schema: Schema{Type: functionData.ReturnValue}}, // Example for a string response, modify as needed
					// For number: Schema{Type: "number"}
					// For bool: Schema{Type: "boolean"}
				},
			},
		}
	}

	openAPISpec := OpenAPI{
		Openapi: "3.0.0",
		Info: Info{
			Title:       functionData.Name,
			Version:     "1.0.0",
			Description: functionData.Description,
		},
		Servers: []Server{
			{},
		},
		Paths: map[string]Path{
			"/": {
				Post: &Operation{
					Summary:     utils.TruncateString(functionData.Description),
					Description: functionData.Description,
					RequestBody: &requestBody,
					Responses:   responses,
				},
			},
		},
	}
	openAPISpecBytes, err := json.Marshal(openAPISpec)
	if err != nil {
		fmt.Println("Error marshalling OpenAPI Specification JSON: ", err.Error())
	}
	return string(openAPISpecBytes), nil
}

func generateSchema(dataString string) Schema {
	schema := Schema{
		Properties: make(map[string]Schema), // Initialize Properties map
	}

	if utils.IsJSONObject(dataString) {
		dataMap, err := utils.JsonToMap(dataString)
		if err != nil {
			fmt.Println("error converting Json string to Map", err.Error())
			return schema // Return empty schema on error
		}

		for key, inputType := range dataMap {
			inputTypeStr := fmt.Sprintf("%v", inputType)
			strType, err := isValidOpenAPIType(inputTypeStr)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			if strType != "object" {
				schema.Properties[key] = Schema{Type: strType}
			} else {
				// Recursively handle object type
				subSchema := generateSchema(inputTypeStr)
				schema.Properties[key] = subSchema
			}
		}
	} else {
		strType, err := isValidOpenAPIType(dataString)
		if err != nil {
			fmt.Println("error checking open api type")
			return schema // Return empty schema on error
		}

		schema.Type = strType
	}

	return schema
}
