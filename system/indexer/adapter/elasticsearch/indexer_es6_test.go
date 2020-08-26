package elasticsearch

import (
	"testing"
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
