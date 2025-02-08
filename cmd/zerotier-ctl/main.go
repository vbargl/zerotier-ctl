package main

import (
	"os"

	cmd "github.com/vbargl/zerotier-ctl/cmd/zerotier-ctl/cmd"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: zerotier-ctl <command>")
		os.Exit(1)
	}
	cmd.Execute()
}
