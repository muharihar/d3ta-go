package elasticsearch

import (
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/estransport"
)

func TestIndexerES8_Methods(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES8 := ConfigParserES8(cfg.Indexers.DataIndexer.Configurations)

	indexer, err := NewIndexerES8(cfgES8)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES8 (ES8): %s", err.Error())
		return
	}

	testIndexerMethods(indexer, t)
}

func TestIndexerES8_SearchIndexDoc(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES8 := ConfigParserES8(cfg.Indexers.DataIndexer.Configurations)
	cfgES8.Logger = &estransport.ColorLogger{Output: os.Stdout}
	indexer, err := NewIndexerES8(cfgES8)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES8 (ES8): %s", err.Error())
		return
	}

	// search by index
	testIndexerSearchIndexDoc(indexer, t)
}

func TestIndexerES8_Search(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES8 := ConfigParserES8(cfg.Indexers.DataIndexer.Configurations)
	cfgES8.Logger = &estransport.ColorLogger{Output: os.Stdout}

	indexer, err := NewIndexerES8(cfgES8)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES8 (ES8): %s", err.Error())
		return
	}

	testIndexerSearch(indexer, t)
}
