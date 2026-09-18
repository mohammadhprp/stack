# stack

`stack` is a Go CLI and TUI that installs two things into a project for
multiple coding-agent harnesses: **skills** and **MCP servers**. It is the
multi-harness successor to the older OpenCode-only `system-prompt` CLI.

It supports several harnesses at once (opencode, claude, codex, cursor, gemini,
windsurf, amp) so one selection can be wired into every agent the developer uses.

## Stack

- Language: Go (1.25+)
- CLI: `spf13/cobra`
- TUI: `charmbracelet/bubbletea` (+ `bubbles`, `lipgloss`)
- Content: embedded in the binary with `go:embed` (`framework/skills`, `framework/mcps`)
- Distribution: GitHub Releases + `install.sh` (no npm)

## Commands

- Build: `go build ./...`
- Test: `go test ./...`
- Vet: `go vet ./...`
- Format: `gofmt -l .`
- Run: `go run . list`

## Checks

- `go build ./...`, `go vet ./...`, `gofmt -l .`, `go test ./...` must all pass.
- Install behaviour is verified against a temporary directory, never the repo.
- Harness config formats are verified against each harness's current docs
  before the adapter is written; formats change, so do not guess.

## Conventions

- Package layout: `main.go`, `cmd/` (cobra), `internal/models`, `internal/config`,
  `internal/harness`, `internal/install`, `internal/tui`.
- Tests live under `tests/`, mirroring the source tree, as external `_test`
  packages; never place `_test.go` files in `internal/` or `cmd/`.
- One file per harness adapter in `internal/harness/`; register with `init()`.
- Comments: write **none by default**. Comment only a non-obvious *why* — a
  constraint, gotcha, or external-format detail the code cannot express (for
  example the `go:embed` `..` limitation). Never restate the code, document
  obvious fields or methods, or add section-divider/narration comments. Do not
  add doc comments to exported identifiers unless the comment carries real
  information. If a comment would survive only by being obvious, delete it.
- Never overwrite a target file the user changed without `--force`.
- Keep the diff minimal; no speculative abstractions.
- Read `CONTEXT.md` before changing anything.

## Acceptance bar

- A change is done when `go test ./...` passes **and** the behaviour was
  exercised with a real command whose output is captured (an install into a
  temp dir, a `list`, a `doctor`, or a TUI render).
- Installing twice must be safe and idempotent.
- Source existing is not enough: show the command and its output.

## Team Mate workers

- Work only inside this project; do not commit, merge, push, publish, deploy, or
  delete.
- Read this file and `CONTEXT.md` first (`load-project-context`).
- Stay inside the assignment; ask with `raise-blocker` instead of guessing.
- End with the `handoff-report` shape (`report-result`) written to
  `.teammate-report.md` in the project root.
