//go:build unit
// +build unit

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	testJSONFilePath   = "/tmp/events.json"
	testOpenAIAPIKey   = "test-key"
	testSMTPHost       = "smtp.example.com"
	testSMTPPort       = "587"
	testSMTPUser       = "user@example.com"
	testSMTPPassword   = "password"
	testSMTPReceipents = "to1@example.com,to2@example.com"
)

type ConfigSuite struct {
	suite.Suite
	originalEnv map[string]string
}

func (s *ConfigSuite) SetupTest() {
	s.originalEnv = map[string]string{
		EnvJSONFilePath:   os.Getenv(EnvJSONFilePath),
		EnvOpenAIAPIKey:   os.Getenv(EnvOpenAIAPIKey),
		EnvSMTPHost:       os.Getenv(EnvSMTPHost),
		EnvSMTPPort:       os.Getenv(EnvSMTPPort),
		EnvSMTPUser:       os.Getenv(EnvSMTPUser),
		EnvSMTPPassword:   os.Getenv(EnvSMTPPassword),
		EnvSMTPReceipents: os.Getenv(EnvSMTPReceipents),
	}
}

func (s *ConfigSuite) TearDownTest() {
	for k, v := range s.originalEnv {
		os.Setenv(k, v)
	}
}

func (s *ConfigSuite) setAllEnv() {
	os.Setenv(EnvJSONFilePath, testJSONFilePath)
	os.Setenv(EnvOpenAIAPIKey, testOpenAIAPIKey)
	os.Setenv(EnvSMTPHost, testSMTPHost)
	os.Setenv(EnvSMTPPort, testSMTPPort)
	os.Setenv(EnvSMTPUser, testSMTPUser)
	os.Setenv(EnvSMTPPassword, testSMTPPassword)
	os.Setenv(EnvSMTPReceipents, testSMTPReceipents)
}

func (s *ConfigSuite) TestLoadConfig_Success() {
	s.setAllEnv()
	cfg, err := LoadConfig()
	s.NoError(err)
	s.Equal(testJSONFilePath, cfg.JSONFilePath)
	s.Equal(testOpenAIAPIKey, cfg.OpenAIAPIKey)
	s.Equal(testSMTPHost, cfg.SMTPHost)
	s.Equal(testSMTPPort, cfg.SMTPPort)
	s.Equal(testSMTPUser, cfg.SMTPUser)
	s.Equal(testSMTPPassword, cfg.SMTPPassword)
	s.Equal([]string{"to1@example.com", "to2@example.com"}, cfg.SMTPReceipents)
}

func (s *ConfigSuite) TestLoadConfig_MissingEnv() {
	os.Unsetenv(EnvJSONFilePath)
	os.Setenv(EnvOpenAIAPIKey, testOpenAIAPIKey)
	os.Setenv(EnvSMTPHost, testSMTPHost)
	os.Setenv(EnvSMTPPort, testSMTPPort)
	os.Setenv(EnvSMTPUser, testSMTPUser)
	os.Setenv(EnvSMTPPassword, testSMTPPassword)
	os.Setenv(EnvSMTPReceipents, testSMTPReceipents)
	cfg, err := LoadConfig()
	s.Nil(cfg)
	s.Error(err)
	s.Contains(err.Error(), EnvJSONFilePath)
}

func (s *ConfigSuite) TestLoadConfig_MissingReceipents() {
	s.setAllEnv()
	os.Unsetenv(EnvSMTPReceipents)
	cfg, err := LoadConfig()
	s.Nil(cfg)
	s.Error(err)
	s.Contains(err.Error(), EnvSMTPReceipents)
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}
