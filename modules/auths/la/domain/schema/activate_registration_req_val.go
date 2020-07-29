package schema

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate ActivateRegistrationRequest
func (r *ActivateRegistrationRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ActivationCode, validation.Required))
}
