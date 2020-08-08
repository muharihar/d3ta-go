package feature

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewOpenAPI new FOpenAPI
func NewOpenAPI(h *handler.Handler) (*FOpenAPI, error) {

	f := new(FOpenAPI)
	f.handler = h

	return f, nil
}

// FOpenAPI represent FOpenAPI
type FOpenAPI struct {
	BaseFeature
}

// SwaggerUI generate SwaggerUI html page
func (f *FOpenAPI) SwaggerUI(c echo.Context) error {

	cfg, err := f.handler.GetConfig()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"htmlTitle": cfg.Applications.Servers.RestAPI.Name,
	}

	return c.Render(http.StatusOK, "openapis/swagger-ui", data)
}

// GenOpenAPI generate openapi definition
func (f *FOpenAPI) GenOpenAPI(c echo.Context) error {
	cfg, err := f.handler.GetConfig()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"info.Title":              cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Title,
		"info.Description":        cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Description,
		"info.Contact.Name":       cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Contact.Name,
		"info.Contact.URL":        cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Contact.URL,
		"info.Contact.Email":      cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Contact.Email,
		"info.License.Name":       cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.License.Name,
		"info.License.URL":        cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.License.URL,
		"info.Version":            cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Version,
		"server.URL.Host.Default": c.Request().Host,
	}

	// return c.Blob(http.StatusOK, "text/plain; charset=utf-8", data)
	return c.Render(http.StatusOK, "openapi.yaml", data)
}
