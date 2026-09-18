package cmd_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDoctorNoLock(t *testing.T) {
	dir := t.TempDir()
	out, _, err := execute(t, loadCatalog(t), "doctor", "--target", dir)
	if err != nil {
		t.Fatalf("doctor: %v", err)
	}
	if !strings.Contains(out, "nothing installed") {
		t.Errorf("unexpected output:\n%s", out)
	}
}

func TestDoctorClassifiesOKModifiedMissing(t *testing.T) {
	dir := t.TempDir()
	if _, _, err := execute(t, loadCatalog(t), "install",
		"--harness", "opencode",
		"--skills", "commit",
		"--mcp", "playwright-mcp",
		"--target", dir,
	); err != nil {
		t.Fatalf("install: %v", err)
	}

	out, _, err := execute(t, loadCatalog(t), "doctor", "--target", dir)
	if err != nil {
		t.Fatalf("doctor after install: %v\n%s", err, out)
	}
	if !strings.Contains(out, "0 modified") || !strings.Contains(out, "0 missing") {
		t.Errorf("clean install should have no problems:\n%s", out)
	}

	skillFile := filepath.Join(dir, ".opencode", "skills", "commit", "SKILL.md")
	if err := os.WriteFile(skillFile, []byte("user edit\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(dir, ".opencode", "skills", "commit", "examples.md")); err != nil {
		t.Fatal(err)
	}

	out, _, err = execute(t, loadCatalog(t), "doctor", "--target", dir)
	if err == nil {
		t.Fatalf("expected a non-zero exit for modified/missing files:\n%s", out)
	}
	if !strings.Contains(out, "modified") {
		t.Errorf("output missing a modified classification:\n%s", out)
	}
	if !strings.Contains(out, "missing") {
		t.Errorf("output missing a missing classification:\n%s", out)
	}
}
