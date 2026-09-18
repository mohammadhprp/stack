package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/x/term"
	"github.com/spf13/cobra"

	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/tui"
)

func NewRootCommand(cat *install.Catalog) *cobra.Command {
	root := &cobra.Command{
		Use:   "stack",
		Short: "Install skills and MCP servers for coding-agent harnesses",
		Long: "stack installs two things into a project for one or more coding-agent\n" +
			"harnesses: skills and MCP servers.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
			}
			if !isTerminal(cmd.InOrStdin()) || !isTerminal(cmd.OutOrStdout()) {
				return cmd.Help()
			}
			return tui.Run(cat, ".")
		},
	}
	root.AddCommand(
		newInstallCommand(cat),
		newListCommand(cat),
		newDoctorCommand(cat),
		newVersionCommand(),
	)
	return root
}

func isTerminal(v any) bool {
	f, ok := v.(*os.File)
	return ok && term.IsTerminal(f.Fd())
}
