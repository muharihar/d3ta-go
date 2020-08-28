package country

import "encoding/json"

// RefreshCountryIndexerResponse type
type RefreshCountryIndexerResponse struct {
	Status      string `json:"status"`
	ProcessType string `json:"processType"`
}

// ToJSON covert to JSON
func (r *RefreshCountryIndexerResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
