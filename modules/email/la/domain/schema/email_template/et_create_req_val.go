package emailtemplate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate ETCreateRequest
func (r *ETCreateRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Required, validation.Length(10, 100)),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.IsActive, validation.NotNil),
		validation.Field(&r.EmailFormat, validation.Required, validation.In("TEXT", "HTML")),
		validation.Field(&r.Template, validation.Required),
	)
}

// Validate ETCreateVersion
func (r *ETCreateVersion) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.SubjectTpl, validation.Required, validation.Length(10, 255)),
		validation.Field(&r.BodyTpl, validation.Required, validation.Length(10, 20000)),
	)
}
