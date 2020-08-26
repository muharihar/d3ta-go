package indexer

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/utils"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func testCreateIndex(indexer *Indexer, index string, t *testing.T) {
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
            "fieldSpatialShape": {
                "type": "shape"
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
                "type": "text",
                "fields": {
                    "length": { 
                        "type":     "token_count",
                        "analyzer": "standard"
                    }
                }
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

	err := indexer.CreateIndex(index, strings.NewReader(mappings))
	if err != nil {
		t.Errorf("Error while creating Index: %s", err.Error())
		return
	}
}

func testIndexerMethods(indexer *Indexer, index string, t *testing.T) {
	engine := indexer.GetEngine()
	if engine == nil {
		t.Error("Should not be nil")
	}

	// Create Doc
	id := utils.GenerateUUID()
	d1 := time.Now().Format("2006-01-02 15:04:05")
	d2 := time.Now().Format(time.RFC3339Nano)
	d3 := time.Now().AddDate(1, 0, 0).Format("2006-01-02")

	_doc := fmt.Sprintf(`{
		"fieldText": "Value of fieldText",
		"fieldKeyword": "Value of fieldKeyword",
		"fieldWildcard": "We can query this field by wildcard",
		"fieldNumLong": 2147483648,
		"fieldNumInteger": 2147483647,
		"fieldNumShort": 32767,
		"fieldNumByte": 127,
		"fieldNumDouble": 2147483648.2147483648,
		"fieldNumFloat": 2147483648.2147483648,
		"fieldNumHalfFloat": 65504,
		"fieldNumScaledFloat": 2147483648.2147483648,
		"fieldDate": "%s",
		"fieldDateNanoSecords": "%s",
		"fieldBoolean": "false",
		"fieldBinary": "%s",
		"fieldRangeInteger": {
			"gte" : -2147483648,
    		"lte" : 2147483647
		},
		"fieldRangeFloat": {
			"gte" : -2147483648.2147483648,
    		"lte" : 2147483648.2147483648
		},
		"fieldRangeLong": {
			"gte" : -2147483648,
    		"lte" : 2147483648
		},
		"fieldRangeDouble": {
			"gte" : -2147483648.2147483648,
    		"lte" : 2147483648.2147483648
		},
		"fieldRangeDate": {
			"gte" : "%s",
			"lte" : "%s"
		},
		"fieldRangeIP": "%s",
		"fieldIP": "%s",
		"fieldNested": [
			{
				"fieldNested01": "Value of fieldNested01.01",
				"fieldNested02": "Value of fieldNested02.01"
			},
			{
				"fieldNested01": "Value of fieldNested01.02",
				"fieldNested02": "Value of fieldNested02.02"
			}
		],
		"fieldObject": {
			"fieldObjectText": "Value of fieldObjectText",
			"fieldObjectKeyword": "Value of fieldObjectKeyword",
			"fieldObjectWildcard": "We can query this field by wildcard",
			"fieldObjectNumLong": 2147483648,
			"fieldObjectNumInteger": 2147483647
		},
		"fieldSpatialGeoPoint": {
			"lat": -7.7829,
			"lon": 110.36708
		},
		"fieldSpatialGeoShape": {
			"type" : "multipolygon",
			"coordinates" : [
				[ [[102.0, 2.0], [103.0, 2.0], [103.0, 3.0], [102.0, 3.0], [102.0, 2.0]] ],
				[ [[100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0]],
				  [[100.2, 0.2], [100.8, 0.2], [100.8, 0.8], [100.2, 0.8], [100.2, 0.2]] ]
			]
		},
		"fieldSpatialPoint": {
			"x": 41.12,
    		"y": -71.34
		},
		"fieldSpatialShape": {
			"type": "geometrycollection",
			"geometries": [
				{
					"type": "point",
					"coordinates": [1000.0, 100.0]
				},
				{
					"type": "linestring",
					"coordinates": [ [1001.0, 100.0], [1002.0, 100.0] ]
				}
			]
		},
		"fieldArrayInteger": [1,2,3,4,5,6,7,8,9,0],
		"fieldPercolator": {
			"match": {
				"fieldKeyword": "Value of fieldKeyword"
			}
		},
		"fieldCompletion": [
			{
			"input": "Nevermind",
			"weight": 10
			},
			{
			"input": "Nirvana",
			"weight": 3
			}
		],
		"fieldTokenCount": "length of text (token count)",
		"fieldRankFeature": 8,
		"fieldRankFeatures": {
			"politics": 20,
    		"economics": 50.8
		},
		"fieldDenseVector": [0.5, 10, 6],
		"fieldSparseVector": {"1": 0.5, "5": -0.5,  "100": 1},
		"fieldSearchAsYouType": "quick brown fox jump lazy dog",
		"fieldFlattened": {
			"priority": "urgent",
			"release": ["v1.2.5", "v1.3.0"],
			"timestamp": {
			"created": 1541458026,
			"closed": 1541457010
			}
		},
		"fieldHistogram": {
			"values" : [0.1, 0.2, 0.3, 0.4, 0.5], 
			"counts" : [3, 7, 23, 12, 6] 
		},
		"fieldConstantKeyword": "ConstantKeywordValue"
	}
	`,
		d1,
		d2,
		base64.StdEncoding.EncodeToString([]byte(`binary-in-base64`)),
		d1,
		d3,
		fmt.Sprintf("%s/24", utils.GetCurrentIP()),
		utils.GetCurrentIP(),
	)

	if err := indexer.CreateDoc(index, id, strings.NewReader(_doc)); err != nil {
		t.Errorf("Error while creating document: CreateDoc(), %s", err.Error())
	}

	// Doc Exist
	exist, err := indexer.DocExist(index, id)
	if err != nil {
		t.Errorf("Error while checking document: DocExist(), %s", err.Error())
	}
	t.Logf("exist: %v\n", exist)

	// Get Doc
	doc, err := indexer.GetDoc(index, id)
	if err != nil {
		t.Errorf("Error while getting document: GetDoc(), %s", err.Error())
	}
	t.Logf("Doc: %s\n", string(doc))

	// Update Doc
	du1 := time.Now().Format("2006-01-02 15:04:05")
	du2 := time.Now().Format(time.RFC3339Nano)
	du3 := time.Now().AddDate(1, 0, 0).Format("2006-01-02")

	_docUpdate := fmt.Sprintf(`{
		"doc": {
			"fieldText": "Value of fieldText-Updated",
			"fieldKeyword": "Value of fieldKeyword",
			"fieldWildcard": "We can query this field by wildcard",
			"fieldNumLong": 2147483648,
			"fieldNumInteger": 2147483647,
			"fieldNumShort": 32767,
			"fieldNumByte": 127,
			"fieldNumDouble": 2147483648.2147483648,
			"fieldNumFloat": 2147483648.2147483648,
			"fieldNumHalfFloat": 65504,
			"fieldNumScaledFloat": 2147483648.2147483648,
			"fieldDate": "%s",
			"fieldDateNanoSecords": "%s",
			"fieldBoolean": "true",
			"fieldBinary": "%s",
			"fieldRangeInteger": {
				"gte" : -2147483648,
				"lte" : 2147483647
			},
			"fieldRangeFloat": {
				"gte" : -2147483648.2147483648,
				"lte" : 2147483648.2147483648
			},
			"fieldRangeLong": {
				"gte" : -2147483648,
				"lte" : 2147483648
			},
			"fieldRangeDouble": {
				"gte" : -2147483648.2147483648,
				"lte" : 2147483648.2147483648
			},
			"fieldRangeDate": {
				"gte" : "%s",
				"lte" : "%s"
			},
			"fieldRangeIP": "%s",
			"fieldIP": "%s"
		}
	}
	`,
		du1,
		du2,
		base64.StdEncoding.EncodeToString([]byte(`binary-in-base64-update`)),
		du1,
		du3,
		fmt.Sprintf("%s/24", utils.GetCurrentIP()),
		utils.GetCurrentIP(),
	)

	if err := indexer.UpdateDoc(index, id, strings.NewReader(_docUpdate)); err != nil {
		t.Errorf("Error while updating document: UpdateDoc(), %s", err.Error())
	}

	// Delete Doc
	if err := indexer.DeleteDoc(index, id); err != nil {
		t.Errorf("Error while deleting document: DeleteDoc(), %s", err.Error())
	}

	// Doc Exist
	exist, err = indexer.DocExist(index, id)
	if err != nil {
		t.Errorf("Error while checking document: DocExist(), %s", err.Error())
	}
	t.Logf("exist: %v\n", exist)
}
