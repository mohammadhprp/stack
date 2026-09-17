package install

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	LockFile    = ".stack-lock.json"
	LockVersion = 1

	KindSkill = "skill"
	KindMCP   = "mcp"
)

var ErrNoLock = errors.New("no lockfile")

type Lock struct {
	Version int        `json:"version"`
	Items   []LockItem `json:"items"`
}

type LockItem struct {
	Harness string            `json:"harness"`
	Kind    string            `json:"kind"`
	ID      string            `json:"id"`
	Source  string            `json:"source"`
	Files   map[string]string `json:"files"`
}

func Hash(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// LoadLock returns ErrNoLock when the file does not exist.
func LoadLock(target string) (*Lock, error) {
	data, err := os.ReadFile(filepath.Join(target, LockFile))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, ErrNoLock
	}
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", LockFile, err)
	}
	var lock Lock
	if err := json.Unmarshal(data, &lock); err != nil {
		return nil, fmt.Errorf("parse %s: %w", LockFile, err)
	}
	if lock.Version != LockVersion {
		return nil, fmt.Errorf("%s: unsupported version %d (want %d)", LockFile, lock.Version, LockVersion)
	}
	return &lock, nil
}

func SaveLock(target string, lock *Lock) error {
	if lock.Version == 0 {
		lock.Version = LockVersion
	}
	data, err := json.MarshalIndent(lock, "", "  ")
	if err != nil {
		return fmt.Errorf("encode %s: %w", LockFile, err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(filepath.Join(target, LockFile), data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", LockFile, err)
	}
	return nil
}

func (l *Lock) Validate(target string) error {
	if l.Version != LockVersion {
		return fmt.Errorf("%s: unsupported version %d (want %d)", LockFile, l.Version, LockVersion)
	}
	for _, item := range l.Items {
		for rel, want := range item.Files {
			data, err := os.ReadFile(filepath.Join(target, filepath.FromSlash(rel)))
			if err != nil {
				return fmt.Errorf("%s/%s %s: %w", item.Harness, item.ID, rel, err)
			}
			if got := Hash(data); got != want {
				return fmt.Errorf("%s/%s %s: modified (locked %s, found %s)", item.Harness, item.ID, rel, want, got)
			}
		}
	}
	return nil
}

// index fails when the same path is recorded with two different hashes.
func (l *Lock) index() (map[string]string, error) {
	hashes := make(map[string]string)
	for _, item := range l.Items {
		for rel, hash := range item.Files {
			if existing, ok := hashes[rel]; ok && existing != hash {
				return nil, fmt.Errorf("%s: inconsistent hashes for %s", LockFile, rel)
			}
			hashes[rel] = hash
		}
	}
	return hashes, nil
}

// clone deep-copies so updates never mutate the loaded lock.
func (l *Lock) clone() *Lock {
	out := &Lock{Version: l.Version}
	for _, item := range l.Items {
		files := make(map[string]string, len(item.Files))
		for k, v := range item.Files {
			files[k] = v
		}
		item.Files = files
		out.Items = append(out.Items, item)
	}
	return out
}

// upsert keys on harness+kind+id.
func (l *Lock) upsert(item LockItem) {
	for i, existing := range l.Items {
		if existing.Harness == item.Harness && existing.Kind == item.Kind && existing.ID == item.ID {
			l.Items[i] = item
			return
		}
	}
	l.Items = append(l.Items, item)
}

func (l *Lock) ManagedFiles() []string {
	seen := map[string]bool{}
	var files []string
	for _, item := range l.Items {
		for rel := range item.Files {
			if !seen[rel] {
				seen[rel] = true
				files = append(files, rel)
			}
		}
	}
	return files
}
