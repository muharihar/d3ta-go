package indexer

import (
	"fmt"
	"testing"
	"time"

	"github.com/muharihar/d3ta-go/system/indexer/adapter"
	ieES "github.com/muharihar/d3ta-go/system/indexer/adapter/elasticsearch"
)

func TestEngine_Methods(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
		return
	}

	ie, it, err := NewIndexerEngine(ElasticSearchIndexer,
		adapter.IEOptions{
			Engine:  adapter.IEType(cfg.Indexers.DataIndexer.Engine),
			Version: cfg.Indexers.DataIndexer.Version,
			Options: ieES.ConfigParserES7(cfg.Indexers.DataIndexer.Configurations),
		},
	)

	i, err := NewIndexer(it, ie)
	if err != nil {
		t.Errorf("Error while creating NewIndexer: %s", err.Error())
	}
	i.Context = "system"
	i.Container = "indexer"
	i.Component = "test"

	// index := fmt.Sprintf("test-index-%s", utils.GenerateUUID())
	index := fmt.Sprintf("test-index-%s", time.Now().Format("2006-01-02"))
	t.Logf("Index: `%s`", index)

	// check index
	exist, err := i.IndexExist([]string{index})
	if err != nil {
		t.Errorf("Error while checking Index: %s", err.Error())
		return
	}
	if !exist {
		// create index
		testCreateIndex(i, index, t)
	}

	// test index/document operation
	testIndexerMethods(i, index, t)

	// drop index
	time.Sleep(20 * time.Second)
	err = i.DropIndex([]string{index})
	if err != nil {
		t.Errorf("Error while droping Index: %s", err.Error())
		return
	}
}
