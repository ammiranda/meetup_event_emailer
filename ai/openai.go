package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Completer interface {
	GenerateCompletion(ctx context.Context, prompt string) (string, error)
}

type OpenAI struct {
	Completer Completer
}

type RealCompleter struct {
	client *openai.Client
}

func NewOpenAI(ctx context.Context, apiKey string) *OpenAI {
	client := openai.NewClient(option.WithAPIKey(apiKey))
	return &OpenAI{
		Completer: &RealCompleter{client: &client},
	}
}

func (r *RealCompleter) GenerateCompletion(ctx context.Context, prompt string) (string, error) {
	params := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT3_5Turbo,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Seed: openai.Int(0),
	}
	resp, err := r.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

// GenerateEmail generates email markup based on the data provided
func (o *OpenAI) GenerateEmailBody(ctx context.Context, data string) (string, error) {
	prompt := generatePrompt(ctx, data)
	return o.Completer.GenerateCompletion(ctx, prompt)
}

// generatePrompt generates a prompt for the OpenAI API
func generatePrompt(ctx context.Context, data string) string {
	prompt := fmt.Sprintf(`
		Please compose an email that contains the relevant events in this data: %s
		
		The email is for a male software engineer looking for networking and business opportunities.
		Focus on technology focused events that can be technical in nature or socializing with others
		in the tech industry.
		
		Requirements:
		- Include event links to make navigation easy
		- List events in chronological order based on when they will occur and include the date and time for each event included
		- Use HTML tags where relevant, (there is no Markdown parsing) like for the links as an example
		- Provide brief explanations for why each event was included
		- Keep the tone professional but friendly
		- Do not include a closing that needs a name or if you do use "Your friendly AI"
	`, data)

	return prompt
}
