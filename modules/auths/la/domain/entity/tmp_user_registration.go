package entity

import "time"

// TmpUserRegistration type
type TmpUserRegistration struct {
	ID uint64 `json:"ID" gorm:"primary_key;column:id"`

	UUID     string `json:"uuid" gorm:"column:uuid;size:255;unique;not null" sql:"index"`
	Username string `json:"userName" gorm:"column:username;size:255;unique;not null" sql:"index"`
	Password string `json:"-" gorm:"column:password;size:255;not null" sql:"index"`
	NickName string `json:"nickName" gorm:"column:nick_name;size:255;not null"`
	Email    string `json:"email" gorm:"column:email;size:255;unique;not null;" sql:"index"`

	IsActivated    bool       `json:"isActivated" gorm:"column:is_activated;index" sql:"index"`
	ActivationCode string     `json:"activationCode" gorm:"column:activation_code;size:255" sql:"index"`
	ActivatedAt    *time.Time `json:"activateddAt,omitempty" gorm:"column:sys_activated_at" sql:"index"`

	BaseEntity
}

// TableName get real database table name
func (t TmpUserRegistration) TableName() string {
	return "iam_tmp_user_registrations"
}
