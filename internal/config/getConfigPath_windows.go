//go:build windows

package config

import "os"

const (
	configDir = "$APPDATA/zerotier-ctl/config.toml"
)

func getConfigDir() string {
	return os.ExpandEnv(configDir)
}
