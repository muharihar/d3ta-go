package elasticsearch

import (
	"fmt"
	"strings"
	"testing"
	"time"

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

	// index
	index := fmt.Sprintf("test-%s", utils.GenerateUUID())
	_mapping := `{
		"mappings": {
			"properties": {
				"title": {
					"type": "text"
				}
			}
		}
	}
	`
	// check index exist
	exist, err := indexer.IndexExist([]string{index})
	if err != nil {
		t.Errorf("Error while checking existing index is exist: %s", err.Error())
		return
	}
	if !exist {
		// create index
		if err := indexer.CreateIndex(index, strings.NewReader(_mapping)); err != nil {
			t.Errorf("Error while creating index is exist: %s", err.Error())
			return
		}
	}

	id := fmt.Sprintf("test-id-%s", utils.GenerateUUID())

	// check doc exist
	exist, err = indexer.DocExist(index, id)
	if err != nil {
		t.Errorf("Error while checking index is exist: %s", err.Error())
	}
	t.Log("Exist: ", exist)

	// create docs
	newDocs := `{
		  "title": "create-title-data"
		  }`
	err = indexer.CreateDoc(index, id, strings.NewReader(newDocs))
	if err != nil {
		t.Errorf("Error while creating document: %s", err.Error())
	}

	// check doc exist
	exist, err = indexer.DocExist(index, id)
	if err != nil {
		t.Errorf("Error while checking index is exist: %s", err.Error())
	}
	t.Log("Exist: ", exist)

	// get docs
	resp, err := indexer.GetDoc(index, id)
	if err != nil {
		t.Errorf("Error while getting index document value: %s", err.Error())
	}
	t.Logf("resp: %s", string(resp))

	// update docs
	updateDocs := `{
		  "doc": {
			  "title": "create-title-data-updated"
		  }
		}`
	err = indexer.UpdateDoc(index, id, strings.NewReader(updateDocs))
	if err != nil {
		t.Errorf("Error while creating document: %s", err.Error())
	}

	// get docs
	resp, err = indexer.GetDoc(index, id)
	if err != nil {
		t.Errorf("Error while getting index document value: %s", err.Error())
	}
	t.Logf("resp: %s", string(resp))

	// delete docs
	err = indexer.DeleteDoc(index, id)
	if err != nil {
		t.Errorf("Error while deleting document: %s", err.Error())
	}

	// check doc exist
	exist, err = indexer.DocExist(index, id)
	if err != nil {
		t.Errorf("Error while checking index  document is exist: %s", err.Error())
	}
	t.Log("Exist: ", exist)

	// drop index
	time.Sleep(20 * time.Second)
	if err := indexer.DropIndex([]string{index}); err != nil {
		t.Errorf("Error while droping index is exist: %s", err.Error())
	}
}

func testIndexerSearch(indexer adapter.IIndexerEngine, t *testing.T) {
	// search
	// req: kibana sample data (kibana_sample_data_flights)
	query := fmt.Sprintf(`
{
    "size": 2,
    "query": {
        "bool": {
            "must": [],
            "filter": [
                {
                    "match": {
                        "_index": "kibana_sample_data_flights"
                    }
				},
				{
                    "range": {
                        "timestamp": {
                            "gte": "2020-08-25T23:48:58.658Z",
                            "lte": "%s",
                            "format": "strict_date_optional_time"
                        }
                    }
                }
			],
			"should": [],
            "must_not": []
        }
	},
	"sort": [
        {
            "timestamp": {
                "order": "desc",
                "unmapped_type": "boolean"
            }
        }
    ],
	"script_fields": {
        "hour_of_day": {
            "script": {
                "source": "doc['timestamp'].value.hour",
                "lang": "painless"
            }
        }
    },
	"docvalue_fields": [
        {
            "field": "timestamp",
            "format": "date_time"
        }
    ],
    "_source": {
        "excludes": []
    }
}`, time.Now().Format("2006-01-02T15:04:05.006Z"))

	t.Log("query:", query)
	res, err := indexer.Search(strings.NewReader(query), true)
	if err != nil {
		t.Errorf("Error while Searching: Search() [%s]", err.Error())
	}
	t.Logf("Res: %s", string(res))
}

func testIndexerSearchIndexDoc(indexer adapter.IIndexerEngine, t *testing.T) {
	// search by index
	// req: kibana sample data (kibana_sample_data_flights)
	query := fmt.Sprintf(`
{
    "query": {
        "bool": {
            "must": [],
            "filter": [
                {
                    "match_all": {}
                },
                {
                    "range": {
                        "timestamp": {
                            "gte": "2020-08-25T23:48:58.658Z",
                            "lte": "%s",
                            "format": "strict_date_optional_time"
                        }
                    }
                }
            ],
            "should": [],
            "must_not": []
        }
    },
    "sort": [
        {
            "timestamp": {
                "order": "desc",
                "unmapped_type": "boolean"
            }
        }
    ],
    "script_fields": {
        "hour_of_day": {
            "script": {
                "source": "doc['timestamp'].value.hour",
                "lang": "painless"
            }
        }
    },
    "docvalue_fields": [
        {
            "field": "timestamp",
            "format": "date_time"
        }
    ],
    "_source": {
        "excludes": []
    }
}`, time.Now().Format("2006-01-02T15:04:05.006Z"))

	t.Log("query:", query)
	res, err := indexer.SearchIndexDoc("kibana_sample_data_flights", strings.NewReader(query), 2, true)
	if err != nil {
		t.Errorf("Error while Searching: SearchIndexDoc() [%s]", err.Error())
	}
	t.Logf("Res: %s", string(res))
}
