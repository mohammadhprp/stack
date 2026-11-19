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
		"adhd", "agent-browser", "before-and-after", "diagram-design", "docker", "docs-writer",
		"find-docs", "find-skills", "frontend-design", "grill-with-docs", "humanizer",
		"laravel-best-practices", "merge-risk", "perf-web-optimization", "ponytail",
		"resolve-merge-conflicts", "sentry", "show-me",
		"stack-diagram", "tdd", "unslop", "yolo",
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
		for _, file := range []string{
			"grill-with-docs/references/glossary-format.md",
			"grill-with-docs/references/adr-format.md",
			"diagram-design/assets/sample-sequence.puml",
			"diagram-design/scripts/plantuml_extract.py",
			"diagram-design/scripts/verify_geometry.py",
			"stack-diagram/scripts/fetch_icons.py",
			"yolo/licenses/mattpocock-skills.txt",
			"yolo/licenses/ponytail.txt",
			"yolo/licenses/pstack.txt",
			"yolo/licenses/superpowers.txt",
		} {
			if _, err := os.Stat(filepath.Join(target, prefix, "skills", filepath.FromSlash(file))); err != nil {
				t.Errorf("%s missing bundled resource %s: %v", prefix, file, err)
			}
		}
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
