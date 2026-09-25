package install_test

import (
	"testing"
)

func TestLoadCountsAndSorting(t *testing.T) {
	cat := loadCatalog(t)

	if got, want := len(cat.Skills()), 37; got != want {
		t.Errorf("skills: got %d, want %d", got, want)
	}
	if got, want := len(cat.MCPs()), 5; got != want {
		t.Errorf("mcps: got %d, want %d", got, want)
	}

	for i := 1; i < len(cat.Skills()); i++ {
		if cat.Skills()[i-1].ID > cat.Skills()[i].ID {
			t.Fatalf("skills are not sorted at %d: %q > %q", i, cat.Skills()[i-1].ID, cat.Skills()[i].ID)
		}
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
