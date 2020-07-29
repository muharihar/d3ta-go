package config

// Config represent Config
type Config struct {
	Applications      Applications          `json:"applications" yaml:"applications"`
	IAM               IAM                   `json:"IAM" yaml:"IAM"`
	Securities        Securities            `json:"securities" yaml:"securities"`
	DirLocations      DirLocations          `json:"dirLocations" yaml:"dirLocations"`
	Databases         Databases             `json:"databases" yaml:"databases"`
	Connectors        Connectors            `json:"connectors" yaml:"connectors"`
	SMTPServers       SMTPServers           `json:"SMTPServers" yaml:"SMTPServers"`
	OpenAPIDefinition RootOpenAPIDefinition `json:"openAPIDefinition" yaml:"openAPIDefinition"`
}

// Applications represent Applications
type Applications struct {
	Name        string  `json:"name" yaml:"name"`
	Version     string  `json:"version" yaml:"version"`
	Description string  `json:"description" yaml:"description"`
	Options     Options `json:"options" yaml:"options"`
}

// Options represent Options
type Options struct {
	ShowEngineHeader bool     `json:"showEngineHeader" yaml:"showEngineHeader"`
	DisplayOpenAPI   bool     `json:"displayOpenAPI" yaml:"displayOpenAPI"`
	Listener         Listener `json:"listener" yaml:"listener"`
}

// Listener represent Listener
type Listener struct {
	Port string `json:"port" yaml:"port"`
}

// Databases represent Databases
type Databases struct {
	IdentityDB Database `json:"identityDB" yaml:"identityDB"`
	MainDB     Database `json:"mainDB" yaml:"mainDB"`
	LogDB      Database `json:"logDB" yaml:"logDB"`
}

// DirLocations represent DirLocations
type DirLocations struct {
	Conf  string `json:"conf" yaml:"conf"`
	WWW   string `json:"www" yaml:"www"`
	Temp  string `json:"temp" yaml:"temp"`
	Log   string `json:"log" yaml:"log"`
	Cache string `json:"cache" yaml:"cache"`
}

// Connectors type
type Connectors struct {
	Identity Identity   `json:"identity" yaml:"identity"`
	Covid19  ConCovid19 `json:"covid19" yaml:"covid19"`
}

// Identity type
type Identity struct {
	EA2M EA2M `json:"ea2m" yaml:"ea2m"`
}

// EA2M Type
type EA2M struct {
	Server            string `json:"server" yaml:"server"`
	ClientAccessKey   string `json:"clientAccessKey" yaml:"clientAccessKey"`
	ClientSecretKey   string `json:"clientSecretKey" yaml:"clientSecretKey"`
	AllowDevToken     bool   `json:"allowDevToken" yaml:"allowDevToken"`
	DevIdentityToken  string `json:"devIdentityToken" yaml:"devIdentityToken"`
	DevIdentityClaims string `json:"devIdentityClaims" yaml:"devIdentityClaims"`
}
