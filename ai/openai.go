package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAI struct {
	Client *openai.Client
}

// NewOpenAI creates a new OpenAI client
func NewOpenAI(ctx context.Context, apiKey string) *OpenAI {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	return &OpenAI{Client: &client}
}

// GenerateEmail generates email markup based on the data provided
func (o *OpenAI) GenerateEmailBody(ctx context.Context, data string) (string, error) {
	prompt := generatePrompt(ctx, data)

	response, err := o.Client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT3_5Turbo,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Seed: openai.Int(0),
	})
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}

// generatePrompt generates a prompt for the OpenAI API
func generatePrompt(ctx context.Context, data string) string {
	prompt := fmt.Sprintf(`
		Please compose an email that contains the relevant events in this data: %s
		
		The email is for an out-of-work software engineer looking for networking and business opportunities.
		
		Requirements:
		- Include event links to make navigation easy
		- List events in chronological order  
		- Provide brief explanations for why each event was included
		- Use ONLY standard ASCII characters (no smart quotes, em dashes, special punctuation)
		- Use regular quotes ("), hyphens (-), and normal spaces
		- Keep the tone professional but friendly
	`, data)

	return prompt
}
