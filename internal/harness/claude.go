package harness

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/mohammadhprp/stack/internal/model"
)

// Claude Code formats, verified against the official docs on 2026-09-17:
//
//   - Skills: .claude/skills/<name>/SKILL.md
//     https://docs.claude.com/en/docs/claude-code/skills
//   - MCP:    .mcp.json -> mcpServers; stdio entries use command/args/env, and
//     a url entry requires an explicit "type" (http is the recommended
//     transport). https://docs.claude.com/en/docs/claude-code/mcp
const (
	claudeID        = "claude"
	claudeName      = "Claude Code"
	claudeSkillsDir = ".claude/skills"
	claudeConfig    = ".mcp.json"
)

func init() { Register(claude{}) }

type claude struct{}

func (claude) ID() string           { return claudeID }
func (claude) Name() string         { return claudeName }
func (claude) SupportsSkills() bool { return true }

func (claude) PlanSkills(_ string, skills []model.Skill, src fs.FS) ([]File, error) {
	var files []File
	for _, skill := range skills {
		if skill.Dir == "" {
			return nil, fmt.Errorf("claude: skill %q has no source directory", skill.ID)
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
				Path:    path.Join(claudeSkillsDir, skill.ID, filepath.ToSlash(rel)),
				Content: data,
				Source:  skill.Dir,
			})
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("claude: plan skill %q: %w", skill.ID, err)
		}
	}
	return files, nil
}

func (claude) PlanMCPs(target string, mcps []model.MCP, _ fs.FS) ([]File, error) {
	entries := make(map[string]any, len(mcps))
	for _, mcp := range mcps {
		entry, err := claudeMCPEntry(mcp)
		if err != nil {
			return nil, err
		}
		entries[mcp.ID] = entry
	}
	return mergeJSONSection(claudeID, target, claudeConfig, "mcpServers", entries)
}

func claudeMCPEntry(mcp model.MCP) (map[string]any, error) {
	switch mcp.Type {
	case "local":
		if len(mcp.Command) == 0 {
			return nil, fmt.Errorf("claude: mcp %q is local but has no command", mcp.Slug)
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
			return nil, fmt.Errorf("claude: mcp %q is remote but has no url", mcp.Slug)
		}
		entry := map[string]any{"type": "http", "url": mcp.URL}
		if len(mcp.Headers) > 0 {
			entry["headers"] = mcp.Headers
		}
		return entry, nil
	default:
		return nil, fmt.Errorf("claude: mcp %q has unknown type %q", mcp.Slug, mcp.Type)
	}
}
