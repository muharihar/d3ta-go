package country

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate AddCountryRequest
func (r *AddCountryRequest) Validate() error {

	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Required),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.ISO2Code, validation.Required, validation.Length(2, 2)),
		validation.Field(&r.ISO3Code, validation.Length(3, 3)),
		validation.Field(&r.WHORegion, validation.Required),
	)
}
