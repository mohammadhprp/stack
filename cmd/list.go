package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mohammadhprp/stack/internal/catalog"
)

func newListCommand(cat *catalog.Catalog) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List available skills and MCP servers",
		Args:    cobra.NoArgs,
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			out := cmd.OutOrStdout()

			skills := cat.Skills()
			fmt.Fprintf(out, "Skills (%d):\n", len(skills))
			for _, skill := range skills {
				fmt.Fprintf(out, "  %-28s %s\n", skill.ID, skill.Description)
			}

			mcps := cat.MCPs()
			fmt.Fprintf(out, "\nMCP servers (%d):\n", len(mcps))
			for _, mcp := range mcps {
				fmt.Fprintf(out, "  %-24s %-22s %s [%s]\n", mcp.Slug, mcp.Name, mcp.Description, mcp.Type)
			}
			return nil
		},
	}
}
