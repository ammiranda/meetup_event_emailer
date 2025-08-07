//go:build unit
// +build unit

package mailer

import (
	"context"
	"errors"
	"testing"

	"github.com/ammiranda/meetup_emailer/mailer/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MailerSuite struct {
	suite.Suite
	mailer     *Mailer
	mockSender *mocks.EmailSender
}

func (s *MailerSuite) SetupTest() {
	s.mockSender = mocks.NewEmailSender(s.T())
	s.mailer = &Mailer{
		SMTPHost:       "smtp.example.com",
		SMTPPort:       "587",
		SMTPUser:       "user@example.com",
		SMTPPassword:   "password",
		SMTPReceipents: []string{"to@example.com"},
		Sender:         s.mockSender,
	}
}

func (s *MailerSuite) TestSendHTMLEmail_Success() {
	s.mockSender.
		On("SendMail", "smtp.example.com:587", mock.Anything, "user@example.com", []string{"to@example.com"}, mock.MatchedBy(func(msg []byte) bool {
			return string(msg) == "To: to@example.com\nSubject: Test Subject\n\n<b>HTML</b>"
		})).
		Return(nil).
		Once()

	err := s.mailer.SendHTMLEmail(context.Background(), "Test Subject", "<b>HTML</b>")
	s.NoError(err)
}

func (s *MailerSuite) TestSendHTMLEmail_Error() {
	s.mockSender.
		On("SendMail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("smtp error")).
		Once()

	err := s.mailer.SendHTMLEmail(context.Background(), "Test Subject", "<b>HTML</b>")
	s.Error(err)
	s.Contains(err.Error(), "smtp error")
}

func TestMailerSuite(t *testing.T) {
	suite.Run(t, new(MailerSuite))
}
