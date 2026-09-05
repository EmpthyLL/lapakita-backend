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

// Helper internal untuk resolusi label persona
func resolvePersonaLabel(persona string) string {
	switch persona {
	case "tenant":
		return "Tenant / Business"
	case "owner":
		return "Stall Owner"
	case "supplier":
		return "Supplier / B2B"
	case "general":
		return "General / Media"
	case "partner":
		return "Strategic Partner"
	default:
		return persona
	}
}

// Helper internal untuk resolusi label inquiry_type berdasarkan data ContactData.ts
func resolveInquiryTypeLabel(inquiryType string) string {
	labels := map[string]string{
		// Partnership Options
		"property_developer": "Property Manager / Developer",
		"event_organizer":    "Event Organizer (EO)",
		"b2b_distributor":    "Main Distributor / FMCG Producer",
		"financial_bank":     "Banking / Payment Gateway",
		"sme_community":      "MSME Community & Institution",
		"other_partner":      "Other Forms of Partnership",

		// Tenant Options
		"rental_inquiry":       "Stall Search & Rental Application",
		"bazaar_event":         "Bazaar / Pop-up Event Registration",
		"lease_contract":       "Lease Contract & Start Date Issues",
		"pos_cashier":          "POS Machine & Staff Cashier Account Support",
		"deposit_refund":       "Escrow Deposit Refund",
		"billing_subscription": "Pro Subscription Billing",
		"supplier_order":       "Supplier Raw Material Order Issues",
		"other_tenant":         "Other (Tenant Issues)",

		// Owner Options
		"stall_listing":       "Stall Listing & Verification Help",
		"bazaar_creation":     "Bazaar / Festival Event Slot Creation",
		"tenant_vetting":      "Tenant Approval / Rejection",
		"deposit_claim":       "Property Damage / Lost Key Claims",
		"payout_disbursement": "Bank Rent Payout Issues",
		"other_owner":         "Other (Owner Issues)",

		// Supplier Options
		"catalog_approval":  "B2B Product Catalog Upload & Verification",
		"moq_pricing":       "MOQ Setup & Wholesale Pricing Tiers",
		"order_fulfillment": "Order & Digital Delivery Slip Issues",
		"buyer_dispute":     "Order Disputes / Product Complaints",
		"other_supplier":    "Other (Supplier Issues)",

		// General Options
		"press_media":         "Press, News, & Media Coverage",
		"career_job":          "Careers & Recruitment Information",
		"feedback_suggestion": "Platform Feedback & Suggestions",
		"other_general":       "Other (General Inquiries)",
	}

	if label, exists := labels[inquiryType]; exists {
		return label
	}
	return inquiryType
}

func (m *Mailer) SendContactInquiryAutoReply(toEmail string, name string, persona string, inquiryType string, message string) error {
	workDir, err := os.Getwd()
	if err != nil {
		m.log.Error("[Mailer] Failed to get working directory", zap.Error(err))
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	tmplPath := filepath.Join(workDir, "pkg", "mailer", "template", "contact_us.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		m.log.Error("[Mailer] Failed to parse email template file",
			zap.String("path", tmplPath),
			zap.Error(err),
		)
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	isPartnership := persona == "partner"

	data := struct {
		Name             string
		PersonaLabel     string
		InquiryTypeLabel string
		Message          string
		IsPartnership    bool
	}{
		Name:             name,
		PersonaLabel:     resolvePersonaLabel(persona),
		InquiryTypeLabel: resolveInquiryTypeLabel(inquiryType),
		Message:          message,
		IsPartnership:    isPartnership,
	}

	subject := "We've Received Your Inquiry - Lapakita"
	if isPartnership {
		subject = "Partnership Inquiry Received - Lapakita Collaboration"
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
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.cfg.SMTPHost,
	}

	if err := dialer.DialAndSend(msg); err != nil {
		m.log.Error("[Mailer] Failed to send contact inquiry auto-reply via SMTP",
			zap.String("to", toEmail),
			zap.Error(err),
		)
		return fmt.Errorf("failed to send auto-reply email via gomail: %w", err)
	}

	m.log.Info("[Mailer] Contact inquiry auto-reply email sent successfully",
		zap.String("to", toEmail),
		zap.Bool("is_partnership", isPartnership),
	)

	return nil
}
