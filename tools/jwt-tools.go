package tools

import (
	"os"
	"fmt"
	"path"
	"time"
	"path/filepath"
)

type JwtTools struct {
	Config     *Config
	ConfigPath string
}

func JwtToolsCreate() *JwtTools {
	configPath := ConfigFindFile()
	config, err := ConfigLoad(configPath)

	if err != nil && !os.IsNotExist(err) {
		CheckError(err)
	}

	tools := JwtTools{
		&config,
		configPath,
	}

	if tools.Config == nil {
		presets := make(map[string]Preset)
		tools.Config = &Config{
			Presets: presets,
		}
	}

	return &tools
}

func (tools *JwtTools) SaveConfig() {
	configDir := filepath.Dir(tools.ConfigPath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0700)
	} else {
		tools.createBackup()
	}

	configJson := tools.Config.ToJson()

	configFile, err := os.OpenFile(tools.ConfigPath, os.O_RDWR|os.O_CREATE, 0644)
	CheckError(err)

	defer configFile.Close()
	configFile.Write(configJson)
}

func (tools *JwtTools) createBackup() {
	dirName := path.Dir(tools.ConfigPath)

	now := time.Now()
	year, month, day := now.Date()

	backupFileName := fmt.Sprintf("config.%d%d%d_%d%d%d.json", year, month, day, now.Hour(), now.Minute(), now.Second())
	fmt.Printf("Creating config backup: %s\n", path.Join(dirName, backupFileName))
	os.Rename(tools.ConfigPath, path.Join(dirName, backupFileName))
}
