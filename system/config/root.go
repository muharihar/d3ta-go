package config

// Config represent Config
type Config struct {
	Applications Applications `json:"applications" yaml:"applications"`
	IAM          IAM          `json:"IAM" yaml:"IAM"`
	Securities   Securities   `json:"securities" yaml:"securities"`
	DirLocations DirLocations `json:"dirLocations" yaml:"dirLocations"`
	Databases    Databases    `json:"databases" yaml:"databases"`
	Connectors   Connectors   `json:"connectors" yaml:"connectors"`
	SMTPServers  SMTPServers  `json:"SMTPServers" yaml:"SMTPServers"`
}

// Databases represent Databases
type Databases struct {
	IdentityDB Database `json:"identityDB" yaml:"identityDB"`
	MainDB     Database `json:"mainDB" yaml:"mainDB"`
	LogDB      Database `json:"logDB" yaml:"logDB"`
	EmailDB    Database `json:"emailDB" yaml:"emailDB"`
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
