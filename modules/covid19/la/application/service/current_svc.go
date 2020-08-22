package service

import (
	"encoding/json"
	"fmt"

	appDto "github.com/muharihar/d3ta-go/modules/covid19/la/application/dto"
	domRepo "github.com/muharihar/d3ta-go/modules/covid19/la/domain/repository"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
	infraRepo "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/repository"
	"github.com/muharihar/d3ta-go/system/cacher"
	sysError "github.com/muharihar/d3ta-go/system/error"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
	"github.com/muharihar/d3ta-go/system/utils"
)

// NewCurrentSvc new CurrentSvc
func NewCurrentSvc(h *handler.Handler) (*CurrentSvc, error) {
	var err error

	svc := new(CurrentSvc)
	svc.handler = h

	if svc.repo, err = infraRepo.NewCurrentRepo(h); err != nil {
		return nil, err
	}

	if svc.ce, err = svc.SetCacher("module", "covid19", "DisplayCurrentDataByCountry"); err != nil {
		return nil, err
	}

	return svc, nil
}

// CurrentSvc represent Current service
type CurrentSvc struct {
	BaseSvc
	repo domRepo.ICurrentRepo
	ce   *cacher.Cacher
}

// DisplayCurrentDataByCountry display Current Data By Country
func (s *CurrentSvc) DisplayCurrentDataByCountry(req *appDto.DisplayCurrentDataByCountryReqDTO, i identity.Identity) (*appDto.DisplayCurrentDataByCountryResDTO, error) {
	// authorization
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	// request domain
	reqDom := domSchema.DisplayCurrentDataByCountryRequest{
		CountryCode: req.CountryCode,
		Providers:   req.Providers,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	// load response from cache (if exist)
	exist, respDTO, err := s._loadFromCache(&reqDom)
	if err != nil {
		return nil, err
	}

	// load response from repo & save to cache
	if !exist {
		respDom, err := s.repo.DisplayCurrentDataByCountry(&reqDom, i)
		if err != nil {
			return nil, err
		}

		respDTO = &appDto.DisplayCurrentDataByCountryResDTO{
			Query: respDom.Query,
			Data:  respDom.Data,
		}

		// save to cache
		if err := s._saveToCache(&reqDom, respDTO); err != nil {
			return nil, err
		}
	}

	return respDTO, nil
}

func (s *CurrentSvc) _loadFromCache(req *domSchema.DisplayCurrentDataByCountryRequest) (bool, *appDto.DisplayCurrentDataByCountryResDTO, error) {
	// s.ce.Component = "DisplayCurrentDataByCountry"

	searchKey := utils.MD5(req.ToJSON())
	if !s.ce.IsExist(searchKey) {
		return false, nil, nil
	}

	val := s.ce.Get(searchKey)
	if val == nil {
		return false, nil, nil
	}

	var resp appDto.DisplayCurrentDataByCountryResDTO
	if err := json.Unmarshal([]byte(val.(string)), &resp); err != nil {
		return false, nil, err
	}

	return true, &resp, nil
}

func (s *CurrentSvc) _saveToCache(req *domSchema.DisplayCurrentDataByCountryRequest, resp *appDto.DisplayCurrentDataByCountryResDTO) error {
	// s.ce.Component = "DisplayCurrentDataByCountry"

	searchKey := utils.MD5(req.ToJSON())

	// 10 * 60 (second) = 10 minute
	if err := s.ce.Put(searchKey, resp.ToJSON(), 10*60); err != nil {
		return err
	}

	return nil
}
