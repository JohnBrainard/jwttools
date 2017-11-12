package tools

import (
	"flag"
	"os"
	"fmt"
	"github.com/SermoDigital/jose/jws"
)

type InfoCommand struct {
	Token string

	Config  *Config
	flagSet *flag.FlagSet
}

func InfoCommandNew(config *Config) *InfoCommand {
	var command = InfoCommand{}

	command.flagSet = flag.NewFlagSet("info", flag.ExitOnError)
	command.flagSet.StringVar(&command.Token, "token", "", "Token")

	command.Config = config

	return &command
}

func (command *InfoCommand) Init() {
	err := command.flagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (command *InfoCommand) Validate() bool {
	return true
}

func (command *InfoCommand) Usage() {
	command.flagSet.Usage()
}

func (command *InfoCommand) Execute() {
	token, err := jws.ParseJWT([]byte(command.Token))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	expiration, _ := token.Claims().Expiration()
	fmt.Printf("Expiration: %s\n", expiration)

	fmt.Println("Claims:")
	for key, value := range token.Claims() {
		fmt.Printf("  %s = %s\n", key, value)
	}
}
