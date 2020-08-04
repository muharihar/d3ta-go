package feature

import "github.com/muharihar/d3ta-go/system/handler"

// NewFeature new Feature
func NewFeature(h *handler.Handler) (*Feature, error) {
	var err error

	f := new(Feature)
	f.handler = h

	if f.System, err = NewSystem(h); err != nil {
		return nil, err
	}

	if f.OpenAPI, err = NewOpenAPI(h); err != nil {
		return nil, err
	}

	if f.Auths, err = NewFAuths(h); err != nil {
		return nil, err
	}

	if f.Covid19, err = NewFCovid19(h); err != nil {
		return nil, err
	}

	if f.GeoLocation, err = NewFGeoLocation(h); err != nil {
		return nil, err
	}

	if f.Email, err = NewFEmail(h); err != nil {
		return nil, err
	}

	return f, nil
}

// Feature represet Feature
type Feature struct {
	BaseFeature

	System      *FSystem
	OpenAPI     *FOpenAPI
	Auths       *FAuths
	Covid19     *FCovid19
	GeoLocation *FGeoLocation
	Email       *FEmail
}
