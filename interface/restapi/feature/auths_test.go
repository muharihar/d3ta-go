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

func TestAuths_RegisterUser(t *testing.T) {
	// variables
	reqDTO := `{
		"username" : "admin.d3tago", 
		"password" : "P4s$W0rd!@!",
		"email" : "admin.d3tago@email.tld",
		"nickName" : "Hari",
		"captcha": "just-capthcha-value",
		"captchaID": "just-chaptcha-id"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auths/register", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		panic(err)
	}

	auths, err := NewFAuths(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, auths.RegisterUser(c)) {
		// assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.auths.RegisterUser: %s", res.Body.String())
	}
}

func TestAuths_ActivateRegistration(t *testing.T) {
	// variables
	// via url path

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/auths/registration/activate/:activationCode", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("activationCode")
	c.SetParamValues("a70112cc-bca6-45c2-9bb6-cf3a56daf566")

	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		panic(err)
	}

	auths, err := NewFAuths(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, auths.ActivateRegistration(c)) {
		// assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.auths.ActivateRegistration: %s", res.Body.String())
	}
}

func TestAuths_Login(t *testing.T) {
	// variables

	reqDTO := `{
		"username" : "admin.d3tago", 
		"password" : "P4s$W0rd!@!",
		"captcha": "just-capthcha-value",
		"captchaID": "just-chaptcha-id"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auths/login", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		panic(err)
	}

	auths, err := NewFAuths(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, auths.Login(c)) {
		// assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.auths.Login: %s", res.Body.String())
	}
}

func TestAuths_LoginApp(t *testing.T) {
	// variables

	reqDTO := `{
		"clientKey" : "53102ba5-b6b2-47ad-a68d-682463a8be29", 
		"secretKey" : "OTk5ZDlmYjJlZGUyMjAxNTZkZThiNmNkMmJmNDI1NjdiNTYzMzcxNDEwNzNiNDBjM2NhZmIxOWY3NzZmYzhmNg=="
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auths/login-app", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		panic(err)
	}

	auths, err := NewFAuths(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, auths.LoginApp(c)) {
		// assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.auths.LoginApp: %s", res.Body.String())
	}
}
