package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"lapakita-backend/config"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	cfg *config.Config
}

func NewMailer(cfg *config.Config) *Mailer {
	return &Mailer{cfg: cfg}
}

func (m *Mailer) SendOTPEmail(toEmail string, otpCode string, mode string) error {
	var templateFile string
	var subject string

	if mode == "register" {
		templateFile = "otp_register.html"
		subject = "Verify Your Lapakita Registration"
	} else {
		templateFile = "otp_reset_password.html"
		subject = "Reset Your Lapakita Password"
	}

	tmplPath := filepath.Join("pkg", "mailer", "templates", templateFile)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	data := struct {
		OTPCode string
		Email   string
	}{
		OTPCode: otpCode,
		Email:   toEmail,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.cfg.SMTPSender)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body.String())

	dialer := gomail.NewDialer(m.cfg.SMTPHost, m.cfg.SMTPPort, m.cfg.SMTPUser, m.cfg.SMTPPassword)

	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email via gomail: %w", err)
	}

	return nil
}
