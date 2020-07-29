package service

import (
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	cGormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewCasbinEnforcer new casbin enforcer
func NewCasbinEnforcer(h *handler.Handler, modelPath string) (*casbin.Enforcer, error) {
	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}

	dbCon, err := h.GetGormDB(cfg.Databases.IdentityDB.ConnectionName)
	if err != nil {
		return nil, err
	}

	if modelPath == "" {
		modelPath = cfg.IAM.Casbin.ModelPath
	}

	modelPath, err = GetFileConfPath("", modelPath, h)
	if err != nil {
		return nil, err
	}

	/*
		// enf first then adapter
		enf, err := casbin.NewEnforcer(modelPath)
		if err != nil {
			return nil, err
		}
		enf.AddFunction("ParamsMatch", paramsMatchFunc)
	*/

	cAdp, err := cGormAdapter.NewAdapterByDBUsePrefix(dbCon, "iam_")
	if err != nil {
		return nil, err
	}

	/*
		cAdp.LoadPolicy(enf.GetModel())
		if err != nil {
			return nil, err
		}
	*/

	/*
		err = cAdp.LoadFilteredPolicy(enf.GetModel(), cGormAdapter.Filter{
			V0: []string{"role:admin"},
		})
		if err != nil {
			return nil, err
		}
	*/

	enf, err := casbin.NewEnforcer(modelPath, cAdp)
	if err != nil {
		return nil, err
	}
	enf.AddFunction("ParamsMatch", paramsMatchFunc)

	/*
		err = enf.LoadFilteredPolicy(cGormAdapter.Filter{
			V0: []string{"role:admin"},
		})
		if err != nil {
			return nil, err
		}
	*/

	return enf, nil
}

func paramsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}

func paramsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return paramsMatch(name1, name2), nil
}
