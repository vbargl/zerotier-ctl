//go:build !darwin && !windows

package config

import "os"

const (
	configDir = "$XDG_CONFIG_HOME/zerotier-ctl"
)

func getConfigDir() string {
	return os.ExpandEnv(configDir)
}
