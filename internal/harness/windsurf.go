package harness

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/mohammadhprp/stack/internal/models"
)

// Windsurf formats, verified against the official docs on 2026-09-18:
//
//   - Skills: .windsurf/skills/<name>/SKILL.md (workspace scope)
//     https://docs.windsurf.com/windsurf/cascade/skills
//   - MCP: user-level only, with no project-level file, so PlanMCPs plans
//     nothing and reports why instead.
//     https://docs.windsurf.com/windsurf/cascade/mcp
const (
	windsurfID        = "windsurf"
	windsurfName      = "Windsurf"
	windsurfSkillsDir = ".windsurf/skills"
	windsurfMCPNote   = "MCP config is user-level only (~/.codeium/windsurf/mcp_config.json); no project-level file to write"
)

func init() { Register(windsurf{}) }

type windsurf struct{}

func (windsurf) ID() string           { return windsurfID }
func (windsurf) Name() string         { return windsurfName }
func (windsurf) SupportsSkills() bool { return true }

func (windsurf) MCPConfigNote() string { return windsurfMCPNote }

func (windsurf) PlanMCPs(string, []models.MCP, fs.FS) ([]File, error) { return nil, nil }

func (windsurf) PlanSkills(_ string, skills []models.Skill, src fs.FS) ([]File, error) {
	var files []File
	for _, skill := range skills {
		if skill.Dir == "" {
			return nil, fmt.Errorf("windsurf: skill %q has no source directory", skill.ID)
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
				Path:    path.Join(windsurfSkillsDir, skill.ID, filepath.ToSlash(rel)),
				Content: data,
				Source:  skill.Dir,
			})
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("windsurf: plan skill %q: %w", skill.ID, err)
		}
	}
	return files, nil
}
