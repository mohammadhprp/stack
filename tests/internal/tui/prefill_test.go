package tui_test

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/models"
	"github.com/mohammadhprp/stack/internal/tui"
)

func selectedIDs(opts []tui.Option) []string {
	var ids []string
	for _, o := range opts {
		if o.Selected {
			ids = append(ids, o.ID)
		}
	}
	sort.Strings(ids)
	return ids
}

func assertSelected(t *testing.T, opts []tui.Option, want ...string) {
	t.Helper()
	got := selectedIDs(opts)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("selected: got %v, want %v", got, want)
	}
}

func installInto(t *testing.T, dir string, harnessIDs, skillIDs, mcpSlugs []string) {
	t.Helper()
	cat := loadCatalog(t)

	var adapters []harness.Adapter
	for _, id := range harnessIDs {
		adapter, ok := harness.Get(id)
		if !ok {
			t.Fatalf("harness %q not registered", id)
		}
		adapters = append(adapters, adapter)
	}
	var skills []models.Skill
	for _, id := range skillIDs {
		skill, ok := cat.Skill(id)
		if !ok {
			t.Fatalf("skill %q not found", id)
		}
		skills = append(skills, skill)
	}
	var mcps []models.MCP
	for _, slug := range mcpSlugs {
		mcp, ok := cat.MCP(slug)
		if !ok {
			t.Fatalf("mcp %q not found", slug)
		}
		mcps = append(mcps, mcp)
	}

	if _, err := install.Run(install.Request{
		Target:   dir,
		Adapters: adapters,
		Skills:   skills,
		MCPs:     mcps,
		Source:   cat.FS(),
	}); err != nil {
		t.Fatalf("install.Run: %v", err)
	}
}

func snapshotFiles(t *testing.T, dir string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(dir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}
		out[rel] = install.Hash(data)
		return nil
	})
	if err != nil {
		t.Fatalf("snapshot: %v", err)
	}
	return out
}

func TestNewPrefillsFromLock(t *testing.T) {
	dir := t.TempDir()
	installInto(t, dir, []string{"opencode", "claude"}, []string{"commit"}, []string{"playwright-mcp"})

	m := tui.New(loadCatalog(t), dir)
	t.Logf("lock prefill: harnesses=%v skills=%v mcps=%v", selectedIDs(m.Harnesses), selectedIDs(m.Skills), selectedIDs(m.MCPs))
	assertSelected(t, m.Harnesses, "claude", "opencode")
	assertSelected(t, m.Skills, "commit")
	assertSelected(t, m.MCPs, "playwright-mcp")
}

func TestNewPrefillsFromDiskWithoutLock(t *testing.T) {
	dir := t.TempDir()
	installInto(t, dir, []string{"opencode"}, []string{"commit"}, []string{"playwright-mcp"})
	if err := os.Remove(filepath.Join(dir, install.LockFile)); err != nil {
		t.Fatal(err)
	}

	m := tui.New(loadCatalog(t), dir)
	t.Logf("disk fallback: harnesses=%v skills=%v mcps=%v", selectedIDs(m.Harnesses), selectedIDs(m.Skills), selectedIDs(m.MCPs))
	assertSelected(t, m.Harnesses, "opencode")
	assertSelected(t, m.Skills, "commit")
	assertSelected(t, m.MCPs, "playwright-mcp")
}

func TestNewFallbackAttributesSharedAgentsSkills(t *testing.T) {
	dir := t.TempDir()
	installInto(t, dir, []string{"codex"}, []string{"commit"}, nil)
	if err := os.Remove(filepath.Join(dir, install.LockFile)); err != nil {
		t.Fatal(err)
	}

	m := tui.New(loadCatalog(t), dir)
	t.Logf("shared fallback: harnesses=%v skills=%v mcps=%v", selectedIDs(m.Harnesses), selectedIDs(m.Skills), selectedIDs(m.MCPs))
	assertSelected(t, m.Harnesses, "amp", "codex", "cursor")
	assertSelected(t, m.Skills, "commit")
	assertSelected(t, m.MCPs)
}

func TestNewEmptyTargetSelectsNothing(t *testing.T) {
	m := tui.New(loadCatalog(t), t.TempDir())
	t.Logf("empty target: harnesses=%v skills=%v mcps=%v", selectedIDs(m.Harnesses), selectedIDs(m.Skills), selectedIDs(m.MCPs))
	assertSelected(t, m.Harnesses)
	assertSelected(t, m.Skills)
	assertSelected(t, m.MCPs)
}

func TestNewCorruptLockSelectsNothing(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, install.LockFile), []byte("{ not json"), 0o644); err != nil {
		t.Fatal(err)
	}

	m := tui.New(loadCatalog(t), dir)
	assertSelected(t, m.Harnesses)
	assertSelected(t, m.Skills)
	assertSelected(t, m.MCPs)
}

func TestNewPrefillDoesNotWrite(t *testing.T) {
	dir := t.TempDir()
	installInto(t, dir, []string{"opencode"}, []string{"commit"}, []string{"playwright-mcp"})
	before := snapshotFiles(t, dir)

	m := tui.New(loadCatalog(t), dir)
	m, _ = send(m, "q")
	if m.Stage != tui.StageCancelled {
		t.Fatalf("stage: got %v, want cancelled", m.Stage)
	}

	after := snapshotFiles(t, dir)
	if !reflect.DeepEqual(before, after) {
		t.Errorf("prefill/cancel changed the target:\nbefore: %v\nafter:  %v", before, after)
	}
}
