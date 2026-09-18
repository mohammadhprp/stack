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

// Amp formats, verified against the official docs on 2026-09-18:
//
//   - Skills: .agents/skills/<name>/SKILL.md (project scope)
//     https://ampcode.com/docs/customize/skills
//   - MCP:    .amp/settings.json -> "amp.mcpServers" (the literal key includes
//     the "amp." prefix); local entries use command/args/env, and a remote
//     entry uses url/headers. https://ampcode.com/docs/customize/mcp
const (
	ampID        = "amp"
	ampName      = "Amp"
	ampSkillsDir = ".agents/skills"
	ampConfig    = ".amp/settings.json"
)

func init() { Register(amp{}) }

type amp struct{}

func (amp) ID() string           { return ampID }
func (amp) Name() string         { return ampName }
func (amp) SupportsSkills() bool { return true }

func (amp) PlanSkills(_ string, skills []models.Skill, src fs.FS) ([]File, error) {
	var files []File
	for _, skill := range skills {
		if skill.Dir == "" {
			return nil, fmt.Errorf("amp: skill %q has no source directory", skill.ID)
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
				Path:    path.Join(ampSkillsDir, skill.ID, filepath.ToSlash(rel)),
				Content: data,
				Source:  skill.Dir,
			})
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("amp: plan skill %q: %w", skill.ID, err)
		}
	}
	return files, nil
}

func (amp) PlanMCPs(target string, mcps []models.MCP, _ fs.FS) ([]File, error) {
	if len(mcps) == 0 {
		return nil, nil
	}
	entries := make(map[string]any, len(mcps))
	for _, mcp := range mcps {
		entry, err := ampMCPEntry(mcp)
		if err != nil {
			return nil, err
		}
		entries[mcp.ID] = entry
	}
	out, err := config.MergeJSONSection(filepath.Join(target, filepath.FromSlash(ampConfig)), "amp.mcpServers", entries)
	if err != nil {
		return nil, err
	}
	return []File{{
		Path:    ampConfig,
		Content: out,
		Merge:   true,
		Source:  "mcps",
	}}, nil
}

func ampMCPEntry(mcp models.MCP) (map[string]any, error) {
	switch mcp.Type {
	case "local":
		if len(mcp.Command) == 0 {
			return nil, fmt.Errorf("amp: mcp %q is local but has no command", mcp.Slug)
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
			return nil, fmt.Errorf("amp: mcp %q is remote but has no url", mcp.Slug)
		}
		entry := map[string]any{"url": mcp.URL}
		if len(mcp.Headers) > 0 {
			entry["headers"] = mcp.Headers
		}
		return entry, nil
	default:
		return nil, fmt.Errorf("amp: mcp %q has unknown type %q", mcp.Slug, mcp.Type)
	}
}
