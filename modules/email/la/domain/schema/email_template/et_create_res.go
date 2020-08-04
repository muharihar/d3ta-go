package emailtemplate

import "encoding/json"

// ETCreateResponse type
type ETCreateResponse struct {
	Code    string `json:"code"`
	Version string `json:"version"`
}

// ToJSON covert to JSON
func (r *ETCreateResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
