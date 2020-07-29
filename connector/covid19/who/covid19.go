package covid19goid

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	c19data "github.com/muharihar/d3ta-go/connector/covid19/who/data"
	c19type "github.com/muharihar/d3ta-go/connector/covid19/who/types"
)

// NewCovid19 new Covid19 Client SDK
func NewCovid19(config Config, httpClient *http.Client) *Covid19 {

	if httpClient == nil {
		netTransport := &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Timeout: time.Second * 10, Transport: netTransport}
	}

	countryGeometry := c19data.NewCountryGeometry()

	return &Covid19{Server: config.Server, httpClient: httpClient, CountryGeometry: countryGeometry}
}

// Covid19 Client SDK
type Covid19 struct {
	Server          string
	httpClient      *http.Client
	CountryGeometry c19data.CountryGeometry
}

func (c *Covid19) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte(`{}`), err
	}
	defer res.Body.Close()

	if res.StatusCode > 200 {
		return nil, fmt.Errorf("ERROR:[%v] %s; URL: %s", res.StatusCode, res.Status, req.URL)
	}

	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func (c *Covid19) getRequest(module string, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.Server, module)

	req, err := http.NewRequest("GET", url, body)
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetCountryByRegion get GetCountryByRegion
func (c *Covid19) GetCountryByRegion(whoRegion string, countryCode string) (*c19type.GetCountryResponse, error) {
	url := fmt.Sprintf("page-data/region/%s/country/%s/page-data.json", strings.ToLower(whoRegion), strings.ToLower(countryCode))

	resp, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := c19type.GetCountryResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

// GetCountry get GetCountry.json
func (c *Covid19) GetCountry(req *c19type.GetCountryRequest) (*c19type.GetCountryResponse, error) {

	info := c.CountryGeometry.FindByCountryCode(req.CountryCode)
	if info == nil {
		return nil, fmt.Errorf("ERROR: [%s] Country Not Found", req.CountryCode)
	}

	respObj, err := c.GetCountryByRegion(string(info.WHORegion), info.ISO2Code)
	if err != nil {
		return nil, err
	}
	respObj.Result.PageContext.Feature.Properties.ISO3Code = info.ISO3Code
	respObj.Result.PageContext.Feature.Properties.WHORegion = string(info.WHORegion)

	return respObj, nil
}
