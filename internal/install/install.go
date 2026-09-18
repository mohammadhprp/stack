package install

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/models"
)

type Action string

// ActionMerge is a first-time merge into an existing config; ActionConflict is a
// user-modified file and is refused without Force; ActionRemove deletes an
// installed item no longer selected (only when pruning).
const (
	ActionCreate   Action = "create"
	ActionUpdate   Action = "update"
	ActionMerge    Action = "merge"
	ActionKeep     Action = "keep"
	ActionConflict Action = "conflict"
	ActionRemove   Action = "remove"
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
	// Prune removes installed items that are no longer selected. It is opt-in
	// for the CLI and always on for the TUI, whose selection is the end state.
	Prune bool
}

type Change struct {
	Action    Action
	Path      string
	Item      string
	Harnesses []string
	Source    string
}

type Report struct {
	Changes  []Change
	Warnings []string
	Lock     *Lock
}

type plannedFile struct {
	path      string
	content   []byte
	merge     bool
	reconcile bool
	source    string
}

type plannedItem struct {
	harness string
	kind    string
	id      string
	source  string
	files   []harness.File
}

type deletion struct {
	path string
	hash string
}

// Run never overwrites a file whose content differs from its locked hash
// unless Force is set. Removals require Prune and follow the same rule.
func Run(req Request) (*Report, error) {
	if req.Target == "" {
		return nil, errors.New("install: target directory is required")
	}
	if req.Source == nil {
		return nil, errors.New("install: catalog filesystem is required")
	}
	report := &Report{}

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

	planned := map[string]*plannedFile{}
	owners := map[string][]string{}
	var order []string
	var items []plannedItem

	addPlanned := func(owner string, f harness.File, reconcile bool) error {
		if f.Path == "" {
			return fmt.Errorf("install: %s planned an empty path", owner)
		}
		if existing, ok := planned[f.Path]; ok {
			if !bytes.Equal(existing.content, f.Content) {
				return fmt.Errorf("install: conflicting plans for %s", f.Path)
			}
			existing.merge = existing.merge || f.Merge
			existing.reconcile = existing.reconcile || reconcile
		} else {
			planned[f.Path] = &plannedFile{
				path:      f.Path,
				content:   f.Content,
				merge:     f.Merge,
				reconcile: reconcile,
				source:    f.Source,
			}
			order = append(order, f.Path)
		}
		owners[f.Path] = appendUnique(owners[f.Path], owner)
		return nil
	}
	addFiles := func(a harness.Adapter, kind, id, source string, files []harness.File) error {
		for _, f := range files {
			if err := addPlanned(a.ID(), f, false); err != nil {
				return err
			}
		}
		items = append(items, plannedItem{harness: a.ID(), kind: kind, id: id, source: source, files: files})
		return nil
	}

	// Desired keys for the selection, and the MCPs each selected harness keeps.
	desired := map[string]bool{}
	selectedHarness := map[string]bool{}
	selectedMCPs := map[string][]models.MCP{}
	for _, a := range req.Adapters {
		selectedHarness[a.ID()] = true
		if a.SupportsSkills() {
			for _, s := range req.Skills {
				desired[itemKey(a.ID(), KindSkill, s.ID)] = true
			}
		}
		if len(req.MCPs) > 0 {
			selectedMCPs[a.ID()] = req.MCPs
			for _, m := range req.MCPs {
				desired[itemKey(a.ID(), KindMCP, m.Slug)] = true
			}
		}
	}

	// Installed items that are no longer selected.
	removed := map[string]bool{}
	var skillItems []LockItem
	removedMCPItems := map[string][]LockItem{}
	if req.Prune {
		for _, item := range lock.Items {
			k := itemKey(item.Harness, item.Kind, item.ID)
			if desired[k] {
				continue
			}
			removed[k] = true
			switch item.Kind {
			case KindSkill:
				skillItems = append(skillItems, item)
			case KindMCP:
				removedMCPItems[item.Harness] = append(removedMCPItems[item.Harness], item)
			}
		}
	}

	mcpBySlug := map[string]models.MCP{}
	if len(removedMCPItems) > 0 {
		if mcps, err := loadMCPs(req.Source); err == nil {
			for _, m := range mcps {
				mcpBySlug[m.Slug] = m
			}
		}
	}
	resolveMCP := func(item LockItem) models.MCP {
		if m, ok := mcpBySlug[item.ID]; ok {
			return m
		}
		return models.MCP{Slug: item.ID, MCPSpec: models.MCPSpec{ID: item.ID}}
	}
	dropFor := func(h string) []models.MCP {
		var drop []models.MCP
		for _, item := range removedMCPItems[h] {
			drop = append(drop, resolveMCP(item))
		}
		return drop
	}

	var deletions []deletion
	var removalChanges []Change

	// An item is removed atomically: if any existing managed file is modified it
	// is kept whole (files and lock entry) unless Force; missing files never
	// block removal.
	itemModified := func(files map[string]string) string {
		for p, hash := range files {
			dest, ok := managedDest(req.Target, p)
			if !ok {
				continue
			}
			current, err := os.ReadFile(dest)
			if err != nil {
				continue
			}
			if Hash(current) != hash {
				return p
			}
		}
		return ""
	}

	removeItemFiles := func(item LockItem, protected map[string]bool) {
		paths := make([]string, 0, len(item.Files))
		for p := range item.Files {
			paths = append(paths, p)
		}
		sort.Strings(paths)
		for _, p := range paths {
			if protected[p] {
				continue
			}
			dest, ok := managedDest(req.Target, p)
			if !ok {
				continue
			}
			if _, err := os.Stat(dest); err != nil {
				continue
			}
			deletions = append(deletions, deletion{path: p, hash: item.Files[p]})
			removalChanges = append(removalChanges, Change{
				Action:    ActionRemove,
				Path:      p,
				Item:      item.ID,
				Harnesses: []string{item.Harness},
				Source:    item.Source,
			})
		}
	}

	recordServerRemovals := func(h string, lockItems []LockItem) {
		for _, item := range lockItems {
			if !removed[itemKey(item.Harness, item.Kind, item.ID)] {
				continue
			}
			mcp := resolveMCP(item)
			for p := range item.Files {
				removalChanges = append(removalChanges, Change{
					Action:    ActionRemove,
					Path:      p,
					Item:      mcp.ID,
					Harnesses: []string{h},
					Source:    item.Source,
				})
			}
		}
	}

	// A config that would become empty is deleted; if it is modified it is kept
	// whole and the dropped items stay locked.
	recordConfigDeletes := func(h string, deletes []string) {
		for _, p := range deletes {
			dest, ok := managedDest(req.Target, p)
			if !ok {
				continue
			}
			var owners []LockItem
			for _, item := range lock.Items {
				if item.Harness != h {
					continue
				}
				if _, ok := item.Files[p]; ok {
					owners = append(owners, item)
				}
			}
			if len(owners) == 0 {
				continue
			}
			current, err := os.ReadFile(dest)
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			if err != nil || (!req.Force && Hash(current) != owners[0].Files[p]) {
				for _, item := range owners {
					delete(removed, itemKey(item.Harness, item.Kind, item.ID))
				}
				report.Warnings = append(report.Warnings,
					fmt.Sprintf("kept modified %s %s: %s", owners[0].Kind, owners[0].ID, p))
				continue
			}
			deletions = append(deletions, deletion{path: p, hash: owners[0].Files[p]})
		}
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

		drop := dropFor(a.ID())
		if len(drop) > 0 {
			remover, ok := a.(harness.MCPRemover)
			if !ok {
				report.Warnings = append(report.Warnings,
					fmt.Sprintf("%s: adapter does not support MCP removal; left unchanged", a.Name()))
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
				continue
			}
			rewrites, deletes, err := remover.PlanMCPRemoval(req.Target, req.MCPs, drop, req.Source)
			if err != nil {
				return report, err
			}
			for _, f := range rewrites {
				if err := addPlanned(a.ID(), f, true); err != nil {
					return report, err
				}
			}
			for _, mcp := range req.MCPs {
				if err := addFiles(a, KindMCP, mcp.Slug, mcp.Dir, rewrites); err != nil {
					return report, err
				}
			}
			recordConfigDeletes(a.ID(), deletes)
			recordServerRemovals(a.ID(), removedMCPItems[a.ID()])
			continue
		}

		if len(req.MCPs) > 0 {
			note := ""
			if noted, ok := a.(harness.MCPConfigReporter); ok {
				note = noted.MCPConfigNote()
			}
			if note != "" {
				report.Warnings = append(report.Warnings, fmt.Sprintf("%s: %s", a.Name(), note))
			} else {
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
	}

	// Harnesses that were dropped entirely still need their MCP config cleaned.
	harnessKeys := make([]string, 0, len(removedMCPItems))
	for h := range removedMCPItems {
		harnessKeys = append(harnessKeys, h)
	}
	sort.Strings(harnessKeys)
	for _, h := range harnessKeys {
		if selectedHarness[h] {
			continue
		}
		a, ok := harness.Get(h)
		if !ok {
			report.Warnings = append(report.Warnings, fmt.Sprintf("cannot remove MCP config for unknown harness %q", h))
			continue
		}
		remover, ok := a.(harness.MCPRemover)
		if !ok {
			report.Warnings = append(report.Warnings,
				fmt.Sprintf("%s: adapter does not support MCP removal; left unchanged", a.Name()))
			continue
		}
		rewrites, deletes, err := remover.PlanMCPRemoval(req.Target, nil, dropFor(h), req.Source)
		if err != nil {
			return report, err
		}
		for _, f := range rewrites {
			if err := addPlanned(h, f, true); err != nil {
				return report, err
			}
		}
		recordConfigDeletes(h, deletes)
		recordServerRemovals(h, removedMCPItems[h])
	}

	// Skill items: remove atomically, protecting paths the selection still needs.
	if req.Prune && len(skillItems) > 0 {
		protected := map[string]bool{}
		for p := range planned {
			protected[p] = true
		}
		for _, item := range skillItems {
			if modified := itemModified(item.Files); modified != "" && !req.Force {
				delete(removed, itemKey(item.Harness, item.Kind, item.ID))
				report.Warnings = append(report.Warnings,
					fmt.Sprintf("kept modified %s %s: %s", item.Kind, item.ID, modified))
				continue
			}
			removeItemFiles(item, protected)
		}
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
			case pf.reconcile:
				// Reconcile preserves unrelated keys, so a modified config is safe.
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

	sort.Slice(removalChanges, func(i, j int) bool {
		if removalChanges[i].Path != removalChanges[j].Path {
			return removalChanges[i].Path < removalChanges[j].Path
		}
		return removalChanges[i].Item < removalChanges[j].Item
	})
	removalChanges = dedupeRemovals(removalChanges)
	report.Changes = append(report.Changes, removalChanges...)
	deletions = dedupeDeletions(deletions)

	newLock := lock.clone()
	if req.Prune {
		kept := make([]LockItem, 0, len(newLock.Items))
		for _, item := range newLock.Items {
			if removed[itemKey(item.Harness, item.Kind, item.ID)] {
				continue
			}
			kept = append(kept, item)
		}
		newLock.Items = kept
	}
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

	for _, d := range deletions {
		dest, ok := managedDest(req.Target, d.path)
		if !ok {
			continue
		}
		if err := os.Remove(dest); err != nil && !errors.Is(err, fs.ErrNotExist) {
			return report, fmt.Errorf("install: remove %s: %w", d.path, err)
		}
		pruneEmptyDirs(req.Target, filepath.Dir(dest))
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

func itemKey(h, kind, id string) string { return h + "\x00" + kind + "\x00" + id }

func dedupeRemovals(in []Change) []Change {
	seen := map[string]int{}
	out := in[:0]
	for _, c := range in {
		k := c.Path + "\x00" + c.Item
		if i, ok := seen[k]; ok {
			out[i].Harnesses = mergeUnique(out[i].Harnesses, c.Harnesses)
			continue
		}
		seen[k] = len(out)
		out = append(out, c)
	}
	return out
}

func dedupeDeletions(in []deletion) []deletion {
	seen := map[string]bool{}
	out := in[:0]
	for _, d := range in {
		if seen[d.path] {
			continue
		}
		seen[d.path] = true
		out = append(out, d)
	}
	return out
}

func mergeUnique(values, extra []string) []string {
	for _, v := range extra {
		values = appendUnique(values, v)
	}
	return values
}

func managedDest(target, rel string) (string, bool) {
	clean := filepath.Clean(filepath.FromSlash(rel))
	if filepath.IsAbs(clean) || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return "", false
	}
	dest := filepath.Join(target, clean)
	inside, err := filepath.Rel(target, dest)
	if err != nil || inside == ".." || strings.HasPrefix(inside, ".."+string(filepath.Separator)) {
		return "", false
	}
	return dest, true
}

func pruneEmptyDirs(target, dir string) {
	targetAbs, err := filepath.Abs(target)
	if err != nil {
		return
	}
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return
	}
	for dirAbs != targetAbs {
		entries, err := os.ReadDir(dirAbs)
		if err != nil || len(entries) > 0 {
			return
		}
		if err := os.Remove(dirAbs); err != nil {
			return
		}
		parent := filepath.Dir(dirAbs)
		if parent == dirAbs {
			return
		}
		dirAbs = parent
	}
}

func appendUnique(values []string, value string) []string {
	for _, v := range values {
		if v == value {
			return values
		}
	}
	return append(values, value)
}
