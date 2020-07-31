package entity

import (
	"time"

	"gorm.io/gorm"
)

// BaseEntity base type
type BaseEntity struct {
	CreatedBy string         `json:"createdBy,omitempty" gorm:"column:sys_created_by;size:255" `
	CreatedAt *time.Time     `json:"createdAt,omitempty" gorm:"column:sys_created_at" sql:"index"`
	UpdatedBy string         `json:"updatedBy,omitempty" gorm:"column:sys_updated_by;size:255"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty" gorm:"column:sys_updated_at" sql:"index"`
	DeletedBy string         `json:"deletedBy,omitempty" gorm:"column:sys_deleted_by;size:255"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;index:idx_delete_at" sql:"index"`
	//DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"column:sys_deleted_at" sql:"index"`
}
