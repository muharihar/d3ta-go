package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_GenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	uuidNext := GenerateUUID()

	if assert.NotEqual(t, uuid, uuidNext) {
		return
	}
}

func TestGenerate_GeneratePassword(t *testing.T) {
	pwd := GeneratePassword()
	pwdNext := GeneratePassword()

	if assert.NotEqual(t, pwd, pwdNext) {
		return
	}
}

func TestGenerate_GenerateClientKey(t *testing.T) {
	clientKey := GenerateClientKey()
	clientKeyNext := GenerateClientKey()

	if assert.NotEqual(t, clientKey, clientKeyNext) {
		return
	}
}

func TestGenerate_GenerateSecretKey(t *testing.T) {
	secretKey := GenerateSecretKey()
	secretKeyNext := GenerateSecretKey()

	if assert.NotEqual(t, secretKey, secretKeyNext) {
		return
	}
}

func TestGenerate_GenerateRegistrationActivationCode(t *testing.T) {
	activationCode := GenerateRegistrationActivationCode()
	activationCodeNext := GenerateRegistrationActivationCode()

	if assert.NotEqual(t, activationCode, activationCodeNext) {
		return
	}
}

func TestGenerator_GenSemVersion(t *testing.T) {
	// first
	v := GenSemVersion("")
	if !assert.Equal(t, "1.0.0", v) {
		return
	}

	// next patch
	v = GenSemVersion(v)
	if !assert.Equal(t, "1.0.1", v) {
		return
	}

	// next minor
	v = "1.0.9"
	v = GenSemVersion(v)
	if !assert.Equal(t, "1.1.0", v) {
		return
	}

	// next major
	v = "1.9.9"
	v = GenSemVersion(v)
	if !assert.Equal(t, "2.0.0", v) {
		return
	}
}
