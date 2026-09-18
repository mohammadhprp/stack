package cmd

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/models"
)

type installOptions struct {
	harnesses []string
	skills    []string
	mcps      []string
	target    string
	all       bool
	dryRun    bool
	force     bool
	prune     bool
}

func newInstallCommand(cat *install.Catalog) *cobra.Command {
	opts := &installOptions{}
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install skills and MCP servers into a project",
		Long: "Install selected skills and MCP servers into a target project for one\n" +
			"or more harnesses. Re-running is idempotent; a file that differs from the\n" +
			"lockfile is only overwritten with --force.",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runInstall(cmd, cat, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringSliceVar(&opts.harnesses, "harness", nil, "harnesses to install for (comma-separated)")
	flags.StringSliceVar(&opts.skills, "skills", nil, "skill ids to install (comma-separated)")
	flags.StringSliceVar(&opts.mcps, "mcp", nil, "MCP slugs to install (comma-separated)")
	flags.StringVar(&opts.target, "target", ".", "target project directory")
	flags.BoolVar(&opts.all, "all", false, "install every skill and MCP")
	flags.BoolVar(&opts.dryRun, "dry-run", false, "print the plan without writing anything")
	flags.BoolVar(&opts.force, "force", false, "overwrite files that differ from the lockfile")
	flags.BoolVar(&opts.prune, "prune", false, "remove installed items that are no longer selected")
	return cmd
}

func runInstall(cmd *cobra.Command, cat *install.Catalog, opts *installOptions) error {
	if strings.TrimSpace(opts.target) == "" {
		return errors.New("--target must not be empty")
	}

	adapters, err := resolveHarnesses(opts.harnesses)
	if err != nil {
		return err
	}
	skills, mcps, err := resolveSelection(cat, opts)
	if err != nil {
		return err
	}

	report, runErr := install.Run(install.Request{
		Target:   opts.target,
		Adapters: adapters,
		Skills:   skills,
		MCPs:     mcps,
		Source:   cat.FS(),
		Force:    opts.force,
		DryRun:   opts.dryRun,
		Prune:    opts.prune,
	})

	out := cmd.OutOrStdout()
	for _, warning := range reportWarnings(report) {
		fmt.Fprintf(cmd.ErrOrStderr(), "warning: %s\n", warning)
	}
	if report != nil {
		printPlan(out, opts, report)
	}
	if runErr != nil {
		return runErr
	}
	if opts.dryRun {
		fmt.Fprintln(out, "Dry run: no files written.")
	} else {
		fmt.Fprintf(out, "Installed %d file(s) into %s\n", countWrites(report), opts.target)
		if removed := countRemovals(report); removed > 0 {
			fmt.Fprintf(out, "Removed %d item(s) from %s\n", removed, opts.target)
		}
	}
	return nil
}

func reportWarnings(report *install.Report) []string {
	if report == nil {
		return nil
	}
	return report.Warnings
}

func printPlan(out io.Writer, opts *installOptions, report *install.Report) {
	if report == nil || len(report.Changes) == 0 {
		fmt.Fprintln(out, "Plan: nothing to install.")
		return
	}
	fmt.Fprintf(out, "Plan for %s:\n", opts.target)
	for _, change := range report.Changes {
		path := change.Path
		if change.Item != "" {
			path = fmt.Sprintf("%s (%s)", change.Path, change.Item)
		}
		fmt.Fprintf(out, "  %-8s %s\n", change.Action, path)
	}
}

func countWrites(report *install.Report) int {
	if report == nil {
		return 0
	}
	count := 0
	for _, change := range report.Changes {
		switch change.Action {
		case install.ActionCreate, install.ActionUpdate, install.ActionMerge:
			count++
		}
	}
	return count
}

func countRemovals(report *install.Report) int {
	if report == nil {
		return 0
	}
	count := 0
	for _, change := range report.Changes {
		if change.Action == install.ActionRemove {
			count++
		}
	}
	return count
}

func resolveHarnesses(ids []string) ([]harness.Adapter, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("no harness selected; pass --harness (available: %s)", availableHarnesses())
	}
	seen := map[string]bool{}
	var adapters []harness.Adapter
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id == "" || seen[id] {
			continue
		}
		adapter, ok := harness.Get(id)
		if !ok {
			return nil, fmt.Errorf("unknown harness %q; available: %s", id, availableHarnesses())
		}
		seen[id] = true
		adapters = append(adapters, adapter)
	}
	if len(adapters) == 0 {
		return nil, fmt.Errorf("no harness selected; pass --harness (available: %s)", availableHarnesses())
	}
	return adapters, nil
}

func availableHarnesses() string {
	all := harness.All()
	ids := make([]string, 0, len(all))
	for _, adapter := range all {
		ids = append(ids, adapter.ID())
	}
	return strings.Join(ids, ", ")
}

func resolveSelection(cat *install.Catalog, opts *installOptions) ([]models.Skill, []models.MCP, error) {
	var skills []models.Skill
	var mcps []models.MCP

	if opts.all {
		skills = cat.Skills()
		mcps = cat.MCPs()
	}
	for _, id := range opts.skills {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		skill, ok := cat.Skill(id)
		if !ok {
			return nil, nil, fmt.Errorf("unknown skill %q", id)
		}
		skills = append(skills, skill)
	}
	for _, slug := range opts.mcps {
		slug = strings.TrimSpace(slug)
		if slug == "" {
			continue
		}
		mcp, ok := cat.MCP(slug)
		if !ok {
			return nil, nil, fmt.Errorf("unknown MCP %q", slug)
		}
		mcps = append(mcps, mcp)
	}

	skills = dedupeSkills(skills)
	mcps = dedupeMCPs(mcps)
	if len(skills) == 0 && len(mcps) == 0 {
		return nil, nil, errors.New("nothing selected; pass --skills, --mcp, or --all")
	}
	return skills, mcps, nil
}

func dedupeSkills(in []models.Skill) []models.Skill {
	seen := map[string]bool{}
	out := in[:0]
	for _, skill := range in {
		if seen[skill.ID] {
			continue
		}
		seen[skill.ID] = true
		out = append(out, skill)
	}
	return out
}

func dedupeMCPs(in []models.MCP) []models.MCP {
	seen := map[string]bool{}
	out := in[:0]
	for _, mcp := range in {
		if seen[mcp.Slug] {
			continue
		}
		seen[mcp.Slug] = true
		out = append(out, mcp)
	}
	return out
}
