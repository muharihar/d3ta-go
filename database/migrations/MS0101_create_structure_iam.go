package migrations

import (
	"github.com/muharihar/d3ta-go/database/base"
	iamEntt "github.com/muharihar/d3ta-go/modules/auths/la/domain/entity"
	"github.com/muharihar/d3ta-go/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewMS0101CreateStructureIAM new MS0101CreateStructureIAM
func NewMS0101CreateStructureIAM(h *handler.Handler) base.IMigration {
	m := new(MS0101CreateStructureIAM)
	m.handler = h

	return m
}

// MS0101CreateStructureIAM type
type MS0101CreateStructureIAM struct {
	BaseMigration
}

// Run migration
func (ms01 *MS0101CreateStructureIAM) Run(db *gorm.DB) error {
	// DB Identity
	if err := db.AutoMigrate(&iamEntt.SysUser{}, &iamEntt.SysUserClientApps{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&iamEntt.TmpUserRegistration{}); err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *MS0101CreateStructureIAM) Rollback(db *gorm.DB) error {
	var count int64

	if db.Migrator().HasTable(&iamEntt.SysUserClientApps{}) {
		db.Find(&iamEntt.SysUserClientApps{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&iamEntt.SysUserClientApps{}); err != nil {
				return err
			}
		}
	}

	if db.Migrator().HasTable(&iamEntt.SysUser{}) {
		db.Find(&iamEntt.SysUser{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&iamEntt.SysUser{}); err != nil {
				return err
			}
		}
	}

	if db.Migrator().HasTable(&iamEntt.TmpUserRegistration{}) {
		db.Find(&iamEntt.TmpUserRegistration{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&iamEntt.TmpUserRegistration{}); err != nil {
				return err
			}
		}
	}

	return nil
}
