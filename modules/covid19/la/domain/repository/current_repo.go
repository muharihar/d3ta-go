package repository

import (
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
	"github.com/muharihar/d3ta-go/system/identity"
)

// ICurrentRepo interface
type ICurrentRepo interface {
	DisplayCurrentDataByCountry(req *domSchema.DisplayCurrentDataByCountryRequest, i identity.Identity) (*domSchema.DisplayCurrentDataByCountryResponse, error)
	// GetAvailableProviders()
}
