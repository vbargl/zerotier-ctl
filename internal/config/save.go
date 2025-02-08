package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func Save(cfg *Config) error {
	configPath := filepath.Join(getConfigDir(), configFile)
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	return encoder.Encode(cfg)
}
