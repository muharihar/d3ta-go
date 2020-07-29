package country

import "encoding/json"

// UpdateCountryResponse type
type UpdateCountryResponse struct {
	Query interface{} `json:"query"`
	Data  *Country    `json:"data"`
}

// ToJSON covert to JSON
func (r *UpdateCountryResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
