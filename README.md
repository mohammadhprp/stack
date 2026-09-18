# Stack

One CLI to install **skills** and **MCP servers** into a project for multiple
coding-agent harnesses

## Install

```sh
curl -fsSL https://raw.githubusercontent.com/mohammadhprp/stack/main/install.sh | sh
```

The script detects your OS and architecture, downloads the matching binary from
the latest [release](https://github.com/mohammadhprp/stack/releases), verifies
its checksum, and installs `stack` to `~/.local/bin`.

## Usage

```sh
stack            # interactive TUI: pick harnesses, skills, and MCPs
stack list       # list available skills and MCPs
stack install --harness opencode,claude --skills commit,why --mcp playwright-mcp
stack install --all --harness claude
stack doctor     # check what is installed
```

## Supported harnesses

Every harness supports **skills**; most support **MCP** too.

| Harness  | Skills | MCP config |
|----------|--------|------------|
| opencode | yes    | `opencode.json` |
| claude   | yes    | `.mcp.json` |
| codex    | yes    | `.codex/config.toml` |
| cursor   | yes    | `.cursor/mcp.json` |
| gemini   | yes    | `.gemini/settings.json` |
| amp      | yes    | `.amp/settings.json` |
| windsurf | yes    | — (MCP config is user-level only; `stack` won't write outside the project) |

Skills go to the location each harness reads: `.opencode/skills/`,
`.claude/skills/`, `.gemini/skills/`, `.windsurf/skills/`, and the shared
`.agents/skills/` used by codex, cursor, and amp.

## License

[MIT](LICENSE)
