package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func printCompletition(cmd *cobra.Command, args []string) {
	toComplete := ""
	if len(args) > 0 {
		toComplete = args[len(args)-1]
	}

	completion, _ := cmd.ValidArgsFunction(cmd, args, toComplete)
	for _, v := range completion {
		fmt.Println(v)
	}
	os.Exit(0)
}
