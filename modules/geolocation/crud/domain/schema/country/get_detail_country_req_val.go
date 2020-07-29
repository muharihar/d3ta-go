package country

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate GetDetailCountryRequest
func (r *GetDetailCountryRequest) Validate() error {

	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Required),
	)
}
