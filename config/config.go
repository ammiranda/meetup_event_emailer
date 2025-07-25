package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	JSONFilePath   string   `json:"json_file_path"`
	OpenAIAPIKey   string   `json:"openai_api_key"`
	SMTPHost       string   `json:"smtp_host"`
	SMTPPort       string   `json:"smtp_port"`
	SMTPUser       string   `json:"smtp_user"`
	SMTPPassword   string   `json:"smtp_password"`
	SMTPReceipents []string `json:"smtp_receipents"`
}

func LoadConfig() (*Config, error) {
	jsonFilePath, ok := os.LookupEnv("JSON_FILE_PATH")
	if !ok || jsonFilePath == "" {
		return nil, fmt.Errorf("JSON_FILE_PATH environment variable is not set")
	}

	openAIAPIKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok || openAIAPIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	smtpHost, ok := os.LookupEnv("SMTP_HOST")
	if !ok || smtpHost == "" {
		return nil, fmt.Errorf("SMTP_HOST environment variable is not set")
	}

	smtpPort, ok := os.LookupEnv("SMTP_PORT")
	if !ok || smtpPort == "" {
		return nil, fmt.Errorf("SMTP_PORT environment variable is not set")
	}

	smtpUser, ok := os.LookupEnv("SMTP_USER")
	if !ok || smtpUser == "" {
		return nil, fmt.Errorf("SMTP_USER environment variable is not set")
	}

	smtpPassword, ok := os.LookupEnv("SMTP_PASSWORD")
	if !ok || smtpPassword == "" {
		return nil, fmt.Errorf("SMTP_PASSWORD environment variable is not set")
	}

	smtpReceipentsRaw, ok := os.LookupEnv("SMTP_RECEIPENTS")
	if !ok || smtpReceipentsRaw == "" {
		return nil, fmt.Errorf("SMTP_RECEIPENTS environment variable is not set")
	}

	smtpReceipents := strings.Split(smtpReceipentsRaw, ",")

	return &Config{
		JSONFilePath:   jsonFilePath,
		OpenAIAPIKey:   openAIAPIKey,
		SMTPHost:       smtpHost,
		SMTPPort:       smtpPort,
		SMTPUser:       smtpUser,
		SMTPPassword:   smtpPassword,
		SMTPReceipents: smtpReceipents,
	}, nil
}
