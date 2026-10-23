package install_test

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/models"
)

func TestPonytailInstalledResources(t *testing.T) {
	cat := loadCatalog(t)
	skill, ok := cat.Skill("ponytail")
	if !ok {
		t.Fatal("ponytail missing from catalog")
	}
	for _, tc := range []struct {
		id     string
		prefix string
	}{
		{"opencode", ".opencode"},
		{"claude", ".claude"},
		{"codex", ".agents"},
		{"cursor", ".agents"},
		{"gemini", ".gemini"},
		{"windsurf", ".windsurf"},
		{"amp", ".agents"},
	} {
		t.Run(tc.id, func(t *testing.T) {
			adapter, ok := harness.Get(tc.id)
			if !ok {
				t.Fatalf("missing harness %s", tc.id)
			}
			target := t.TempDir()
			req := install.Request{
				Target:   target,
				Adapters: []harness.Adapter{adapter},
				Skills:   []models.Skill{skill},
				Source:   cat.FS(),
			}
			if _, err := install.Run(req); err != nil {
				t.Fatal(err)
			}
			dir := filepath.Join(target, tc.prefix, "skills", "ponytail")
			body, err := os.ReadFile(filepath.Join(dir, "SKILL.md"))
			if err != nil {
				t.Fatal(err)
			}
			resources := []string{"examples.md", "LICENSE"}
			for _, match := range regexp.MustCompile("`(references/[^`]+)`").FindAllSubmatch(body, -1) {
				resources = append(resources, string(match[1]))
			}
			if len(resources) < 10 {
				t.Fatal("missing ponytail operation references")
			}
			for _, resource := range resources {
				if _, err := os.Stat(filepath.Join(dir, filepath.FromSlash(resource))); err != nil {
					t.Errorf("installed resource %s: %v", resource, err)
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
		})
	}
}
