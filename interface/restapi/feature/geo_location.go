package feature

import (
	"net/http"

	"github.com/labstack/echo/v4"
	response "github.com/muharihar/d3ta-go/interface/restapi/response"
	appGeoLoc "github.com/muharihar/d3ta-go/modules/geolocation/crud/application"
	appGeoLocDTO "github.com/muharihar/d3ta-go/modules/geolocation/crud/application/dto"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewFGeoLocation new FGeoLocation
func NewFGeoLocation(h *handler.Handler) (*FGeoLocation, error) {
	var err error

	f := new(FGeoLocation)
	f.handler = h

	if f.appGeoLocation, err = appGeoLoc.NewGeoLocationApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FGeoLocation represent FGeoLocation
type FGeoLocation struct {
	BaseFeature
	appGeoLocation *appGeoLoc.GeoLocationApp
}

// ListAllCountry list AllCountry
func (f *FGeoLocation) ListAllCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	resp, err := f.appGeoLocation.CountrySvc.ListAll(i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// RefreshCountryIndexer refresh Country Indexer
func (f *FGeoLocation) RefreshCountryIndexer(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appGeoLocDTO.RefreshCountryIndexerReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appGeoLocation.CountrySvc.RefreshIndexer(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// SearchCountryIndexer search Country Indexer
func (f *FGeoLocation) SearchCountryIndexer(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appGeoLocDTO.SearchCountryIndexerReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appGeoLocation.CountrySvc.SearchIndexer(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// GetCountry get Country
func (f *FGeoLocation) GetCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// params
	code := c.Param("code")

	req := new(appGeoLocDTO.GetCountryReqDTO)
	req.Code = code

	resp, err := f.appGeoLocation.CountrySvc.GetDetail(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// AddCountry add Country
func (f *FGeoLocation) AddCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appGeoLocDTO.AddCountryReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appGeoLocation.CountrySvc.Add(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// UpdateCountry update Country
func (f *FGeoLocation) UpdateCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// params
	code := c.Param("code")

	reqKeys := new(appGeoLocDTO.UpdateCountryKeysDTO)
	reqKeys.Code = code

	reqData := new(appGeoLocDTO.UpdateCountryDataDTO)
	if err := c.Bind(reqData); err != nil {
		return f.translateErrorMessage(err, c)
	}

	req := new(appGeoLocDTO.UpdateCountryReqDTO)
	req.Keys = reqKeys
	req.Data = reqData

	resp, err := f.appGeoLocation.CountrySvc.Update(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// DeleteCountry delete Country
func (f *FGeoLocation) DeleteCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// params
	code := c.Param("code")

	req := new(appGeoLocDTO.DeleteCountryReqDTO)
	req.Code = code

	resp, err := f.appGeoLocation.CountrySvc.Delete(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}
