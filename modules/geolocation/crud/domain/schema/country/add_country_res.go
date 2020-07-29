package country

import "encoding/json"

// AddCountryResponse type
type AddCountryResponse struct {
	Query interface{} `json:"query"`
	Data  *Country    `json:"data"`
}

// ToJSON covert to JSON
func (r *AddCountryResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
