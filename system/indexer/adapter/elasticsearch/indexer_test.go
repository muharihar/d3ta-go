package elasticsearch

import (
	"testing"
)

func TestIndexer_Methods(t *testing.T) {

	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	cfgES7 := ConfigParserES7(cfg.Indexers.DataIndexer.Configurations)
	t.Logf("cfgES7: %#v\n", cfgES7)

	indexer, err := NewIndexer(ESVersion7, cfgES7)
	if err != nil {
		t.Errorf("Error while creating NewIndexer (ES7): %s", err.Error())
	}

	testIndexerMethods(indexer, t)
	t.Error("SHOWTEST")
}
