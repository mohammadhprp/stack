// Package harness registers one adapter per harness with init(); there is no
// central switch.
package harness

import (
	"errors"
	"fmt"
	"io/fs"
	"sort"

	"github.com/mohammadhprp/stack/internal/models"
)

var ErrNotImplemented = errors.New("adapter not implemented")

type File struct {
	// Path uses forward slashes and is relative to the target directory.
	Path    string
	Content []byte
	// Merge marks a config file whose Content was produced by merging into the
	// existing file rather than replacing it. Merge files may be created over
	// an existing untracked file; non-merge files may not.
	Merge  bool
	Source string
}

// Adapter only plans files; the install engine applies them.
type Adapter interface {
	ID() string
	Name() string
	SupportsSkills() bool
	PlanSkills(target string, skills []models.Skill, src fs.FS) ([]File, error)
	// PlanMCPs must preserve unrelated keys in the existing config.
	PlanMCPs(target string, mcps []models.MCP, src fs.FS) ([]File, error)
}

// MCPConfigReporter is implemented by adapters whose MCP config cannot live
// inside the target (for example a harness with only a user-level config file).
// The install engine records the note as a warning and plans no MCP files.
type MCPConfigReporter interface {
	MCPConfigNote() string
}

var registry = map[string]Adapter{}

func Register(a Adapter) {
	if a == nil {
		panic("harness: Register(nil)")
	}
	if a.ID() == "" {
		panic("harness: adapter with empty ID")
	}
	if _, dup := registry[a.ID()]; dup {
		panic(fmt.Sprintf("harness: duplicate adapter id %q", a.ID()))
	}
	registry[a.ID()] = a
}

func Get(id string) (Adapter, bool) {
	a, ok := registry[id]
	return a, ok
}

func All() []Adapter {
	adapters := make([]Adapter, 0, len(registry))
	for _, a := range registry {
		adapters = append(adapters, a)
	}
	sort.Slice(adapters, func(i, j int) bool { return adapters[i].ID() < adapters[j].ID() })
	return adapters
}
