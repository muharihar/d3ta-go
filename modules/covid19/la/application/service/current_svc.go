package service

import (
	"fmt"

	appDto "github.com/muharihar/d3ta-go/modules/covid19/la/application/dto"
	domRepo "github.com/muharihar/d3ta-go/modules/covid19/la/domain/repository"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
	infraRepo "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/repository"
	sysError "github.com/muharihar/d3ta-go/system/error"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

// NewCurrentSvc new CurrentSvc
func NewCurrentSvc(h *handler.Handler) (*CurrentSvc, error) {
	var err error

	svc := new(CurrentSvc)
	svc.handler = h

	if svc.repo, err = infraRepo.NewCurrentRepo(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// CurrentSvc represent Current service
type CurrentSvc struct {
	BaseSvc
	repo domRepo.ICurrentRepo
}

// DisplayCurrentDataByCountry display Current Data By Country
func (s *CurrentSvc) DisplayCurrentDataByCountry(req *appDto.DisplayCurrentDataByCountryReqDTO, i identity.Identity) (*appDto.DisplayCurrentDataByCountryResDTO, error) {
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	reqDom := domSchema.DisplayCurrentDataByCountryRequest{
		CountryCode: req.CountryCode,
		Providers:   req.Providers,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	respDom, err := s.repo.DisplayCurrentDataByCountry(&reqDom, i)
	if err != nil {
		return nil, err
	}

	respDTO := appDto.DisplayCurrentDataByCountryResDTO{
		Query: respDom.Query,
		Data:  respDom.Data,
	}

	return &respDTO, nil
}
