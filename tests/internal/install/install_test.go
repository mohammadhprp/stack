package install_test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/mohammadhprp/stack/internal/catalog"
	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/model"
)

func loadCatalog(t *testing.T) *catalog.Catalog {
	t.Helper()
	cat, err := catalog.Load(os.DirFS("../../../framework"))
	if err != nil {
		t.Fatalf("catalog.Load: %v", err)
	}
	return cat
}

func opencodeAdapter(t *testing.T) harness.Adapter {
	t.Helper()
	adapter, ok := harness.Get("opencode")
	if !ok {
		t.Fatal("opencode adapter not registered")
	}
	return adapter
}

func readJSON(t *testing.T, path string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var out map[string]any
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}
	return out
}

func TestRunInstallsOpenCode(t *testing.T) {
	cat := loadCatalog(t)
	skill, _ := cat.Skill("commit")
	mcp, _ := cat.MCP("playwright-mcp")
	dir := t.TempDir()

	if _, err := install.Run(install.Request{
		Target:   dir,
		Adapters: []harness.Adapter{opencodeAdapter(t)},
		Skills:   []model.Skill{skill},
		MCPs:     []model.MCP{mcp},
		Source:   cat.FS(),
	}); err != nil {
		t.Fatalf("Run: %v", err)
	}

	for _, rel := range []string{
		".opencode/skills/commit/SKILL.md",
		".opencode/skills/commit/examples.md",
		"opencode.json",
		install.LockFile,
	} {
		if _, err := os.Stat(filepath.Join(dir, filepath.FromSlash(rel))); err != nil {
			t.Errorf("expected %s: %v", rel, err)
		}
	}

	config := readJSON(t, filepath.Join(dir, "opencode.json"))
	mcpSection, ok := config["mcp"].(map[string]any)
	if !ok {
		t.Fatalf("opencode.json has no mcp section: %v", config)
	}
	if _, ok := mcpSection["playwright"]; !ok {
		t.Errorf("mcp section missing the canonical server key: %v", mcpSection)
	}

	lock, err := install.LoadLock(dir)
	if err != nil {
		t.Fatalf("LoadLock: %v", err)
	}
	if len(lock.Items) != 2 {
		t.Fatalf("lock items: got %d, want 2", len(lock.Items))
	}
	if err := lock.Validate(dir); err != nil {
		t.Fatalf("Validate after install: %v", err)
	}

	second, err := install.Run(install.Request{
		Target:   dir,
		Adapters: []harness.Adapter{opencodeAdapter(t)},
		Skills:   []model.Skill{skill},
		MCPs:     []model.MCP{mcp},
		Source:   cat.FS(),
	})
	if err != nil {
		t.Fatalf("second Run: %v", err)
	}
	for _, change := range second.Changes {
		if change.Action != install.ActionKeep {
			t.Errorf("re-run changed %s: %s", change.Path, change.Action)
		}
	}
	if err := second.Lock.Validate(dir); err != nil {
		t.Fatalf("Validate after re-run: %v", err)
	}
}

func TestRunDryRunWritesNothing(t *testing.T) {
	cat := loadCatalog(t)
	skill, _ := cat.Skill("commit")
	dir := t.TempDir()

	report, err := install.Run(install.Request{
		Target:   dir,
		Adapters: []harness.Adapter{opencodeAdapter(t)},
		Skills:   []model.Skill{skill},
		Source:   cat.FS(),
		DryRun:   true,
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(report.Changes) == 0 {
		t.Fatal("dry run produced no plan")
	}
	assertEmptyDir(t, dir)
}

func TestRunRefusesUserModifiedFileThenForce(t *testing.T) {
	cat := loadCatalog(t)
	skill, _ := cat.Skill("commit")
	dir := t.TempDir()
	req := install.Request{
		Target:   dir,
		Adapters: []harness.Adapter{opencodeAdapter(t)},
		Skills:   []model.Skill{skill},
		Source:   cat.FS(),
	}
	if _, err := install.Run(req); err != nil {
		t.Fatalf("first Run: %v", err)
	}

	dest := filepath.Join(dir, ".opencode", "skills", "commit", "SKILL.md")
	original, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dest, []byte("user edit\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	report, err := install.Run(req)
	if err == nil {
		t.Fatal("expected a refusal for a user-modified file")
	}
	var sawConflict bool
	for _, change := range report.Changes {
		if change.Action == install.ActionConflict {
			sawConflict = true
		}
	}
	if !sawConflict {
		t.Error("report did not include a conflict change")
	}

	req.Force = true
	if _, err := install.Run(req); err != nil {
		t.Fatalf("forced Run: %v", err)
	}
	restored, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if string(restored) != string(original) {
		t.Error("forced run did not restore the managed content")
	}
}

func TestRunMergesExistingConfigPreservingKeys(t *testing.T) {
	cat := loadCatalog(t)
	mcp, _ := cat.MCP("playwright-mcp")
	dir := t.TempDir()
	existing := `{
  "$schema": "https://opencode.ai/config.json",
  "theme": "opencode",
  "mcp": {
    "existing": { "type": "remote", "url": "https://example.com/mcp" }
  }
}`
	if err := os.WriteFile(filepath.Join(dir, "opencode.json"), []byte(existing), 0o644); err != nil {
		t.Fatal(err)
	}

	report, err := install.Run(install.Request{
		Target:   dir,
		Adapters: []harness.Adapter{opencodeAdapter(t)},
		MCPs:     []model.MCP{mcp},
		Source:   cat.FS(),
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if report.Changes[0].Action != install.ActionMerge {
		t.Errorf("first install into an existing config: got %s, want merge", report.Changes[0].Action)
	}

	config := readJSON(t, filepath.Join(dir, "opencode.json"))
	if config["theme"] != "opencode" {
		t.Errorf("unrelated key theme was dropped: %v", config["theme"])
	}
	if config["$schema"] != "https://opencode.ai/config.json" {
		t.Errorf("unrelated key $schema was dropped: %v", config["$schema"])
	}
	mcpSection := config["mcp"].(map[string]any)
	if _, ok := mcpSection["existing"]; !ok {
		t.Error("existing MCP entry was dropped")
	}
	if _, ok := mcpSection["playwright"]; !ok {
		t.Error("playwright MCP was not added")
	}
}

// testAdapter exercises engine behaviour without a real harness format.
type testAdapter struct {
	supportsSkills bool
}

func (a testAdapter) ID() string           { return "test" }
func (a testAdapter) Name() string         { return "Test Harness" }
func (a testAdapter) SupportsSkills() bool { return a.supportsSkills }

func (a testAdapter) PlanSkills(string, []model.Skill, fs.FS) ([]harness.File, error) {
	return nil, nil
}

func (a testAdapter) PlanMCPs(string, []model.MCP, fs.FS) ([]harness.File, error) {
	return []harness.File{{Path: "test.json", Content: []byte("{}\n"), Merge: true, Source: "mcps"}}, nil
}

func TestRunWarnsAndSkipsSkillsForUnsupportedHarness(t *testing.T) {
	cat := loadCatalog(t)
	skill, _ := cat.Skill("commit")
	mcp, _ := cat.MCP("playwright-mcp")
	dir := t.TempDir()

	report, err := install.Run(install.Request{
		Target:   dir,
		Adapters: []harness.Adapter{testAdapter{supportsSkills: false}},
		Skills:   []model.Skill{skill},
		MCPs:     []model.MCP{mcp},
		Source:   cat.FS(),
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(report.Warnings) != 1 {
		t.Fatalf("warnings: got %d, want 1", len(report.Warnings))
	}
	if _, err := os.Stat(filepath.Join(dir, ".opencode")); !os.IsNotExist(err) {
		t.Error("skills were written for a harness that does not support them")
	}
}

func assertEmptyDir(t *testing.T, dir string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		names := make([]string, 0, len(entries))
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
		t.Fatalf("dry run wrote files: %v", names)
	}
}
