package tools

import (
	"time"
	"os"
	"os/user"
	path2 "path"
	"io/ioutil"
	json2 "encoding/json"
)

type Preset struct {
	Claims      map[string]interface{} `json:"claims"`
	Key         string                 `json:"key"`
	Expires     string                 `json:"expires"`
	Description string                 `json:"info"`
}

type Config struct {
	Presets map[string]Preset `json:"presets"`
}

func ConfigLoadDefault() (config Config, err error) {
	currentUser, _ := user.Current()

	path := path2.Join(currentUser.HomeDir, ".jwttools", "config.json")

	return ConfigLoad(path)
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

func (config Config) GetPreset(name string) Preset {
	return config.Presets[name]
}

func (config Config) HasPreset(name string) bool {
	_, ok := config.Presets[name]
	return ok
}

func (preset Preset) GetExpiration() (expires time.Duration, err error) {
	expires, err = time.ParseDuration(preset.Expires)
	return
}
