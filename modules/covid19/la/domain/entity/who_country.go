package entity

// WHOCountry represent WHOCountry
type WHOCountry struct {
	ID        string `json:"ID"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	ISO2Code  string `json:"ISO2code"`
	ISO3Code  string `json:"ISO3Code"`
	WHORegion string `json:"WHORegion"`
}
