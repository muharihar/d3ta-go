package data

import (
	domEmailEtt "github.com/muharihar/d3ta-go/modules/email/la/domain/entity"
	"github.com/muharihar/d3ta-go/system/utils"
)

// EmailTemplate03 data (TEXT)
func EmailTemplate03() domEmailEtt.EmailTemplate {
	return domEmailEtt.EmailTemplate{
		UUID:        utils.GenerateUUID(),
		Code:        "account-activation-plaintext",
		Name:        "Account Activation Email (TEXT)",
		IsActive:    true,
		EmailFormat: "TEXT",
	}
}

// EmailTemplate03Version data
func EmailTemplate03Version() domEmailEtt.EmailTemplateVersion {
	return domEmailEtt.EmailTemplateVersion{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Account Activation",
		BodyTpl: `{{define "T"}}Dear {{index . "Header.Name"}},

Conglatulation! your account has been activated.

If you didn't attempt to verify your email address with our service, delete this email.

Cheers,

{{index . "Footer.Name"}}
{{end}}`,
	}
}
