package covid19goid

import "testing"

func newCovid19(t *testing.T) *Covid19 {

	config := Config{Server: "https://data.covid19.go.id"}
	return NewCovid19(config, nil)
}

func TestGetUpdate(t *testing.T) {

	c := newCovid19(t)
	resp, err := c.GetUpdate()
	if err != nil {
		t.Error(err)
	}

	if resp != nil {
		t.Logf("RESPONSE.GetUpdate.data: %#v", resp.Data)
	}
}

func TestGetData(t *testing.T) {

	c := newCovid19(t)
	resp, err := c.GetData()
	if err != nil {
		t.Error(err)
	}

	if resp != nil {
		t.Logf("RESPONSE.GetData.LastUpdate: %#v", resp.LastUpdate)
	}
}

func TestGetProv(t *testing.T) {

	c := newCovid19(t)
	resp, err := c.GetProv()
	if err != nil {
		t.Error(err)
	}

	if resp != nil {
		t.Logf("RESPONSE.GetProv.LastUpdate: %#v", resp.LastDate)
	}
}
