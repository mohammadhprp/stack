package install_test

import (
	"io/fs"
	"os"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/mohammadhprp/stack/internal/install"
)

func TestLoadCountsAndSorting(t *testing.T) {
	cat := loadCatalog(t)

	for _, category := range []string{"skills", "mcps"} {
		entries, err := fs.ReadDir(os.DirFS("../../../framework"), category)
		if err != nil {
			t.Fatal(err)
		}
		var want, got []string
		for _, entry := range entries {
			if entry.IsDir() {
				want = append(want, entry.Name())
			}
		}
		if len(want) == 0 {
			t.Fatalf("%s source catalog is empty", category)
		}
		if category == "skills" {
			for _, skill := range cat.Skills() {
				got = append(got, skill.ID)
			}
		} else {
			for _, mcp := range cat.MCPs() {
				got = append(got, mcp.Slug)
			}
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s IDs: got %v, want %v", category, got, want)
		}
	}
}

func TestLoadRejectsMalformedCatalog(t *testing.T) {
	for _, tc := range []struct {
		name string
		path string
		data string
		want string
	}{
		{"missing skill file", "skills/broken/notes.md", "notes", `skill "broken": read SKILL.md`},
		{"missing frontmatter", "skills/broken/SKILL.md", "body", "missing YAML frontmatter"},
		{"unterminated frontmatter", "skills/broken/SKILL.md", "---\nname: broken\n", "unterminated YAML frontmatter"},
		{"invalid yaml", "skills/broken/SKILL.md", "---\nname: [\n---\n", "parse frontmatter"},
		{"missing spec", "mcps/broken/notes.md", "notes", `mcp "broken": read spec.json`},
		{"invalid json", "mcps/broken/spec.json", "{", "parse spec.json"},
		{"invalid mcp type", "mcps/broken/spec.json", `{"id":"broken","type":"other"}`, `invalid type "other"`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			files := fstest.MapFS{
				"skills": {Mode: fs.ModeDir},
				"mcps":   {Mode: fs.ModeDir},
				tc.path:  {Data: []byte(tc.data)},
			}
			_, err := install.Load(files)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("Load error: got %v, want containing %q", err, tc.want)
			}
		})
	}
}

func TestLoadNormalizesMetadataAndCollectsNestedFiles(t *testing.T) {
	files := fstest.MapFS{
		"skills/example/SKILL.md":            {Data: []byte("---\r\ndescription: >\r\n  First line\r\n  second line\r\n---\r\nBody\r\n")},
		"skills/example/references/guide.md": {Data: []byte("guide")},
		"skills/README.md":                   {Data: []byte("not a skill")},
		"mcps/README.md":                     {Data: []byte("not a server")},
		"mcps/example-mcp/spec.json":         {Data: []byte(`{"id":"server","name":"Example MCP","type":"remote","url":"https://example.com/mcp"}`)},
	}
	cat, err := install.Load(files)
	if err != nil {
		t.Fatal(err)
	}
	skill, ok := cat.Skill("example")
	if !ok || len(cat.Skills()) != 1 || skill.Name != "example" || skill.Description != "First line second line" {
		t.Fatalf("unexpected skill metadata: %+v (found=%v)", skill, ok)
	}
	if want := []string{"SKILL.md", "references/guide.md"}; !reflect.DeepEqual(skill.Files, want) {
		t.Errorf("files: got %v, want %v", skill.Files, want)
	}
	mcp, ok := cat.MCP("example-mcp")
	if !ok || len(cat.MCPs()) != 1 || mcp.ID != "server" || mcp.Description != "Example MCP" {
		t.Fatalf("unexpected MCP metadata: %+v (found=%v)", mcp, ok)
	}
	if _, ok := cat.MCP("server"); ok {
		t.Error("MCP lookup must use the catalog slug, not the server key")
	}
	if _, ok := cat.Skill("missing"); ok {
		t.Error("missing skill reported as found")
	}
}

func TestLoadSkillFrontmatterAndFiles(t *testing.T) {
	cat := loadCatalog(t)

	skill, ok := cat.Skill("commit")
	if !ok {
		t.Fatal("skill commit not found")
	}
	if skill.Name != "commit" {
		t.Errorf("name: got %q, want %q", skill.Name, "commit")
	}
	if skill.Description == "" {
		t.Error("description is empty")
	}
	if skill.Dir != "skills/commit" {
		t.Errorf("dir: got %q, want %q", skill.Dir, "skills/commit")
	}
	want := map[string]bool{"SKILL.md": true, "examples.md": true}
	for _, file := range skill.Files {
		delete(want, file)
	}
	if len(want) != 0 {
		t.Errorf("skill files missing: %v", want)
	}
}

func TestLoadWalkthroughSkill(t *testing.T) {
	skill, ok := loadCatalog(t).Skill("walkthrough")
	if !ok {
		t.Fatal("skill walkthrough not found")
	}
	if skill.Name != "walkthrough" || skill.Description == "" {
		t.Fatalf("unexpected walkthrough metadata: %+v", skill)
	}
	want := map[string]bool{
		"SKILL.md": true, "examples.md": true, "references/schema.md": true,
		"scripts/validate_walkthrough_json.py": true, "scripts/render_walkthrough.py": true,
	}
	for _, file := range skill.Files {
		delete(want, file)
	}
	if len(want) != 0 {
		t.Errorf("walkthrough files missing: %v", want)
	}
}

func TestLoadSkillFoldedDescription(t *testing.T) {
	cat := loadCatalog(t)

	skill, ok := cat.Skill("ponytail")
	if !ok {
		t.Fatal("skill ponytail not found")
	}
	if skill.Description == "" {
		t.Fatal("folded description was not parsed")
	}
	if len(skill.Files) == 0 {
		t.Fatal("skill files were not collected")
	}
}

func TestCategoryReadmeIsIgnored(t *testing.T) {
	cat := loadCatalog(t)

	if _, ok := cat.Skill("README.md"); ok {
		t.Error("skills/README.md was loaded as a skill")
	}
	if _, ok := cat.MCP("README.md"); ok {
		t.Error("mcps/README.md was loaded as an MCP")
	}
}

func TestLoadMCPLocalAndRemote(t *testing.T) {
	cat := loadCatalog(t)

	playwright, ok := cat.MCP("playwright-mcp")
	if !ok {
		t.Fatal("mcp playwright-mcp not found")
	}
	if playwright.Slug != "playwright-mcp" {
		t.Errorf("slug: got %q", playwright.Slug)
	}
	if playwright.ID != "playwright" {
		t.Errorf("canonical id: got %q, want %q", playwright.ID, "playwright")
	}
	if playwright.Type != "local" {
		t.Errorf("type: got %q, want local", playwright.Type)
	}
	if got := len(playwright.Command); got != 2 {
		t.Errorf("command length: got %d, want 2", got)
	}
	if playwright.Description == "" {
		t.Error("description is empty")
	}

	supabase, ok := cat.MCP("supabase-mcp")
	if !ok {
		t.Fatal("mcp supabase-mcp not found")
	}
	if supabase.Type != "remote" {
		t.Errorf("type: got %q, want remote", supabase.Type)
	}
	if supabase.URL == "" {
		t.Error("remote url is empty")
	}
}
