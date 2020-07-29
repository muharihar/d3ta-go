package country

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate UpdateCountryRequest
func (r *UpdateCountryRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Keys, validation.Required),
		validation.Field(&r.Data, validation.Required),
	)
}

// Validate UpdateCountryKeys
func (r *UpdateCountryKeys) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Required, validation.Length(2, 2)),
	)
}

// Validate UpdateCountryData
func (r *UpdateCountryData) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.ISO2Code, validation.Required, validation.Length(2, 2)),
		validation.Field(&r.ISO3Code, validation.Length(3, 3)),
		validation.Field(&r.WHORegion, validation.Required),
	)
}
