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
		"fieldIP": "%s"
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
	t.Logf("_doc: %s", _doc)

	if err := indexer.CreateDoc(index, id, strings.NewReader(_doc)); err != nil {
		t.Errorf("Error while creating document: CreateDoc(), %s", err.Error())
	}

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

	t.Logf("_docUpdate: %s", _docUpdate)
	if err := indexer.UpdateDoc(index, id, strings.NewReader(_docUpdate)); err != nil {
		t.Errorf("Error while updating document: UpdateDoc(), %s", err.Error())
	}
}
