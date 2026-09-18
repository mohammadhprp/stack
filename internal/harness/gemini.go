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

// Gemini CLI formats, verified against the official docs on 2026-09-18:
//
//   - Skills: .gemini/skills/<name>/SKILL.md (workspace scope)
//     https://geminicli.com/docs/cli/skills
//   - MCP:    .gemini/settings.json -> "mcpServers"; stdio entries use
//     command/args/env, and a remote entry needs the Streamable HTTP key
//     "httpUrl" plus "headers" ("url" is the legacy SSE transport).
//     https://geminicli.com/docs/tools/mcp-server
const (
	geminiID        = "gemini"
	geminiName      = "Gemini CLI"
	geminiSkillsDir = ".gemini/skills"
	geminiConfig    = ".gemini/settings.json"
)

func init() { Register(gemini{}) }

type gemini struct{}

func (gemini) ID() string           { return geminiID }
func (gemini) Name() string         { return geminiName }
func (gemini) SupportsSkills() bool { return true }

func (gemini) PlanSkills(_ string, skills []models.Skill, src fs.FS) ([]File, error) {
	var files []File
	for _, skill := range skills {
		if skill.Dir == "" {
			return nil, fmt.Errorf("gemini: skill %q has no source directory", skill.ID)
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
				Path:    path.Join(geminiSkillsDir, skill.ID, filepath.ToSlash(rel)),
				Content: data,
				Source:  skill.Dir,
			})
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("gemini: plan skill %q: %w", skill.ID, err)
		}
	}
	return files, nil
}

func (gemini) PlanMCPs(target string, mcps []models.MCP, _ fs.FS) ([]File, error) {
	if len(mcps) == 0 {
		return nil, nil
	}
	entries := make(map[string]any, len(mcps))
	for _, mcp := range mcps {
		entry, err := geminiMCPEntry(mcp)
		if err != nil {
			return nil, err
		}
		entries[mcp.ID] = entry
	}
	out, err := config.MergeJSONSection(filepath.Join(target, filepath.FromSlash(geminiConfig)), "mcpServers", entries)
	if err != nil {
		return nil, err
	}
	return []File{{
		Path:    geminiConfig,
		Content: out,
		Merge:   true,
		Source:  "mcps",
	}}, nil
}

func geminiMCPEntry(mcp models.MCP) (map[string]any, error) {
	switch mcp.Type {
	case "local":
		if len(mcp.Command) == 0 {
			return nil, fmt.Errorf("gemini: mcp %q is local but has no command", mcp.Slug)
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
			return nil, fmt.Errorf("gemini: mcp %q is remote but has no url", mcp.Slug)
		}
		entry := map[string]any{"httpUrl": mcp.URL}
		if len(mcp.Headers) > 0 {
			entry["headers"] = mcp.Headers
		}
		return entry, nil
	default:
		return nil, fmt.Errorf("gemini: mcp %q has unknown type %q", mcp.Slug, mcp.Type)
	}
}
