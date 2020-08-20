package initialize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
