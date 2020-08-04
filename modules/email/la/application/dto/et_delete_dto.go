package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema/email_template"
)

// ETDeleteReqDTO type
type ETDeleteReqDTO struct {
	domSchema.ETDeleteRequest
}

// ETDeleteResDTO type
type ETDeleteResDTO struct {
	domSchema.ETDeleteResponse
}

// ToJSON covert to JSON
func (r *ETDeleteResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
