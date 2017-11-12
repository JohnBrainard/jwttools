package main

import (
	"fmt"
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
	config, err := tools.ConfigLoadDefault()
	if os.IsNotExist(err) {
		fmt.Printf("Cannot find config, skipping loading presets")
	} else {
		checkError(err)
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	commands := map[string]Command{
		"generate": tools.GenerateCommandNew(&config),
		"info":     tools.InfoCommandNew(&config),
		"presets":  tools.PresetsCommandNew(&config),
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
