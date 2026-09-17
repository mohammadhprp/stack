package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Overridden at build time with
// -ldflags "-X github.com/mohammadhprp/stack/cmd.Version=<version>".
var Version = "dev"

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the stack version",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "stack %s\n", Version)
			return nil
		},
	}
}
