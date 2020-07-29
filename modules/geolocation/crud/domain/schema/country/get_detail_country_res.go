package country

import "encoding/json"

// GetDetailCountryResponse type
type GetDetailCountryResponse struct {
	Query interface{} `json:"query"`
	Data  *Country    `json:"data"`
}

// ToJSON covert to JSON
func (r *GetDetailCountryResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
