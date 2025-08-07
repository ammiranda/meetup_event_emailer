//go:build unit
// +build unit

// Unit tests for OpenAI logic

package ai

import (
	"context"
	"testing"

	"github.com/ammiranda/meetup_emailer/ai/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type OpenAISuite struct {
	suite.Suite
	ai            *OpenAI
	mockCompleter *mocks.Completer
}

func (s *OpenAISuite) SetupTest() {
	s.mockCompleter = mocks.NewCompleter(s.T())
	s.ai = &OpenAI{
		Completer: s.mockCompleter,
	}
}

func (s *OpenAISuite) TestGeneratePrompt() {
	ctx := context.Background()
	data := `[{"title": "Test Event"}]`
	prompt := generatePrompt(ctx, data)

	s.NotEmpty(prompt)
	s.Contains(prompt, "Test Event")
	s.Contains(prompt, "Please compose an email")
}

func (s *OpenAISuite) TestGenerateEmailBody_Mock() {
	s.mockCompleter.
		On("GenerateCompletion", mock.Anything, mock.AnythingOfType("string")).
		Return("MOCKED EMAIL BODY", nil).
		Once()
	result, err := s.ai.GenerateEmailBody(context.Background(), "test data")
	s.NoError(err)
	s.Equal("MOCKED EMAIL BODY", result)
	s.mockCompleter.AssertExpectations(s.T())
}

func TestOpenAISuite(t *testing.T) {
	suite.Run(t, new(OpenAISuite))
}
