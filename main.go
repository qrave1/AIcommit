package main

import (
	"os"

	"github.com/qrave1/AIcommit/cmd"
	"github.com/qrave1/AIcommit/cmd/version"
)

var Version = "dev"

func main() {
	version.Version = Version

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
