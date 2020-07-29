package repository

import (
	"testing"

	"github.com/muharihar/d3ta-go/modules/auths/la/domain/repository"
	"github.com/muharihar/d3ta-go/modules/auths/la/domain/schema"
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

func newRepo(t *testing.T) (repository.IAuthenticationRepo, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabase(h); err != nil {
		return nil, nil, err
	}

	r, err := NewAuthenticationRepo(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	i.Claims.Username = "test.d3tago"
	i.RequestInfo.Host = "127.0.0.1:2020"

	return i
}

func TestAuthRepo_Register(t *testing.T) {
	repo, h, err := newRepo(t)
	if err != nil {
		t.Errorf("Error.newRepo: %s", err.Error())
	}

	req := &schema.RegisterRequest{
		Username:  "admin.d3tago",
		Password:  "P4s$W0rd!@!",
		Email:     "admin.d3tago@email.tld",
		NickName:  "Hari",
		Captcha:   "just-capthcha-value", // validation on interface
		CaptchaID: "just-chaptcha-id",    // validation on interface
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Rew.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Register(req, i)
	if err != nil {
		t.Errorf("Error.AuthRepo.Register: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.AuthRepo.Register: %s", string(respJSON))
	}
}

func TestAuthRepo_ActivateRegistration(t *testing.T) {
	repo, h, err := newRepo(t)
	if err != nil {
		t.Errorf("Error.newRepo: %s", err.Error())
	}

	req := &schema.ActivateRegistrationRequest{
		ActivationCode: "a70112cc-bca6-45c2-9bb6-cf3a56daf566",
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Rew.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.ActivateRegistration(req, i)
	if err != nil {
		t.Errorf("Error.AuthRepo.ActivateRegistration: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.AuthRepo.ActivateRegistration: %s", string(respJSON))
	}
}

func TestAuthRepo_Login(t *testing.T) {
	repo, h, err := newRepo(t)
	if err != nil {
		t.Errorf("Error.newRepo: %s", err.Error())
	}

	req := &schema.LoginRequest{
		Username:  "admin.d3tago",
		Password:  "P4s$W0rd!@!",
		Captcha:   "just-capthcha-value", // validation on interface
		CaptchaID: "just-chaptcha-id",    // validation on interface
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Rew.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Login(req, i)
	if err != nil {
		t.Errorf("Error.AuthRepo.Login: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.AuthRepo.Login: %s", string(respJSON))
	}
}

func TestAuthRepo_LoginApp(t *testing.T) {
	repo, h, err := newRepo(t)
	if err != nil {
		t.Errorf("Error.newRepo: %s", err.Error())
	}

	req := &schema.LoginAppRequest{
		ClientKey: "53102ba5-b6b2-47ad-a68d-682463a8be29",
		SecretKey: "OTk5ZDlmYjJlZGUyMjAxNTZkZThiNmNkMmJmNDI1NjdiNTYzMzcxNDEwNzNiNDBjM2NhZmIxOWY3NzZmYzhmNg==",
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.LoginApp(req, i)
	if err != nil {
		t.Errorf("Error.AuthRepo.LoginApp: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.AuthRepo.LoginApp: %s", string(respJSON))
	}
}
