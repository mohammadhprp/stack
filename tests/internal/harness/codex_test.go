package harness_test

import (
	"bytes"
	"io/fs"
	"os"
	"testing"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/models"
)

func TestCodexPlanSkills(t *testing.T) {
	adapter, _ := harness.Get("codex")
	skill := models.Skill{ID: "commit", Dir: "skills/commit"}
	src := os.DirFS("../../../framework")

	files, err := adapter.PlanSkills(t.TempDir(), []models.Skill{skill}, src)
	if err != nil {
		t.Fatalf("PlanSkills: %v", err)
	}
	got := map[string]bool{}
	for _, f := range files {
		got[f.Path] = true
		if f.Source != "skills/commit" {
			t.Errorf("%s: source got %q", f.Path, f.Source)
		}
	}
	for _, want := range []string{
		".agents/skills/commit/SKILL.md",
		".agents/skills/commit/examples.md",
	} {
		if !got[want] {
			t.Errorf("missing planned file %s", want)
		}
	}
	want, err := fs.ReadFile(src, "skills/commit/SKILL.md")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if f.Path == ".agents/skills/commit/SKILL.md" && !bytes.Equal(f.Content, want) {
			t.Error("SKILL.md content does not match the catalog source")
		}
	}
}

func TestCodexPlanMCPsPathUnchanged(t *testing.T) {
	adapter, _ := harness.Get("codex")
	mcps := []models.MCP{{Slug: "playwright-mcp", MCPSpec: models.MCPSpec{ID: "playwright", Type: "local", Command: []string{"npx", "@playwright/mcp@latest"}}}}
	files, err := adapter.PlanMCPs(t.TempDir(), mcps, nil)
	if err != nil {
		t.Fatalf("PlanMCPs: %v", err)
	}
	if len(files) != 1 || files[0].Path != ".codex/config.toml" {
		t.Fatalf("codex MCP path changed: %+v", files)
	}
}
