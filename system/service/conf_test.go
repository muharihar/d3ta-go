package service

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/utils"
	"github.com/stretchr/testify/assert"
)

func TestService_GetFileConfPath(t *testing.T) {
	f, err := GetFileConfPath("../../", "conf/casbin/casbin_rbac_rest_model.conf", nil)
	if assert.NoError(t, err, "Error while get file conf: GetFileConfPath error") {
		assert.Equal(t, "../../conf/casbin/casbin_rbac_rest_model.conf", f)
	}

	e, err := utils.FileIsExist(f)
	if assert.NoError(t, err) {
		assert.Equal(t, true, e)
	}
}
