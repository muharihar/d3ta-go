package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/feature"
	internalMiddleware "github.com/muharihar/d3ta-go/interface/restapi/middleware"
)

// SetCovid19 set Covid19 Router
func SetCovid19(eg *echo.Group, f *feature.FCovid19) {

	gc := eg.Group("/covid19")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("/current/by-country", f.DisplayCurrentDataByCountry)
}
