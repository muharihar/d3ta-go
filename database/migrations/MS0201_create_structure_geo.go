package migrations

import (
	"github.com/muharihar/d3ta-go/database/base"
	geoModel "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/model"
	"github.com/muharihar/d3ta-go/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewMS0201CreateStructureGeo new MS0201CreateStructureGeo
func NewMS0201CreateStructureGeo(h *handler.Handler) base.IMigration {
	m := new(MS0201CreateStructureGeo)
	m.handler = h

	return m
}

// MS0201CreateStructureGeo type
type MS0201CreateStructureGeo struct {
	BaseMigration
}

// Run migration
func (ms01 *MS0201CreateStructureGeo) Run(db *gorm.DB) error {
	// DB GeoLocation
	if err := db.AutoMigrate(&geoModel.Country{}); err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *MS0201CreateStructureGeo) Rollback(db *gorm.DB) error {
	var count int64

	if db.Migrator().HasTable(&geoModel.Country{}) {
		db.Find(&geoModel.Country{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&geoModel.Country{}); err != nil {
				return err
			}
		}
	}

	return nil
}
