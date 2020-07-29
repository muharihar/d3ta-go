package schema

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
