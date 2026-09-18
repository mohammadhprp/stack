package install

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/models"
)

type Action string

// ActionMerge is a first-time merge into an existing config; ActionConflict is
// a user-modified file and is refused without Force.
const (
	ActionCreate   Action = "create"
	ActionUpdate   Action = "update"
	ActionMerge    Action = "merge"
	ActionKeep     Action = "keep"
	ActionConflict Action = "conflict"
)

type Request struct {
	Target   string
	Adapters []harness.Adapter
	Skills   []models.Skill
	MCPs     []models.MCP
	// Source is the catalog filesystem, rooted at framework/.
	Source fs.FS
	Force  bool
	DryRun bool
}

type Change struct {
	Action    Action
	Path      string
	Harnesses []string
	Source    string
}

type Report struct {
	Changes  []Change
	Warnings []string
	Lock     *Lock
}

type plannedFile struct {
	path    string
	content []byte
	merge   bool
	source  string
}

type plannedItem struct {
	harness string
	kind    string
	id      string
	source  string
	files   []harness.File
}

// Run never overwrites a file whose content differs from its locked hash
// unless Force is set.
func Run(req Request) (*Report, error) {
	if req.Target == "" {
		return nil, errors.New("install: target directory is required")
	}
	if req.Source == nil {
		return nil, errors.New("install: catalog filesystem is required")
	}
	report := &Report{}

	planned := map[string]*plannedFile{}
	owners := map[string][]string{}
	var order []string
	var items []plannedItem

	addFiles := func(a harness.Adapter, kind, id, source string, files []harness.File) error {
		for _, f := range files {
			if f.Path == "" {
				return fmt.Errorf("install: %s %s planned an empty path", a.ID(), id)
			}
			if existing, ok := planned[f.Path]; ok {
				if !bytes.Equal(existing.content, f.Content) {
					return fmt.Errorf("install: conflicting plans for %s", f.Path)
				}
				existing.merge = existing.merge || f.Merge
			} else {
				planned[f.Path] = &plannedFile{
					path:    f.Path,
					content: f.Content,
					merge:   f.Merge,
					source:  f.Source,
				}
				order = append(order, f.Path)
			}
			owners[f.Path] = appendUnique(owners[f.Path], a.ID())
		}
		items = append(items, plannedItem{harness: a.ID(), kind: kind, id: id, source: source, files: files})
		return nil
	}

	for _, a := range req.Adapters {
		if len(req.Skills) > 0 {
			if !a.SupportsSkills() {
				report.Warnings = append(report.Warnings,
					fmt.Sprintf("%s does not support skills; skipped %d skill(s)", a.Name(), len(req.Skills)))
			} else {
				for _, skill := range req.Skills {
					files, err := a.PlanSkills(req.Target, []models.Skill{skill}, req.Source)
					if err != nil {
						return report, err
					}
					if err := addFiles(a, KindSkill, skill.ID, skill.Dir, files); err != nil {
						return report, err
					}
				}
			}
		}
		if len(req.MCPs) > 0 {
			files, err := a.PlanMCPs(req.Target, req.MCPs, req.Source)
			if err != nil {
				return report, err
			}
			for _, mcp := range req.MCPs {
				if err := addFiles(a, KindMCP, mcp.Slug, mcp.Dir, files); err != nil {
					return report, err
				}
			}
		}
	}

	lock, err := LoadLock(req.Target)
	if err != nil && !errors.Is(err, ErrNoLock) {
		return report, err
	}
	if lock == nil {
		lock = &Lock{Version: LockVersion}
	}
	locked, err := lock.index()
	if err != nil {
		return report, err
	}

	actions := make(map[string]Action, len(order))
	var conflicts []string
	for _, path := range order {
		pf := planned[path]
		dest := filepath.Join(req.Target, filepath.FromSlash(pf.path))
		current, err := os.ReadFile(dest)
		var action Action
		switch {
		case errors.Is(err, fs.ErrNotExist):
			action = ActionCreate
		case err != nil:
			return report, fmt.Errorf("install: read %s: %w", pf.path, err)
		case bytes.Equal(current, pf.content):
			action = ActionKeep
		default:
			lockedHash, tracked := locked[pf.path]
			switch {
			case tracked && Hash(current) == lockedHash:
				action = ActionUpdate
			case pf.merge && !tracked:
				action = ActionMerge
			case req.Force:
				action = ActionUpdate
			default:
				action = ActionConflict
				conflicts = append(conflicts, pf.path)
			}
		}
		actions[path] = action
		report.Changes = append(report.Changes, Change{
			Action:    action,
			Path:      pf.path,
			Harnesses: owners[path],
			Source:    pf.source,
		})
	}

	if len(conflicts) > 0 {
		return report, fmt.Errorf("install: refusing to overwrite %d modified file(s) without --force: %s",
			len(conflicts), strings.Join(conflicts, ", "))
	}

	newLock := lock.clone()
	for _, item := range items {
		files := make(map[string]string, len(item.files))
		for _, f := range item.files {
			if pf, ok := planned[f.Path]; ok {
				files[f.Path] = Hash(pf.content)
			}
		}
		newLock.upsert(LockItem{
			Harness: item.harness,
			Kind:    item.kind,
			ID:      item.id,
			Source:  item.source,
			Files:   files,
		})
	}
	report.Lock = newLock

	if req.DryRun {
		return report, nil
	}

	for _, path := range order {
		if actions[path] == ActionKeep {
			continue
		}
		pf := planned[path]
		dest := filepath.Join(req.Target, filepath.FromSlash(pf.path))
		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return report, fmt.Errorf("install: create directory for %s: %w", pf.path, err)
		}
		if err := os.WriteFile(dest, pf.content, 0o644); err != nil {
			return report, fmt.Errorf("install: write %s: %w", pf.path, err)
		}
	}
	if err := SaveLock(req.Target, newLock); err != nil {
		return report, err
	}
	return report, nil
}

func appendUnique(values []string, value string) []string {
	for _, v := range values {
		if v == value {
			return values
		}
	}
	return append(values, value)
}
