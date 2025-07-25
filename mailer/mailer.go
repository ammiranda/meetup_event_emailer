package mailer

import (
	"context"
	"fmt"
	"net/smtp"
)

type Mailer struct {
	SMTPHost       string
	SMTPPort       string
	SMTPUser       string
	SMTPPassword   string
	SMTPReceipents []string
}

func NewMailer(
	smtpHost string,
	smtpPort string,
	smtpUser string,
	smtpPassword string,
	smtpReceipents []string,
) *Mailer {
	return &Mailer{
		SMTPHost:       smtpHost,
		SMTPPort:       smtpPort,
		SMTPUser:       smtpUser,
		SMTPPassword:   smtpPassword,
		SMTPReceipents: smtpReceipents,
	}
}

// SendHTMLEmail sends an HTML compliant email.
func (m *Mailer) SendHTMLEmail(ctx context.Context, to string, subject string, htmlBody string) error {
	smtpHost := m.SMTPHost
	smtpPort := m.SMTPPort
	smtpUser := m.SMTPUser
	smtpPassword := m.SMTPPassword
	smtpReceipents := m.SMTPReceipents

	message := fmt.Appendf([]byte{}, "To: %s\nSubject: %s\n\n%s", to, subject, htmlBody)

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, smtpUser, smtpReceipents, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
