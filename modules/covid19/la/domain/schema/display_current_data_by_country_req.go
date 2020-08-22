package schema

import "encoding/json"

// DisplayCurrentDataByCountryRequest represent DisplayCurrentDataByCountryRequest
type DisplayCurrentDataByCountryRequest struct {
	CountryCode string      `json:"countryCode"`
	Providers   []*Provider `json:"providers"`
}

// ProviderList represent Provider List
type ProviderList []*Provider

// Provider represent Provider
type Provider struct {
	Code string `json:"code"`
}

// ToJSON covert to JSON
func (r *DisplayCurrentDataByCountryRequest) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
