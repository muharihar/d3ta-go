package base

import "gorm.io/gorm"

// IMigration interface
type IMigration interface {
	Run(db *gorm.DB) error
	Rollback(db *gorm.DB) error
}
