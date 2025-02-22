package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	res "github.com/vbargl/zerotier-ctl/internal/res"
	zt "github.com/vbargl/zerotier-ctl/internal/zerotier"
)

var createCmd = &cobra.Command{
	Use:   "create [resource]",
	Short: "Create a new resource",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		resPath := strings.Split(toComplete, sep)
		resource, _ := shift(resPath)

		switch resource {
		case "network":
			return nil, cobra.ShellCompDirectiveNoFileComp

		default:
			return []string{"network"}, cobra.ShellCompDirectiveNoFileComp
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := newClient()
		handleError(cmd, err)

		resPath := strings.Split(args[0], sep)
		resource, _ := shift(resPath)

		switch resource {
		case "network":
			resp, err := res.CreateNetwork(ctx, client, zt.ControllerNetworkRequest{})
			handleError(cmd, err)
			handleOutput(cmd, resp, func(w io.Writer, net *zt.ControllerNetwork) error {
				_, err := fmt.Fprintf(w, "Network ID: %s\n", net.Id)
				return err
			})

		default:
			handleError(cmd, fmt.Errorf("invalid resource: %s", resource))
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	flags := createCmd.Flags()
	flags.StringP("file", "f", "", "file to use for creating the resource")

	err := createCmd.MarkFlagFilename("file")
	if err != nil {
		panic(err)
	}
}
