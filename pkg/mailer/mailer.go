package mailer

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"lapakita-backend/config"
	"lapakita-backend/pkg/logger"

	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	cfg *config.Config
	log *logger.Logger
}

func NewMailer(cfg *config.Config, log *logger.Logger) *Mailer {
	return &Mailer{
		cfg: cfg,
		log: log,
	}
}

func (m *Mailer) SendOTPEmail(toEmail string, otpCode string, mode string) error {
	var templateFile string
	var subject string

	if mode == "register" {
		templateFile = "register_otp.html"
		subject = "Verify Your Lapakita Registration"
	} else {
		templateFile = "reset_password_otp.html"
		subject = "Reset Your Lapakita Password"
	}

	workDir, err := os.Getwd()
	if err != nil {
		m.log.Error("[Mailer] Failed to get working directory", zap.Error(err))
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	tmplPath := filepath.Join(workDir, "pkg", "mailer", "template", templateFile)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		m.log.Error("[Mailer] Failed to parse template file",
			zap.String("path", tmplPath),
			zap.Error(err),
		)
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
		m.log.Error("[Mailer] Failed to execute template", zap.Error(err))
		return fmt.Errorf("failed to execute template: %w", err)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.cfg.SMTPSender)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body.String())

	dialer := gomail.NewDialer(m.cfg.SMTPHost, m.cfg.SMTPPort, m.cfg.SMTPUser, m.cfg.SMTPPassword)

	// Izinkan TLS Handshake tanpa terblokir sertifikat lokal
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.cfg.SMTPHost,
	}

	if err := dialer.DialAndSend(msg); err != nil {
		m.log.Error("[Mailer] Failed to send email via SMTP",
			zap.String("to", toEmail),
			zap.String("host", m.cfg.SMTPHost),
			zap.Int("port", m.cfg.SMTPPort),
			zap.String("user", m.cfg.SMTPUser),
			zap.Error(err),
		)
		return fmt.Errorf("failed to send email via gomail: %w", err)
	}

	m.log.Info("[Mailer] OTP email sent successfully",
		zap.String("to", toEmail),
		zap.String("mode", mode),
	)

	return nil
}
