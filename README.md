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

## License

[MIT](LICENSE)
