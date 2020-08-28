package feature

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/system/initialize"
	"github.com/stretchr/testify/assert"
)

func TestFGeoLocation_ListAllCountry(t *testing.T) {

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/geolocation/countries/list-all", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := generateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.ListAllCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.geoLoc.ListAllCountry: %s", res.Body.String())
	}
}

func TestFGeoLocation_RefreshCountryIndexer(t *testing.T) {
	// variables
	reqDTO := `{
		"processType":"SYNC"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/geolocation/countries/indexer/refresh", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(handler); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := generateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.RefreshCountryIndexer(c)) {
		assert.Equal(t, http.StatusCreated, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.geoLoc.RefreshCountryIndexer: %s", res.Body.String())
	}
}

func TestFGeoLocation_SearchCountryIndexer(t *testing.T) {
	// variables
	reqDTO := `{
		"name":"IND"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/geolocation/countries/indexer/search", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(handler); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := generateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.SearchCountryIndexer(c)) {
		assert.Equal(t, http.StatusCreated, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.geoLoc.SearchCountryIndexer: %s", res.Body.String())
	}
}

func TestFGeoLocation_AddCountry(t *testing.T) {
	// variables
	reqDTO := `{
		"code":"XX", 
		"name": "XX COUNTRY", 
		"ISO2Code": "XX", 
		"ISO3Code": "", 
		"WHORegion": "WPRO"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/geolocation/country", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := generateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.AddCountry(c)) {
		assert.Equal(t, http.StatusCreated, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.geoLoc.AddCountry: %s", res.Body.String())
	}
}

func TestFGeoLocation_GetCountry(t *testing.T) {

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/geolocation/country/:code", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	// c.SetPath("/api/v1/geolocation/country/:code")
	c.SetParamNames("code")
	c.SetParamValues("XX")

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := generateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature

	geoLoc, err := NewFGeoLocation(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.GetCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.geoLoc.GetCountry: %s", res.Body.String())
	}
}

func TestFGeoLocation_UpdateCountry(t *testing.T) {
	// variables
	reqDTO := `{
		"name": "XX COUNTRY UPDATED",
		"ISO2Code": "XX",
		"ISO3Code": "",
		"WHORegion": "WPRO"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/api/v1/geolocation/country/:code", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	// c.SetPath("/api/v1/geolocation/country/:code")
	c.SetParamNames("code")
	c.SetParamValues("XX")

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := generateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.UpdateCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.geoLoc.UpdateCountry: %s", res.Body.String())
	}
}

func TestFGeoLocation_DeleteCountry(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/geolocation/country/:code", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	// c.SetPath("/api/v1/geolocation/country/:code")
	c.SetParamNames("code")
	c.SetParamValues("XX")

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := generateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.DeleteCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.geoLoc.DeleteCountry: %s", res.Body.String())
	}
}
