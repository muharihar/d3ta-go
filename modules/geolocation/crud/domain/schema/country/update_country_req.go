package country

// UpdateCountryRequest type
type UpdateCountryRequest struct {
	Keys *UpdateCountryKeys `json:"keys"`
	Data *UpdateCountryData `json:"data"`
}

// UpdateCountryKeys type
type UpdateCountryKeys struct {
	Code string `json:"code"`
}

// UpdateCountryData type
type UpdateCountryData struct {
	Name      string `json:"name"`
	ISO2Code  string `json:"ISO2Code"`
	ISO3Code  string `json:"ISO3Code"`
	WHORegion string `json:"WHORegion"`
}
