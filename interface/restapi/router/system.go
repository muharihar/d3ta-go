package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/feature"
)

// SetSystem set FSystem Router
func SetSystem(eg *echo.Group, f *feature.FSystem) {
	eg.GET("/health", f.HealthCheck)
	eg.GET("/captcha/generate", f.GenerateCaptcha)
	eg.GET("/captcha/image/:captchaID", f.GenerateCaptchaImage)
}
