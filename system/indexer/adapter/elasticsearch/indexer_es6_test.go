package elasticsearch

import (
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v6/estransport"
)

func TestIndexerES6_Methods(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES6 := ConfigParserES6(cfg.Indexers.LogIndexer.Configurations)

	indexer, err := NewIndexerES6(cfgES6)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES6 (ES6): %s", err.Error())
		return
	}

	testIndexerMethods(indexer, t)
}

func TestIndexerES6_SearchIndexDoc(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES6 := ConfigParserES6(cfg.Indexers.DataIndexer.Configurations)
	cfgES6.Logger = &estransport.ColorLogger{Output: os.Stdout}
	indexer, err := NewIndexerES6(cfgES6)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES6 (ES6): %s", err.Error())
		return
	}

	// search by index
	testIndexerSearchIndexDoc(indexer, t)
}

func TestIndexerES6_Search(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES6 := ConfigParserES6(cfg.Indexers.DataIndexer.Configurations)
	cfgES6.Logger = &estransport.ColorLogger{Output: os.Stdout}

	indexer, err := NewIndexerES6(cfgES6)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES6 (ES6): %s", err.Error())
		return
	}

	testIndexerSearch(indexer, t)
}
