package tui_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/tui"
)

func loadCatalog(t *testing.T) *install.Catalog {
	t.Helper()
	cat, err := install.Load(os.DirFS("../../../framework"))
	if err != nil {
		t.Fatalf("install.Load: %v", err)
	}
	return cat
}

func key(name string) tea.KeyMsg {
	switch name {
	case "enter":
		return tea.KeyMsg{Type: tea.KeyEnter}
	case "esc":
		return tea.KeyMsg{Type: tea.KeyEsc}
	case "up":
		return tea.KeyMsg{Type: tea.KeyUp}
	case "down":
		return tea.KeyMsg{Type: tea.KeyDown}
	default:
		return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(name)}
	}
}

func send(m *tui.Model, keys ...string) (*tui.Model, tea.Cmd) {
	var model tea.Model = m
	var last tea.Cmd
	for _, k := range keys {
		model, last = model.Update(key(k))
	}
	return model.(*tui.Model), last
}

func TestNewStartsAtHarnessStage(t *testing.T) {
	m := tui.New(loadCatalog(t), t.TempDir())
	if m.Stage != tui.StageHarness {
		t.Fatalf("stage: got %v, want harness", m.Stage)
	}
	if len(m.Harnesses) != 4 {
		t.Fatalf("harnesses: got %d, want 4", len(m.Harnesses))
	}
	if len(m.Skills) == 0 || len(m.MCPs) == 0 {
		t.Fatal("catalog selections were not populated")
	}
}

func TestCodexOnlySkipsSkillStage(t *testing.T) {
	// harness.All() is sorted: claude, codex, cursor, opencode.
	m := tui.New(loadCatalog(t), t.TempDir())
	m, _ = send(m, "down", " ", "enter")
	if m.Stage != tui.StageMCPs {
		t.Fatalf("stage: got %v, want MCPs (skills are not supported by codex)", m.Stage)
	}
}

func TestSkillsStageNotesSkippingHarness(t *testing.T) {
	m := tui.New(loadCatalog(t), t.TempDir())
	m, _ = send(m, " ", "down", " ", "enter")
	if m.Stage != tui.StageSkills {
		t.Fatalf("stage: got %v, want skills", m.Stage)
	}
	if !strings.Contains(m.View(), "will be skipped") {
		t.Errorf("skills view missing the codex/cursor note:\n%s", m.View())
	}
}

func TestFullFlowInstalls(t *testing.T) {
	dir := t.TempDir()
	m := tui.New(loadCatalog(t), dir)

	m, _ = send(m, "down", "down", "down", " ") // select opencode
	m, _ = send(m, "enter")
	if m.Stage != tui.StageSkills {
		t.Fatalf("stage after harness: got %v, want skills", m.Stage)
	}

	m, _ = send(m, " ")     // select the first skill
	m, _ = send(m, "enter") // to MCPs
	m, _ = send(m, " ")     // select the first MCP
	m, _ = send(m, "enter")
	if m.Stage != tui.StageConfirm {
		t.Fatalf("stage before install: got %v, want confirm", m.Stage)
	}

	m, _ = send(m, "enter")
	if m.Stage != tui.StageDone {
		t.Fatalf("stage after install: got %v, want done", m.Stage)
	}
	if m.Err != nil {
		t.Fatalf("install error: %v", m.Err)
	}
	if !strings.Contains(m.Summary, "Installed") {
		t.Errorf("summary missing install result:\n%s", m.Summary)
	}
	if _, err := os.Stat(filepath.Join(dir, "opencode.json")); err != nil {
		t.Errorf("opencode.json not written: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".stack-lock.json")); err != nil {
		t.Errorf("lockfile not written: %v", err)
	}
}

func TestQuitCancelsWithoutWriting(t *testing.T) {
	dir := t.TempDir()
	m := tui.New(loadCatalog(t), dir)
	m, cmd := send(m, "q")
	if m.Stage != tui.StageCancelled {
		t.Fatalf("stage: got %v, want cancelled", m.Stage)
	}
	if cmd == nil {
		t.Error("expected a quit command")
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Errorf("cancelling wrote files: %v", entries)
	}
}

func TestRunProgramHeadless(t *testing.T) {
	dir := t.TempDir()
	in, writer := io.Pipe()
	p := tea.NewProgram(tui.New(loadCatalog(t), dir),
		tea.WithInput(in),
		tea.WithOutput(io.Discard),
		tea.WithoutRenderer(),
	)

	var out bytes.Buffer
	done := make(chan error, 1)
	go func() { done <- tui.RunProgram(p, &out) }()

	keys := []string{
		"\x1b[B", "\x1b[B", "\x1b[B", " ", "\r", // select opencode, next
		" ", "\r", // select first skill, next
		" ", "\r", // select first MCP, next
		"\r", // install
		"\r", // quit
	}
	for _, k := range keys {
		if _, err := writer.Write([]byte(k)); err != nil {
			t.Fatalf("write input: %v", err)
		}
	}
	writer.Close()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("RunProgram: %v", err)
		}
	case <-time.After(15 * time.Second):
		t.Fatal("RunProgram did not exit")
	}

	if !strings.Contains(out.String(), "Installed") {
		t.Errorf("summary was not printed after the TUI exited:\n%s", out.String())
	}
	if _, err := os.Stat(filepath.Join(dir, "opencode.json")); err != nil {
		t.Errorf("opencode.json not written: %v", err)
	}
}
