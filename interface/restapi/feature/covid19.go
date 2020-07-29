package feature

import (
	"net/http"

	"github.com/labstack/echo/v4"
	response "github.com/muharihar/d3ta-go/interface/restapi/response"
	appCovid19 "github.com/muharihar/d3ta-go/modules/covid19/la/application"
	appCovid19DTO "github.com/muharihar/d3ta-go/modules/covid19/la/application/dto"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewFCovid19 new FCovid19
func NewFCovid19(h *handler.Handler) (*FCovid19, error) {
	var err error

	f := new(FCovid19)
	f.handler = h

	if f.appCovid19, err = appCovid19.NewCovid19App(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FCovid19 represent FCovid19
type FCovid19 struct {
	BaseFeature
	appCovid19 *appCovid19.Covid19App
}

// DisplayCurrentDataByCountry display CurrentDataByCountry
func (f *FCovid19) DisplayCurrentDataByCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appCovid19DTO.DisplayCurrentDataByCountryReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appCovid19.CurrentSvc.DisplayCurrentDataByCountry(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}
