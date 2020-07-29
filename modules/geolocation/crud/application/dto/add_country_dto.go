package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/schema/country"
)

// AddCountryReqDTO represent AddCountryReqDTO
type AddCountryReqDTO struct {
	domSchema.AddCountryRequest
}

// AddCountryResDTO represent AddCountryResDTO
type AddCountryResDTO struct {
	Query interface{}        `json:"query"`
	Data  *domSchema.Country `json:"data"`
}

// ToJSON covert to JSON
func (r *AddCountryResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
