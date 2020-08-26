package indexer

import (
	"fmt"

	es6 "github.com/elastic/go-elasticsearch/v6"
	es7 "github.com/elastic/go-elasticsearch/v7"
	es8 "github.com/elastic/go-elasticsearch/v8"

	"github.com/muharihar/d3ta-go/system/indexer/adapter"
	ieES "github.com/muharihar/d3ta-go/system/indexer/adapter/elasticsearch"
)

// NewIndexerEngine new IndexerEngine
func NewIndexerEngine(indexerType IndexerType, ieOptions adapter.IEOptions) (adapter.IIndexerEngine, IndexerType, error) {
	tOpt := fmt.Sprintf("%T", ieOptions.Options)

	switch ieOptions.Engine {
	case adapter.IEElasticSearch:
		if tOpt != "elasticsearch.Config" {
			return nil, "", fmt.Errorf("Invalid ElasticSearch options (should be `elasticsearch.Config`)")
		}

		var ie adapter.IIndexerEngine
		var err error

		switch ieOptions.Version {
		case "6":
			cfg := ieOptions.Options.(es6.Config)
			ie, err = ieES.NewIndexer(ieES.ESVersion6, cfg)
		case "7":
			cfg := ieOptions.Options.(es7.Config)
			ie, err = ieES.NewIndexer(ieES.ESVersion7, cfg)
		case "8":
			cfg := ieOptions.Options.(es8.Config)
			ie, err = ieES.NewIndexer(ieES.ESVersion8, cfg)
		default:
			err = fmt.Errorf("Invalid adapter.IEElasticSearch Version: %s", ieOptions.Version)
		}

		if err != nil {
			return nil, "", err
		}
		return ie, indexerType, nil

	default:
		return nil, "", fmt.Errorf("Invalid Indexer Engine")
	}
}
