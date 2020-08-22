package elasticsearch

import (
	"github.com/muharihar/d3ta-go/system/indexer/adapter"

	es6 "github.com/elastic/go-elasticsearch/v6"
	es7 "github.com/elastic/go-elasticsearch/v7"
	es8 "github.com/elastic/go-elasticsearch/v8"
)

// NewIndexer new Elastic Search Indexer
func NewIndexer() (adapter.IIndexerEngine, error) {
	return nil, nil
}

// Indexer type
type Indexer struct {
	engineES6 *es6.Client
	engineES7 *es7.Client
	engineES8 *es8.Client
}

// ESVersion represent Elastic Version Type
type ESVersion string

const (
	ESVersion6 ESVersion = "6"
	ESVersion7 ESVersion = "7"
	ESVersion8 ESVersion = "8"
)
