package install_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/models"
)

func TestUpdatedSkillsInstallWithPortableResources(t *testing.T) {
	cat := loadCatalog(t)
	var skills []models.Skill
	for _, id := range []string{
		"agent-browser", "architect", "arena", "before-and-after", "diagram-design",
		"docs-writer", "frontend-design", "grilling", "how", "humanizer", "sentry",
		"unslop", "why",
	} {
		skill, ok := cat.Skill(id)
		if !ok {
			t.Fatalf("missing skill %s", id)
		}
		skills = append(skills, skill)
	}
	var adapters []harness.Adapter
	for _, id := range []string{"opencode", "claude", "codex", "cursor", "gemini", "windsurf", "amp"} {
		adapter, ok := harness.Get(id)
		if !ok {
			t.Fatalf("missing harness %s", id)
		}
		adapters = append(adapters, adapter)
	}
	target := t.TempDir()
	req := install.Request{Target: target, Adapters: adapters, Skills: skills, Source: cat.FS()}
	if _, err := install.Run(req); err != nil {
		t.Fatal(err)
	}
	links := regexp.MustCompile(`\]\(([^)\s]+)\)`)
	for _, prefix := range []string{".opencode", ".claude", ".agents", ".gemini", ".windsurf"} {
		for _, skill := range skills {
			dir := filepath.Join(target, prefix, "skills", skill.ID)
			body, err := os.ReadFile(filepath.Join(dir, "SKILL.md"))
			if err != nil {
				t.Fatal(err)
			}
			for _, match := range links.FindAllSubmatch(body, -1) {
				link := strings.SplitN(string(match[1]), "#", 2)[0]
				if link == "" || strings.Contains(link, ":") {
					continue
				}
				if _, err := os.Stat(filepath.Join(dir, filepath.FromSlash(link))); err != nil {
					t.Errorf("%s/%s unresolved resource %s: %v", prefix, skill.ID, link, err)
				}
			}
		}
	}
	second, err := install.Run(req)
	if err != nil {
		t.Fatal(err)
	}
	for _, change := range second.Changes {
		if change.Action != install.ActionKeep {
			t.Errorf("repeat install changed %s: %s", change.Path, change.Action)
		}
	}
}
