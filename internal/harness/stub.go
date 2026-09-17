package harness

import (
	"fmt"
	"io/fs"

	"github.com/mohammadhprp/stack/internal/model"
)

// stub reports real metadata so listing works but fails loudly when planning.
type stub struct {
	id             string
	name           string
	supportsSkills bool
}

func (s stub) ID() string           { return s.id }
func (s stub) Name() string         { return s.name }
func (s stub) SupportsSkills() bool { return s.supportsSkills }

func (s stub) PlanSkills(string, []model.Skill, fs.FS) ([]File, error) {
	return nil, fmt.Errorf("%s: %w", s.id, ErrNotImplemented)
}

func (s stub) PlanMCPs(string, []model.MCP, fs.FS) ([]File, error) {
	return nil, fmt.Errorf("%s: %w", s.id, ErrNotImplemented)
}
