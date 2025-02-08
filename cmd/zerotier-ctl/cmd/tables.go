package cmd

import (
	"github.com/vbargl/zerotier-ctl/internal/prettyprint"
	zt "github.com/vbargl/zerotier-ctl/internal/zerotier"
)

var networkTable = prettyprint.NewTable[*zt.ControllerNetwork]()

func init() {
	networkTable.AddColumn("ID", func(n *zt.ControllerNetwork) prettyprint.Value {
		return prettyprint.Single(n.Id)
	})
	networkTable.AddColumn("Name", func(n *zt.ControllerNetwork) prettyprint.Value {
		return prettyprint.Single(n.Name)
	})
	networkTable.AddColumn("Private", func(n *zt.ControllerNetwork) prettyprint.Value {
		var private string
		if n.Private {
			private = "[x]"
		} else {
			private = "[ ]"
		}
		return prettyprint.Single(private)
	})
	networkTable.AddColumn("IP-Pool", func(n *zt.ControllerNetwork) prettyprint.Value {
		return prettyprint.Seq(func(yield func(string) bool) {
			for _, pool := range n.IpAssignmentPools {
				if !yield(pool.String()) {
					break
				}
			}
		})
	})
	networkTable.AddColumn("Routes", func(n *zt.ControllerNetwork) prettyprint.Value {
		return prettyprint.Seq(func(yield func(string) bool) {
			for _, route := range n.Routes {
				if !yield(route.String()) {
					break
				}
			}
		})
	})
}

var memberTable = prettyprint.NewTable[*zt.ControllerNetworkMember]()

func init() {
	memberTable.AddColumn("ID", func(m *zt.ControllerNetworkMember) prettyprint.Value {
		return prettyprint.Single(m.Id)
	})
	memberTable.AddColumn("Name", func(m *zt.ControllerNetworkMember) prettyprint.Value {
		name := ""
		if m.Name != nil {
			name = *m.Name
		}
		return prettyprint.Single(name)
	})
	memberTable.AddColumn("Authorized", func(m *zt.ControllerNetworkMember) prettyprint.Value {
		var authorized string
		if m.Authorized != nil && *m.Authorized {
			authorized = "[x]"
		} else {
			authorized = "[ ]"
		}
		return prettyprint.Single(authorized)
	})
	memberTable.AddColumn("IP-Addresses", func(m *zt.ControllerNetworkMember) prettyprint.Value {
		return prettyprint.Seq(func(yield func(string) bool) {
			for _, ip := range *m.IpAssignments {
				if !yield(ip.String()) {
					break
				}
			}
		})
	})

}
