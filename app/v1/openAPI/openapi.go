package openapi

import (
	"fmt"
	"os"
	"os/exec"
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
	Tags        []string            `json:"tags,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Description string              `json:"description,omitempty"`
	OperationID string              `json:"operationId,omitempty"`
	Deprecated  bool                `json:"deprecated"`
	Responses   map[string]Response `json:"responses"`
	// Add other Operation fields as needed
}

// Response represents a response in the OpenAPI specification.
type Response struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content,omitempty"`
	// Add other Response fields as needed
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

// Components represents the components section in the OpenAPI specification.
type Components struct {
	Schemas map[string]interface{} `json:"schemas,omitempty"`
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

// func CreateOpenAPISpec(outDirPath string) OpenAPI {
// 	openAPISpec := OpenAPI{
// 		Openapi: "3.0.0",
// 		Info: Info{
// 			Title:   "",
// 			Version: "1.0.0",
// 		},
// 		Servers: []Server{
// 			{},
// 		},
// 		Paths: {
// 			{"/", Path{}},
// 		},
// 	}
// 	return openAPISpec
// }
