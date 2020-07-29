package service

import (
	"fmt"

	appDTO "github.com/muharihar/d3ta-go/modules/geolocation/crud/application/dto"
	domSchema "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/schema/country"
	domSVC "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/service"
	sysError "github.com/muharihar/d3ta-go/system/error"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

// NewCountrySvc new CountrySvc
func NewCountrySvc(h *handler.Handler) (*CountrySvc, error) {
	svc := new(CountrySvc)
	svc.handler = h

	domSvc, err := domSVC.NewCountrySvc(h)
	if err != nil {
		return nil, err
	}
	svc.domSvc = domSvc

	return svc, nil
}

// CountrySvc represent CountrySvc
type CountrySvc struct {
	BaseSvc
	domSvc *domSVC.CountrySvc
}

// ListAll list All Country
func (s *CountrySvc) ListAll(i identity.Identity) (*appDTO.ListCountryResDTO, error) {
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	res, err := s.domSvc.ListAll(i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.ListCountryResDTO)
	resDTO.Query = res.Query
	resDTO.Data = res.Data

	return resDTO, nil
}

// GetDetail get Detail Country
func (s *CountrySvc) GetDetail(req *appDTO.GetCountryReqDTO, i identity.Identity) (*appDTO.GetCountryResDTO, error) {
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	reqDom := &domSchema.GetDetailCountryRequest{
		Code: req.Code,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.domSvc.GetDetail(reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.GetCountryResDTO)
	resDTO.Query = req
	resDTO.Data = res.Data

	return resDTO, nil
}

// Add add Country
func (s *CountrySvc) Add(req *appDTO.AddCountryReqDTO, i identity.Identity) (*appDTO.AddCountryResDTO, error) {
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	reqDom := &domSchema.AddCountryRequest{
		Code:      req.Code,
		Name:      req.Name,
		ISO2Code:  req.ISO2Code,
		ISO3Code:  req.ISO3Code,
		WHORegion: req.WHORegion,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.domSvc.Add(reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.AddCountryResDTO)
	resDTO.Query = req
	resDTO.Data = res.Data

	return resDTO, nil
}

// Update update Country
func (s *CountrySvc) Update(req *appDTO.UpdateCountryReqDTO, i identity.Identity) (*appDTO.UpdateCountryResDTO, error) {
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	reqDom := &domSchema.UpdateCountryRequest{
		Keys: &domSchema.UpdateCountryKeys{Code: req.Keys.Code},
		Data: &domSchema.UpdateCountryData{
			Name:      req.Data.Name,
			ISO2Code:  req.Data.ISO2Code,
			ISO3Code:  req.Data.ISO3Code,
			WHORegion: req.Data.WHORegion,
		},
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.domSvc.Update(reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.UpdateCountryResDTO)
	resDTO.Query = req
	resDTO.Data = res.Data

	return resDTO, nil
}

// Delete delete Country
func (s *CountrySvc) Delete(req *appDTO.DeleteCountryReqDTO, i identity.Identity) (*appDTO.DeleteCountryResDTO, error) {
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	reqDom := &domSchema.DeleteCountryRequest{
		Code: req.Code,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.domSvc.Delete(reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.DeleteCountryResDTO)
	resDTO.Query = req
	resDTO.Data = res.Data

	return resDTO, nil
}
