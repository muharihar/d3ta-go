package schema

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate is a validation method for DisplayCurrentDataByCountryRequest
func (s *DisplayCurrentDataByCountryRequest) Validate() error {

	return validation.ValidateStruct(s,
		validation.Field(&s.CountryCode, validation.Required),
		validation.Field(&s.Providers, validation.Required),
	)
}
