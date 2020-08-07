package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema"
)

// SendEmailReqDTO type
type SendEmailReqDTO struct {
	TemplateCode   string                 `json:"templateCode"`
	From           *MailAddressDTO        `json:"from"`
	To             *MailAddressDTO        `json:"to"`
	CC             []*MailAddressDTO      `json:"cc"`
	BCC            []*MailAddressDTO      `json:"bcc"`
	TemplateData   map[string]interface{} `json:"templateData"`
	ProcessingType string                 `json:"processingType"`
	// domSchema.SendEmailRequest
}

// ConvertCC2Domain convert to domSchema
func (r *SendEmailReqDTO) ConvertCC2Domain() []*domSchema.MailAddress {
	var ms []*domSchema.MailAddress
	for _, v := range r.CC {
		ms = append(ms, &domSchema.MailAddress{Email: v.Email, Name: v.Name})
	}
	return ms
}

// ConvertBCC2Domain convert to domSchema
func (r *SendEmailReqDTO) ConvertBCC2Domain() []*domSchema.MailAddress {
	var ms []*domSchema.MailAddress
	for _, v := range r.BCC {
		ms = append(ms, &domSchema.MailAddress{Email: v.Email, Name: v.Name})
	}
	return ms
}

// SendEmailResDTO type
type SendEmailResDTO struct {
	domSchema.SendEmailResponse
}

// MailAddressDTO type
type MailAddressDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ToJSON covert to JSON
func (r *SendEmailReqDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
