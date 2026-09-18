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
