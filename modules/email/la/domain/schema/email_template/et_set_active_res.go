package emailtemplate

import "encoding/json"

// ETSetActiveResponse type
type ETSetActiveResponse struct {
	Code     string `json:"code"`
	IsActive bool   `json:"isActive"`
}

// ToJSON covert to JSON
func (r *ETSetActiveResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
