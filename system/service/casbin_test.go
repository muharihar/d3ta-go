package service

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/initialize"
)

func newCasbinEnforcer(t *testing.T) (*casbin.Enforcer, error) {
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

	enf, err := NewCasbinEnforcer(h, "../../conf/casbin/casbin_rbac_rest_model.conf")

	return enf, err
}

func TestCasbin_NewCasbinEnforcer(t *testing.T) {
	enf, err := newCasbinEnforcer(t)
	if err != nil {
		t.Errorf("newCasbinEnforcer: %s", err.Error())
	}

	ng := enf.GetNamedGroupingPolicy("g")
	t.Logf("named.groups.all: %#v", ng)

	ngf := enf.GetFilteredNamedGroupingPolicy("g", 0, "group:admin")
	t.Logf("named.groups.Filtered: %#v", ngf)

	r, err := enf.GetRolesForUser("group:admin")
	if err != nil {
		t.Errorf("GetRolesForUser: %s", err.Error())
	}
	t.Logf("user.roles: %#v", r)

	ir, err := enf.GetImplicitRolesForUser("group:admin")
	if err != nil {
		t.Errorf("GetImplicitRolesForUser: %s", err.Error())
	}
	t.Logf("user.roles.implisit: %#v", ir)

	pu := enf.GetPermissionsForUser("group:admin")
	t.Logf("user.permissions: %#v", pu)

	ipu, err := enf.GetImplicitPermissionsForUser("group:admin")
	if err != nil {
		t.Errorf("GetImplicitPermissionsForUser: %s", err.Error())
	}
	t.Logf("user.permissions.implisit: %#v", ipu)

	can := false
	for _, role := range ir {
		t.Log("role:", role)
		sub := role                                 // "group:admin"
		obj := "/api/v1/covid19/current/by-country" //"/api/v1/geolocation/country"
		act := "GET"                                // "POST"
		can, err = enf.Enforce(sub, obj, act)
		if err != nil {
			t.Errorf("Enforce: %s", err.Error())
		}

		t.Logf("can: %v", can)
	}

	p := enf.GetPolicy()
	t.Logf("GetPolicy: %#v", p)

	t.Errorf("%v", can)
}
