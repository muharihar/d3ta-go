package entity

// SysUserClientApps type
type SysUserClientApps struct {
	ID uint64 `gorm:"primary_key;column:id"`

	UUID          string `json:"uuid" gorm:"column:uuid;size:255;unique;not null" sql:"index"`
	ClientAppCode string `json:"clientAppCode" gorm:"column:client_app_code;size:255;not null;unique" sql:"index"`
	ClientAppName string `json:"clientAppName" gorm:"column:client_app_name;size:255;not null"`
	ClientAppDesc string `json:"ClientAppDesc" gorm:"column:client_app_desc;size:255;"`
	ClientKey     string `json:"clientKey" gorm:"column:client_key;size:255;" sql:"index"`
	SecretKey     string `json:"secretKey" gorm:"column:secret_key;size:500;" sql:"index"`
	IsActive      bool   `json:"isActive" gorm:"column:is_active;index" sql:"index"`

	User   SysUser `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:ID;"`
	UserID uint64  `json:"userID" gorm:"column:user_id;" sql:"index"`

	BaseEntity
}

// TableName get real database table name
func (t SysUserClientApps) TableName() string {
	return "iam_sys_user_client_apps"
}
