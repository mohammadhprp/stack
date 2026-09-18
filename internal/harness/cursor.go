package harness

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/mohammadhprp/stack/internal/config"
	"github.com/mohammadhprp/stack/internal/models"
)

// Cursor's project MCP format, verified against the official docs on
// 2026-09-17: .cursor/mcp.json -> mcpServers, where a stdio entry uses
// command/args/env and a remote entry uses url/headers.
// https://cursor.com/docs/mcp
const (
	cursorID     = "cursor"
	cursorName   = "Cursor"
	cursorConfig = ".cursor/mcp.json"
)

func init() { Register(cursor{}) }

type cursor struct{}

func (cursor) ID() string           { return cursorID }
func (cursor) Name() string         { return cursorName }
func (cursor) SupportsSkills() bool { return false }

func (cursor) PlanSkills(string, []models.Skill, fs.FS) ([]File, error) { return nil, nil }

func (cursor) PlanMCPs(target string, mcps []models.MCP, _ fs.FS) ([]File, error) {
	if len(mcps) == 0 {
		return nil, nil
	}
	entries := make(map[string]any, len(mcps))
	for _, mcp := range mcps {
		entry, err := cursorMCPEntry(mcp)
		if err != nil {
			return nil, err
		}
		entries[mcp.ID] = entry
	}
	out, err := config.MergeJSONSection(filepath.Join(target, cursorConfig), "mcpServers", entries)
	if err != nil {
		return nil, err
	}
	return []File{{
		Path:    cursorConfig,
		Content: out,
		Merge:   true,
		Source:  "mcps",
	}}, nil
}

func cursorMCPEntry(mcp models.MCP) (map[string]any, error) {
	switch mcp.Type {
	case "local":
		if len(mcp.Command) == 0 {
			return nil, fmt.Errorf("cursor: mcp %q is local but has no command", mcp.Slug)
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
			return nil, fmt.Errorf("cursor: mcp %q is remote but has no url", mcp.Slug)
		}
		entry := map[string]any{"url": mcp.URL}
		if len(mcp.Headers) > 0 {
			entry["headers"] = mcp.Headers
		}
		return entry, nil
	default:
		return nil, fmt.Errorf("cursor: mcp %q has unknown type %q", mcp.Slug, mcp.Type)
	}
}
