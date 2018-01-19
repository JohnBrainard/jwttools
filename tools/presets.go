package tools

import (
	"flag"
	"os"
	"fmt"
	"encoding/json"
)

type PresetsCommand struct {
	Config *Config

	keys    bool
	preset  string
	verbose bool

	flagSet *flag.FlagSet
}

func PresetsCommandNew(config *Config) *PresetsCommand {
	var command = PresetsCommand{}

	command.flagSet = flag.NewFlagSet("presets", flag.ExitOnError)
	command.flagSet.BoolVar(&command.keys, "keys", false, "Show keys [default false]")
	command.flagSet.StringVar(&command.preset, "preset", "", "Display preset")
	command.flagSet.BoolVar(&command.verbose, "verbose", false, "Display extra information")

	command.Config = config

	return &command
}

func (command *PresetsCommand) Init() {
	err := command.flagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (command *PresetsCommand) Validate() bool {
	if len(command.preset) > 0 && !command.Config.HasPreset(command.preset) {
		fmt.Printf("Preset not found: %s\n", command.preset)
		return false
	}

	return true
}

func (command *PresetsCommand) Usage() {
	command.flagSet.Usage()
}

func (command *PresetsCommand) Execute() {
	var presets map[string]Preset

	if len(command.preset) > 0 {
		presets = map[string]Preset{
			command.preset: command.Config.GetPreset(command.preset),
		}
	} else {
		presets = command.Config.Presets
	}

	for key, config := range presets {
		fmt.Println(config.Description)
		fmt.Printf("  Preset: %s\n", key)
		fmt.Printf("  Issuer: %s\n", config.Claims["iss"])
		fmt.Printf("  Expires: %s\n", config.Expires)

		if command.keys {
			fmt.Printf("  Key: %s\n", config.Key)
		}

		if command.verbose {
			claimsJson, _ := json.MarshalIndent(config.Claims, "  ", "  ")
			fmt.Printf("  Claims:\n  %s\n", claimsJson)
		}

		fmt.Println()
	}
}
