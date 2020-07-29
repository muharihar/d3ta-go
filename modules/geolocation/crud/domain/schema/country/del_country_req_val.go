package country

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate DeleteCountryRequest
func (r *DeleteCountryRequest) Validate() error {

	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Required),
	)
}
