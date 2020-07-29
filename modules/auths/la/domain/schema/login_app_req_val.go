package schema

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate LoginAppRequest
func (r *LoginAppRequest) Validate() error {

	return validation.ValidateStruct(r,
		validation.Field(&r.ClientKey, validation.Required),
		validation.Field(&r.SecretKey, validation.Required),
	)
}
