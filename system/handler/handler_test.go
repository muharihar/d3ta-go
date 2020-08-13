package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	h, err := NewHandler()
	if !assert.NoError(t, err, "Error while create new Handler: NewHandler") {
		return
	}

	// Config
	h.SetConfig(nil)
	cfg, err := h.GetConfig()
	if assert.Error(t, err, "Should be Error while getting config from Handler: h.GetConfig()") {
		assert.Nil(t, cfg)
	}

	// Gorm DB
	dbCon, err := h.GetGormDB("not-found")
	if assert.Error(t, err, "Should be Error while getting GormDB from Handler: h.GetGormDB()") {
		assert.Nil(t, dbCon)
	}

	h.SetGormDB("nil-value", nil)
	dbCon2, err2 := h.GetGormDB("nil-value")
	if assert.NoError(t, err2, "Error while getting GormDB from Handler: h.GetGormDB()") {
		assert.Nil(t, dbCon2)
	}

	// Casbin Enforcer
	ce, err := h.GetCasbinEnforcer("not-found")
	if assert.Error(t, err, "Should be Error while getting Casbin Enforcer from Handler: h.GetCasbinEnforcer()") {
		assert.Nil(t, ce)
	}

	h.SetCasbinEnforcer("nil-value", nil)
	ce2, err2 := h.GetCasbinEnforcer("nil-value")
	if assert.NoError(t, err2, "Error while getting Casbin Enforcer from Handler: h.GetCasbinEnforcer()") {
		assert.Nil(t, ce2)
	}
}
