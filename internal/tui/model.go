package tui

import (
	"errors"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mohammadhprp/stack/internal/catalog"
	"github.com/mohammadhprp/stack/internal/harness"
	"github.com/mohammadhprp/stack/internal/install"
	"github.com/mohammadhprp/stack/internal/models"
)

type Stage int

const (
	StageHarness Stage = iota
	StageSkills
	StageMCPs
	StageConfirm
	StageDone
	StageCancelled
)

type Option struct {
	ID       string
	Label    string
	Selected bool
}

type Model struct {
	catalog   *catalog.Catalog
	target    string
	Stage     Stage
	Cursor    int
	Harnesses []Option
	Skills    []Option
	MCPs      []Option
	Summary   string
	Err       error
}

var (
	styleTitle  = lipgloss.NewStyle().Bold(true)
	styleCursor = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	styleHelp   = lipgloss.NewStyle().Faint(true)
	styleError  = lipgloss.NewStyle().Foreground(lipgloss.Color("203"))
)

func New(cat *catalog.Catalog, target string) *Model {
	m := Model{catalog: cat, target: target, Stage: StageHarness}
	for _, a := range harness.All() {
		caps := "mcp"
		if a.SupportsSkills() {
			caps = "skills, mcp"
		}
		m.Harnesses = append(m.Harnesses, Option{ID: a.ID(), Label: fmt.Sprintf("%s (%s)", a.Name(), caps)})
	}
	for _, s := range cat.Skills() {
		m.Skills = append(m.Skills, Option{ID: s.ID, Label: s.ID})
	}
	for _, mcp := range cat.MCPs() {
		m.MCPs = append(m.MCPs, Option{ID: mcp.Slug, Label: mcp.Name})
	}
	return &m
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	if m.Stage == StageDone {
		switch key.String() {
		case "ctrl+c", "q", "enter", "esc":
			return m, tea.Quit
		}
		return m, nil
	}

	switch key.String() {
	case "ctrl+c", "q":
		m.Stage = StageCancelled
		return m, tea.Quit
	case "up", "k":
		m.moveCursor(-1)
	case "down", "j":
		m.moveCursor(1)
	case " ":
		m.toggle()
	case "enter":
		return m.advance()
	case "esc":
		return m.back()
	}
	return m, nil
}

func (m *Model) View() string {
	switch m.Stage {
	case StageHarness:
		return m.viewList("Select harnesses", m.Harnesses)
	case StageSkills:
		view := m.viewList("Select skills", m.Skills)
		if note := m.skillSkipNote(); note != "" {
			view += "\n\n" + note
		}
		return view
	case StageMCPs:
		return m.viewList("Select MCP servers", m.MCPs)
	case StageConfirm:
		return m.viewConfirm()
	case StageDone:
		if m.Err != nil {
			return styleError.Render("Error: "+m.Err.Error()) + "\n\n" + styleHelp.Render("enter/q quit")
		}
		return m.Summary + "\n" + styleHelp.Render("enter/q quit")
	case StageCancelled:
		return "Cancelled; nothing written.\n"
	}
	return ""
}

func (m *Model) moveCursor(delta int) {
	n := m.activeLen()
	if n == 0 {
		return
	}
	m.Cursor = (m.Cursor + delta + n) % n
}

func (m *Model) toggle() {
	opts := m.activeOptions()
	if m.Cursor < 0 || m.Cursor >= len(opts) {
		return
	}
	opts[m.Cursor].Selected = !opts[m.Cursor].Selected
}

func (m *Model) advance() (tea.Model, tea.Cmd) {
	switch m.Stage {
	case StageHarness:
		if len(m.selectedIDs(m.Harnesses)) == 0 {
			return m, nil
		}
		if m.skillsRelevant() {
			m.Stage = StageSkills
		} else {
			m.Stage = StageMCPs
		}
	case StageSkills:
		m.Stage = StageMCPs
	case StageMCPs:
		m.Stage = StageConfirm
	case StageConfirm:
		m.runInstall()
	}
	m.Cursor = 0
	return m, nil
}

func (m *Model) back() (tea.Model, tea.Cmd) {
	switch m.Stage {
	case StageSkills:
		m.Stage = StageHarness
	case StageMCPs:
		if m.skillsRelevant() {
			m.Stage = StageSkills
		} else {
			m.Stage = StageHarness
		}
	case StageConfirm:
		m.Stage = StageMCPs
	}
	m.Cursor = 0
	return m, nil
}

func (m *Model) runInstall() {
	m.Stage = StageDone

	adapters := make([]harness.Adapter, 0)
	for _, id := range m.selectedIDs(m.Harnesses) {
		if a, ok := harness.Get(id); ok {
			adapters = append(adapters, a)
		}
	}
	skills := make([]models.Skill, 0)
	for _, id := range m.selectedIDs(m.Skills) {
		if s, ok := m.catalog.Skill(id); ok {
			skills = append(skills, s)
		}
	}
	mcps := make([]models.MCP, 0)
	for _, id := range m.selectedIDs(m.MCPs) {
		if mcp, ok := m.catalog.MCP(id); ok {
			mcps = append(mcps, mcp)
		}
	}

	if len(skills) == 0 && len(mcps) == 0 {
		m.Err = errors.New("nothing selected; choose at least one skill or MCP")
		return
	}

	report, err := install.Run(install.Request{
		Target:   m.target,
		Adapters: adapters,
		Skills:   skills,
		MCPs:     mcps,
		Source:   m.catalog.FS(),
	})
	if err != nil {
		m.Err = err
		return
	}
	m.Summary = summarize(report, m.target)
}

func (m Model) activeLen() int {
	switch m.Stage {
	case StageHarness:
		return len(m.Harnesses)
	case StageSkills:
		return len(m.Skills)
	case StageMCPs:
		return len(m.MCPs)
	}
	return 0
}

func (m Model) activeOptions() []Option {
	switch m.Stage {
	case StageHarness:
		return m.Harnesses
	case StageSkills:
		return m.Skills
	case StageMCPs:
		return m.MCPs
	}
	return nil
}

func (m Model) selectedIDs(opts []Option) []string {
	var ids []string
	for _, o := range opts {
		if o.Selected {
			ids = append(ids, o.ID)
		}
	}
	return ids
}

func (m Model) skillsRelevant() bool {
	for _, id := range m.selectedIDs(m.Harnesses) {
		if a, ok := harness.Get(id); ok && a.SupportsSkills() {
			return true
		}
	}
	return false
}

func (m Model) skillSkipNote() string {
	var names []string
	for _, id := range m.selectedIDs(m.Harnesses) {
		if a, ok := harness.Get(id); ok && !a.SupportsSkills() {
			names = append(names, a.Name())
		}
	}
	if len(names) == 0 {
		return ""
	}
	return fmt.Sprintf("Note: %s do not support skills; skills will be skipped for them.", strings.Join(names, ", "))
}

func (m Model) viewList(title string, opts []Option) string {
	var b strings.Builder
	b.WriteString(styleTitle.Render(title) + "\n\n")
	for i, o := range opts {
		cursor := "  "
		if i == m.Cursor {
			cursor = "> "
		}
		mark := "[ ]"
		if o.Selected {
			mark = "[x]"
		}
		line := fmt.Sprintf("%s%s %s", cursor, mark, o.Label)
		if i == m.Cursor {
			line = styleCursor.Render(line)
		}
		b.WriteString(line + "\n")
	}
	b.WriteString("\n" + styleHelp.Render("space select · enter next · esc back · q quit"))
	return b.String()
}

func (m Model) viewConfirm() string {
	var b strings.Builder
	b.WriteString(styleTitle.Render("Confirm install") + "\n\n")
	b.WriteString("Harnesses: " + strings.Join(m.selectedIDs(m.Harnesses), ", ") + "\n")
	b.WriteString(fmt.Sprintf("Skills:    %s\n", joinOrNone(m.selectedIDs(m.Skills))))
	b.WriteString(fmt.Sprintf("MCPs:      %s\n", joinOrNone(m.selectedIDs(m.MCPs))))
	b.WriteString("\nTarget: " + m.target + "\n")
	b.WriteString("\n" + styleHelp.Render("enter install · esc back · q quit"))
	return b.String()
}

func joinOrNone(ids []string) string {
	if len(ids) == 0 {
		return "(none)"
	}
	return strings.Join(ids, ", ")
}

func summarize(report *install.Report, target string) string {
	var b strings.Builder
	writes := 0
	for _, c := range report.Changes {
		if c.Action != install.ActionKeep {
			writes++
		}
	}
	fmt.Fprintf(&b, "Installed %d file(s) into %s\n", writes, target)
	for _, c := range report.Changes {
		fmt.Fprintf(&b, "  %-8s %s\n", c.Action, c.Path)
	}
	for _, w := range report.Warnings {
		fmt.Fprintf(&b, "warning: %s\n", w)
	}
	return b.String()
}
