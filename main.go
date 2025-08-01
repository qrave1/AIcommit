package main

import (
	"os"

	"github.com/qrave1/AIcommit/cmd"
)

var Version = "dev"

func main() {
	cmd.Version = Version

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
