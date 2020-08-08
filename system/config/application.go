package config

// Applications represent Applications
type Applications struct {
	Servers Servers `json:"servers" yaml:"servers"`
}

// Servers represent all available application server
type Servers struct {
	RestAPI RestAPI `json:"restapi" yaml:"restapi"`
}

// RestAPI server
type RestAPI struct {
	Name        string  `json:"name" yaml:"name"`
	Version     string  `json:"version" yaml:"version"`
	Description string  `json:"description" yaml:"description"`
	Options     Options `json:"options" yaml:"options"`
}

// Options represent Options
type Options struct {
	ShowEngineHeader  bool                  `json:"showEngineHeader" yaml:"showEngineHeader"`
	DisplayOpenAPI    bool                  `json:"displayOpenAPI" yaml:"displayOpenAPI"`
	Listener          Listener              `json:"listener" yaml:"listener"`
	Middlewares       Middlewares           `json:"middlewares" yaml:"middlewares"`
	OpenAPIDefinition RootOpenAPIDefinition `json:"openAPIDefinition" yaml:"openAPIDefinition"`
}

// Listener represent Listener
type Listener struct {
	Port string `json:"port" yaml:"port"`
}

// Middlewares type
type Middlewares struct {
	Logger MiddlewareLogger `json:"logger" yaml:"logger"`
}

// MiddlewareLogger type
type MiddlewareLogger struct {
	Enable bool `json:"enable" yaml:"enable"`
}
