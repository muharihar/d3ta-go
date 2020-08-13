package service

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/initialize"
	"github.com/stretchr/testify/assert"
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
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, err
	}

	enf, err := NewCasbinEnforcer(h, "../../conf/casbin/casbin_rbac_rest_model.conf")

	return enf, err
}

func TestCasbin_NewCasbinEnforcer(t *testing.T) {
	enf, err := newCasbinEnforcer(t)
	if assert.NoError(t, err, "Error while creating Casbin Enforcer: newCasbinEnforcer") {
		// use case: user -> group -> role -> permission

		// get named grouping policy
		ng := enf.GetNamedGroupingPolicy("g")
		assert.NotEmpty(t, ng)
		t.Logf("named.groups.policy.all: %#v", ng)

		// get filtered named grouping policy
		ngf := enf.GetFilteredNamedGroupingPolicy("g", 0, "group:admin")
		assert.NotEmpty(t, ngf)
		t.Logf("named.groups.policy.Filtered: %#v", ngf)

		// get roles for user (group of users)
		r, err := enf.GetRolesForUser("group:admin")
		if assert.NoError(t, err, "Error while getting roles for user/group of user: GetRolesForUser") {
			assert.NotEmpty(t, r)
			t.Logf("user/group.roles: %#v", r)
		}

		// get implicit roles for user (group of user)
		ir, err := enf.GetImplicitRolesForUser("group:admin")
		if assert.NoError(t, err, "Error whilw getting implicit roles for user/group of users: GetImplicitRolesForUser") {
			assert.NotEmpty(t, ir)
			t.Logf("user.roles.implisit: %#v", ir)
		}

		// get permissions for user
		// we are using casbin model: user -> group -> role -> permission
		pu := enf.GetPermissionsForUser("group:admin")
		// assert.NotEmpty(t, pu)
		t.Logf("user/group.permissions: %#v", pu)

		// get implicit permissions for user
		ipu, err := enf.GetImplicitPermissionsForUser("group:admin")
		if assert.NoError(t, err, "Error while getting implicit permission for user/group of user: GetImplicitPermissionsForUser") {
			assert.NotEmpty(t, ipu)
			t.Logf("user.permissions.implisit: %#v", ipu)
		}

		// check user permission (enforcer)
		can := false
		for _, role := range ir {
			sub := role                                 // "group:admin"
			obj := "/api/v1/covid19/current/by-country" //"/api/v1/geolocation/country"
			act := "GET"                                // "POST"
			can, err = enf.Enforce(sub, obj, act)
			if assert.NoError(t, err, "Error while checking user permission: Enforce") {
				t.Logf("can (role) `%s` access `%s.%s`: `%v`", sub, obj, act, can)
			}
		}

		// get all policy
		p := enf.GetPolicy()
		assert.NotEmpty(t, p)
		t.Logf("GetPolicy: %#v", p)
	}
}
