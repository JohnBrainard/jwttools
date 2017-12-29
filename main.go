package main

import (
	"github.com/JohnBrainard/jwttools/tools"
	"os"
	"flag"
)

type Command interface {
	Init()
	Validate() bool
	Usage()
	Execute()
}

func main() {
	jwtTools := tools.JwtToolsCreate()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	commands := map[string]Command{
		"edit":     tools.EditCommandNew(jwtTools),
		"generate": tools.GenerateCommandNew(jwtTools.Config),
		"info":     tools.InfoCommandNew(jwtTools.Config),
		"init":     tools.InitCommandNew(jwtTools),
		"presets":  tools.PresetsCommandNew(jwtTools.Config),
	}

	command, exists := commands[os.Args[1]]
	if exists {
		command.Init()
		if !command.Validate() {
			command.Usage()
			os.Exit(1)
		}

		command.Execute()
	} else {
		flag.Usage()
	}
}
