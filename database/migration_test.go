package database

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/initialize"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newMigration(t *testing.T) (*Migration, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabase(h); err != nil {
		return nil, nil, err
	}

	db, err := h.GetGormDB(c.Databases.IdentityDB.ConnectionName)
	if err != nil {
		return nil, nil, err
	}

	mig := NewMigration(db)

	return mig, h, nil
}

type MigTest struct {
}

type TestTable struct {
	ID       uint64 `gorm:"primary_key;column:id"`
	Column01 string `gorm:"column:column_01;size:200;"`
	Column02 string `gorm:"column:column_02;size:255;"`
}

func (ms *MigTest) Run(db *gorm.DB) error {
	if err := db.AutoMigrate(&TestTable{}); err != nil {
		return err
	}
	return nil
}

func (ms *MigTest) Rollback(db *gorm.DB) error {
	if db.Migrator().HasTable(&TestTable{}) {
		if err := db.Migrator().DropTable(&TestTable{}); err != nil {
			return err
		}
	}
	return nil
}

func TestMigration_Run(t *testing.T) {
	mig, _, err := newMigration(t)
	if err != nil {
		t.Errorf("newMigration: %s", err.Error())
	}

	if err := mig.Run(&MigTest{}); err != nil {
		t.Errorf("Run: %s", err.Error())
	}
}

func TestMigration_Rollback(t *testing.T) {
	mig, _, err := newMigration(t)
	if err != nil {
		t.Errorf("newMigration: %s", err.Error())
	}

	if err := mig.Rollback(&MigTest{}); err != nil {
		t.Errorf("Rollback: %s", err.Error())
	}
}
