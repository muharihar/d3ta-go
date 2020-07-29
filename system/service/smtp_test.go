package service

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
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

func TestEmail_Send(t *testing.T) {

	smtp, err := newEmail(t)
	if err != nil {
		t.Errorf("newEmail: %s", err.Error())
		return
	}

	err = smtp.SendEmail("admin.d3tago@email.tld", "Subject Test Email", "Body Test Email")
	if err != nil {
		t.Errorf("SendEmail: %s", err.Error())
		return
	}
}
