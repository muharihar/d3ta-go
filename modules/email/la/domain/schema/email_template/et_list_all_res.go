package emailtemplate

import "encoding/json"

// ETListAllResponse type
type ETListAllResponse struct {
	Count int64            `json:"count"`
	Data  []*EmailTemplate `json:"data"`
}

// ToJSON covert to JSON
func (r *ETListAllResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
