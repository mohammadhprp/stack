package config_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/BurntSushi/toml"

	"github.com/mohammadhprp/stack/internal/config"
)

func TestReconcileJSONSectionKeepsUnrelatedKeys(t *testing.T) {
	path := filepath.Join(t.TempDir(), "settings.json")
	seeded := `{"theme":"dark","mcpServers":{"a":{"command":"x"},"b":{"command":"y"}}}`
	if err := os.WriteFile(path, []byte(seeded), 0o644); err != nil {
		t.Fatal(err)
	}

	out, remove, err := config.ReconcileJSONSection(path, "mcpServers", map[string]any{"b": map[string]any{"command": "y2"}}, []string{"a"})
	if err != nil {
		t.Fatalf("ReconcileJSONSection: %v", err)
	}
	if remove {
		t.Fatal("unexpected remove=true")
	}
	var root map[string]any
	if err := json.Unmarshal(out, &root); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if root["theme"] != "dark" {
		t.Errorf("unrelated key theme dropped: %v", root["theme"])
	}
	section := root["mcpServers"].(map[string]any)
	if _, ok := section["a"]; ok {
		t.Error("dropped entry a is still present")
	}
	if section["b"].(map[string]any)["command"] != "y2" {
		t.Errorf("kept entry b not updated: %v", section["b"])
	}
}

func TestReconcileJSONSectionEmptyRemoves(t *testing.T) {
	path := filepath.Join(t.TempDir(), "settings.json")
	if err := os.WriteFile(path, []byte(`{"mcpServers":{"a":{"command":"x"}}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	_, remove, err := config.ReconcileJSONSection(path, "mcpServers", nil, []string{"a"})
	if err != nil {
		t.Fatalf("ReconcileJSONSection: %v", err)
	}
	if !remove {
		t.Error("expected remove=true when the file is left empty")
	}
}

func TestReconcileTOMLTableKeepsUnrelatedKeys(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.toml")
	seeded := "model = \"o3\"\n\n[mcp_servers.a]\ncommand = \"x\"\n\n[mcp_servers.b]\ncommand = \"y\"\n"
	if err := os.WriteFile(path, []byte(seeded), 0o644); err != nil {
		t.Fatal(err)
	}

	out, remove, err := config.ReconcileTOMLTable(path, "mcp_servers", map[string]any{"b": map[string]any{"command": "y2"}}, []string{"a"})
	if err != nil {
		t.Fatalf("ReconcileTOMLTable: %v", err)
	}
	if remove {
		t.Fatal("unexpected remove=true")
	}
	var root map[string]any
	if err := toml.Unmarshal(out, &root); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if root["model"] != "o3" {
		t.Errorf("unrelated key model dropped: %v", root["model"])
	}
	section := root["mcp_servers"].(map[string]any)
	if _, ok := section["a"]; ok {
		t.Error("dropped table a is still present")
	}
	if section["b"].(map[string]any)["command"] != "y2" {
		t.Errorf("kept table b not updated: %v", section["b"])
	}
}

func TestReconcileTOMLTableEmptyRemoves(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(path, []byte("[mcp_servers.a]\ncommand = \"x\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, remove, err := config.ReconcileTOMLTable(path, "mcp_servers", nil, []string{"a"})
	if err != nil {
		t.Fatalf("ReconcileTOMLTable: %v", err)
	}
	if !remove {
		t.Error("expected remove=true when the file is left empty")
	}
}
