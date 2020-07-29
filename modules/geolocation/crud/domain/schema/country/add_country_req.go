package country

// AddCountryRequest type
type AddCountryRequest struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	ISO2Code  string `json:"ISO2Code"`
	ISO3Code  string `json:"ISO3Code"`
	WHORegion string `json:"WHORegion"`
}
