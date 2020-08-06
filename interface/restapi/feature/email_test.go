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

func TestEmail_ListAllEmailTemplate(t *testing.T) {
	// client request
	// none

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/email/templates/list-all", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabase: %s", err.Error())
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
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.ListAllEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("RESPONSE.Email.ListAllEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_FindEmailTemplateByCode(t *testing.T) {
	// client request
	// --> set on context param [http method = GET]

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/email/template/:code", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("code")
	c.SetParamValues("test.code")

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabase: %s", err.Error())
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
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.FindEmailTemplateByCode(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("RESPONSE.Email.FindEmailTemplateByCode: %s", res.Body.String())
	}
}

func TestEmail_CreateEmailTemplate(t *testing.T) {
	// client request
	reqDTO := `{
	"code": "test.code.04",
	"name": "Template Name 04",
	"isActive": true,
	"emailFormat": "TEXT",
	"template": {
		"subjectTpl": "Subject Template 04",
		"bodyTpl": "{{define \"T\"}}Body Template 04{{end}}"
	}
}`

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/email/template", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabase: %s", err.Error())
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
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.CreateEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("RESPONSE.Email.CreateEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_UpdateEmailTemplate(t *testing.T) {
	// client request
	reqDTO := `{
	"name": "Template Name 04 Updated",
	"isActive": true,
	"emailFormat": "TEXT",
	"template": {
		"subjectTpl": "Subject Template 04 Updated",
		"bodyTpl": "{{define \"T\"}}Body Template 04 Updated{{end}}"
	}
}`

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/api/v1/email/template/update/:code", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("code")
	c.SetParamValues("test.code.04")

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabase: %s", err.Error())
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
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.UpdateEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("RESPONSE.Email.UpdateEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_SetActiveEmailTemplate(t *testing.T) {
	// client request
	reqDTO := `{
	"isActive": true
}`

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/api/v1/email/template/set-active/:code", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("code")
	c.SetParamValues("test.code.04")

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabase: %s", err.Error())
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
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.SetActiveEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("RESPONSE.Email.SetActiveEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_DeleteEmailTemplate(t *testing.T) {
	// client request
	// none

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/email/template/:code", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("code")
	c.SetParamValues("test.code.04")

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabase: %s", err.Error())
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
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.DeleteEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("RESPONSE.Email.DeleteEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_SendEmail(t *testing.T) {
	// client request
	reqDTO := `{
    "templateCode": "activate-registration-html",
    "from": { "email": "d3tago.from@domain.tld", "name": "D3TA Golang" },
    "to": { "email": "d3tago.test@outlook.com", "name": "D3TAgo Test (Outlook)" },
    "cc": [
        { "email": "d3tago.test@protonmail.com", "name": "D3TAgo Test CC 1 (Protonmail)" },
        { "email": "d3tago.test.cc@tutanota.com", "name": "D3TAgo Test CC 2 (Tutanota)" }
    ],
    "bcc": [
        { "email": "d3tago.test@tutanota.com", "name": "D3TAgo Test BCC 1 (Tutanota)" },
		{ "email": "d3tago.test.bcc@outlook.com", "name": "D3TAgo Test BCC 2 (Outlook)" }
    ],
    "templateData": {
		"Header.Name": "John Doe",
		"Body.UserAccount": "john.doe",
		"Body.ActivationURL": "https://google.com",
        "Footer.Name": "Customer Service"
	},
	"processingType": "SYNC"
}`

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/email/send", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := newHandler()
	if err := initialize.LoadAllDatabase(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabase: %s", err.Error())
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
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.SendEmail(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("RESPONSE.Email.SendEmail: %s", res.Body.String())
	}
}
