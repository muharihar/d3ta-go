package database

import (
	"github.com/muharihar/d3ta-go/database/base"
	"gorm.io/gorm"
)

// NewMigration new Migration
func NewMigration(db *gorm.DB) *Migration {

	mig := new(Migration)
	mig.DB = db

	return mig
}

// Migration type
type Migration struct {
	DB *gorm.DB
}

// Run func
func (m *Migration) Run(mig base.IMigration) error {
	return mig.Run(m.DB)
}

// Rollback func
func (m *Migration) Rollback(mig base.IMigration) error {
	return mig.Rollback(m.DB)
}
