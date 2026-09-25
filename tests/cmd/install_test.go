package cmd_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mohammadhprp/stack/cmd"
	"github.com/mohammadhprp/stack/internal/install"
)

func loadCatalog(t *testing.T) *install.Catalog {
	t.Helper()
	cat, err := install.Load(os.DirFS("../../framework"))
	if err != nil {
		t.Fatalf("install.Load: %v", err)
	}
	return cat
}

func execute(t *testing.T, cat *install.Catalog, args ...string) (string, string, error) {
	t.Helper()
	root := cmd.NewRootCommand(cat)
	var out, errOut bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&errOut)
	root.SetArgs(args)
	err := root.Execute()
	return out.String(), errOut.String(), err
}

func TestListCommandShowsEveryEntry(t *testing.T) {
	out, _, err := execute(t, loadCatalog(t), "list")
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	for _, want := range []string{"Skills (37):", "MCP servers (5):", "playwright-mcp", "Playwright MCP"} {
		if !strings.Contains(out, want) {
			t.Errorf("list output missing %q", want)
		}
	}
}

func TestInstallCommandDryRunWritesNothing(t *testing.T) {
	dir := t.TempDir()
	out, _, err := execute(t, loadCatalog(t), "install",
		"--harness", "opencode",
		"--skills", "commit",
		"--mcp", "playwright-mcp",
		"--target", dir,
		"--dry-run",
	)
	if err != nil {
		t.Fatalf("install --dry-run: %v", err)
	}
	if !strings.Contains(out, "create") || !strings.Contains(out, "Dry run: no files written.") {
		t.Errorf("unexpected dry-run output:\n%s", out)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Errorf("dry run wrote files: %v", entries)
	}
}

func TestInstallCommandWritesSkillAndMergesConfig(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "opencode.json"), []byte(`{"theme":"opencode"}`), 0o644); err != nil {
		t.Fatal(err)
	}

	out, errOut, err := execute(t, loadCatalog(t), "install",
		"--harness", "opencode",
		"--skills", "commit",
		"--mcp", "playwright-mcp",
		"--target", dir,
	)
	if err != nil {
		t.Fatalf("install: %v\nstderr: %s", err, errOut)
	}
	if !strings.Contains(out, "Installed 3 file(s)") {
		t.Errorf("unexpected install output:\n%s", out)
	}

	if _, err := os.Stat(filepath.Join(dir, ".opencode", "skills", "commit", "SKILL.md")); err != nil {
		t.Errorf("skill not installed: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".stack-lock.json")); err != nil {
		t.Errorf("lockfile not written: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "opencode.json"))
	if err != nil {
		t.Fatal(err)
	}
	var config struct {
		Theme string `json:"theme"`
		MCP   map[string]struct {
			Type    string   `json:"type"`
			Command []string `json:"command"`
		} `json:"mcp"`
	}
	if err := json.Unmarshal(data, &config); err != nil {
		t.Fatalf("decode opencode.json: %v", err)
	}
	if config.Theme != "opencode" {
		t.Errorf("unrelated key dropped: theme=%q", config.Theme)
	}
	entry, ok := config.MCP["playwright"]
	if !ok {
		t.Fatalf("playwright mcp not merged: %v", config.MCP)
	}
	if entry.Type != "local" || len(entry.Command) == 0 {
		t.Errorf("bad playwright entry: %+v", entry)
	}
}

func TestInstallPruneFlagRemovesDeselected(t *testing.T) {
	dir := t.TempDir()
	if _, _, err := execute(t, loadCatalog(t), "install",
		"--harness", "claude,codex", "--skills", "commit", "--target", dir); err != nil {
		t.Fatalf("first install: %v", err)
	}

	out, _, err := execute(t, loadCatalog(t), "install",
		"--harness", "claude", "--skills", "commit", "--target", dir)
	if err != nil {
		t.Fatalf("install without --prune: %v", err)
	}
	if strings.Contains(out, "remove") {
		t.Errorf("install without --prune must not remove anything:\n%s", out)
	}
	if _, err := os.Stat(filepath.Join(dir, ".agents", "skills", "commit", "SKILL.md")); err != nil {
		t.Error("codex skill should remain without --prune")
	}

	out, _, err = execute(t, loadCatalog(t), "install",
		"--harness", "claude", "--skills", "commit", "--prune", "--target", dir)
	if err != nil {
		t.Fatalf("install --prune: %v", err)
	}
	if !strings.Contains(out, "remove") || !strings.Contains(out, "Removed") {
		t.Errorf("--prune output should show removals:\n%s", out)
	}
	if _, err := os.Stat(filepath.Join(dir, ".agents", "skills", "commit", "SKILL.md")); !os.IsNotExist(err) {
		t.Error("codex skill should be removed with --prune")
	}
}
