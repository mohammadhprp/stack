package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mohammadhprp/stack/internal/catalog"
)

func NewRootCommand(cat *catalog.Catalog) *cobra.Command {
	root := &cobra.Command{
		Use:   "stack",
		Short: "Install skills and MCP servers for coding-agent harnesses",
		Long: "stack installs two things into a project for one or more coding-agent\n" +
			"harnesses: skills and MCP servers.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(
		newInstallCommand(cat),
		newListCommand(cat),
		newVersionCommand(),
	)
	return root
}
