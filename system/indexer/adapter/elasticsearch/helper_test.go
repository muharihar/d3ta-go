package elasticsearch

import (
	"fmt"
	"strings"
	"testing"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/indexer/adapter"
	"github.com/muharihar/d3ta-go/system/utils"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func testIndexerMethods(indexer adapter.IIndexerEngine, t *testing.T) {
	engine := indexer.GetEngine()
	if engine == nil {
		t.Error("Should not be nil")
	}

	id := fmt.Sprintf("test-id-%s", utils.GenerateUUID())

	// check exist
	exist, err := indexer.DocExist("test", id)
	if err != nil {
		t.Errorf("Error while checking index is exist: %s", err.Error())
	}
	t.Log("Exist: ", exist)

	// create docs
	newDocs := `{
		  "title": "create-title-data"
		  }`
	err = indexer.CreateDoc("test", id, strings.NewReader(newDocs))
	if err != nil {
		t.Errorf("Error while creating document: %s", err.Error())
	}

	// check exist
	exist, err = indexer.DocExist("test", id)
	if err != nil {
		t.Errorf("Error while checking index is exist: %s", err.Error())
	}
	t.Log("Exist: ", exist)

	// get docs
	resp, err := indexer.GetDoc("test", id)
	if err != nil {
		t.Errorf("Error while getting index value: %s", err.Error())
	}
	t.Logf("resp: %s", string(resp))

	// update docs
	updateDocs := `{
		  "doc": {
			  "title": "create-title-data-updated"
		  }
		}`
	err = indexer.UpdateDoc("test", id, strings.NewReader(updateDocs))
	if err != nil {
		t.Errorf("Error while creating document: %s", err.Error())
	}

	// get docs
	resp, err = indexer.GetDoc("test", id)
	if err != nil {
		t.Errorf("Error while getting index value: %s", err.Error())
	}
	t.Logf("resp: %s", string(resp))

	// delete docs
	err = indexer.DeleteDoc("test", id)
	if err != nil {
		t.Errorf("Error while deleting document: %s", err.Error())
	}

	// check exist
	exist, err = indexer.DocExist("test", id)
	if err != nil {
		t.Errorf("Error while checking index is exist: %s", err.Error())
	}
	t.Log("Exist: ", exist)
}
