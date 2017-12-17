package tools

import (
	"os"
	"fmt"
	"path"
	"time"
)

type JwtTools struct {
	Config     *Config
	ConfigPath string
}

func JwtToolsCreate() *JwtTools {
	configPath := ConfigFindFile()
	config, err := ConfigLoad(configPath)

	if os.IsNotExist(err) {
		fmt.Printf("Cannot find config. Skipping presets")
	} else {
		CheckError(err)
	}

	return &JwtTools{
		&config,
		configPath,
	}
}

func (tools *JwtTools) SaveConfig() {
	tools.createBackup()

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
