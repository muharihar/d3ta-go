package country

import "encoding/json"

// DeleteCountryResponse type
type DeleteCountryResponse struct {
	Query interface{} `json:"query"`
	Data  *Country    `json:"data"`
}

// ToJSON covert to JSON
func (r *DeleteCountryResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
