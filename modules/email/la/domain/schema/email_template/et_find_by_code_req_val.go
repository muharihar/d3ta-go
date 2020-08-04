package emailtemplate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate ETFindByCodeRequest
func (r *ETFindByCodeRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Required),
	)
}
