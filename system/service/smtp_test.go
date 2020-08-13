package service

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/stretchr/testify/assert"
)

func newConfig(t *testing.T) (*config.Config, error) {

	c, _, err := config.NewConfig("../../conf")
	if err != nil {
		return nil, err
	}

	return c, nil
}

func newEmail(t *testing.T) (*SMTPSender, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetConfig(c)

	r, err := NewSMTPSender(h)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func TestSMTP_SendEmail(t *testing.T) {

	smtp, err := newEmail(t)
	if assert.NoError(t, err, "Error while create smtp: newEmail error") {
		err = smtp.SendEmail("admin.d3tago@email.tld", "Subject Test Email", "Body Test Email")
		if assert.NoError(t, err, "Error while sending email via smtp: smtp.SendEmail error") {
			return
		}
	}
}

func TestSMTP_SendEmails(t *testing.T) {
	smtp, err := newEmail(t)
	if assert.NoError(t, err, "error while create smtp: newEmail error") {

		toEmail := []string{"to.d3tago@email.tld", "cc.d3tago@email.tld", "bcc.d3tago@email.tld"}
		hFromEmail := "From Name <from.d3tago@email.tld>"
		hToEmail := "To Email <to.d3tago@email.tld>"
		hCcEmail := "CC Email <cc.d3tago@email.tld>"
		hBccEmail := "BCC Email <bcc.d3tago@gmail.tld>"

		err = smtp.SendEmails(toEmail, hFromEmail, hToEmail, hCcEmail, hBccEmail, "Subject Test Eemails", "Body Test Emails (cc, bcc)", TEXTEmail)
		if assert.NoError(t, err, "error while sending email via smtp: smtp.SendEmails error") {
			return
		}
	}
}
