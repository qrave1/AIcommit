package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "AIcommit",
	Short:   "Generate git commit messages via LLM",
	Run:     CommitCmdRun,
	Version: Version,
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	flags := RootCmd.Flags()
	flags.StringP(
		"config",
		"c",
		"aicommit_config.json",
		"Path to config file",
	)
}
