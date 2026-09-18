package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func MergeJSONSection(path, section string, entries map[string]any) ([]byte, error) {
	if len(entries) == 0 {
		return nil, nil
	}

	root := map[string]json.RawMessage{}
	switch data, err := os.ReadFile(path); {
	case err == nil:
		if len(bytes.TrimSpace(data)) > 0 {
			if err := json.Unmarshal(data, &root); err != nil {
				return nil, fmt.Errorf("parse %s: %w", path, err)
			}
		}
	case errors.Is(err, fs.ErrNotExist):
	default:
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	existing := map[string]json.RawMessage{}
	if raw, ok := root[section]; ok && len(bytes.TrimSpace(raw)) > 0 {
		if err := json.Unmarshal(raw, &existing); err != nil {
			return nil, fmt.Errorf("parse %s section in %s: %w", section, path, err)
		}
	}
	for name, entry := range entries {
		raw, err := json.Marshal(entry)
		if err != nil {
			return nil, fmt.Errorf("encode entry %q: %w", name, err)
		}
		existing[name] = raw
	}

	rawSection, err := json.Marshal(existing)
	if err != nil {
		return nil, fmt.Errorf("encode %s section: %w", section, err)
	}
	root[section] = rawSection

	out, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encode %s: %w", path, err)
	}
	out = append(out, '\n')
	return out, nil
}
