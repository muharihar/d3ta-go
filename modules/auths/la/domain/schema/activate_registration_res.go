package schema

import "encoding/json"

// ActivateRegistrationResponse type
type ActivateRegistrationResponse struct {
	Email       string `json:"email"`
	NickName    string `json:"NickName"`
	DefaultRole string `json:"defaultRole"`
}

// ToJSON covert to JSON
func (r *ActivateRegistrationResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
