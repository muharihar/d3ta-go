package schema

import "encoding/json"

// LoginAppResponse type
type LoginAppResponse struct {
	TokenType     string `json:"tokenType"`
	ClientAppCode string `json:"clientAppCode"`
	ClientAppName string `json:"clientAppName"`
	Token         string `json:"token"`
	ExpiredAt     int64  `json:"expiredAt"`
}

// ToJSON covert to JSON
func (r *LoginAppResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
