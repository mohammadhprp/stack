package harness_test

import (
	"os"
	"testing"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/models"
)

func TestWindsurfPlanMCPsPlansNothing(t *testing.T) {
	adapter, _ := harness.Get("windsurf")
	mcps := []models.MCP{
		{Slug: "playwright-mcp", MCPSpec: models.MCPSpec{ID: "playwright", Type: "local", Command: []string{"npx", "@playwright/mcp@latest"}}},
		{Slug: "supabase-mcp", MCPSpec: models.MCPSpec{ID: "supabase", Type: "remote", URL: "https://mcp.supabase.com/mcp"}},
	}

	files, err := adapter.PlanMCPs(t.TempDir(), mcps, nil)
	if err != nil {
		t.Fatalf("PlanMCPs: %v", err)
	}
	if len(files) != 0 {
		t.Fatalf("Windsurf MCP is user-level only; PlanMCPs planned %d file(s): %+v", len(files), files)
	}
}

func TestWindsurfReportsGlobalOnlyMCP(t *testing.T) {
	adapter, _ := harness.Get("windsurf")
	noted, ok := adapter.(harness.MCPConfigReporter)
	if !ok {
		t.Fatal("windsurf must report its user-level-only MCP config")
	}
	if noted.MCPConfigNote() == "" {
		t.Error("MCP config note is empty")
	}
}

func TestWindsurfPlanSkills(t *testing.T) {
	adapter, _ := harness.Get("windsurf")
	skill := models.Skill{ID: "commit", Dir: "skills/commit"}
	files, err := adapter.PlanSkills(t.TempDir(), []models.Skill{skill}, os.DirFS("../../../framework"))
	if err != nil {
		t.Fatalf("PlanSkills: %v", err)
	}
	got := map[string]bool{}
	for _, f := range files {
		got[f.Path] = true
	}
	for _, want := range []string{
		".windsurf/skills/commit/SKILL.md",
		".windsurf/skills/commit/examples.md",
	} {
		if !got[want] {
			t.Errorf("missing planned file %s", want)
		}
	}
}
