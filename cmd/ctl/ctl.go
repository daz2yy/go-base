package main

import (
	"os"

	"github.com/daz2yy/go-base/internal/ctl/cmd"
)

func main() {
	command := cmd.NewDefaultCtlCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
