package country

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate RefreshCountryIndexerRequest
func (r *RefreshCountryIndexerRequest) Validate() error {

	return validation.ValidateStruct(r,
		validation.Field(&r.ProcessType, validation.Required, validation.In("SYNC", "ASYNC")),
	)
}
