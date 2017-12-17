package tools

import (
	"time"
	"os"
	"os/user"
	path2 "path"
	"io/ioutil"
	json2 "encoding/json"
	"fmt"
)

type Preset struct {
	Description string                 `json:"info"`
	Claims      map[string]interface{} `json:"claims"`
	Key         string                 `json:"key"`
	Expires     string                 `json:"expires"`
}

type Config struct {
	Presets map[string]Preset `json:"presets"`
}

func ConfigFindFile() string {
	currentUser, _ := user.Current()
	return path2.Join(currentUser.HomeDir, ".jwttools", "config.json")
}

func ConfigLoad(path string) (config Config, err error) {
	if _, err = os.Stat(path); err != nil {
		return
	}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	json2.Unmarshal(contents, &config)
	return
}

func (config *Config) GetPreset(name string) Preset {
	return config.Presets[name]
}

func (config *Config) CreatePreset(name string) Preset {
	preset := Preset{}

	preset.Claims = make(map[string]interface{})

	preset.Claims["iss"] = "issuer"
	preset.Expires = "24h"
	preset.Description = fmt.Sprintf("Token: %s", name)

	return preset
}

func (config *Config) ParsePreset(json []byte) *Preset {
	preset := Preset{}
	err := json2.Unmarshal(json, &preset)
	CheckError(err)

	return &preset
}

func (config *Config) HasPreset(name string) bool {
	_, ok := config.Presets[name]
	return ok
}

func (config *Config) ToJson() []byte {
	bytes, err := json2.MarshalIndent(config, "", "\t")
	CheckError(err)

	return bytes
}

func (preset *Preset) GetExpiration() (expires time.Duration, err error) {
	expires, err = time.ParseDuration(preset.Expires)
	return
}

func (preset *Preset) ToJson() []byte {
	bytes, err := json2.MarshalIndent(preset, "", "\t")
	CheckError(err)

	return bytes
}
