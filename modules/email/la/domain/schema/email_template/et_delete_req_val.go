package emailtemplate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate ETDeleteRequest
func (r *ETDeleteRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Length(10, 100), validation.Required),
	)
}
