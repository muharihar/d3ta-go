package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash_MD5(t *testing.T) {
	expected := "dcc7c9b47b7b3e278a982773feb767f5"
	md5 := MD5([]byte("please-make-me-as-md5"))
	if !assert.Equal(t, expected, md5) {
		return
	}
}

func TestHash_AsSha256(t *testing.T) {
	expected := "d4d72eb3fa0a7f80ede9bb91f136ec027d912533c832566187fa55724ad85a33"
	sha256 := AsSha256("plese-make-me-as-sha256")
	if !assert.Equal(t, expected, sha256) {
		return
	}
}
