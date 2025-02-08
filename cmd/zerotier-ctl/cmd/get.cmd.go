package cmd

import (
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vbargl/zerotier-ctl/internal/prettyprint"
	res "github.com/vbargl/zerotier-ctl/internal/res"
	zt "github.com/vbargl/zerotier-ctl/internal/zerotier"
)

var getCmd = &cobra.Command{
	Use:   "get [resource]",
	Short: "Get details of a resource",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		respath := strings.Split(toComplete, sep)
		resource, args := shift(respath)

		client, err := newClient()
		handleError(cmd, err)

		switch resource {
		case "node":
			return nil, cobra.ShellCompDirectiveNoFileComp

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
			return []string{"node", "network/", "member/"}, cobra.ShellCompDirectiveNoFileComp
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := newClient()
		handleError(cmd, err)

		respath, _ := shift(args)
		resource, args := shift(strings.Split(respath, sep))
		ppCfg := prettyprint.TablePrintConfig{}

		switch resource {
		case "node":
			cmd.Println("Getting node details...")

		case "network":
			network, _ := shift(args)
			switch {
			case network != "":
				network, err := res.GetNetwork(ctx, client, network)
				handleError(cmd, err)
				_ = networkTable.Print(ppCfg, notNil(slices.Values([]*zt.ControllerNetwork{network})))

			default:
				networks, err := res.ListNetworks(ctx, client)
				handleError(cmd, err)
				_ = networkTable.Print(ppCfg, notNil(networks))
			}

		case "member":
			network, args := shift(args)
			member, _ := shift(args)

			if network == "" {
				cmd.PrintErrln("Network ID is required")
				os.Exit(1)
			}

			switch {
			case member != "":
				member, err := res.GetMember(ctx, client, network, member)
				handleError(cmd, err)
				_ = memberTable.Print(ppCfg, notNil(slices.Values([]*zt.ControllerNetworkMember{member})))

			default:
				members, err := res.ListMembers(ctx, client, network)
				handleError(cmd, err)
				_ = memberTable.Print(ppCfg, notNil(members))
			}

		default:
			cmd.Printf("Unknown resource: %s\n", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
