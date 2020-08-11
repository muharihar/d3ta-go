package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFS_FileIsExist(t *testing.T) {
	exist, err := FileIsExist("./file_system.go")
	if assert.NoError(t, err) {
		assert.Equal(t, true, exist)
	}

	notExist, err := FileIsExist("./file_system" + GenerateUUID() + ".go")
	if assert.Error(t, err) {
		assert.Equal(t, false, notExist)
	}
}
