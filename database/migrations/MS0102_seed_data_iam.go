package migrations

import (
	"github.com/muharihar/d3ta-go/database/base"
	iamEntt "github.com/muharihar/d3ta-go/modules/auths/la/domain/entity"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/utils"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewMS0102SeedDataIAM new MS0102SeedDataIAM
func NewMS0102SeedDataIAM(h *handler.Handler) base.IMigration {
	m := new(MS0102SeedDataIAM)
	m.handler = h

	return m
}

// MS0102SeedDataIAM type
type MS0102SeedDataIAM struct {
	BaseMigration
}

// Run migration
func (ms01 *MS0102SeedDataIAM) Run(db *gorm.DB) error {
	// DB Identity
	cfg, err := ms01.handler.GetConfig()
	if err != nil {
		return err
	}

	// add default user: super admin
	superAdmin := iamEntt.SysUser{
		UUID:        utils.GenerateUUID(),
		Username:    cfg.IAM.DefaultAdmin.Username,
		Password:    utils.MD5([]byte(cfg.IAM.DefaultAdmin.Password)),
		NickName:    cfg.IAM.DefaultAdmin.NickName,
		Email:       cfg.IAM.DefaultAdmin.Email,
		IsActive:    true,
		AuthorityID: cfg.IAM.DefaultAdmin.AuthorityID,
	}
	superAdmin.CreatedBy = "system.d3tago@installation"

	if err := db.Create(&superAdmin).Error; err != nil {
		return err
	}

	// add default user client app: super admin
	superAdminApp := iamEntt.SysUserClientApps{
		UUID:          utils.GenerateUUID(),
		ClientAppCode: "super-admin-app",
		ClientAppName: "Super Admin Client Application",
		ClientAppDesc: "Default Installation",
		ClientKey:     utils.GenerateClientKey(),
		SecretKey:     utils.GenerateSecretKey(),
		IsActive:      true,
		UserID:        superAdmin.ID,
	}
	superAdminApp.CreatedBy = "system.d3tago@installation"

	if err := db.Create(&superAdminApp).Error; err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *MS0102SeedDataIAM) Rollback(db *gorm.DB) error {
	return nil
}
