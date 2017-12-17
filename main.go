package main

import (
	"github.com/JohnBrainard/jwttools/tools"
	"os"
	"flag"
	"fmt"
)

type Command interface {
	Init()
	Validate() bool
	Usage()
	Execute()
}

func main() {
	jwtTools := tools.JwtToolsCreate()

	fmt.Printf("%s\n", jwtTools.Config.ToJson())

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	commands := map[string]Command{
		"generate": tools.GenerateCommandNew(jwtTools.Config),
		"info":     tools.InfoCommandNew(jwtTools.Config),
		"presets":  tools.PresetsCommandNew(jwtTools.Config),
		"edit":     tools.EditCommandNew(jwtTools),
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
