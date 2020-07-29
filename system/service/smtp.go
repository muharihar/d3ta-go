package service

import (
	"fmt"
	"net/smtp"

	"github.com/muharihar/d3ta-go/system/handler"
)

// NewSMTPSender new SMTPSender
func NewSMTPSender(h *handler.Handler) (*SMTPSender, error) {
	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}

	smtp := new(SMTPSender)
	smtp.handler = h

	smtp.server = cfg.SMTPServers.DefaultSMTP.Server
	smtp.port = cfg.SMTPServers.DefaultSMTP.Port
	smtp.username = cfg.SMTPServers.DefaultSMTP.Username
	smtp.password = cfg.SMTPServers.DefaultSMTP.Password
	smtp.sender = cfg.SMTPServers.DefaultSMTP.Sender

	return smtp, nil
}

// SMTPSender type
type SMTPSender struct {
	handler  *handler.Handler
	server   string
	port     string
	username string
	password string
	sender   string
}

// SendEmail send email using SMTP
func (s *SMTPSender) SendEmail(toEmail string, subject string, body string) error {
	subjectBody := fmt.Sprintf("Subject: %s\n\n %s", subject, body)

	auth := smtp.PlainAuth("", s.username, s.password, s.server)
	err := smtp.SendMail(fmt.Sprintf("%s:%s", s.server, s.port), auth, s.sender, []string{toEmail}, []byte(subjectBody))
	if err != nil {
		return fmt.Errorf("Error from SMTP Server: %s", err.Error())
	}

	return nil
}
