package feature

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// variables

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/system/health", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	handler := newHandler()
	system, _ := NewSystem(handler)

	// Assertions
	if assert.NoError(t, system.HealthCheck(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// t.Logf("RESPONSE.system.HealthCheck: %s", res.Body.String())
	}
}
