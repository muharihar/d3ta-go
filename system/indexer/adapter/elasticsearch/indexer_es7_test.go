package elasticsearch

import (
	"testing"
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
