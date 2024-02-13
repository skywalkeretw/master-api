package asyncapi

// AsyncAPI represents the AsyncAPI specification.
type AsyncAPI struct {
	Asyncapi     string                 `json:"asyncapi"`
	Id           string                 `json:"id,omitempty"` //"urn:example:com:smartylighting:streetlights:server"
	Info         AsyncAPIInfo           `json:"info"`
	Servers      map[string]Server      `json:"servers,omitempty"`
	Channels     map[string]Channel     `json:"channels"`
	Components   AsyncAPIComponents     `json:"components,omitempty"`
	ExternalDocs ExternalDocs           `json:"externalDocs,omitempty"`
	Tags         []Tag                  `json:"tags,omitempty"`
	Extensions   map[string]interface{} `json:"x-extensions,omitempty"`
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
	URL         string                    `json:"url"`
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
	Address   string     `json:"address,omitempty"`
	Subscribe Operation  `json:"subscribe,omitempty"`
	Publish   Operation  `json:"publish,omitempty"`
	Request   *Operation `json:"request,omitempty"`
	Response  *Operation `json:"response,omitempty"`
}

// AsyncAPIComponents represents the components section in the AsyncAPI specification.
type AsyncAPIComponents struct {
	Schemas map[string]interface{} `json:"schemas,omitempty"`
	// Add other AsyncAPIComponents fields as needed
}

// Tag represents a tag in the AsyncAPI specification.
type Tag struct {
	Name         string       `json:"name"`
	Description  string       `json:"description,omitempty"`
	ExternalDocs ExternalDocs `json:"externalDocs,omitempty"`
}

// Operation represents an operation in the AsyncAPI specification.
type Operation struct {
	Tags        []string               `json:"tags,omitempty"`
	Summary     string                 `json:"summary,omitempty"`
	Description string                 `json:"description,omitempty"`
	OperationID string                 `json:"operationId,omitempty"`
	Bindings    map[string]interface{} `json:"bindings,omitempty"`
	// Add other Operation fields as needed
}

// ExternalDocs represents external documentation in the AsyncAPI specification.
type ExternalDocs struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
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

func CreateAsyncAPISpec(functionData AsyncAPISpecData) (AsyncAPI, error) {
	asyncAPISpec := AsyncAPI{
		Asyncapi: "3.0.0",
		Info: AsyncAPIInfo{
			Title:       functionData.Name,
			Version:     "1.0.0",
			Description: functionData.Description,
		},
		Servers: map[string]Server{
			"local": {
				URL:      "",
				Protocol: "amqp",
			},
		},
		Channels: map[string]Channel{
			"func-xxx-input-channel": {
				Address: "",
			},
			"func-xxx-output-channel": {},
		},
	}

	return asyncAPISpec, nil
}
