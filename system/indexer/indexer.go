package indexer

import "github.com/muharihar/d3ta-go/system/indexer/adapter"

// IndexerType represent IndexerType
type IndexerType string

const (
	ElasticSearchIndexer IndexerType = "ElasticSearch"
)

// Indexer type
type Indexer struct {
	indexerEngine adapter.IIndexerEngine
}
