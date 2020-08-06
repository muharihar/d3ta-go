package repository

import (
	"testing"

	"github.com/muharihar/d3ta-go/modules/email/la/domain/repository"
	"github.com/muharihar/d3ta-go/modules/email/la/domain/schema"
	domSchemaET "github.com/muharihar/d3ta-go/modules/email/la/domain/schema/email_template"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/initialize"
)

func newEmailRepo(t *testing.T) (repository.IEmailRepo, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabase(h); err != nil {
		return nil, nil, err
	}

	r, err := NewEmailRepo(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func TestEMailRepo_Send(t *testing.T) {
	repo, h, err := newEmailRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailRepo: %s", err.Error())
	}

	req := &schema.SendEmailRequest{
		TemplateCode: "activate-registration-html",
		// TemplateCode: "account-activation-html"
		From: &schema.MailAddress{Email: "d3tago.from@domain.tld", Name: "D3TA Golang"},
		To:   &schema.MailAddress{Email: "d3tago.test@outlook.com", Name: "D3TAgo Test (Outlook)"},
		CC: []*schema.MailAddress{
			{Email: "d3tago.test@protonmail.com", Name: "D3TAgo Test CC 1 (Protonmail)"},
			{Email: "d3tago.test.cc@tutanota.com", Name: "D3TAgo Test CC 2 (Tutanota)"}},
		BCC: []*schema.MailAddress{
			{Email: "d3tago.test@tutanota.com", Name: "D3TAgo Test BCC 1 (Tutanota)"},
			{Email: "d3tago.test.bcc@outlook.com", Name: "D3TAgo Test BCC 2 (Outlook)"}},
		TemplateData: map[string]interface{}{
			"Header.Name":        "John Doe",
			"Body.UserAccount":   "john.doe",
			"Body.ActivationURL": "https://google.com",
			"Footer.Name":        "Customer Service",
		},
		ProcessingType: "SYNC",
		Template: &domSchemaET.ETFindByCodeData{
			EmailTemplate: domSchemaET.EmailTemplate{
				ID:          1,
				EmailFormat: "HTML",
			},
			DefaultTemplateVersion: domSchemaET.EmailTemplateVersion{
				SubjectTpl: "Subject Email Template",
				BodyTpl: `{{define "T"}}<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
	</head>
	<body>
		<p>
			Dear {{index . "Header.Name"}}, 
		</p>
		<p>
			Velit consequat pariatur nisi anim. Enim mollit mollit officia voluptate eiusmod tempor dolor proident aliqua officia occaecat dolor.
			<strong>Laboris consequat ex eu ad et tempor tempor ullamco sunt est. </strong>
			Aliquip sunt cillum aute exercitation nisi non. Ad quis adipisicing voluptate laboris est elit duis. 
			Nostrud fugiat enim qui qui incididunt irure. Et elit laborum aliquip ullamco enim veniam aliqua ad veniam do ipsum.
		</p>
		<p>
			<a href="{{index . "Body.URL"}}">Eiusmod fugiat sunt elit deserunt sint labore ut dolor sint enim.</a>
		</p>
		<p>
			Regards, 
		</p>
		<p>
			{{index . "Footer.Signature"}}
		</p>

	</body>
</html>
{{end}}`,
			},
		},
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Send(req, i)
	if err != nil {
		t.Errorf("Error.EmailRepo.Send: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailRepo.Send: %s", string(respJSON))
	}
}
