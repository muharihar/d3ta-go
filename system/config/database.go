package config

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
