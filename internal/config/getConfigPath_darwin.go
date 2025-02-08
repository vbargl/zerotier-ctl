//go:build darwin

package config

import "os"

const (
	configDir = "$HOME/Library/Application Support/zerotier-ctl"
)

func getConfigDir() string {
	return os.ExpandEnv(configDir)
}
