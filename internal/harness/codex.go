package harness

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"

	"github.com/mohammadhprp/stack/internal/model"
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

func (codex) PlanSkills(string, []model.Skill, fs.FS) ([]File, error) { return nil, nil }

func (codex) PlanMCPs(target string, mcps []model.MCP, _ fs.FS) ([]File, error) {
	if len(mcps) == 0 {
		return nil, nil
	}

	root := map[string]any{}
	switch data, err := os.ReadFile(filepath.Join(target, filepath.FromSlash(codexConfig))); {
	case err == nil:
		if len(bytes.TrimSpace(data)) > 0 {
			if err := toml.Unmarshal(data, &root); err != nil {
				return nil, fmt.Errorf("codex: parse %s: %w", codexConfig, err)
			}
		}
	case errors.Is(err, fs.ErrNotExist):
	default:
		return nil, fmt.Errorf("codex: read %s: %w", codexConfig, err)
	}

	servers := map[string]any{}
	if raw, ok := root["mcp_servers"]; ok {
		existing, ok := raw.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("codex: mcp_servers in %s is not a table", codexConfig)
		}
		servers = existing
	}
	for _, mcp := range mcps {
		entry, err := codexMCPEntry(mcp)
		if err != nil {
			return nil, err
		}
		servers[mcp.ID] = entry
	}
	root["mcp_servers"] = servers

	out, err := toml.Marshal(root)
	if err != nil {
		return nil, fmt.Errorf("codex: encode %s: %w", codexConfig, err)
	}
	if len(out) == 0 || out[len(out)-1] != '\n' {
		out = append(out, '\n')
	}

	return []File{{Path: codexConfig, Content: out, Merge: true, Source: "mcps"}}, nil
}

func codexMCPEntry(mcp model.MCP) (map[string]any, error) {
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
