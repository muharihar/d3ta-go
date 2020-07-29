package service

import (
	"fmt"
	"net/http"
	"strings"

	domModel "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/model"
	domSchema "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/schema/country"
	sysErr "github.com/muharihar/d3ta-go/system/error"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

// NewCountrySvc new CountrySvc
func NewCountrySvc(h *handler.Handler) (*CountrySvc, error) {
	svc := new(CountrySvc)
	svc.SetHandler(h)

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	svc.SetDBConnectionName(cfg.Databases.MainDB.ConnectionName)

	return svc, nil
}

// CountrySvc represent CountrySvc
type CountrySvc struct {
	BaseSvc
}

// ListAll list All Country
func (s *CountrySvc) ListAll(i identity.Identity) (*domSchema.ListCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// find
	var countryList []domModel.Country
	if err := dbCon.Order("code asc").Find(&countryList).Error; err != nil {
		return nil, err
	}

	// response
	var listCountry []*domSchema.Country
	for _, rec := range countryList {
		tmp := new(domSchema.Country)

		tmp.ID = rec.ID
		tmp.Code = rec.Code
		tmp.Name = rec.Name
		tmp.ISO2Code = rec.ISO2Code
		tmp.ISO3Code = rec.ISO3Code
		tmp.WHORegion = rec.WHORegion

		listCountry = append(listCountry, tmp)
	}

	res := new(domSchema.ListCountryResponse)
	res.Data = listCountry

	return res, nil
}

// GetDetail get Detail Country by Code
func (s *CountrySvc) GetDetail(req *domSchema.GetDetailCountryRequest, i identity.Identity) (*domSchema.GetDetailCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// find
	var countryMdl domModel.Country
	if err := dbCon.Where("code = ?", req.Code).First(&countryMdl).Error; err != nil || countryMdl.Code == "" {
		return nil, &sysErr.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Data not found (for Code=%s)", req.Code)}
	}

	// response
	country := &domSchema.Country{
		ID:        countryMdl.ID,
		Code:      countryMdl.Code,
		Name:      countryMdl.Name,
		ISO2Code:  countryMdl.ISO2Code,
		ISO3Code:  countryMdl.ISO3Code,
		WHORegion: countryMdl.WHORegion,
	}

	res := new(domSchema.GetDetailCountryResponse)
	res.Query = req
	res.Data = country

	return res, nil
}

// Add Country
func (s *CountrySvc) Add(req *domSchema.AddCountryRequest, i identity.Identity) (*domSchema.AddCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// save
	countryMdl := &domModel.Country{
		Code:      req.Code,
		Name:      req.Name,
		ISO2Code:  req.ISO2Code,
		ISO3Code:  req.ISO3Code,
		WHORegion: req.WHORegion,
	}
	countryMdl.CreatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Create(countryMdl).Error; err != nil {
		if strings.Index(err.Error(), "Error 1062: Duplicate entry") > -1 {
			return nil, &sysErr.SystemError{StatusCode: http.StatusConflict, Err: err}
		}
		return nil, err
	}

	// response
	country := &domSchema.Country{
		ID:        countryMdl.ID,
		Code:      countryMdl.Code,
		Name:      countryMdl.Name,
		ISO2Code:  countryMdl.ISO2Code,
		ISO3Code:  countryMdl.ISO3Code,
		WHORegion: countryMdl.WHORegion,
	}

	res := new(domSchema.AddCountryResponse)
	res.Query = req
	res.Data = country

	return res, nil
}

// Update Country
func (s *CountrySvc) Update(req *domSchema.UpdateCountryRequest, i identity.Identity) (*domSchema.UpdateCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// find
	var countryMdl domModel.Country
	if err := dbCon.Where("code = ?", req.Keys.Code).First(&countryMdl).Error; err != nil || countryMdl.Code == "" {
		return nil, &sysErr.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Data not found (for Code=%s)", req.Keys.Code)}
	}

	// update
	countryMdl.Name = req.Data.Name
	countryMdl.ISO2Code = req.Data.ISO2Code
	countryMdl.ISO3Code = req.Data.ISO3Code
	countryMdl.WHORegion = req.Data.WHORegion

	countryMdl.UpdatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Save(&countryMdl).Error; err != nil {
		if strings.Index(err.Error(), "Error 1062: Duplicate entry") > -1 {
			return nil, &sysErr.SystemError{StatusCode: http.StatusConflict, Err: err}
		}
		return nil, err
	}

	// response
	country := &domSchema.Country{
		ID:        countryMdl.ID,
		Code:      countryMdl.Code,
		Name:      countryMdl.Name,
		ISO2Code:  countryMdl.ISO2Code,
		ISO3Code:  countryMdl.ISO3Code,
		WHORegion: countryMdl.WHORegion,
	}

	res := new(domSchema.UpdateCountryResponse)
	res.Query = req
	res.Data = country

	return res, nil
}

// Delete delete Country
func (s *CountrySvc) Delete(req *domSchema.DeleteCountryRequest, i identity.Identity) (*domSchema.DeleteCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// find
	var countryMdl domModel.Country
	if err := dbCon.Where("code = ?", req.Code).First(&countryMdl).Error; err != nil || countryMdl.Code == "" {
		return nil, &sysErr.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Data not found (for Code=%s)", req.Code)}
	}

	// update deleted by user
	countryMdl.DeletedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Save(&countryMdl).Error; err != nil {
		return nil, err
	}

	// delete data
	if err := dbCon.Delete(&countryMdl).Error; err != nil {
		return nil, err
	}

	// response
	country := &domSchema.Country{
		ID:        countryMdl.ID,
		Code:      countryMdl.Code,
		Name:      countryMdl.Name,
		ISO2Code:  countryMdl.ISO2Code,
		ISO3Code:  countryMdl.ISO3Code,
		WHORegion: countryMdl.WHORegion,
	}

	res := new(domSchema.DeleteCountryResponse)
	res.Query = req
	res.Data = country

	return res, nil
}
