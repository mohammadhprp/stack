# CONTEXT.md

`stack` installs **skills** and **MCP servers** into a project for one or more
coding-agent harnesses. It is a Go rewrite and multi-harness successor to the
OpenCode-only `system-prompt` CLI. Only skills and MCPs are supported — agents,
commands, plugins, styles, modes, standards, templates, and memory are
intentionally out of scope.

## Supported harnesses (v1)

| Harness  | Skills | MCP config                                            | Skills path            |
|----------|--------|-------------------------------------------------------|------------------------|
| opencode | yes    | `opencode.json` → `mcp`                               | `.opencode/skills/`    |
| claude   | yes    | `.mcp.json` → `mcpServers`                            | `.claude/skills/`      |
| codex    | yes    | `.codex/config.toml` → `[mcp_servers.<name>]`         | `.agents/skills/`      |
| cursor   | yes    | `.cursor/mcp.json` → `mcpServers`                     | `.agents/skills/`      |
| gemini   | yes    | `.gemini/settings.json` → `mcpServers`                | `.gemini/skills/`      |
| windsurf | yes    | — (user-level only)                                   | `.windsurf/skills/`    |
| amp      | yes    | `.amp/settings.json` → `amp.mcpServers`               | `.agents/skills/`      |

Every current harness supports skills. When a future adapter reports
`SupportsSkills` false, skills are skipped for it with a warning rather than
translated. `amp`, `codex`, and `cursor` all read `.agents/skills/`, so selecting
them together installs a single copy of each skill.

Windsurf's MCP config is user-level only
(`~/.codeium/windsurf/mcp_config.json`); it has no project-level file, so
`stack` plans no Windsurf MCP config and warns instead of writing one. Windsurf
skills are project-level (`.windsurf/skills/`).

**Always verify a harness's current on-disk format against its official docs
before implementing or changing an adapter.** Formats drift.

## Project structure

- `main.go` — entry point; builds the cobra root command.
- `cmd/` — cobra commands: `root`, `install`, `list`, `doctor`, `version`.
- `internal/models/` — domain types, one file per topic: `skill.go` (`Skill`),
  `mcp.go` (`MCPSpec`, `MCP`), `harness.go` (`Harness`).
- `internal/config/` — harness-agnostic config-file I/O: `json.go`
  (`MergeJSONSection`), `toml.go` (`MergeTOMLTable`). It knows paths and
  sections, never harnesses.
- `internal/harness/` — the adapter interface, the registry, and one file per
  harness (`opencode.go`, `claude.go`, `codex.go`, `cursor.go`); each adapter
  renders its entries and delegates the file merge to `internal/config`.
- `internal/install/` — catalog loading (`catalog.go`), plan → diff → apply
  engine, lockfile, content hashing.
- `internal/tui/` — bubbletea wizard used by bare `stack`.
- `tests/` — **every Go test lives here**, mirroring the source tree
  (`tests/internal/harness/opencode_test.go`, `tests/cmd/install_test.go`).
- `framework/skills/<id>/SKILL.md` — skill content (frontmatter + body).
- `framework/mcps/<id>/spec.json` — MCP content. **Minimal by design:** one
  canonical JSON file per MCP, no README/install/troubleshooting markdown and no
  per-harness config directory. If richer docs are ever needed, they belong
  outside the shipped content.
- `install.sh` — curl-able installer that fetches the latest GitHub Release.
- `.goreleaser.yaml` — cross-platform release build.
- `.github/workflows/` — CI and release workflows.

## Embedded content

`framework/` stays at the repository root as the editable content root.
`go:embed` patterns cannot use `..`, so the embed lives in the root package
(`main.go`: `//go:embed all:framework`) and the resulting `fs.FS`, sub-rooted at
`framework/`, is passed down to `internal/install`. Mirrored tests load the same
tree with `os.DirFS("../../../framework")`.

## Harness adapter contract

Every adapter implements one interface: identity (`ID`, `Name`), capability
(`SupportsSkills`), and planning/writing of skills and MCP entries for a target
directory. Adapters must merge into existing config files without dropping
unrelated keys, and must be idempotent.

## MCP canonical spec

Each MCP is exactly one file: `framework/mcps/<dir>/spec.json`, where `<dir>` is
the catalog item id. `spec.id` is the MCP **server key** written into harness
configs (it can differ from the directory name, e.g. `playwright-mcp/` →
`playwright`).

```json
{
  "id": "playwright",
  "name": "Playwright MCP",
  "description": "Cross-browser automation with accessibility snapshots",
  "type": "local",
  "command": ["npx", "@playwright/mcp@latest"],
  "env": {},
  "url": null,
  "headers": {}
}
```

Adapters render this into their harness's format. `type` is `local` (uses
`command`/`env`) or `remote` (uses `url`/`headers`). `name` and `description`
are for `stack list` and the TUI. Keep this file the single source of MCP data.

## Testing

All Go tests live under `tests/`, mirroring the source tree:

- `internal/harness/opencode.go` → `tests/internal/harness/opencode_test.go`
- `cmd/install.go` → `tests/cmd/install_test.go`

Rules:

- No `_test.go` files inside `internal/` or `cmd/`; tests never sit beside the
  code they exercise.
- Test files are **external test packages** (for example `package harness_test`)
  and import the real package by its module path. They can only use exported
  identifiers; design the packages with exported seams where tests need them.
- Fixtures and `testdata/` are mirrored under `tests/` too.
- `go test ./...` must discover and run the whole mirrored suite.
- Coverage of the real packages needs `-coverpkg` (the mirrored test packages
  measure themselves otherwise).

## Lockfile

Installs write `.stack-lock.json` in the target project recording, per item, the
source, target paths, and a SHA-256 of every managed file. It powers idempotent
re-runs, change detection (`stack doctor`), and safe removal. Ported from the
`system-prompt` lock schema (version 1) but scoped to skills and mcps.

The TUI prefills its selections from that lockfile: a re-run starts with the
previously installed harnesses, skills, and MCPs already selected. When there is
no lock it falls back to detecting managed files already on disk (planning each
candidate through its adapter and checking the planned files match), so an
existing install is review-and-adjust. A fresh target starts empty.

## Removal and pruning

The TUI prunes by default: its prefilled selection is the desired end state, so
deselecting a harness, skill, or MCP removes it on confirm, after the confirm
screen lists the removals. The CLI prunes only with `--prune`; without it the CLI
deletes nothing.

Removal only ever touches lock-managed paths inside the target: a file is
deleted only while its hash still matches the lock (a user-modified file is kept
and warned about unless `--force`); shared `.agents/skills/` paths are kept while
a selected harness still needs them; MCP entries are dropped through the optional
`harness.MCPRemover` interface, which preserves unrelated config keys and other
servers and signals a config that becomes empty for deletion; empty directories
are pruned up to the target root; and removed items are dropped from the lock so
a later install recreates them cleanly.

## Build & Test

- `go build ./...`
- `go test ./...`
- `go vet ./...`
- `gofmt -l .` (must print nothing)

## Distribution

No npm. Releases are `stack_<os>_<arch>.tar.gz` plus `checksums.txt` attached to
a GitHub Release. `install.sh` detects OS/arch, downloads the latest release,
verifies the checksum, and installs the binary to `~/.local/bin` by default.

Release process: tag `vX.Y.Z`, push the tag, the release workflow runs
goreleaser, then the tag's assets become installable via
`https://github.com/mohammadhprp/stack/releases/latest/download/install.sh`.

## Engineering conduct

1. Safety before speed — never overwrite a user-modified file without `--force`.
2. Verify formats from docs before coding an adapter.
3. Prefer simplicity — minimal dependencies, no speculative abstraction.
4. Verify before concluding — run the real command and capture output.
5. Small, reversible changes.
6. **Comments are the exception, not the norm.** Write no comments by default.
   Keep one only when it records a non-obvious constraint or *why* the code is
   shaped that way; never restate code or document obvious fields, methods, or
   exported identifiers. No narration or section-divider comments.
