package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [resource] [key]=[value]",
	Short: "Set a property of a resource",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		resource := args[0]
		id := args[1]
		keyValue := args[2]
		switch resource {
		case "network":
			fmt.Printf("Setting %s for network %s...\n", keyValue, id)
			// Add logic to set network property
		case "member":
			if len(args) > 2 {
				netID := args[1]
				memberID := args[2]
				keyValue := args[3]
				fmt.Printf("Setting %s for member %s in network %s...\n", keyValue, memberID, netID)
				// Add logic to set member property
			} else {
				fmt.Println("Invalid arguments for member resource")
			}
		default:
			fmt.Printf("Unknown resource: %s\n", resource)
		}
	},
}

func init() {
	// rootCmd.AddCommand(setCmd)
}
