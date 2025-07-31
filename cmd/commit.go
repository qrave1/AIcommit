package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"

	"github.com/qrave1/AIcommit/config"
)

var ErrNoChangesFound = fmt.Errorf("no staged changes found")

func commitCmdRun(cmd *cobra.Command, _ []string) {
	ctx := cmd.Context()

	cfg, err := config.LoadJSONConfig(cfgPath)
	if err != nil {
		log.Fatal("failed to load json config:", err)
	}

	diff, err := getGitDiff()
	switch {
	case err == nil:
	case errors.Is(err, ErrNoChangesFound):
		fmt.Println("no staged changes found. Exiting.")
		return
	default:
		log.Fatal("failed to get git diff:", err)
	}

	prompt, err := preparePrompt(diff, cfg)

	llm, err := createLLM(cfg)
	if err != nil {
		log.Fatal("failed to create OpenAI LLM:", err)
	}

	llmAnswer, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		prompt,
		llms.WithMaxTokens(cfg.OpenAI.MaxTokens),
		llms.WithTemperature(cfg.OpenAI.Temperature),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Suggested commit message:\n\n" + llmAnswer + "\n")
	fmt.Print("Commit? [y/N]: ")

	var userAgreement string

	_, err = fmt.Scanln(&userAgreement)
	if err != nil {
		log.Fatalf("failed to read user agreement: %v", err)
	}

	if strings.ToLower(userAgreement) == "y" {
		err = exec.Command("git", "commit", "-m", llmAnswer).Run()
		if err != nil {
			log.Fatal("git commit:", err)
		}
		fmt.Println("Committed!")
	} else {
		fmt.Println("Aborted.")
	}
}

func getGitDiff() (string, error) {
	diffByte, err := exec.Command("git", "diff", "--staged").Output()
	if err != nil {
		return "", fmt.Errorf("failed to exec git diff: %w", err)
	}

	if strings.TrimSpace(string(diffByte)) == "" {
		return "", ErrNoChangesFound
	}

	return string(diffByte), nil
}

func preparePrompt(diff string, cfg *config.Config) (string, error) {
	promptTemplate := prompts.NewPromptTemplate(cfg.Style.Template, []string{"diff, max_length"})

	prompt, err := promptTemplate.Format(
		map[string]any{
			"max_length": cfg.Style.MaxLength,
			"diff":       diff,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to format prompt: %w", err)
	}

	return prompt, nil
}

func createLLM(cfg *config.Config) (llms.Model, error) {
	apiToken := os.Getenv(cfg.OpenAI.APIKeyEnv)

	if apiToken == "" {
		return nil, fmt.Errorf("API key environment variable %q is not set", cfg.OpenAI.APIKeyEnv)
	}

	opts := []openai.Option{
		openai.WithToken(apiToken),
		openai.WithModel(cfg.OpenAI.Model),
	}

	if cfg.OpenAI.URL != "" {
		opts = append(opts, openai.WithBaseURL(cfg.OpenAI.URL))
	}

	llm, err := openai.New(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAI LLM: %w", err)
	}

	return llm, nil
}
