package tools

import (
	"os"
)

type InitCommand struct {
	JwtTools *JwtTools
}

func InitCommandNew(tools *JwtTools) *InitCommand {
	var command = InitCommand{
		JwtTools: tools,
	}

	return &command
}

func (command *InitCommand) Init() {
}

func (command *InitCommand) Validate() bool {
	_, err := os.Stat(command.JwtTools.ConfigPath)
	if os.IsNotExist(err) {
		return true
	}

	return false
}

func (command *InitCommand) Usage() {
}

func (command *InitCommand) Execute() {
	command.JwtTools.SaveConfig()
}
