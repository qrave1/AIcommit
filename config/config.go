package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	OpenAI struct {
		URL         string  `json:"url"`
		APIKeyEnv   string  `json:"api_key_env"`
		Model       string  `json:"model"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float64 `json:"temperature"`
	} `json:"openai"`

	Style struct {
		Template  string `json:"template"`
		MaxLength int    `json:"max_length"`
	} `json:"style"`
}

func LoadJSONConfig(path string) (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(filepath.Join(homeDir, path))
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err = json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
