package who

import (
	"fmt"

	conC19Who "github.com/muharihar/d3ta-go/connector/covid19/who"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
	infC19Adp "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/adapter/covid19"
	mapper "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/adapter/covid19/who/mapping"

	"github.com/muharihar/d3ta-go/system/handler"
)

// NewCovid19WHOAdapter New Covid19WHOAdapter
func NewCovid19WHOAdapter(h *handler.Handler) (infC19Adp.ICovid19Adapter, *infC19Adp.Covid19AdapterInfo, error) {
	var err error

	adp := new(Covid19WHOAdapter)
	adp.handler = h

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, nil, err
	}

	info := infC19Adp.Covid19AdapterInfo{
		Code:   cfg.Connectors.Covid19.Covid19WHO.Code,
		Name:   cfg.Connectors.Covid19.Covid19WHO.Name,
		Server: cfg.Connectors.Covid19.Covid19WHO.Server,
		Enable: cfg.Connectors.Covid19.Covid19WHO.Enable,
	}
	adp.SetInfo(info)

	config := conC19Who.Config{
		Code:   cfg.Connectors.Covid19.Covid19WHO.Code,
		Name:   cfg.Connectors.Covid19.Covid19WHO.Name,
		Server: cfg.Connectors.Covid19.Covid19WHO.Server,
		Enable: cfg.Connectors.Covid19.Covid19WHO.Enable,
	}

	adp.connector = conC19Who.NewCovid19(config, nil)

	return adp, &info, err
}

// Covid19WHOAdapter represent Covid19WHOAdapter Adapter
type Covid19WHOAdapter struct {
	infC19Adp.BaseCovid19Adapter

	handler   *handler.Handler
	connector *conC19Who.Covid19
}

// DisplayCurrentDataByCountry display CurrentDataByCountry
func (a *Covid19WHOAdapter) DisplayCurrentDataByCountry(req *domSchema.DisplayCurrentDataByCountryRequest) (*domSchema.TotalCountryProviderData, error) {
	resp := new(domSchema.TotalCountryProviderData)
	if a.GetInfo().Enable {
		connReq, err := mapper.MapDisplayCurrentDataByCountryReq(req)
		if err != nil {
			return nil, err
		}

		connResp, err := a.connector.GetCountry(connReq)
		if err != nil {
			return nil, err
		}

		resp, err = mapper.MapDisplayCurrentDataByCountryRes(connResp)
		if err != nil {
			return nil, err
		}
	}
	resp.Provider = a.GetInfo().Code
	resp.Information = fmt.Sprintf("[Enable: %v] %s [%s]", a.GetInfo().Enable, a.GetInfo().Name, a.GetInfo().Server)

	return resp, nil
}
