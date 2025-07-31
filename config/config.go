package config

import (
	"encoding/json"
	"os"
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
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err = json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
