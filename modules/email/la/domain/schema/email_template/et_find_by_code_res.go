package emailtemplate

import "encoding/json"

// ETFindByCodeResponse type
type ETFindByCodeResponse struct {
	Query ETFindByCodeRequest `json:"query"`
	Data  ETFindByCodeData    `json:"data"`
}

// ETFindByCodeData type
type ETFindByCodeData struct {
	EmailTemplate
	DefaultTemplateVersion EmailTemplateVersion `json:"defaultTemplate"`
}

// ToJSON covert to JSON
func (r *ETFindByCodeResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
