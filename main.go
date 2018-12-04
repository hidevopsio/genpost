package main

import (
	"github.com/hidevopsio/genpost/cmd"
	"hidevops.io/hiboot/pkg/app/cli"
	"hidevops.io/hiboot/pkg/starter/logging"
)

func main() {
	// create new cli application and run it
	cli.NewApplication(cmd.NewRootCommand).
		SetProperty(logging.Level, logging.LevelWarn).
		Run()
}