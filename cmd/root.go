package cmd

import (
	"github.com/spf13/cobra"
)

var cfgPath string

var rootCmd = &cobra.Command{
	Use:   "ai-commiter",
	Short: "Generate git commit messages via LLM",
	Run:   commitCmdRun,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&cfgPath,
		"config",
		"config.json",
		"Path to config file",
	)
}
