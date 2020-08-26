package indexer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/muharihar/d3ta-go/system/indexer/adapter"
	ieES "github.com/muharihar/d3ta-go/system/indexer/adapter/elasticsearch"
	"github.com/muharihar/d3ta-go/system/utils"
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

	index := fmt.Sprintf("test-index-%s", utils.GenerateUUID())
	// index := fmt.Sprintf("test-index-%s", time.Now().Format("2006-01-02"))
	t.Logf("Index: `%s`", index)
	mappings := `
{
    "aliases": {
        "test-index-alias": {}
    },
    "mappings": {
        "properties": {
            "fieldText": {
                "type": "text"
            },
            "fieldKeyword": {
                "type": "keyword"
            },
            "fieldWildcard": {
                "type": "wildcard"
            },
            "fieldNumLong": {
                "type": "long"
            },
            "fieldNumInteger": {
                "type": "integer"
            },
            "fieldNumShort": {
                "type": "short"
            },
            "fieldNumByte": {
                "type": "byte"
            },
            "fieldNumDouble": {
                "type": "double"
            },
            "fieldNumFloat": {
                "type": "float"
            },
            "fieldNumHalfFloat": {
                "type": "half_float"
            },
            "fieldNumScaledFloat": {
                "type": "scaled_float",
                "scaling_factor": 100
            },
            "fieldDate": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
            },
            "fieldDateNanoSecords": {
                "type": "date_nanos"
            },
            "fieldBoolean": {
                "type": "boolean"
            },
            "fieldBinary": {
                "type": "binary"
            },
            "fieldRangeInteger": {
                "type": "integer_range"
            },
            "fieldRangeFloat": {
                "type": "float_range"
            },
            "fieldRangeLong": {
                "type": "long_range"
            },
            "fieldRangeDouble": {
                "type": "double_range"
            },
            "fieldRangeDate": {
                "type": "date_range",
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
            },
            "fieldRangeIP": {
                "type": "ip_range"
            },
            "fieldIP": {
                "type": "ip"
            },
            "fieldNested": {
                "type": "nested"
            },
            "fieldObject": {
                "properties": {
                    "fieldObjectText": {
                        "type": "text"
                    },
                    "fieldObjectKeyword": {
                        "type": "keyword"
                    },
                    "fieldObjectWildcard": {
                        "type": "wildcard"
                    },
                    "fieldObjectNumLong": {
                        "type": "long"
                    },
                    "fieldObjectNumInteger": {
                        "type": "integer"
                    }
                }
			},
			"fieldSpatialGeoPoint": {
				"type": "geo_point"
			},
			"fieldSpatialGeoShape": {
				"type": "geo_shape"
			},
			"fieldSpatialPoint": {
				"type": "point"
			},
			"fieldArray": {
				"type": "integer"
			},
			"fieldPercolator": {
				"type": "percolator"
			},
			"fieldCompletion": {
				"type": "completion"
			},
			"fieldTokenCount": {
				"type": "token_count",
				"analyzer": "standard"
			},
			"fieldRankFeature": {
				"type": "rank_feature"
			},
			"fieldRankFeatures": {
				"type": "rank_features"
			},
			"fieldDenseVector": {
				"type": "dense_vector",
				"dims": 3
			},
			"fieldSparseVector": {
				"type": "sparse_vector"
			},
			"fieldSearchAsYouType": {
				"type": "search_as_you_type"
			},
			"fieldAlias": {
				"type": "alias",
				"path": "fieldKeyword" 
			},
			"fieldFlattened": {
				"type": "flattened"
			},
			"fieldHistogram" : {
				"type" : "histogram"
			},
			"fieldConstantKeyword": {
				"type": "constant_keyword",
				"value": "ConstantKeywordValue"
			},
			"fieldJoin": { 
				"type": "join",
				"relations": {
					"parent-join": "child-join" 
				}
			}
        }
    },
    "settings": {
        "index": {
            "number_of_shards": 1,
            "number_of_replicas": 1
        }
    }
}`

	err = i.CreateIndex(index, strings.NewReader(mappings))
	if err != nil {
		t.Errorf("Error while creating Index: %s", err.Error())
		//return
	}

	testIndexerMethods(i, index, t)

	/*
		time.Sleep(20 * time.Second)
		err = i.DropIndex([]string{index})
		if err != nil {
			t.Errorf("Error while droping Index: %s", err.Error())
		}
	*/

	t.Error("SHOWTEST")
}
