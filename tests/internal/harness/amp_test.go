package harness_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/models"
)

func TestAmpPlanMCPsRendersLocalAndRemote(t *testing.T) {
	adapter, _ := harness.Get("amp")
	mcps := []models.MCP{
		{
			Slug: "playwright-mcp",
			MCPSpec: models.MCPSpec{
				ID:      "playwright",
				Type:    "local",
				Command: []string{"npx", "@playwright/mcp@latest"},
				Env:     map[string]string{"TOKEN": "abc"},
			},
		},
		{
			Slug: "supabase-mcp",
			MCPSpec: models.MCPSpec{
				ID:      "supabase",
				Type:    "remote",
				URL:     "https://mcp.supabase.com/mcp",
				Headers: map[string]string{"X-Key": "v"},
			},
		},
	}

	files, err := adapter.PlanMCPs(t.TempDir(), mcps, nil)
	if err != nil {
		t.Fatalf("PlanMCPs: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("planned files: got %d, want 1", len(files))
	}
	if files[0].Path != ".amp/settings.json" {
		t.Errorf("path: got %q", files[0].Path)
	}
	if !files[0].Merge {
		t.Error("config file should be marked Merge")
	}

	var root struct {
		MCP map[string]map[string]any `json:"amp.mcpServers"`
	}
	if err := json.Unmarshal(files[0].Content, &root); err != nil {
		t.Fatalf("decode config: %v", err)
	}
	local := root.MCP["playwright"]
	if local["command"] != "npx" {
		t.Errorf("local command: got %v", local["command"])
	}
	remote := root.MCP["supabase"]
	if remote["url"] != "https://mcp.supabase.com/mcp" {
		t.Errorf("remote url: got %v", remote["url"])
	}
	headers, ok := remote["headers"].(map[string]any)
	if !ok || headers["X-Key"] != "v" {
		t.Errorf("remote headers: got %v", remote["headers"])
	}
}

func TestAmpPlanMCPsPreservesExistingConfig(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, ".amp"), 0o755); err != nil {
		t.Fatal(err)
	}
	existing := `{
  "amp.other": 1,
  "amp.mcpServers": {
    "keepme": { "command": "echo" }
  }
}`
	if err := os.WriteFile(filepath.Join(dir, ".amp", "settings.json"), []byte(existing), 0o644); err != nil {
		t.Fatal(err)
	}

	adapter, _ := harness.Get("amp")
	mcps := []models.MCP{{Slug: "playwright-mcp", MCPSpec: models.MCPSpec{ID: "playwright", Type: "local", Command: []string{"npx", "@playwright/mcp@latest"}}}}
	files, err := adapter.PlanMCPs(dir, mcps, nil)
	if err != nil {
		t.Fatalf("PlanMCPs: %v", err)
	}

	var root map[string]any
	if err := json.Unmarshal(files[0].Content, &root); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if root["amp.other"] != float64(1) {
		t.Errorf("unrelated key amp.other was dropped: %v", root["amp.other"])
	}
	section := root["amp.mcpServers"].(map[string]any)
	if _, ok := section["keepme"]; !ok {
		t.Error("existing MCP entry was dropped")
	}
	if _, ok := section["playwright"]; !ok {
		t.Error("playwright MCP was not added")
	}
}

func TestAmpPlanMCPsIdempotent(t *testing.T) {
	dir := t.TempDir()
	adapter, _ := harness.Get("amp")
	mcps := []models.MCP{{Slug: "playwright-mcp", MCPSpec: models.MCPSpec{ID: "playwright", Type: "local", Command: []string{"npx", "@playwright/mcp@latest"}}}}

	first, err := adapter.PlanMCPs(dir, mcps, nil)
	if err != nil {
		t.Fatalf("first PlanMCPs: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(dir, ".amp"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".amp", "settings.json"), first[0].Content, 0o644); err != nil {
		t.Fatal(err)
	}
	second, err := adapter.PlanMCPs(dir, mcps, nil)
	if err != nil {
		t.Fatalf("second PlanMCPs: %v", err)
	}
	if !bytes.Equal(first[0].Content, second[0].Content) {
		t.Errorf("re-plan changed the config:\nfirst:\n%s\nsecond:\n%s", first[0].Content, second[0].Content)
	}
}

func TestAmpPlanSkills(t *testing.T) {
	adapter, _ := harness.Get("amp")
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
		".agents/skills/commit/SKILL.md",
		".agents/skills/commit/examples.md",
	} {
		if !got[want] {
			t.Errorf("missing planned file %s", want)
		}
	}
}
