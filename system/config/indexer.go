package config

// Indexers represent Caches
type Indexers struct {
	DataIndexer Indexer `json:"dataIndexer" yaml:"dataIndexer"`
	LogIndexer  Indexer `json:"logIndexer" yaml:"logIndexer"`
}

// Indexer represent Indexer
type Indexer struct {
	ConnectionName string      `json:"connectionName" yaml:"connectionName"`
	Engine         string      `json:"engine" yaml:"engine"`
	Driver         string      `json:"driver" yaml:"driver"`
	Version        string      `json:"version" yaml:"version"`
	Configurations interface{} `json:"configurations" yaml:"configurations"`
}
