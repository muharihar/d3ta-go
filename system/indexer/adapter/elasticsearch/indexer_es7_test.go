package elasticsearch

import (
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/estransport"
)

func TestIndexerES7_Methods(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES7 := ConfigParserES7(cfg.Indexers.DataIndexer.Configurations)
	indexer, err := NewIndexerES7(cfgES7)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES7 (ES7): %s", err.Error())
		return
	}

	testIndexerMethods(indexer, t)
}

func TestIndexerES7_SearchIndexDoc(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES7 := ConfigParserES7(cfg.Indexers.DataIndexer.Configurations)
	cfgES7.Logger = &estransport.ColorLogger{Output: os.Stdout}
	indexer, err := NewIndexerES7(cfgES7)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES7 (ES7): %s", err.Error())
		return
	}

	// search by index
	testIndexerSearchIndexDoc(indexer, t)
}

func TestIndexerES7_Search(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES7 := ConfigParserES7(cfg.Indexers.DataIndexer.Configurations)
	cfgES7.Logger = &estransport.ColorLogger{Output: os.Stdout}

	indexer, err := NewIndexerES7(cfgES7)
	if err != nil {
		t.Errorf("Error while creating NewIndexerES7 (ES7): %s", err.Error())
		return
	}

	testIndexerSearch(indexer, t)
}
