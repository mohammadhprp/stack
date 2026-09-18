// Package catalog roots its filesystem at framework/, holding
// "skills/<id>/SKILL.md" and "mcps/<slug>/spec.json" at the top level.
package catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/mohammadhprp/stack/internal/models"
)

type Catalog struct {
	skills []models.Skill
	mcps   []models.MCP
	fsys   fs.FS
}

// Load ignores README.md at the skills/ and mcps/ roots; they are
// documentation, not entries.
func Load(fsys fs.FS) (*Catalog, error) {
	skills, err := loadSkills(fsys)
	if err != nil {
		return nil, err
	}
	mcps, err := loadMCPs(fsys)
	if err != nil {
		return nil, err
	}
	return &Catalog{skills: skills, mcps: mcps, fsys: fsys}, nil
}

func (c *Catalog) Skills() []models.Skill { return c.skills }

func (c *Catalog) MCPs() []models.MCP { return c.mcps }

func (c *Catalog) FS() fs.FS { return c.fsys }

func (c *Catalog) Skill(id string) (models.Skill, bool) {
	for _, s := range c.skills {
		if s.ID == id {
			return s, true
		}
	}
	return models.Skill{}, false
}

func (c *Catalog) MCP(slug string) (models.MCP, bool) {
	for _, m := range c.mcps {
		if m.Slug == slug {
			return m, true
		}
	}
	return models.MCP{}, false
}

func loadSkills(fsys fs.FS) ([]models.Skill, error) {
	entries, err := fs.ReadDir(fsys, "skills")
	if err != nil {
		return nil, fmt.Errorf("read skills directory: %w", err)
	}
	var skills []models.Skill
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		id := entry.Name()
		dir := "skills/" + id
		raw, err := fs.ReadFile(fsys, dir+"/SKILL.md")
		if err != nil {
			return nil, fmt.Errorf("skill %q: read SKILL.md: %w", id, err)
		}
		name, description, err := frontmatter(raw)
		if err != nil {
			return nil, fmt.Errorf("skill %q: %w", id, err)
		}
		if name == "" {
			name = id
		}
		files, err := dirFiles(fsys, dir)
		if err != nil {
			return nil, fmt.Errorf("skill %q: %w", id, err)
		}
		skills = append(skills, models.Skill{
			ID:          id,
			Name:        name,
			Description: description,
			Dir:         dir,
			Files:       files,
		})
	}
	sort.Slice(skills, func(i, j int) bool { return skills[i].ID < skills[j].ID })
	return skills, nil
}

func loadMCPs(fsys fs.FS) ([]models.MCP, error) {
	entries, err := fs.ReadDir(fsys, "mcps")
	if err != nil {
		return nil, fmt.Errorf("read mcps directory: %w", err)
	}
	var mcps []models.MCP
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		slug := entry.Name()
		dir := "mcps/" + slug
		raw, err := fs.ReadFile(fsys, dir+"/spec.json")
		if err != nil {
			return nil, fmt.Errorf("mcp %q: read spec.json: %w", slug, err)
		}
		var spec models.MCPSpec
		if err := json.Unmarshal(raw, &spec); err != nil {
			return nil, fmt.Errorf("mcp %q: parse spec.json: %w", slug, err)
		}
		if spec.Type != "local" && spec.Type != "remote" {
			return nil, fmt.Errorf("mcp %q: invalid type %q", slug, spec.Type)
		}
		if spec.Description == "" {
			spec.Description = spec.Name
		}
		mcps = append(mcps, models.MCP{Slug: slug, MCPSpec: spec, Dir: dir})
	}
	sort.Slice(mcps, func(i, j int) bool { return mcps[i].Slug < mcps[j].Slug })
	return mcps, nil
}

func dirFiles(fsys fs.FS, dir string) ([]string, error) {
	var files []string
	err := fs.WalkDir(fsys, dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		files = append(files, strings.TrimPrefix(p, dir+"/"))
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	return files, nil
}

func frontmatter(data []byte) (name, description string, err error) {
	block, err := splitFrontmatter(data)
	if err != nil {
		return "", "", err
	}
	var fm struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
	}
	if err := yaml.Unmarshal(block, &fm); err != nil {
		return "", "", fmt.Errorf("parse frontmatter: %w", err)
	}
	return strings.TrimSpace(fm.Name), strings.Join(strings.Fields(fm.Description), " "), nil
}

func splitFrontmatter(data []byte) ([]byte, error) {
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	if !strings.HasPrefix(text, "---\n") {
		return nil, errors.New("missing YAML frontmatter")
	}
	rest := text[len("---\n"):]
	idx := strings.Index(rest, "\n---")
	if idx < 0 {
		return nil, errors.New("unterminated YAML frontmatter")
	}
	return []byte(rest[:idx]), nil
}
