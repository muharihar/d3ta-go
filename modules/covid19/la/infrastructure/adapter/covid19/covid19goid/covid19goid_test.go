package who

import (
	"encoding/json"
	"testing"

	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
	infC19Adp "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/adapter/covid19"
	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
)

func newConfig(t *testing.T) (*config.Config, error) {

	c, _, err := config.NewConfig("../../../../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newCovid19goidAdapter(t *testing.T) (infC19Adp.ICovid19Adapter, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetConfig(c)

	adp, _, err := NewCovid19goidAdapter(h)
	if err != nil {
		return nil, err
	}

	return adp, nil
}

func TestCovid19goidAdapter_DisplayCurrentDataByCountry(t *testing.T) {
	adp, err := newCovid19goidAdapter(t)
	if err != nil {
		t.Errorf("newCovid19goidAdapter: %s", err.Error())
	}

	reqJSON := []byte(`{ "countryCode": "ID", "providers": [ { "code": "COVID19COID" } ] }`)

	var req domSchema.DisplayCurrentDataByCountryRequest
	if err := json.Unmarshal(reqJSON, &req); err != nil {
		t.Errorf("Unmarshal: %s", err.Error())
	}

	res, err := adp.DisplayCurrentDataByCountry(&req)
	if err != nil {
		t.Errorf("Request: [%s]", err.Error())
	}

	if res == nil {
		t.Fail()
	}

	if res != nil {
		resJSON, err := json.Marshal(res)
		if err != nil {
			t.Errorf("json.Marshal: [%s]", err.Error())
		}
		t.Logf("Resp: %s", string(resJSON))
	}
}
