package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"

	"github.com/mohammadhprp/stack/internal/catalog"
	"github.com/mohammadhprp/stack/internal/install"
)

type doctorStatus string

const (
	statusOK       doctorStatus = "ok"
	statusModified doctorStatus = "modified"
	statusMissing  doctorStatus = "missing"
)

func newDoctorCommand(cat *catalog.Catalog) *cobra.Command {
	var target string
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Check installed skills and MCP configs for changes",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runDoctor(cmd, cat, target)
		},
	}
	cmd.Flags().StringVar(&target, "target", ".", "target project directory")
	return cmd
}

func runDoctor(cmd *cobra.Command, cat *catalog.Catalog, target string) error {
	lock, err := install.LoadLock(target)
	if errors.Is(err, install.ErrNoLock) {
		fmt.Fprintf(cmd.OutOrStdout(), "No %s in %s; nothing installed.\n", install.LockFile, target)
		return nil
	}
	if err != nil {
		return err
	}

	out := cmd.OutOrStdout()
	ok, modified, missing, unknown := 0, 0, 0, 0
	for _, item := range lock.Items {
		fmt.Fprintf(out, "%s/%s/%s\n", item.Harness, item.Kind, item.ID)
		if !catalogHas(cat, item.Kind, item.ID) {
			unknown++
			fmt.Fprintln(out, "  ! catalog item no longer exists")
		}
		paths := make([]string, 0, len(item.Files))
		for rel := range item.Files {
			paths = append(paths, rel)
		}
		sort.Strings(paths)
		for _, rel := range paths {
			status := classify(target, rel, item.Files[rel])
			switch status {
			case statusOK:
				ok++
			case statusModified:
				modified++
			case statusMissing:
				missing++
			}
			fmt.Fprintf(out, "  %-8s %s\n", status, rel)
		}
	}

	fmt.Fprintf(out, "Summary: %d ok, %d modified, %d missing, %d unknown\n", ok, modified, missing, unknown)
	if modified+missing > 0 {
		return fmt.Errorf("doctor: %d problem(s) found", modified+missing)
	}
	return nil
}

func catalogHas(cat *catalog.Catalog, kind, id string) bool {
	switch kind {
	case install.KindSkill:
		_, ok := cat.Skill(id)
		return ok
	case install.KindMCP:
		_, ok := cat.MCP(id)
		return ok
	default:
		return false
	}
}

func classify(target, rel, want string) doctorStatus {
	data, err := os.ReadFile(filepath.Join(target, filepath.FromSlash(rel)))
	if errors.Is(err, fs.ErrNotExist) || err != nil {
		return statusMissing
	}
	if install.Hash(data) == want {
		return statusOK
	}
	return statusModified
}
