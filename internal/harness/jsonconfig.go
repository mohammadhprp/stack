package harness

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// mergeJSONSection merges rendered entries into the named section of a JSON
// config file, preserving every unrelated key. The file is marked Merge so the
// install engine may create it over an existing untracked config.
func mergeJSONSection(harnessID, target, configFile, section string, entries map[string]any) ([]File, error) {
	if len(entries) == 0 {
		return nil, nil
	}

	root := map[string]json.RawMessage{}
	switch data, err := os.ReadFile(filepath.Join(target, filepath.FromSlash(configFile))); {
	case err == nil:
		if len(bytes.TrimSpace(data)) > 0 {
			if err := json.Unmarshal(data, &root); err != nil {
				return nil, fmt.Errorf("%s: parse %s: %w", harnessID, configFile, err)
			}
		}
	case errors.Is(err, fs.ErrNotExist):
	default:
		return nil, fmt.Errorf("%s: read %s: %w", harnessID, configFile, err)
	}

	existing := map[string]json.RawMessage{}
	if raw, ok := root[section]; ok && len(bytes.TrimSpace(raw)) > 0 {
		if err := json.Unmarshal(raw, &existing); err != nil {
			return nil, fmt.Errorf("%s: parse %s section in %s: %w", harnessID, section, configFile, err)
		}
	}
	for name, entry := range entries {
		raw, err := json.Marshal(entry)
		if err != nil {
			return nil, fmt.Errorf("%s: encode mcp %q: %w", harnessID, name, err)
		}
		existing[name] = raw
	}

	rawSection, err := json.Marshal(existing)
	if err != nil {
		return nil, fmt.Errorf("%s: encode %s section: %w", harnessID, section, err)
	}
	root[section] = rawSection

	out, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("%s: encode %s: %w", harnessID, configFile, err)
	}
	out = append(out, '\n')

	return []File{{Path: configFile, Content: out, Merge: true, Source: "mcps"}}, nil
}
