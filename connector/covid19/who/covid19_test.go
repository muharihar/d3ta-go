package covid19goid

import (
	"testing"

	c19type "github.com/muharihar/d3ta-go/connector/covid19/who/types"
)

func newCovid19(t *testing.T) *Covid19 {

	config := Config{Server: "https://covid19.who.int"}
	return NewCovid19(config, nil)
}

func TestGetCountryBySearo(t *testing.T) {

	c := newCovid19(t)
	resp, err := c.GetCountry(
		&c19type.GetCountryRequest{
			CountryCode: "id",
		})
	if err != nil {
		t.Error(err)
	}

	if resp != nil {
		t.Logf("RESPONSE.GetCountryBySearo: %#v", resp.Result.PageContext.CountryCode)
	}
}

func TestGetCountryByWpro(t *testing.T) {

	c := newCovid19(t)
	resp, err := c.GetCountry(
		&c19type.GetCountryRequest{
			CountryCode: "my",
		})
	if err != nil {
		t.Error(err)
	}

	if resp != nil {
		t.Logf("RESPONSE.GetCountryByWpro: %#v", resp.Result.PageContext.CountryCode)
	}
}

func TestGetCountry(t *testing.T) {

	c := newCovid19(t)
	resp, err := c.GetCountry(
		&c19type.GetCountryRequest{
			CountryCode: "us",
		})
	if err != nil {
		t.Error(err)
	}

	if resp != nil {
		t.Logf("RESPONSE.GetCountry: %#v", resp.Result.PageContext.CountryCode)
	}
}

func TestGetCountryNotFound(t *testing.T) {

	c := newCovid19(t)
	resp, err := c.GetCountry(
		&c19type.GetCountryRequest{
			CountryCode: "uk",
		})
	if err != nil {
		t.Logf("TestERROR: %s", err)
	}

	if resp != nil {
		t.Logf("RESPONSE.GetCountry: %#v", resp.Result.PageContext.CountryCode)
	}
}
