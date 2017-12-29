package tools

import (
	"flag"
	"os"
	"fmt"
	"io/ioutil"
)

type EditCommand struct {
	Preset string

	JwtTools *JwtTools
	flagSet  *flag.FlagSet
}

func EditCommandNew(tools *JwtTools) *EditCommand {
	var command = EditCommand{}

	command.JwtTools = tools

	command.flagSet = flag.NewFlagSet("edit", flag.ExitOnError)
	command.flagSet.StringVar(&command.Preset, "preset", "", "Preset name")

	return &command
}

func (command *EditCommand) Init() {
	err := command.flagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (command *EditCommand) Validate() bool {
	if len(command.Preset) == 0 {
		fmt.Print("Specify preset name with -preset flag")
		return false
	}

	return true
}

func (command *EditCommand) Usage() {
	command.flagSet.Usage()
}

func (command *EditCommand) Execute() {
	fmt.Printf("Editing %s\n", command.Preset)

	var preset Preset
	if command.JwtTools.Config.HasPreset(command.Preset) {
		preset = command.JwtTools.Config.GetPreset(command.Preset)
	} else {
		preset = command.JwtTools.Config.CreatePreset(command.Preset)
	}

	tokenFile, _ := ioutil.TempFile("", "jwttools-token-")
	tokenFile.Write(preset.ToJson())
	tokenFile.Close()

	defer os.Remove(tokenFile.Name())

	fmt.Printf("Token file: %s\n", tokenFile.Name())
	presetJson := command.editFile(tokenFile.Name())

	fmt.Printf("New preset: %s\n", presetJson)

	newPreset := command.JwtTools.Config.ParsePreset(presetJson)
	command.JwtTools.Config.AddPreset(command.Preset, *newPreset)

	command.JwtTools.SaveConfig()
}

func (command *EditCommand) editFile(fileName string) []byte {
	cmd := GetSystemEditor(fileName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	CheckError(err)

	err = cmd.Wait()
	CheckError(err)

	file, err := os.Open(fileName)
	CheckError(err)
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	CheckError(err)

	return text
}
