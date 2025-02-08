package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/vbargl/zerotier-ctl/internal/config"
	"github.com/vbargl/zerotier-ctl/internal/zerotier"
)

var ctx = context.Background()

func newClient() (*zerotier.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
		},
	}

	ctl := cfg.ActiveController()
	return zerotier.NewClient(ctl.Address,
		zerotier.WithHTTPClient(httpClient),
		zerotier.WithAuthToken(ctl.AuthToken),
	)
}

func handleError(cmd *cobra.Command, err error) {
	wantJsonOutput, _ := cmd.Flags().GetBool("json")

	switch {
	case err != nil && wantJsonOutput:
		_ = json.NewEncoder(os.Stderr).Encode(map[string]error{"error": err})
		os.Exit(1)
	case err != nil:
		_, _ = fmt.Fprintf(os.Stderr, "error during execution of %s %v: %v\n", cmd.Name(), cmd.Args, err)
		os.Exit(1)
	}
}
