package migrations

import (
	"github.com/muharihar/d3ta-go/database/base"
	domEmailEtt "github.com/muharihar/d3ta-go/modules/email/la/domain/entity"
	"github.com/muharihar/d3ta-go/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewMS0301CreateStructureEmail new MS0301CreateStructureEmail
func NewMS0301CreateStructureEmail(h *handler.Handler) base.IMigration {
	m := new(MS0301CreateStructureEmail)
	m.handler = h

	return m
}

// MS0301CreateStructureEmail type
type MS0301CreateStructureEmail struct {
	BaseMigration
}

// Run migration
func (ms01 *MS0301CreateStructureEmail) Run(db *gorm.DB) error {
	// DB GeoLocation
	if err := db.AutoMigrate(&domEmailEtt.Email{}, &domEmailEtt.EmailTemplate{}, &domEmailEtt.EmailTemplateVersion{}); err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *MS0301CreateStructureEmail) Rollback(db *gorm.DB) error {
	var count int64

	if db.Migrator().HasTable(&domEmailEtt.Email{}) {
		db.Find(&domEmailEtt.Email{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&domEmailEtt.Email{}); err != nil {
				return err
			}
		}
	}

	if db.Migrator().HasTable(&domEmailEtt.EmailTemplate{}) {
		db.Find(&domEmailEtt.EmailTemplate{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&domEmailEtt.EmailTemplate{}); err != nil {
				return err
			}
		}
	}

	if db.Migrator().HasTable(&domEmailEtt.EmailTemplateVersion{}) {
		db.Find(&domEmailEtt.EmailTemplateVersion{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&domEmailEtt.EmailTemplateVersion{}); err != nil {
				return err
			}
		}
	}

	return nil
}
