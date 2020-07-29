package entity

// SysUser type
type SysUser struct {
	ID uint64 `json:"ID" gorm:"primary_key;column:id"`

	UUID      string `json:"uuid" gorm:"column:uuid;size:255;unique;not null" sql:"index"`
	Username  string `json:"userName" gorm:"column:username;size:255;unique;not null" sql:"index"`
	Password  string `json:"-" gorm:"column:password;size:255;not null;index" sql:"index"`
	NickName  string `json:"nickName" gorm:"column:nick_name;size:255;not null"`
	Email     string `json:"email" gorm:"column:email;size:255;unique;not null" sql:"index"`
	HeaderImg string `json:"headerImg" gorm:"column:header_img;size:255"`
	IsActive  bool   `json:"isActive" gorm:"column:is_active;index" sql:"index"`

	// Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId;"`
	AuthorityID string `json:"authorityId" gorm:"column:authority_id;size:255;not null;index" sql:"index"`

	BaseEntity
}

// TableName get real database table name
func (t SysUser) TableName() string {
	return "iam_sys_users"
}
