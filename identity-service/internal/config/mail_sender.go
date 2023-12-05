package config

import (
	"bytes"
	"fmt"
	"html/template"
	"identity-service/internal/models"
	"net/smtp"
	"os"
)

type MailConfig struct {
	FromMail string
	Password string
	Host     string
	Port     string
}

func NewMailConfig(FromMail, Password, Host, Port string) *MailConfig {
	return &MailConfig{
		FromMail: os.Getenv("SMTP_MAIL_ID"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
	}
}

func (m *MailConfig) ParseTemplate(templateFileName, subject string, data interface{}) ([]byte, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return []byte{}, err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return []byte{}, err
	}

	body := buf.String()
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	mailSubject := "Subject: " + subject + "!\n"

	return []byte(mailSubject + mime + "\n" + body), nil
}

func (m *MailConfig) SendOtpEmail(subject, to string, data models.OTPMailData) error {
	toList := []string{to}
	address := fmt.Sprintf("%s:%s", m.Host, m.Port)

	auth := smtp.PlainAuth("", m.FromMail, m.Password, m.Host)
	message, err := m.ParseTemplate("", subject, data)

	if err != nil {
		return err
	}

	if err := smtp.SendMail(address, auth, m.FromMail, toList, message); err != nil {
		return err
	}
	return nil
}
