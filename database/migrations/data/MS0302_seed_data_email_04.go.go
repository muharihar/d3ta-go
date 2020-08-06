package data

import (
	domEmailEtt "github.com/muharihar/d3ta-go/modules/email/la/domain/entity"
	"github.com/muharihar/d3ta-go/system/utils"
)

// EmailTemplate04 data (TEXT)
func EmailTemplate04() domEmailEtt.EmailTemplate {
	return domEmailEtt.EmailTemplate{
		UUID:        utils.GenerateUUID(),
		Code:        "account-activation-html",
		Name:        "Account Activation Email (HTML)",
		IsActive:    true,
		EmailFormat: "HTML",
	}
}

// EmailTemplate04Version data
func EmailTemplate04Version() domEmailEtt.EmailTemplateVersion {
	return domEmailEtt.EmailTemplateVersion{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Account Activation",
		BodyTpl: `{{define "T"}}<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
	</head>
	<body>
		<p>
			Dear {{index . "Header.Name"}},
		</p>
		<p>
			Conglatulation! your account has been activated.
		</p>
		<p>
			If you didn't attempt to verify your email address with our service, delete this email.
		</p>
		<p>
			Cheers,
		</p>
		<p>
			{{index . "Footer.Name"}}
		</p>
	</body>
</html>{{end}}`,
	}
}
