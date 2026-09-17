package harness

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mohammadhprp/stack/internal/model"
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

func (opencode) PlanSkills(_ string, skills []model.Skill, src fs.FS) ([]File, error) {
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

func (opencode) PlanMCPs(target string, mcps []model.MCP, _ fs.FS) ([]File, error) {
	if len(mcps) == 0 {
		return nil, nil
	}

	configPath := filepath.Join(target, opencodeConfig)
	root := map[string]json.RawMessage{}
	switch data, err := os.ReadFile(configPath); {
	case err == nil:
		if len(bytes.TrimSpace(data)) > 0 {
			if err := json.Unmarshal(data, &root); err != nil {
				return nil, fmt.Errorf("opencode: parse %s: %w", opencodeConfig, err)
			}
		}
	case errors.Is(err, fs.ErrNotExist):
	default:
		return nil, fmt.Errorf("opencode: read %s: %w", opencodeConfig, err)
	}

	section := map[string]json.RawMessage{}
	if raw, ok := root["mcp"]; ok && len(bytes.TrimSpace(raw)) > 0 {
		if err := json.Unmarshal(raw, &section); err != nil {
			return nil, fmt.Errorf("opencode: parse mcp section in %s: %w", opencodeConfig, err)
		}
	}
	for _, mcp := range mcps {
		entry, err := renderOpenCodeMCP(mcp)
		if err != nil {
			return nil, err
		}
		section[mcp.ID] = entry
	}

	rawSection, err := json.Marshal(section)
	if err != nil {
		return nil, fmt.Errorf("opencode: encode mcp section: %w", err)
	}
	root["mcp"] = rawSection

	out, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("opencode: encode %s: %w", opencodeConfig, err)
	}
	out = append(out, '\n')

	return []File{{
		Path:    opencodeConfig,
		Content: out,
		Merge:   true,
		Source:  "mcps",
	}}, nil
}

func renderOpenCodeMCP(mcp model.MCP) (json.RawMessage, error) {
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
	return json.Marshal(entry)
}
