package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/feature"
	internalMiddleware "github.com/muharihar/d3ta-go/interface/restapi/middleware"
)

// SetGeoLocation set GeoLocation Router
func SetGeoLocation(eg *echo.Group, f *feature.FGeoLocation) {

	gc := eg.Group("/geolocation")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.GET("/country/list-all", f.ListAllCountry)
	gc.GET("/country/:code", f.GetCountry)
	gc.POST("/country", f.AddCountry)
	gc.PUT("/country/:code", f.UpdateCountry)
	gc.DELETE("/country/:code", f.DeleteCountry)
}
