package main

import (
	"os"

	"github.com/qrave1/AIcommit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
