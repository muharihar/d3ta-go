package feature

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFCovid19_DisplayCurrentDataByCountry(t *testing.T) {
	// variables
	reqDTO := `{"countryCode":"ID", "providers": [{"code":"_ALL_"}]}`
	// resDTO := `{"status":"OK","response":{"message":"Operation succeeded","data":{"status":"OK","data":null}},"serverInfo":{"serverTime":"2020-07-18T11:26:35.377625+07:00"}}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/covid19/current/by-country", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	handler := newHandler()
	covid19, err := NewFCovid19(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, covid19.DisplayCurrentDataByCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.covid19.DisplayCurrentDataByCountry: %s", res.Body.String())
	}
}
