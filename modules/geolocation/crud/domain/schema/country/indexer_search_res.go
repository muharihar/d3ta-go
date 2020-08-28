package country

import "encoding/json"

// SearchCountryIndexerResponse type
type SearchCountryIndexerResponse struct {
	Query interface{} `json:"query"`
	Data  []*Country  `json:"data"`
}

// ToJSON covert to JSON
func (r *SearchCountryIndexerResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
