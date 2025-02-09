package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	res "github.com/vbargl/zerotier-ctl/internal/res"
	zt "github.com/vbargl/zerotier-ctl/internal/zerotier"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [resource]",
	Short: "Delete a resource",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		respath := strings.Split(toComplete, sep)
		resource, args := shift(respath)

		client, err := newClient()
		handleError(cmd, err)

		switch resource {
		case "network":
			networks, err := res.ListNetworkIds(ctx, client)
			handleError(cmd, err)
			return presuf(networks, resource+sep, ""), cobra.ShellCompDirectiveNoFileComp

		case "member":
			network, _ := shift(args)

			switch {
			case network == "" || len(respath) == 2:
				networks, err := res.ListNetworkIds(ctx, client)
				handleError(cmd, err)
				return presuf(networks, resource+sep, sep), cobra.ShellCompDirectiveNoFileComp

			default:
				members, err := res.ListMemberIds(ctx, client, network)
				handleError(cmd, err)
				return presuf(members, resource+sep+network+sep, ""), cobra.ShellCompDirectiveNoFileComp
			}

		default:
			return []string{"network/", "member/"}, cobra.ShellCompDirectiveNoFileComp
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := newClient()
		handleError(cmd, err)

		respath, _ := shift(args)
		resource, args := shift(strings.Split(respath, sep))

		switch resource {
		case "network":
			network, _ := shift(args)
			if network == "" {
				cmd.PrintErrln("Network ID is required")
			}

			resp, err := res.DeleteNetwork(ctx, client, network)
			handleError(cmd, err)
			handleOutput(cmd, resp, func(w io.Writer, net *zt.ControllerNetwork) error {
				_, err := fmt.Fprintf(w, "Network with Id '%s' deleted\n", network)
				return err
			})

		case "member":
			network, args := shift(args)
			member, _ := shift(args)

			if network == "" {
				cmd.PrintErrln("Network ID is required")
				os.Exit(1)
			}

			if member == "" {
				cmd.PrintErrln("Member ID is required")
				os.Exit(1)
			}

			resp, err := res.DeleteMember(ctx, client, network, member)
			handleError(cmd, err)
			handleOutput(cmd, resp, func(w io.Writer, net *zt.ControllerNetworkMember) error {
				_, err := fmt.Fprintf(w, "Member with Id '%s' deleted from network '%s'\n", member, network)
				return err
			})

		default:
			cmd.Printf("Unknown resource: %s\n", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
