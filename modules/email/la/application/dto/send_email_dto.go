package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema"
)

// SendEmailReqDTO type
type SendEmailReqDTO struct {
	domSchema.SendEmailRequest
}

// SendEmailResDTO type
type SendEmailResDTO struct {
	domSchema.SendEmailResponse
}

// ToJSON covert to JSON
func (r *SendEmailReqDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
