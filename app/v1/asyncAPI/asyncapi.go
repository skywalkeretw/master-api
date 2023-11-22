package asyncapi

// AsyncAPI represents the AsyncAPI specification.
type AsyncAPI struct {
	Asyncapi     string                 `json:"asyncapi"`
	Id           string                 `json:"id,omitempty"`
	Info         AsyncAPIInfo           `json:"info"`
	Servers      []Server               `json:"servers,omitempty"`
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
	Subscribe Operation `json:"subscribe,omitempty"`
	Publish   Operation `json:"publish,omitempty"`
	// Add other channel operations (parameters, etc.) as needed
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
