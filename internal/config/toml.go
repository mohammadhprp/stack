package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/BurntSushi/toml"
)

func MergeTOMLTable(path, table string, entries map[string]any) ([]byte, error) {
	if len(entries) == 0 {
		return nil, nil
	}

	root := map[string]any{}
	switch data, err := os.ReadFile(path); {
	case err == nil:
		if len(bytes.TrimSpace(data)) > 0 {
			if err := toml.Unmarshal(data, &root); err != nil {
				return nil, fmt.Errorf("parse %s: %w", path, err)
			}
		}
	case errors.Is(err, fs.ErrNotExist):
	default:
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	existing := map[string]any{}
	if raw, ok := root[table]; ok {
		got, ok := raw.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("%s in %s is not a table", table, path)
		}
		existing = got
	}
	for name, entry := range entries {
		existing[name] = entry
	}
	root[table] = existing

	out, err := toml.Marshal(root)
	if err != nil {
		return nil, fmt.Errorf("encode %s: %w", path, err)
	}
	if len(out) == 0 || out[len(out)-1] != '\n' {
		out = append(out, '\n')
	}
	return out, nil
}

// ReconcileTOMLTable sets the named table to exactly keep, removes the drop
// keys, and preserves every other key. It returns remove=true when the file is
// left with no content and should be deleted.
func ReconcileTOMLTable(path, table string, keep map[string]any, drop []string) ([]byte, bool, error) {
	root := map[string]any{}
	switch data, err := os.ReadFile(path); {
	case err == nil:
		if len(bytes.TrimSpace(data)) > 0 {
			if err := toml.Unmarshal(data, &root); err != nil {
				return nil, false, fmt.Errorf("parse %s: %w", path, err)
			}
		}
	case errors.Is(err, fs.ErrNotExist):
	default:
		return nil, false, fmt.Errorf("read %s: %w", path, err)
	}

	existing := map[string]any{}
	if raw, ok := root[table]; ok {
		got, ok := raw.(map[string]any)
		if !ok {
			return nil, false, fmt.Errorf("%s in %s is not a table", table, path)
		}
		existing = got
	}
	for _, name := range drop {
		delete(existing, name)
	}
	for name, entry := range keep {
		existing[name] = entry
	}

	if len(existing) == 0 {
		delete(root, table)
	} else {
		root[table] = existing
	}
	if len(root) == 0 {
		return nil, true, nil
	}

	out, err := toml.Marshal(root)
	if err != nil {
		return nil, false, fmt.Errorf("encode %s: %w", path, err)
	}
	if len(out) == 0 || out[len(out)-1] != '\n' {
		out = append(out, '\n')
	}
	return out, false, nil
}
