package harness

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/mohammadhprp/stack/internal/config"
	"github.com/mohammadhprp/stack/internal/models"
)

// Codex project config, verified against the official reference on 2026-09-17:
// .codex/config.toml -> [mcp_servers.<name>]; stdio entries use command/args/
// env, and a remote entry uses url with http_headers.
// https://developers.openai.com/codex/config-reference
const (
	codexID     = "codex"
	codexName   = "Codex"
	codexConfig = ".codex/config.toml"
)

func init() { Register(codex{}) }

type codex struct{}

func (codex) ID() string           { return codexID }
func (codex) Name() string         { return codexName }
func (codex) SupportsSkills() bool { return false }

func (codex) PlanSkills(string, []models.Skill, fs.FS) ([]File, error) { return nil, nil }

func (codex) PlanMCPs(target string, mcps []models.MCP, _ fs.FS) ([]File, error) {
	if len(mcps) == 0 {
		return nil, nil
	}

	entries := make(map[string]any, len(mcps))
	for _, mcp := range mcps {
		entry, err := codexMCPEntry(mcp)
		if err != nil {
			return nil, err
		}
		entries[mcp.ID] = entry
	}

	out, err := config.MergeTOMLTable(filepath.Join(target, filepath.FromSlash(codexConfig)), "mcp_servers", entries)
	if err != nil {
		return nil, err
	}

	return []File{{Path: codexConfig, Content: out, Merge: true, Source: "mcps"}}, nil
}

func codexMCPEntry(mcp models.MCP) (map[string]any, error) {
	switch mcp.Type {
	case "local":
		if len(mcp.Command) == 0 {
			return nil, fmt.Errorf("codex: mcp %q is local but has no command", mcp.Slug)
		}
		entry := map[string]any{"command": mcp.Command[0]}
		if len(mcp.Command) > 1 {
			entry["args"] = mcp.Command[1:]
		}
		if len(mcp.Env) > 0 {
			entry["env"] = mcp.Env
		}
		return entry, nil
	case "remote":
		if mcp.URL == "" {
			return nil, fmt.Errorf("codex: mcp %q is remote but has no url", mcp.Slug)
		}
		entry := map[string]any{"url": mcp.URL}
		if len(mcp.Headers) > 0 {
			entry["http_headers"] = mcp.Headers
		}
		return entry, nil
	default:
		return nil, fmt.Errorf("codex: mcp %q has unknown type %q", mcp.Slug, mcp.Type)
	}
}
