package service

import (
	"encoding/json"
	"testing"

	"github.com/muharihar/d3ta-go/modules/auths/la/application/dto"
	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
	"github.com/muharihar/d3ta-go/system/initialize"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newAuthenticationSvc(t *testing.T) (*AuthenticationSvc, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabase(h); err != nil {
		return nil, err
	}

	r, err := NewAuthenticationSvc(h)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func TestAuthenticationService_Register(t *testing.T) {
	svc, err := newAuthenticationSvc(t)
	if err != nil {
		t.Errorf("newAuthenticationSvc: %s", err.Error())
		return
	}

	req := dto.RegisterReqDTO{}
	req.Username = "admin.d3tago"
	req.Password = "P4s$W0rd!@!"
	req.Email = "admin.d3tago@email.tld"
	req.NickName = "Hari"
	req.Captcha = "just-capthcha-value" // validation on interface
	req.CaptchaID = "just-chaptcha-id"  // validation on interface

	resp, err := svc.Register(&req, identity.Identity{})
	if err != nil {
		t.Errorf("Register: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}

func TestAuthenticationService_ActivateRegistration(t *testing.T) {
	svc, err := newAuthenticationSvc(t)
	if err != nil {
		t.Errorf("newAuthenticationSvc: %s", err.Error())
		return
	}

	req := dto.ActivateRegistrationReqDTO{}
	req.ActivationCode = "a70112cc-bca6-45c2-9bb6-cf3a56daf566"

	resp, err := svc.ActivateRegistration(&req, identity.Identity{})
	if err != nil {
		t.Errorf("ActivateRegistration: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}

func TestAuthenticationService_Login(t *testing.T) {
	svc, err := newAuthenticationSvc(t)
	if err != nil {
		t.Errorf("newAuthenticationSvc: %s", err.Error())
		return
	}

	req := dto.LoginReqDTO{}
	req.Username = "admin.d3tago"
	req.Password = "P4s$W0rd!@!"
	req.Captcha = "just-capthcha-value" // validation on interface
	req.CaptchaID = "just-chaptcha-id"  // validation on interface

	resp, err := svc.Login(&req, identity.Identity{})
	if err != nil {
		t.Errorf("Login: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}

func TestAuthenticationService_LoginApp(t *testing.T) {
	svc, err := newAuthenticationSvc(t)
	if err != nil {
		t.Errorf("newAuthenticationSvc: %s", err.Error())
		return
	}

	req := dto.LoginAppReqDTO{}
	req.ClientKey = "53102ba5-b6b2-47ad-a68d-682463a8be29"
	req.SecretKey = "OTk5ZDlmYjJlZGUyMjAxNTZkZThiNmNkMmJmNDI1NjdiNTYzMzcxNDEwNzNiNDBjM2NhZmIxOWY3NzZmYzhmNg=="

	resp, err := svc.LoginApp(&req, identity.Identity{})
	if err != nil {
		t.Errorf("LoginApp: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}
