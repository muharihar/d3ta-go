package restapi

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	feature "github.com/muharihar/d3ta-go/interface/restapi/feature"
	router "github.com/muharihar/d3ta-go/interface/restapi/router"
	"github.com/muharihar/d3ta-go/system/handler"
)

// Template html Template
type Template struct {
	templates *template.Template
}

// Render html template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// SetRouters is a function to ser Echo Routers
func SetRouters(e *echo.Echo, h *handler.Handler) {
	cfg, err := h.GetConfig()

	// set default middleware
	e.Pre(middleware.RemoveTrailingSlash())
	if cfg.Applications.Servers.RestAPI.Options.Middlewares.Logger.Enable {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	// html template
	t := &Template{
		templates: template.Must(template.ParseGlob("www/templates/**/*.*ml")),
	}
	e.Renderer = t

	// Set CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	// features
	features, err := feature.NewFeature(h)
	if err != nil {
		panic(err)
	}

	// System
	gs := e.Group("system")
	router.SetSystem(gs, features.System)

	// OpenApi/swagger-ui
	if cfg.Applications.Servers.RestAPI.Options.DisplayOpenAPI {
		gd := e.Group("openapi")
		router.SetOpenAPI(gd, features.OpenAPI)
	}

	// Group API
	ga := e.Group("api/v1")
	// ga.Use(internalMiddleware.JWTVerifier(h))
	router.SetAuths(ga, features.Auths)
	router.SetCovid19(ga, features.Covid19)
	router.SetGeoLocation(ga, features.GeoLocation)
	router.SetEmail(ga, features.Email)
}
