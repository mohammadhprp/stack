package tui

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/models"
)

type selection struct {
	harnesses map[string]bool
	skills    map[string]bool
	mcps      map[string]bool
}

func prefill(cat *install.Catalog, target string) selection {
	sel := selection{
		harnesses: map[string]bool{},
		skills:    map[string]bool{},
		mcps:      map[string]bool{},
	}
	lock, err := install.LoadLock(target)
	switch {
	case err == nil && lock != nil:
		prefillFromLock(cat, lock, &sel)
	case errors.Is(err, install.ErrNoLock):
		prefillFromDisk(cat, target, &sel)
	}
	return sel
}

func prefillFromLock(cat *install.Catalog, lock *install.Lock, sel *selection) {
	for _, item := range lock.Items {
		sel.harnesses[item.Harness] = true
		switch item.Kind {
		case install.KindSkill:
			if _, ok := cat.Skill(item.ID); ok {
				sel.skills[item.ID] = true
			}
		case install.KindMCP:
			if _, ok := cat.MCP(item.ID); ok {
				sel.mcps[item.ID] = true
			}
		}
	}
}

func prefillFromDisk(cat *install.Catalog, target string, sel *selection) {
	entries, err := os.ReadDir(target)
	if err != nil || len(entries) == 0 {
		return
	}
	for _, adapter := range harness.All() {
		if adapter.SupportsSkills() {
			for _, skill := range cat.Skills() {
				files, err := adapter.PlanSkills(target, []models.Skill{skill}, cat.FS())
				if err != nil || !filesInstalled(target, files) {
					continue
				}
				sel.harnesses[adapter.ID()] = true
				sel.skills[skill.ID] = true
			}
		}
		for _, mcp := range cat.MCPs() {
			files, err := adapter.PlanMCPs(target, []models.MCP{mcp}, cat.FS())
			if err != nil || !filesInstalled(target, files) {
				continue
			}
			sel.harnesses[adapter.ID()] = true
			sel.mcps[mcp.Slug] = true
		}
	}
}

func filesInstalled(target string, files []harness.File) bool {
	if len(files) == 0 {
		return false
	}
	for _, f := range files {
		data, err := os.ReadFile(filepath.Join(target, filepath.FromSlash(f.Path)))
		if err != nil || install.Hash(data) != install.Hash(f.Content) {
			return false
		}
	}
	return true
}
