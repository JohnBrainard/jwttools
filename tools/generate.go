package tools

import (
	"flag"
	"time"
	"os"
	"fmt"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/crypto"
)

type GenerateCommand struct {
	Preset   string
	Key      string
	Issuer   string
	Subject  string
	Audience string
	Claims   string
	Expires  time.Duration

	PrintHeader bool

	Config  *Config
	flagSet *flag.FlagSet
}

func GenerateCommandNew(config *Config) *GenerateCommand {
	var command = GenerateCommand{}

	command.flagSet = flag.NewFlagSet("generate", flag.ExitOnError)
	command.flagSet.StringVar(&command.Preset, "preset", "", "Preset name")
	command.flagSet.StringVar(&command.Key, "key", "", "Key")
	command.flagSet.StringVar(&command.Issuer, "iss", "", "Issuer")
	command.flagSet.StringVar(&command.Audience, "aud", "", "Audience")
	command.flagSet.StringVar(&command.Subject, "sub", "", "Subject")
	command.flagSet.DurationVar(&command.Expires, "exp", 0, "Expires")

	command.flagSet.BoolVar(&command.PrintHeader, "header", false, "Print Header")
	command.flagSet.BoolVar(&command.PrintHeader, "h", false, "Print Header (short version")

	command.Config = config

	return &command
}

func (command *GenerateCommand) Init() {
	err := command.flagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (command *GenerateCommand) Validate() bool {
	if len(command.Preset) > 0 {
		if !command.Config.HasPreset(command.Preset) {
			fmt.Printf("Preset not found: %s\n", command.Preset)
			return false
		}
	}

	return true
}

func (command *GenerateCommand) Usage() {
	command.flagSet.Usage()
}

func (command *GenerateCommand) Execute() {
	var claims = jws.Claims{}

	var key []byte = []byte(command.Key)

	if len(command.Preset) > 0 {
		preset := command.Config.GetPreset(command.Preset)

		for key, value := range preset.Claims {
			claims.Set(key, value)
		}

		expires, _ := preset.GetExpiration()
		claims.SetExpiration(time.Now().Add(expires))

		if len(preset.Key) > 0 {
			key = []byte(preset.Key)
		}
	}

	if len(command.Issuer) > 0 {
		claims.SetIssuer(command.Issuer)
	}

	if len(command.Subject) > 0 {
		claims.SetSubject(command.Subject)
	}

	if len(command.Audience) > 0 {
		claims.SetAudience(command.Audience)
	}

	if command.Expires > 0 {
		claims.SetExpiration(time.Now().Add(command.Expires))
	}

	jwt := jws.NewJWT(claims, crypto.SigningMethodHS256)
	serialized, err := jwt.Serialize(key)
	if err != nil {
		panic(err)
	}

	if command.PrintHeader {
		fmt.Print("Authorization: Bearer ")
	}
	fmt.Printf("%s\n", serialized)
}
