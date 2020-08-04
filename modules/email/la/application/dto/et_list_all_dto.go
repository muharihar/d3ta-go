package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema/email_template"
)

// ETListAllResDTO type
type ETListAllResDTO struct {
	domSchema.ETListAllResponse
}

// ToJSON covert to JSON
func (r *ETListAllResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
