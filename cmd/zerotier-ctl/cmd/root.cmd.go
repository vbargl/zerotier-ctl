package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zerotier-ctl",
	Short: "ZeroTier Controller CLI",
}

func init() {
	globalFlags := rootCmd.PersistentFlags()
	globalFlags.BoolP("json", "j", false, "JSON output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
