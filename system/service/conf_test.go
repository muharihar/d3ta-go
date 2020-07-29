package service

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/utils"
)

func TestService_GetFileConfPath(t *testing.T) {
	f, err := GetFileConfPath("../../", "conf/casbin_rbac_rest_model.conf", nil)
	if err != nil {
		t.Errorf("Error.GetFileConfPath: %s", err.Error())
	}

	e, err := utils.FileIsExist(f)
	if err != nil {
		t.Errorf("Error.FileIsExist: %s", err.Error())
	}

	t.Errorf("File: %s, Exist: %v", f, e)
}
