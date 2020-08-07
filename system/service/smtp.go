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
	smtp.senderEmail = cfg.SMTPServers.DefaultSMTP.SenderEmail
	smtp.senderName = cfg.SMTPServers.DefaultSMTP.SenderName

	return smtp, nil
}

// SMTPSender type
type SMTPSender struct {
	handler     *handler.Handler
	server      string
	port        string
	username    string
	password    string
	senderEmail string
	senderName  string
}

type EmailFormat string

const (
	HTMLEmail EmailFormat = "HTML"
	TEXTEmail EmailFormat = "TEXT"
)

// SendEmail send email using SMTP
func (s *SMTPSender) SendEmail(toEmail string, subject string, body string) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte("To: " + toEmail + "\n" +
		"Subject: " + subject + "\n" + mime + body)

	auth := smtp.PlainAuth("", s.username, s.password, s.server)
	err := smtp.SendMail(fmt.Sprintf("%s:%s", s.server, s.port), auth, s.senderEmail, []string{toEmail}, msg)
	if err != nil {
		return fmt.Errorf("Error from SMTP Server: %s", err.Error())
	}

	return nil
}

// SendEmails send emails with specific format/type
func (s *SMTPSender) SendEmails(toEmails []string, hFromEmail string, hToEmail string, hCcEmail string, hBccEmail string, subject string, body string, format EmailFormat) error {
	mime := ""
	switch format {
	case HTMLEmail:
		mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	case TEXTEmail:
		mime = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	default:
		mime = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	}

	msg := []byte("From: " + hFromEmail + "\n" +
		"To: " + hToEmail + "\n" +
		"Cc: " + hCcEmail + "\n" +
		"Bcc: " + hBccEmail + "\n" +
		"Subject: " + subject + "\n" + mime + body)

	auth := smtp.PlainAuth("", s.username, s.password, s.server)
	err := smtp.SendMail(fmt.Sprintf("%s:%s", s.server, s.port), auth, s.senderEmail, toEmails, msg)

	if err != nil {
		return fmt.Errorf("Error from SMTP Server: %s", err.Error())
	}
	return nil
}
