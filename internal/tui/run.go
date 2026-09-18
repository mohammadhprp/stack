package tui

import (
	"fmt"
	"io"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mohammadhprp/stack/internal/install"
)

func Run(cat *install.Catalog, target string) error {
	return RunProgram(tea.NewProgram(New(cat, target), tea.WithAltScreen()), os.Stdout)
}

func RunProgram(p *tea.Program, out io.Writer) error {
	final, err := p.Run()
	if err != nil {
		return err
	}
	m, ok := final.(*Model)
	if !ok {
		return nil
	}
	if m.Err != nil {
		return m.Err
	}
	if m.Stage == StageCancelled {
		fmt.Fprintln(out, "Cancelled; nothing written.")
		return nil
	}
	if m.Summary != "" {
		fmt.Fprint(out, m.Summary)
	}
	return nil
}
