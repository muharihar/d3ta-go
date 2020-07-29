package schema

import "encoding/json"

// LoginResponse type
type LoginResponse struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}

// ToJSON covert to JSON
func (r *LoginResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
