package config

// Databases represent Databases
type Databases struct {
	IdentityDB Database `json:"identityDB" yaml:"identityDB"`
	MainDB     Database `json:"mainDB" yaml:"mainDB"`
	LogDB      Database `json:"logDB" yaml:"logDB"`
	EmailDB    Database `json:"emailDB" yaml:"emailDB"`
}

// Database represent Database
type Database struct {
	ConnectionName string `json:"connectionName"`
	Driver         string `json:"driver"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	HostName       string `json:"hostName"`
	DBName         string `json:"dbName"`
	Config         string `json:"config"`
	MaxIdleConns   int    `json:"maxIdleConns"`
	MaxOpenConns   int    `json:"maxOpenConns"`
	LogMode        bool   `json:"logMode"`
}
