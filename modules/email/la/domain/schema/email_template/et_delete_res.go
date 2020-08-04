package emailtemplate

import "encoding/json"

// ETDeleteResponse type
type ETDeleteResponse struct {
	Query *ETDeleteRequest      `json:"query"`
	Data  *ETDeleteResponseData `json:"data"`
}

// ETDeleteResponseData type
type ETDeleteResponseData struct {
	EmailTemplate
	VersionCount int64 `json:"versionCount"`
}

// ToJSON covert to JSON
func (r *ETDeleteResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
