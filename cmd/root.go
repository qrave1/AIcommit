package cmd

import (
	"github.com/qrave1/AIcommit/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "AIcommit",
	Short:   "Generate git commit messages via LLM",
	Run:     commitCmdRun,
	Version: version.Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	flags := rootCmd.Flags()
	flags.StringP(
		"config",
		"c",
		"aicommit_config.json",
		"Path to config file",
	)
}
