package data

import (
	"encoding/json"
	"strings"

	c19type "github.com/muharihar/d3ta-go/connector/covid19/who/types"
)

// NewCountryGeometry new NewCountryGeometry
func NewCountryGeometry() CountryGeometry {

	c := CountryGeometry{}
	c.loadData()

	return c
}

// CountryGeometry represent CountryGeometry
type CountryGeometry struct {
	Data *c19type.WHOCountryByGeometry
}

func (c *CountryGeometry) loadData() {
	obj := c19type.WHOCountryByGeometry{}
	err := json.Unmarshal(WHOCountryGeometryData, &obj)
	if err != nil {
		panic(err)
	}
	c.Data = &obj
}

// FindByCountryCode function FindByCountryCode
func (c *CountryGeometry) FindByCountryCode(code string) *c19type.GeometryProperties {

	var result c19type.GeometryProperties

	for _, val := range c.Data.Objects.Countries.Geometries {
		if strings.ToUpper(val.Properties.ISO2Code) == strings.ToUpper(code) {
			result = val.Properties
		}
	}

	if result.ISO2Code == "" {
		return nil
	}

	return &result
}
