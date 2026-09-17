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
| codex    | no     | `.codex/config.toml` → `[mcp_servers.<name>]`         | —                      |
| cursor   | no     | `.cursor/mcp.json` → `mcpServers`                     | —                      |

When a harness has no native skill concept, skills are skipped for it (with a
warning) rather than translated. Codex/Cursor skill adapters are a later
milestone.

**Always verify a harness's current on-disk format against its official docs
before implementing or changing an adapter.** Formats drift.

## Project structure

- `main.go` — entry point; builds the cobra root command.
- `cmd/` — cobra commands: `root`, `install`, `list`, `doctor`, `version`.
- `internal/catalog/` — loads embedded `framework/` content into the model.
- `internal/model/` — `Skill`, `MCP`, `Harness` types.
- `internal/harness/` — the adapter interface, the registry, and one file per
  harness (`opencode.go`, `claude.go`, `codex.go`, `cursor.go`).
- `internal/install/` — plan → diff → apply engine, lockfile, content hashing.
- `internal/tui/` — bubbletea wizard used by bare `stack`.
- `framework/skills/<id>/SKILL.md` — skill content (frontmatter + body).
- `framework/mcps/<id>/spec.json` — MCP content. **Minimal by design:** one
  canonical JSON file per MCP, no README/install/troubleshooting markdown and no
  per-harness config directory. If richer docs are ever needed, they belong
  outside the shipped content.
- `install.sh` — curl-able installer that fetches the latest GitHub Release.
- `.goreleaser.yaml` — cross-platform release build.
- `.github/workflows/` — CI and release workflows.

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

## Lockfile

Installs write `.stack-lock.json` in the target project recording, per item, the
source, target paths, and a SHA-256 of every managed file. It powers idempotent
re-runs, change detection (`stack doctor`), and safe removal. Ported from the
`system-prompt` lock schema (version 1) but scoped to skills and mcps.

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
