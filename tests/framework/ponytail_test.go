package framework_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestPonytailDocumentedDebtScan(t *testing.T) {
	body, err := os.ReadFile("../../framework/skills/ponytail/references/ponytail-debt.md")
	if err != nil {
		t.Fatal(err)
	}
	_, block, ok := strings.Cut(string(body), "```bash\n")
	if !ok {
		t.Fatal("debt scan command missing")
	}
	scan, _, ok := strings.Cut(block, "\n```")
	if !ok {
		t.Fatal("debt scan command block unclosed")
	}
	dir := t.TempDir()
	init := exec.Command("git", "init", "-q", dir)
	if output, err := init.CombinedOutput(); err != nil {
		t.Fatalf("git init: %v\n%s", err, output)
	}
	for name, content := range map[string]string{
		".gitignore":       "cache/\n",
		"current.go":       "// shortcut: global lock, split when throughput matters\n",
		"legacy.py":        "# ponytail: linear scan, index when input grows\n",
		"block.c":          "/* shortcut: one worker, expand when queues grow */\n",
		"cache/ignored.go": "// shortcut: ignored\n",
		".git/ignored.go":  "// shortcut: git internals\n",
		"unmarked.txt":     "The shortcut convention is documented here.\n",
	} {
		path := filepath.Join(dir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	command := exec.Command("sh", "-c", scan)
	command.Dir = dir
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("documented scan: %v\n%s", err, output)
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected three current/legacy markers, got %s", output)
	}
	for _, name := range []string{"current.go:1:", "legacy.py:1:", "block.c:1:"} {
		if !strings.Contains(string(output), name) {
			t.Errorf("scan missing %s: %s", name, output)
		}
	}
}
