package schema

import "encoding/json"

// SendEmailResponse type
type SendEmailResponse struct {
	TemplateCode string `json:"templateCode"`
	Status       string `json:"status"`
}

// ToJSON covert to JSON
func (r *SendEmailResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
