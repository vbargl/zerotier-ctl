package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/vbargl/zerotier-ctl/internal/zerotier"
)

const configFile = "config.toml"

func GetConfigFile() string {
	return filepath.Join(getConfigDir(), configFile)
}

func Load() (*Config, error) {
	config, errFile := LoadFileDefault()
	if errFile == nil {
		return config, nil
	}

	config, errEnv := LoadEnv()
	if errEnv == nil {
		return config, nil
	}

	return nil, fmt.Errorf("could not load config: %v", errors.Join(errFile, errEnv))
}

func LoadFileDefault() (*Config, error) {
	return LoadFile(GetConfigFile())
}

func LoadFile(configFile string) (*Config, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s does not exist", configFile)
	}

	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	var config Config
	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding config file: %v", err)
	}

	if len(config.Controllers) == 0 {
		return nil, fmt.Errorf("no controllers defined in config file")
	}

	if _, ok := config.Controllers[config.Active]; ok {
		return &config, nil
	}

	if _, ok := config.Controllers["default"]; ok {
		config.Active = "default"
		return &config, nil
	}

	return nil, fmt.Errorf("no active controller defined in config file")
}

func LoadEnv() (*Config, error) {
	authSecret, err := os.ReadFile(zerotier.AuthFile)
	if err != nil {
		return nil, fmt.Errorf("could not read auth file %s: %v", zerotier.AuthFile, err)
	}

	ctlAddr := os.Getenv("ZT_CONTROLLER")
	if ctlAddr == "" {
		ctlAddr = "http://127.0.0.1:9993"
	}

	return &Config{
		Active: "env",
		Controllers: map[string]Controller{
			"env": {
				Address:   ctlAddr,
				AuthToken: string(authSecret),
			},
		},
	}, nil
}
