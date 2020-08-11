package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNet_GetCurrentIP(t *testing.T) {
	ip := GetCurrentIP()
	if !assert.NotEqual(t, "", ip) {
		return
	}
}

func TestNet_GetHostName(t *testing.T) {
	hn := GetHostName()
	if !assert.NotEqual(t, "", hn) {
		return
	}
}
