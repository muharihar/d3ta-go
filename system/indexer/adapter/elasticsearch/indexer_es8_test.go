package elasticsearch

import (
	"testing"
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
