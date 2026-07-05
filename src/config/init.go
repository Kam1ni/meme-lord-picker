package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

func getConfigDirPath() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		homeDir := os.Getenv("HOME")
		if homeDir != "" {
			configHome = path.Join(homeDir, ".config")
		}
	}
	if configHome == "" {
		panic("Failed to get config home")
	}
	return path.Join(configHome, "meme-lord-picker")
}

func getConfigFilePath() string {
	return path.Join(getConfigDirPath(), "config.json")
}

func ensureDirExists(dir string) {
	pathParts := strings.Split(dir, "/")
	currentPath := "/"
	for _, part := range pathParts {
		currentPath = path.Join(currentPath, part)
		_, err := os.Stat(currentPath)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.Mkdir(currentPath, os.ModePerm)
				if err != nil {
					panic(fmt.Sprintf("Failed to create directory path %s\n%s", currentPath, err.Error()))
				}
			} else {
				panic(fmt.Sprintf("Failed to ensure directory exists (%s).\nFailed to get stats of %s", dir, currentPath))
			}
		}
	}
}

func ensureConfigFileExists() {
	ensureDirExists(getConfigDirPath())

	_, err := os.Stat(getConfigFilePath())
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		defaultConfigJson, err := json.MarshalIndent(getDefaultConfig(), "", "\t")
		if err != nil {
			panic(fmt.Sprintf("Failed to create default config json\n%s", err.Error()))
		}
		err = os.WriteFile(getConfigFilePath(), defaultConfigJson, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Failed to create config file\n%s", err.Error()))
		}
		fmt.Printf("Created default config file at %s\nPlease configure memeLordApiUrl and memeLordApiToken\n", getConfigFilePath())
		os.Exit(1)
	} else {
		panic(fmt.Sprintf("Failed to get stats of config file %s\n%s", getConfigFilePath(), err.Error()))
	}
}

func loadConfig() {
	ensureConfigFileExists()
	configBytes, err := os.ReadFile(getConfigFilePath())
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file %s\n%s", getConfigFilePath(), err.Error()))
	}
	err = json.Unmarshal(configBytes, &loadedConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal config file %s\n%s", getConfigFilePath(), err.Error()))
	}
	loadedConfig.Window.enforceLimits()
}

func init() {
	loadConfig()
}
