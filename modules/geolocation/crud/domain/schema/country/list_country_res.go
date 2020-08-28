package country

import "encoding/json"

// ListCountryResponse type
type ListCountryResponse struct {
	Query interface{} `json:"query"`
	Data  []*Country  `json:"data"`
}

// ToJSON covert to JSON
func (r *ListCountryResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}

// ListCountry type
type ListCountry []*Country

// Country type
type Country struct {
	ID        int64  `json:"ID"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	ISO2Code  string `json:"ISO2Code"`
	ISO3Code  string `json:"ISO3Code"`
	WHORegion string `json:"WHORegion"`
}

// ToJSON covert to JSON
func (r *Country) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
