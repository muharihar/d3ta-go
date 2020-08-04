package emailtemplate

import "encoding/json"

// ETUpdateResponse type
type ETUpdateResponse struct {
	Code    string `json:"code"`
	Version string `json:"version"`
}

// ToJSON covert to JSON
func (r *ETUpdateResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
