package initialize

import (
	"fmt"
	"strings"
	"testing"

	"github.com/muharihar/d3ta-go/system/utils"
	"github.com/stretchr/testify/assert"
)

func TestOpenAllIndexerConnection(t *testing.T) {
	h, err := newHandler(t)
	if assert.NoError(t, err, "Error while creating handler: newHandler") {
		if !assert.NotNil(t, h) {
			return
		}
	}

	if assert.NoError(t, OpenAllIndexerConnection(h), "Error while opening all indexer connection: OpenAllIndexerConnection") {

		cfg, err := h.GetConfig()
		if !assert.NoError(t, err, "Error while getting config: h.GetConfig") {
			return
		}

		idxCon, err := h.GetIndexer(cfg.Indexers.DataIndexer.ConnectionName)
		if assert.NoError(t, err, "Error while getting Indexer Engine Connection: h.GetIndexer(cfg.Indexers.DataIndexer.ConnectionName)") {

			idxCon.Context = "system"
			idxCon.Container = "indexer"
			idxCon.Component = "test"
			index := fmt.Sprintf("test-index-%s", utils.GenerateUUID())

			if assert.NoError(t, idxCon.CreateDoc(index, "test-indexer-index-id", strings.NewReader(`{ "test": "test value on initialize.indexer"}`))) {
				res, err := idxCon.GetDoc(index, "test-indexer-index-id")
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}

			assert.NoError(t, idxCon.DropIndex([]string{index}))
		}
	}
}
