package feature

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TestSwaggerUI(t *testing.T) {
	// variables

	// Setup
	e := echo.New()

	// html template
	tpl := &Template{
		templates: template.Must(template.ParseGlob("../../../www/templates/**/*.*ml")),
	}
	e.Renderer = tpl

	req := httptest.NewRequest(http.MethodGet, "/openapis/docs/index.html", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextHTML)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	handler := newHandler()
	openapis, _ := NewOpenAPI(handler)

	// Assertions
	if assert.NoError(t, openapis.SwaggerUI(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// t.Logf("RESPONSE.openapis.SwaggerUI: %s", res.Body.String())
	}
}

func TestGenOpenAPI(t *testing.T) {
	// variables

	// Setup
	e := echo.New()

	// html template
	tpl := &Template{
		templates: template.Must(template.ParseGlob("../../../www/templates/**/*.*ml")),
	}
	e.Renderer = tpl

	req := httptest.NewRequest(http.MethodGet, "/openapis/docs/openapi.yaml", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextHTML)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	handler := newHandler()
	openapis, _ := NewOpenAPI(handler)

	// Assertions
	if assert.NoError(t, openapis.GenOpenAPI(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// t.Logf("RESPONSE.openapis.GenOpenAPI: %s", res.Body.String())
	}
}
