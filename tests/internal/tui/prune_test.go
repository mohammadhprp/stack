package tui_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/mohammadhprp/stack/internal/tui"
)

func toggleHarness(m *tui.Model, id string) {
	for i := range m.Harnesses {
		if m.Harnesses[i].ID == id {
			m.Harnesses[i].Selected = !m.Harnesses[i].Selected
		}
	}
}

func TestConfirmListsRemovalsAndCancelWritesNothing(t *testing.T) {
	dir := t.TempDir()
	installInto(t, dir, []string{"opencode", "claude"}, []string{"commit"}, []string{"playwright-mcp"})
	before := snapshotFiles(t, dir)

	m := tui.New(loadCatalog(t), dir)
	toggleHarness(m, "claude")
	m, _ = send(m, "enter", "enter", "enter")
	if m.Stage != tui.StageConfirm {
		t.Fatalf("stage: got %v, want confirm", m.Stage)
	}
	view := m.View()
	t.Logf("confirm view:\n%s", view)
	if !strings.Contains(view, "Will remove:") {
		t.Fatalf("confirm view does not list removals:\n%s", view)
	}
	if !strings.Contains(view, ".claude/skills/commit/SKILL.md") {
		t.Errorf("confirm view missing the claude skill removal:\n%s", view)
	}

	m, _ = send(m, "q")
	if m.Stage != tui.StageCancelled {
		t.Fatalf("stage: got %v, want cancelled", m.Stage)
	}
	after := snapshotFiles(t, dir)
	if !reflect.DeepEqual(before, after) {
		t.Error("cancelling after a prune plan changed the target")
	}
}

func TestConfirmPruneRemovesUnselected(t *testing.T) {
	dir := t.TempDir()
	installInto(t, dir, []string{"opencode", "claude"}, []string{"commit"}, []string{"playwright-mcp"})

	m := tui.New(loadCatalog(t), dir)
	toggleHarness(m, "claude")
	m, _ = send(m, "enter", "enter", "enter", "enter")
	if m.Stage != tui.StageDone {
		t.Fatalf("stage: got %v, want done (err=%v)", m.Stage, m.Err)
	}
	if m.Err != nil {
		t.Fatalf("install error: %v", m.Err)
	}
	if !strings.Contains(m.Summary, "Removed") {
		t.Errorf("summary does not report removals:\n%s", m.Summary)
	}
	if _, err := os.Stat(filepath.Join(dir, ".claude")); !os.IsNotExist(err) {
		t.Error("claude files should have been removed")
	}
	if _, err := os.Stat(filepath.Join(dir, "opencode.json")); err != nil {
		t.Error("opencode config should remain")
	}
}
