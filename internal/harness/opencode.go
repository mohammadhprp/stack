package harness

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/mohammadhprp/stack/internal/config"
	"github.com/mohammadhprp/stack/internal/models"
)

// OpenCode's on-disk formats, verified against the official docs on 2026-09-17:
//
//   - Skills:   .opencode/skills/<name>/SKILL.md
//     https://opencode.ai/docs/skills/
//   - MCP:      opencode.json -> "mcp", each entry has "type" of "local" or
//     "remote"; local entries use "command" and "environment"; remote entries
//     use "url" and "headers".
//     https://opencode.ai/docs/mcp-servers/
const (
	opencodeID        = "opencode"
	opencodeName      = "OpenCode"
	opencodeSkillsDir = ".opencode/skills"
	opencodeConfig    = "opencode.json"
)

func init() { Register(opencode{}) }

type opencode struct{}

func (opencode) ID() string           { return opencodeID }
func (opencode) Name() string         { return opencodeName }
func (opencode) SupportsSkills() bool { return true }

func (opencode) PlanSkills(_ string, skills []models.Skill, src fs.FS) ([]File, error) {
	var files []File
	for _, skill := range skills {
		if skill.Dir == "" {
			return nil, fmt.Errorf("opencode: skill %q has no source directory", skill.ID)
		}
		err := fs.WalkDir(src, skill.Dir, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			data, err := fs.ReadFile(src, p)
			if err != nil {
				return err
			}
			rel := strings.TrimPrefix(p, skill.Dir+"/")
			files = append(files, File{
				Path:    path.Join(opencodeSkillsDir, skill.ID, filepath.ToSlash(rel)),
				Content: data,
				Source:  skill.Dir,
			})
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("opencode: plan skill %q: %w", skill.ID, err)
		}
	}
	return files, nil
}

func (opencode) PlanMCPs(target string, mcps []models.MCP, _ fs.FS) ([]File, error) {
	if len(mcps) == 0 {
		return nil, nil
	}

	entries := make(map[string]any, len(mcps))
	for _, mcp := range mcps {
		entry, err := renderOpenCodeMCP(mcp)
		if err != nil {
			return nil, err
		}
		entries[mcp.ID] = entry
	}

	out, err := config.MergeJSONSection(filepath.Join(target, opencodeConfig), "mcp", entries)
	if err != nil {
		return nil, err
	}

	return []File{{
		Path:    opencodeConfig,
		Content: out,
		Merge:   true,
		Source:  "mcps",
	}}, nil
}

func renderOpenCodeMCP(mcp models.MCP) (map[string]any, error) {
	entry := map[string]any{"enabled": true}
	switch mcp.Type {
	case "local":
		if len(mcp.Command) == 0 {
			return nil, fmt.Errorf("opencode: mcp %q is local but has no command", mcp.Slug)
		}
		entry["type"] = "local"
		entry["command"] = mcp.Command
		if len(mcp.Env) > 0 {
			// OpenCode names this key "environment", not the canonical "env".
			entry["environment"] = mcp.Env
		}
	case "remote":
		if mcp.URL == "" {
			return nil, fmt.Errorf("opencode: mcp %q is remote but has no url", mcp.Slug)
		}
		entry["type"] = "remote"
		entry["url"] = mcp.URL
		if len(mcp.Headers) > 0 {
			entry["headers"] = mcp.Headers
		}
	default:
		return nil, fmt.Errorf("opencode: mcp %q has unknown type %q", mcp.Slug, mcp.Type)
	}
	return entry, nil
}
