package database

import (
	"github.com/muharihar/d3ta-go/database/migrations"
	"github.com/muharihar/d3ta-go/system/handler"
)

// M01FirstInstall func
func M01FirstInstall(h *handler.Handler) error {
	cfg, err := h.GetConfig()
	if err != nil {
		return err
	}

	// identity (IAM)
	iamDB, err := h.GetGormDB(cfg.Databases.IdentityDB.ConnectionName)
	if err != nil {
		return err
	}
	iamMig := NewMigration(iamDB)

	// Create Structure
	if err := iamMig.Run(migrations.NewMS0101CreateStructureIAM(h)); err != nil {
		if err := iamMig.Rollback(migrations.NewMS0101CreateStructureIAM(h)); err != nil {
			return err
		}
	}
	// Seed Data
	if err := iamMig.Run(migrations.NewMS0102SeedDataIAM(h)); err != nil {
		if err := iamMig.Rollback(migrations.NewMS0102SeedDataIAM(h)); err != nil {
			return err
		}
	}
	// Init Casbin
	if err := iamMig.Run(migrations.NewMS0103CreateAndInitCasbin(h)); err != nil {
		if err := iamMig.Rollback(migrations.NewMS0103CreateAndInitCasbin(h)); err != nil {
			return err
		}
	}

	// GeoLocation (MainDB)
	geoDB, err := h.GetGormDB(cfg.Databases.MainDB.ConnectionName)
	if err != nil {
		return err
	}
	geoMig := NewMigration(geoDB)

	// Create Structure
	if err := geoMig.Run(migrations.NewMS0201CreateStructureGeo(h)); err != nil {
		if err := geoMig.Rollback(migrations.NewMS0201CreateStructureGeo(h)); err != nil {
			return err
		}
	}
	// Seed Data
	if err := geoMig.Run(migrations.NewMS0201SeedDataGeo(h)); err != nil {
		if err := geoMig.Rollback(migrations.NewMS0201SeedDataGeo(h)); err != nil {
			return err
		}
	}

	// Email (EmailDB)
	emailDB, err := h.GetGormDB(cfg.Databases.EmailDB.ConnectionName)
	if err != nil {
		return err
	}
	emailMig := NewMigration(emailDB)

	// Create Structure
	if err := emailMig.Run(migrations.NewMS0301CreateStructureEmail(h)); err != nil {
		if err := emailMig.Rollback(migrations.NewMS0301CreateStructureEmail(h)); err != nil {
			return err
		}
	}
	// Seed Data
	if err := emailMig.Run(migrations.NewMS0301SeedDataEmail(h)); err != nil {
		if err := emailMig.Rollback(migrations.NewMS0301SeedDataEmail(h)); err != nil {
			return err
		}
	}

	return nil
}
