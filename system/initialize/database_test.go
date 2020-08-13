package initialize

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/stretchr/testify/assert"
)

func newConfig(t *testing.T) (*config.Config, error) {

	c, _, err := config.NewConfig("../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newHandler(t *testing.T) (*handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetConfig(c)

	return h, nil
}

func TestLoadAllDatabaseConnection(t *testing.T) {

	h, err := newHandler(t)
	if assert.NoError(t, err, "Error while creating handler: newHandler") {
		if !assert.NotNil(t, h) {
			return
		}
	}

	if assert.NoError(t, LoadAllDatabaseConnection(h), "Error while loading all database connection: LoadAllDatabaseConnection") {

		cfg, err := h.GetConfig()
		if !assert.NoError(t, err, "Error while getting config: h.GetConfig") {
			return
		}

		dbCon, err := h.GetGormDB(cfg.Databases.IdentityDB.ConnectionName)
		if assert.NoError(t, err, "Error while getting GORM DB Connection: h.GetGormDB(cfg.Databases.IdentityDB.ConnectionName)") {

			type result struct {
				DBName string //`json:"DBName" gorm:"column:db_name"`
			}
			results := []result{}

			if assert.NoError(t, dbCon.Raw("SELECT DATABASE() db_name FROM DUAL").Scan(&results).Error) {
				if assert.NotEmpty(t, results) {
					assert.Equal(t, cfg.Databases.IdentityDB.DBName, results[0].DBName)
				}
				t.Logf("Result: %#v", results)
			}
		}
	}
}
