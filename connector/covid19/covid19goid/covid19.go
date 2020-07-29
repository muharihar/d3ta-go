package covid19goid

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	c19type "github.com/muharihar/d3ta-go/connector/covid19/covid19goid/types"
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

	return &Covid19{Server: config.Server, httpClient: httpClient}
}

// Covid19 Client SDK
type Covid19 struct {
	Server     string
	httpClient *http.Client
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
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetUpdate get Update.json
func (c *Covid19) GetUpdate() (*c19type.UpdateResponse, error) {
	resp, err := c.getRequest("public/api/update.json", nil)
	if err != nil {
		return nil, err
	}

	respObj := c19type.UpdateResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

// GetData get Data.json
func (c *Covid19) GetData() (*c19type.DataResponse, error) {
	resp, err := c.getRequest("public/api/data.json", nil)
	if err != nil {
		return nil, err
	}

	respObj := c19type.DataResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

// GetProv get Prov.json
func (c *Covid19) GetProv() (*c19type.ProvResponse, error) {
	resp, err := c.getRequest("public/api/prov.json", nil)
	if err != nil {
		return nil, err
	}

	respObj := c19type.ProvResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}
