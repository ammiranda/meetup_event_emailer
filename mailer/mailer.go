package mailer

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"
)

//go:generate mockery --name EmailSender --output ./mocks --outpkg mocks

type EmailSender interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

type SMTPSender struct{}

func (s *SMTPSender) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

type Mailer struct {
	SMTPHost       string
	SMTPPort       string
	SMTPUser       string
	SMTPPassword   string
	SMTPReceipents []string
	Sender         EmailSender
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
		Sender:         &SMTPSender{},
	}
}

// SendHTMLEmail sends an HTML compliant email.
func (m *Mailer) SendHTMLEmail(ctx context.Context, subject string, htmlBody string) error {
	smtpHost := m.SMTPHost
	smtpPort := m.SMTPPort
	smtpUser := m.SMTPUser
	smtpPassword := m.SMTPPassword
	smtpReceipents := m.SMTPReceipents

	smtpReceipentsString := strings.Join(smtpReceipents, ", ")

	message := fmt.Appendf([]byte{}, "To: %s\nSubject: %s\n\n%s", smtpReceipentsString, subject, htmlBody)

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

	err := m.Sender.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, smtpUser, smtpReceipents, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
