package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/schema/country"
)

// ListCountryResDTO represent ListCountryResDTO
type ListCountryResDTO struct {
	Query interface{}          `json:"query"`
	Data  []*domSchema.Country `json:"data"`
}

// ToJSON covert to JSON
func (r *ListCountryResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
