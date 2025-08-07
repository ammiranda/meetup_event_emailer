package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	EnvJSONFilePath   = "JSON_FILE_PATH"
	EnvOpenAIAPIKey   = "OPENAI_API_KEY"
	EnvSMTPHost       = "SMTP_HOST"
	EnvSMTPPort       = "SMTP_PORT"
	EnvSMTPUser       = "SMTP_USER"
	EnvSMTPPassword   = "SMTP_PASSWORD"
	EnvSMTPReceipents = "SMTP_RECEIPENTS"
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
	jsonFilePath, ok := os.LookupEnv(EnvJSONFilePath)
	if !ok || jsonFilePath == "" {
		return nil, fmt.Errorf("%s environment variable is not set", EnvJSONFilePath)
	}

	openAIAPIKey, ok := os.LookupEnv(EnvOpenAIAPIKey)
	if !ok || openAIAPIKey == "" {
		return nil, fmt.Errorf("%s environment variable is not set", EnvOpenAIAPIKey)
	}

	smtpHost, ok := os.LookupEnv(EnvSMTPHost)
	if !ok || smtpHost == "" {
		return nil, fmt.Errorf("%s environment variable is not set", EnvSMTPHost)
	}

	smtpPort, ok := os.LookupEnv(EnvSMTPPort)
	if !ok || smtpPort == "" {
		return nil, fmt.Errorf("%s environment variable is not set", EnvSMTPPort)
	}

	smtpUser, ok := os.LookupEnv(EnvSMTPUser)
	if !ok || smtpUser == "" {
		return nil, fmt.Errorf("%s environment variable is not set", EnvSMTPUser)
	}

	smtpPassword, ok := os.LookupEnv(EnvSMTPPassword)
	if !ok || smtpPassword == "" {
		return nil, fmt.Errorf("%s environment variable is not set", EnvSMTPPassword)
	}

	smtpReceipentsRaw, ok := os.LookupEnv(EnvSMTPReceipents)
	if !ok || smtpReceipentsRaw == "" {
		return nil, fmt.Errorf("%s environment variable is not set", EnvSMTPReceipents)
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
