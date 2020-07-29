package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/feature"
)

// SetOpenAPI set OpenAPI Router
func SetOpenAPI(eg *echo.Group, f *feature.FOpenAPI) {
	eg.GET("/docs/openapi.yaml", f.GenOpenAPI)
	eg.GET("/docs/index.html", f.SwaggerUI)
	eg.Static("/docs/swagger-ui/assets", "./www/public/swagger-ui/assets")
}
