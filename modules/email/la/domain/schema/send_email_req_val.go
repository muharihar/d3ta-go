package schema

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	valIs "github.com/go-ozzo/ozzo-validation/v4/is"
)

// Validate SendEmailRequest
func (r *SendEmailRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.TemplateCode, validation.Required),
		validation.Field(&r.From, validation.Required),
		validation.Field(&r.To, validation.Required),
		validation.Field(&r.TemplateData, validation.Required),
		validation.Field(&r.ProcessingType, validation.Required, validation.In("SYNC", "ASYNC")),
	)
}

// Validate MailAddress
func (r *MailAddress) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email, valIs.Email, validation.Required),
		validation.Field(&r.Name, validation.Required),
	)
}
