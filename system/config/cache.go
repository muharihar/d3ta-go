package config

// Caches represent Caches
type Caches struct {
	SessionCache Cache `json:"sessionCache" yaml:"sessionCache"`
	TmpDataCache Cache `json:"tmpDataCache" yaml:"tmpDataCache"`
}

// Cache represent Cache
type Cache struct {
	ConnectionName string      `json:"connectionName" yaml:"connectionName"`
	Engine         string      `json:"engine" yaml:"engine"`
	Driver         string      `json:"driver" yaml:"driver"`
	Configurations interface{} `json:"configurations" yaml:"configurations"`
}
