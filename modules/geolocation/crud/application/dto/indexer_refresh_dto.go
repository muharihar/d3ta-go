package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/schema/country"
)

// RefreshCountryIndexerReqDTO represent RefreshCountryIndexerReqDTO
type RefreshCountryIndexerReqDTO struct {
	domSchema.RefreshCountryIndexerRequest
}

// RefreshCountryIndexerResDTO represent RefreshCountryIndexerResDTO
type RefreshCountryIndexerResDTO struct {
	domSchema.RefreshCountryIndexerResponse
}

// ToJSON covert to JSON
func (r *RefreshCountryIndexerResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
