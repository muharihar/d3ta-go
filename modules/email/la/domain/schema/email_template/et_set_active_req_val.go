package emailtemplate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate ETSetActiveRequest
func (r *ETSetActiveRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Keys, validation.Required),
		validation.Field(&r.Data, validation.Required),
	)
}

// Validate ETSetActiveKeys
func (r *ETSetActiveKeys) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Length(10, 100), validation.Required),
	)
}

// Validate ETSetActiveData
func (r *ETSetActiveData) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.IsActive, validation.NotNil),
	)
}
