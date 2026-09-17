package harness_test

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/model"
)

func TestRegistry(t *testing.T) {
	all := harness.All()
	wantIDs := []string{"claude", "codex", "cursor", "opencode"}
	if len(all) != len(wantIDs) {
		t.Fatalf("registered adapters: got %d, want %d", len(all), len(wantIDs))
	}
	for i, want := range wantIDs {
		if all[i].ID() != want {
			t.Errorf("adapter %d: got %q, want %q", i, all[i].ID(), want)
		}
	}

	capabilities := map[string]bool{
		"claude":   true,
		"codex":    false,
		"cursor":   false,
		"opencode": true,
	}
	for id, want := range capabilities {
		adapter, ok := harness.Get(id)
		if !ok {
			t.Fatalf("harness %q not registered", id)
		}
		if adapter.SupportsSkills() != want {
			t.Errorf("%s.SupportsSkills(): got %v, want %v", id, adapter.SupportsSkills(), want)
		}
		if adapter.Name() == "" {
			t.Errorf("%s.Name() is empty", id)
		}
	}
}

func TestStubsReportNotImplemented(t *testing.T) {
	for _, id := range []string{"claude", "codex", "cursor"} {
		adapter, _ := harness.Get(id)
		if _, err := adapter.PlanMCPs(t.TempDir(), []model.MCP{{Slug: "x"}}, nil); !errors.Is(err, harness.ErrNotImplemented) {
			t.Errorf("%s.PlanMCPs error: got %v, want ErrNotImplemented", id, err)
		}
		if _, err := adapter.PlanSkills(t.TempDir(), []model.Skill{{ID: "x"}}, nil); !errors.Is(err, harness.ErrNotImplemented) {
			t.Errorf("%s.PlanSkills error: got %v, want ErrNotImplemented", id, err)
		}
	}
}

func TestOpenCodePlanSkills(t *testing.T) {
	adapter, _ := harness.Get("opencode")
	skill := model.Skill{ID: "commit", Dir: "skills/commit"}

	files, err := adapter.PlanSkills(t.TempDir(), []model.Skill{skill}, os.DirFS("../../../framework"))
	if err != nil {
		t.Fatalf("PlanSkills: %v", err)
	}

	got := map[string]bool{}
	for _, f := range files {
		got[f.Path] = true
		if len(f.Content) == 0 {
			t.Errorf("%s: empty content", f.Path)
		}
		if f.Source != "skills/commit" {
			t.Errorf("%s: source got %q", f.Path, f.Source)
		}
	}
	for _, want := range []string{
		".opencode/skills/commit/SKILL.md",
		".opencode/skills/commit/examples.md",
	} {
		if !got[want] {
			t.Errorf("missing planned file %s", want)
		}
	}
}

func TestOpenCodePlanMCPsRendersLocalAndRemote(t *testing.T) {
	adapter, _ := harness.Get("opencode")
	mcps := []model.MCP{
		{
			Slug: "playwright-mcp",
			MCPSpec: model.MCPSpec{
				ID:      "playwright",
				Type:    "local",
				Command: []string{"npx", "@playwright/mcp@latest"},
				Env:     map[string]string{"TOKEN": "abc"},
			},
		},
		{
			Slug: "supabase-mcp",
			MCPSpec: model.MCPSpec{
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
	if files[0].Path != "opencode.json" {
		t.Errorf("path: got %q", files[0].Path)
	}
	if !files[0].Merge {
		t.Error("config file should be marked Merge")
	}

	var root struct {
		MCP map[string]map[string]any `json:"mcp"`
	}
	if err := json.Unmarshal(files[0].Content, &root); err != nil {
		t.Fatalf("decode config: %v", err)
	}
	local := root.MCP["playwright"]
	if local["type"] != "local" {
		t.Errorf("local type: got %v", local["type"])
	}
	if _, ok := local["environment"]; !ok {
		t.Error("local env was not rendered as the OpenCode \"environment\" key")
	}
	if _, ok := local["env"]; ok {
		t.Error("canonical \"env\" key leaked into the OpenCode config")
	}
	remote := root.MCP["supabase"]
	if remote["type"] != "remote" {
		t.Errorf("remote type: got %v", remote["type"])
	}
	if remote["url"] != "https://mcp.supabase.com/mcp" {
		t.Errorf("remote url: got %v", remote["url"])
	}
}

func TestOpenCodePlanMCPsRejectsBadSpec(t *testing.T) {
	adapter, _ := harness.Get("opencode")
	_, err := adapter.PlanMCPs(t.TempDir(), []model.MCP{{Slug: "bad", MCPSpec: model.MCPSpec{ID: "bad", Type: "telepathy"}}}, nil)
	if err == nil {
		t.Fatal("expected an error for an unknown MCP type")
	}
}
