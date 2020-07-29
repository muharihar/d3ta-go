package feature

import (
	"github.com/labstack/echo/v4"
	captcha "github.com/muharihar/d3ta-go/interface/restapi/feature/captcha"
	response "github.com/muharihar/d3ta-go/interface/restapi/response"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewSystem new FSystem
func NewSystem(h *handler.Handler) (*FSystem, error) {

	f := new(FSystem)
	f.handler = h

	return f, nil
}

// FSystem represent FSystem
type FSystem struct {
	BaseFeature
}

// HealthCheck display system health check
func (f *FSystem) HealthCheck(c echo.Context) error {
	data := map[string]interface{}{"serviceStatus": "OK"}
	return response.OkWithData(data, c)
}

// GenerateCaptcha generate Captcha
func (f *FSystem) GenerateCaptcha(c echo.Context) error {
	cfg, err := f.handler.GetConfig()
	if err != nil {
		return err
	}

	resp := captcha.GenerateCaptchaID(cfg, c)

	return response.OkWithData(resp, c)
}

// GenerateCaptchaImage generate CaptchaImage
func (f *FSystem) GenerateCaptchaImage(c echo.Context) error {
	cfg, err := f.handler.GetConfig()
	if err != nil {
		return err
	}

	return captcha.CaptchaServeHTTP(cfg, c)
}
