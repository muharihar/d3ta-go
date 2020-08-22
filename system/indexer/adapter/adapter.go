package adapter

// IIndexerEngine contract/interface for IndexerEngine
type IIndexerEngine interface {
	GetEngine() interface{}
}

// IEType Indexer Engine Type
type IEType string

const (
	// IEElasticSearch Elastic Search Indexer Engine
	IEElasticSearch IEType = "ElasticSearch"
)

// IEOptions represent Indexer Engine Options
type IEOptions struct {
	// Name of Engine adapter. Default is "IEElasticSearch (elastic search)".
	Engine IEType
	// Options denpent on Engine (elasticsearch options, ... options)
	Options interface{}
}
