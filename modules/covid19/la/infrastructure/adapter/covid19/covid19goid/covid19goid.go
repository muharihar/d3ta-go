package who

import (
	"fmt"
	"strings"

	conC19goid "github.com/muharihar/d3ta-go/connector/covid19/covid19goid"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
	infC19Adp "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/adapter/covid19"
	mapper "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/adapter/covid19/covid19goid/mapping"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewCovid19goidAdapter New Covid19goidAdapter
func NewCovid19goidAdapter(h *handler.Handler) (infC19Adp.ICovid19Adapter, *infC19Adp.Covid19AdapterInfo, error) {
	var err error

	adp := new(Covid19goidAdapter)
	adp.handler = h

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, nil, err
	}

	info := infC19Adp.Covid19AdapterInfo{
		Code:   cfg.Connectors.Covid19.Covid19goid.Code,
		Name:   cfg.Connectors.Covid19.Covid19goid.Name,
		Server: cfg.Connectors.Covid19.Covid19goid.Server,
		Enable: cfg.Connectors.Covid19.Covid19goid.Enable,
	}
	adp.SetInfo(info)

	config := conC19goid.Config{
		Code:   cfg.Connectors.Covid19.Covid19goid.Code,
		Name:   cfg.Connectors.Covid19.Covid19goid.Name,
		Server: cfg.Connectors.Covid19.Covid19goid.Server,
		Enable: cfg.Connectors.Covid19.Covid19goid.Enable,
	}

	adp.connector = conC19goid.NewCovid19(config, nil)

	return adp, &info, err
}

// Covid19goidAdapter represent Covid19goidAdapter Adapter
type Covid19goidAdapter struct {
	infC19Adp.BaseCovid19Adapter

	handler   *handler.Handler
	connector *conC19goid.Covid19
}

// DisplayCurrentDataByCountry display CurrentDataByCountry
func (a *Covid19goidAdapter) DisplayCurrentDataByCountry(req *domSchema.DisplayCurrentDataByCountryRequest) (*domSchema.TotalCountryProviderData, error) {
	resp := new(domSchema.TotalCountryProviderData)
	if a.GetInfo().Enable && strings.ToUpper(req.CountryCode) == "ID" {
		connResp, err := a.connector.GetUpdate()
		if err != nil {
			return nil, err
		}

		resp, err = mapper.MapDisplayCurrentDataByCountryRes(connResp)
		if err != nil {
			return nil, err
		}
	}
	resp.Provider = a.GetInfo().Code
	resp.Information = fmt.Sprintf("[Enable: %v] %s [%s] Only For 'ID'", a.GetInfo().Enable, a.GetInfo().Name, a.GetInfo().Server)

	return resp, nil
}
