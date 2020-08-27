package adapter

import "io"

// IIndexerEngine contract/interface for IndexerEngine
type IIndexerEngine interface {
	GetEngine() interface{}

	Search(query io.Reader, prettify bool) ([]byte, error)
	SearchIndexDoc(index string, query io.Reader, size int, prettify bool) ([]byte, error)

	IndexExist(indexs []string) (bool, error)
	CreateIndex(index string, mapping io.Reader) error
	DropIndex(indexs []string) error

	DocExist(index string, id string) (bool, error)
	CreateDoc(index string, id string, body io.Reader) error
	GetDoc(index string, id string) ([]byte, error)
	UpdateDoc(index string, id string, body io.Reader) error
	DeleteDoc(index string, id string) error
}

// IEType Indexer Engine Type
type IEType string

const (
	// IEElasticSearch Elastic Search Indexer Engine
	IEElasticSearch IEType = "elasticsearch"
)

// IEOptions represent Indexer Engine Options
type IEOptions struct {
	// Name of Engine adapter. Default is "IEElasticSearch (elastic search)".
	Engine  IEType
	Version string
	// Options denpent on Engine (elasticsearch options, ... options)
	Options interface{}
}
