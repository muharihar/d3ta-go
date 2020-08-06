package service

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../../conf")
	if err != nil {
		return nil, err
	}
	c.IAM.Casbin.ModelPath = "../../../../../conf/casbin/casbin_rbac_rest_model.conf"
	return c, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	cfg, err := h.GetConfig()
	if err != nil {
		t.Errorf("GetConfig: %s", err.Error())
	}
	if err := i.SetCasbinEnforcer(cfg.IAM.Casbin.ModelPath); err != nil {
		t.Errorf("SetCasbinEnforcer: %s", err.Error())
	}

	i.Claims.Username = "test.d3tago"
	i.Claims.AuthorityID = "group:admin"
	i.RequestInfo.Host = "127.0.0.1:2020"

	return i
}
