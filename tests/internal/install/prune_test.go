package install_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/models"
)

func runInstall(t *testing.T, dir string, harnessIDs, skillIDs, mcpSlugs []string, prune, force, dryRun bool) *install.Report {
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

	report, err := install.Run(install.Request{
		Target:   dir,
		Adapters: adapters,
		Skills:   skills,
		MCPs:     mcps,
		Source:   cat.FS(),
		Prune:    prune,
		Force:    force,
		DryRun:   dryRun,
	})
	if err != nil {
		t.Fatalf("install.Run: %v", err)
	}
	return report
}

func hasRemove(report *install.Report) bool {
	for _, c := range report.Changes {
		if c.Action == install.ActionRemove {
			return true
		}
	}
	return false
}

func hasRemovePath(report *install.Report, path string) bool {
	for _, c := range report.Changes {
		if c.Action == install.ActionRemove && c.Path == path {
			return true
		}
	}
	return false
}

func hasWarning(report *install.Report, substr string) bool {
	for _, w := range report.Warnings {
		if strings.Contains(w, substr) {
			return true
		}
	}
	return false
}

func lockedItem(lock *install.Lock, harness, kind, id string) bool {
	for _, item := range lock.Items {
		if item.Harness == harness && item.Kind == kind && item.ID == id {
			return true
		}
	}
	return false
}

func mustExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		t.Errorf("%s should exist: %v", path, err)
	}
}

func mustNotExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("%s should be removed", path)
	}
}

func TestRunPruneRemovesDeselectedItems(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"claude", "codex"}, []string{"commit", "why"}, []string{"playwright-mcp"}, false, false, false)

	report := runInstall(t, dir, []string{"claude"}, []string{"commit"}, nil, true, false, false)
	if !hasRemove(report) {
		t.Error("prune reported no removals")
	}

	for _, rel := range []string{".codex/config.toml", ".mcp.json", ".agents/skills", ".claude/skills/why"} {
		mustNotExist(t, filepath.Join(dir, filepath.FromSlash(rel)))
	}
	mustExist(t, filepath.Join(dir, ".claude", "skills", "commit", "SKILL.md"))
	mustExist(t, filepath.Join(dir, ".claude", "skills", "commit", "examples.md"))

	lock, err := install.LoadLock(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(lock.Items) != 1 {
		t.Fatalf("lock items after prune: got %d, want 1: %+v", len(lock.Items), lock.Items)
	}
	item := lock.Items[0]
	if item.Harness != "claude" || item.Kind != install.KindSkill || item.ID != "commit" {
		t.Errorf("lock item after prune: %+v", item)
	}
}

func TestRunWithoutPruneKeepsDeselectedItems(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"claude", "codex"}, []string{"commit"}, nil, false, false, false)

	report := runInstall(t, dir, []string{"claude"}, []string{"commit"}, nil, false, false, false)
	if hasRemove(report) {
		t.Error("without --prune nothing may be removed")
	}
	mustExist(t, filepath.Join(dir, ".agents", "skills", "commit", "SKILL.md"))

	lock, err := install.LoadLock(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(lock.Items) != 2 {
		t.Errorf("lock items without prune: got %d, want 2", len(lock.Items))
	}
}

func TestRunPruneKeepsModifiedItemWholeThenForce(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"claude"}, []string{"commit"}, nil, false, false, false)
	skillDir := filepath.Join(dir, ".claude", "skills", "commit")
	skillFile := filepath.Join(skillDir, "SKILL.md")
	if err := os.WriteFile(skillFile, []byte("user edit\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	report := runInstall(t, dir, []string{"claude"}, nil, []string{"playwright-mcp"}, true, false, false)
	if !hasWarning(report, "kept modified skill commit") {
		t.Errorf("expected a warning naming the item and file, got %v", report.Warnings)
	}
	if hasRemovePath(report, ".claude/skills/commit/SKILL.md") || hasRemovePath(report, ".claude/skills/commit/examples.md") {
		t.Error("a modified item must be kept whole; no file may be planned for removal")
	}
	mustExist(t, skillFile)
	mustExist(t, filepath.Join(skillDir, "examples.md"))

	lock, err := install.LoadLock(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !lockedItem(lock, "claude", install.KindSkill, "commit") {
		t.Error("the kept item must remain locked")
	}

	runInstall(t, dir, []string{"claude"}, nil, []string{"playwright-mcp"}, true, true, false)
	mustNotExist(t, skillDir)
	lock, err = install.LoadLock(dir)
	if err != nil {
		t.Fatal(err)
	}
	if lockedItem(lock, "claude", install.KindSkill, "commit") {
		t.Error("a forced prune must drop the item")
	}
}

func TestRunPruneRemovesUnchangedButKeepsModifiedItem(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"claude"}, []string{"commit", "why"}, nil, false, false, false)
	whySkill := filepath.Join(dir, ".claude", "skills", "why", "SKILL.md")
	if err := os.WriteFile(whySkill, []byte("user edit\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	runInstall(t, dir, []string{"claude"}, nil, []string{"playwright-mcp"}, true, false, false)

	mustNotExist(t, filepath.Join(dir, ".claude", "skills", "commit"))
	mustExist(t, whySkill)
	mustExist(t, filepath.Join(dir, ".claude", "skills", "why", "examples.md"))

	lock, err := install.LoadLock(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !lockedItem(lock, "claude", install.KindSkill, "why") {
		t.Error("the modified item must remain locked")
	}
	if lockedItem(lock, "claude", install.KindSkill, "commit") {
		t.Error("the unchanged item must be dropped")
	}
}

func TestRunPruneKeepsModifiedEmptiedConfig(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"codex"}, nil, []string{"playwright-mcp"}, false, false, false)
	configPath := filepath.Join(dir, ".codex", "config.toml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(configPath, append(data, []byte("# user note\n")...), 0o644); err != nil {
		t.Fatal(err)
	}

	report := runInstall(t, dir, []string{"codex"}, []string{"commit"}, nil, true, false, false)
	mustExist(t, configPath)
	if !hasWarning(report, "kept modified mcp") {
		t.Errorf("expected an mcp kept-modified warning, got %v", report.Warnings)
	}
	lock, err := install.LoadLock(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !lockedItem(lock, "codex", install.KindMCP, "playwright-mcp") {
		t.Error("the kept mcp item must remain locked")
	}
}

func TestRunPruneMCPPreservesOtherKeysAndServers(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"claude"}, nil, []string{"playwright-mcp"}, false, false, false)

	seeded := `{
  "theme": "dark",
  "mcpServers": {
    "keepme": { "command": "echo" },
    "playwright": { "command": "npx", "args": ["@playwright/mcp@latest"] },
    "supabase": { "url": "https://mcp.supabase.com/mcp" }
  }
}`
	configPath := filepath.Join(dir, ".mcp.json")
	if err := os.WriteFile(configPath, []byte(seeded), 0o644); err != nil {
		t.Fatal(err)
	}

	runInstall(t, dir, []string{"claude"}, nil, []string{"supabase-mcp"}, true, false, false)

	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	var root map[string]any
	if err := json.Unmarshal(data, &root); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if root["theme"] != "dark" {
		t.Errorf("unrelated key theme was dropped: %v", root["theme"])
	}
	server := root["mcpServers"].(map[string]any)
	if _, ok := server["keepme"]; !ok {
		t.Error("unrelated mcp entry keepme was dropped")
	}
	if _, ok := server["playwright"]; ok {
		t.Error("playwright should have been removed")
	}
	if _, ok := server["supabase"]; !ok {
		t.Error("selected supabase entry was dropped")
	}
}

func TestRunPruneMCPEmptiedConfigDeleted(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"codex"}, nil, []string{"playwright-mcp"}, false, false, false)

	runInstall(t, dir, []string{"codex"}, []string{"commit"}, nil, true, false, false)

	mustNotExist(t, filepath.Join(dir, ".codex", "config.toml"))
	mustNotExist(t, filepath.Join(dir, ".codex"))
	mustExist(t, filepath.Join(dir, ".agents", "skills", "commit", "SKILL.md"))
}

func TestRunPruneSharedAgentsSkillsKeepsNeeded(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"codex", "amp"}, []string{"commit", "why"}, nil, false, false, false)

	runInstall(t, dir, []string{"codex"}, []string{"commit"}, nil, true, false, false)

	mustExist(t, filepath.Join(dir, ".agents", "skills", "commit", "SKILL.md"))
	mustNotExist(t, filepath.Join(dir, ".agents", "skills", "why"))
}

func TestRunPruneNeverDeletesOutsideTarget(t *testing.T) {
	base := t.TempDir()
	dir := filepath.Join(base, "project")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(base, "outside.txt")
	if err := os.WriteFile(outside, []byte("keep me"), 0o644); err != nil {
		t.Fatal(err)
	}

	lock := &install.Lock{Version: install.LockVersion, Items: []install.LockItem{{
		Harness: "claude",
		Kind:    install.KindSkill,
		ID:      "commit",
		Source:  "skills/commit",
		Files:   map[string]string{"../outside.txt": install.Hash([]byte("keep me"))},
	}}}
	if err := install.SaveLock(dir, lock); err != nil {
		t.Fatal(err)
	}

	runInstall(t, dir, []string{"claude"}, nil, []string{"playwright-mcp"}, true, false, false)
	mustExist(t, outside)
}

func TestRunPruneThenReinstallRecreates(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"claude"}, []string{"commit", "why"}, nil, false, false, false)
	runInstall(t, dir, []string{"claude"}, []string{"commit"}, nil, true, false, false)
	runInstall(t, dir, []string{"claude"}, []string{"commit", "why"}, nil, true, false, false)

	mustExist(t, filepath.Join(dir, ".claude", "skills", "why", "SKILL.md"))
	lock, err := install.LoadLock(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(lock.Items) != 2 {
		t.Errorf("lock items after reinstall: got %d, want 2", len(lock.Items))
	}
}

func TestRunPruneDryRunWritesNothing(t *testing.T) {
	dir := t.TempDir()
	runInstall(t, dir, []string{"claude"}, []string{"commit"}, nil, false, false, false)
	lockFile := filepath.Join(dir, install.LockFile)
	before, err := os.ReadFile(lockFile)
	if err != nil {
		t.Fatal(err)
	}

	report := runInstall(t, dir, []string{"claude"}, nil, []string{"playwright-mcp"}, true, false, true)
	if !hasRemove(report) {
		t.Error("dry-run prune should still report removals")
	}
	mustExist(t, filepath.Join(dir, ".claude", "skills", "commit", "SKILL.md"))
	mustNotExist(t, filepath.Join(dir, ".mcp.json"))
	after, err := os.ReadFile(lockFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(before) != string(after) {
		t.Error("dry-run changed the lockfile")
	}
}
