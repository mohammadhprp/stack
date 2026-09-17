package install_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/mohammadhprp/stack/internal/install"
)

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	path := filepath.Join(dir, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadLockMissing(t *testing.T) {
	if _, err := install.LoadLock(t.TempDir()); !errors.Is(err, install.ErrNoLock) {
		t.Fatalf("got %v, want ErrNoLock", err)
	}
}

func TestLoadLockRejectsUnknownVersion(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, install.LockFile, `{"version": 99, "items": []}`)
	if _, err := install.LoadLock(dir); err == nil {
		t.Fatal("expected an error for an unsupported lock version")
	}
}

func TestValidateAcceptsIntactFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.txt", "hello")
	lock := &install.Lock{
		Version: install.LockVersion,
		Items: []install.LockItem{{
			Harness: "opencode",
			Kind:    install.KindSkill,
			ID:      "demo",
			Source:  "skills/demo",
			Files:   map[string]string{"a.txt": install.Hash([]byte("hello"))},
		}},
	}
	if err := lock.Validate(dir); err != nil {
		t.Fatalf("Validate: %v", err)
	}
}

func TestValidateDetectsTamperedFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.txt", "hello")
	lock := &install.Lock{
		Version: install.LockVersion,
		Items: []install.LockItem{{
			Harness: "opencode",
			Kind:    install.KindSkill,
			ID:      "demo",
			Source:  "skills/demo",
			Files:   map[string]string{"a.txt": install.Hash([]byte("hello"))},
		}},
	}
	writeFile(t, dir, "a.txt", "tampered")
	if err := lock.Validate(dir); err == nil {
		t.Fatal("expected Validate to fail on a tampered file")
	}
}

func TestValidateDetectsTamperedLockHash(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.txt", "hello")
	lock := &install.Lock{
		Version: install.LockVersion,
		Items: []install.LockItem{{
			Harness: "opencode",
			Kind:    install.KindSkill,
			ID:      "demo",
			Source:  "skills/demo",
			Files:   map[string]string{"a.txt": install.Hash([]byte("not-hello"))},
		}},
	}
	if err := lock.Validate(dir); err == nil {
		t.Fatal("expected Validate to fail on a tampered lock hash")
	}
}

func TestSaveAndLoadRoundTrip(t *testing.T) {
	dir := t.TempDir()
	want := &install.Lock{
		Version: install.LockVersion,
		Items: []install.LockItem{{
			Harness: "opencode",
			Kind:    install.KindMCP,
			ID:      "playwright-mcp",
			Source:  "mcps/playwright-mcp",
			Files:   map[string]string{"opencode.json": install.Hash([]byte("{}"))},
		}},
	}
	if err := install.SaveLock(dir, want); err != nil {
		t.Fatalf("SaveLock: %v", err)
	}
	got, err := install.LoadLock(dir)
	if err != nil {
		t.Fatalf("LoadLock: %v", err)
	}
	if len(got.Items) != 1 || got.Items[0].ID != "playwright-mcp" {
		t.Fatalf("round trip mismatch: %+v", got)
	}
}
