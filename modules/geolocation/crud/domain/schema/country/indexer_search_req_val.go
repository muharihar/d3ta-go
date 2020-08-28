package country

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate RefreshCountryIndexerRequest
func (r *SearchCountryIndexerRequest) Validate() error {

	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required, validation.Length(2, 10)),
	)
}
