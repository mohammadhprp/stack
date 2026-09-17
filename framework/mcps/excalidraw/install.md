# Excalidraw MCP Installation

## Requirements

- An MCP-compatible AI client that supports MCP Apps (interactive HTML interfaces).
- No API key or authentication needed for the remote server.

## Configuration

Excalidraw MCP can be used remotely or run locally.

### Remote (recommended)

| Field | Value |
| --- | --- |
| Server name | `excalidraw` |
| Transport type | Streamable HTTP |
| URL | `https://mcp.excalidraw.com` |

### Config file examples

Use the examples in [`configs/`](./configs/) as client-specific starting points:

| Client | Example | Typical location |
| --- | --- | --- |
| OpenCode | [`configs/opencode.json`](./configs/opencode.json) | OpenCode MCP configuration. |

### OpenCode

Copy [`configs/opencode.json`](./configs/opencode.json) into your OpenCode project root or user config directory and restart the OpenCode agent.

### Local (build from source)

Clone and build the repository:

```bash
git clone https://github.com/excalidraw/excalidraw-mcp.git
cd excalidraw-mcp
pnpm install && pnpm run build
```

Add to your MCP config:

```json
{
  "mcpServers": {
    "excalidraw": {
      "command": "node",
      "args": ["/path/to/excalidraw-mcp/dist/index.js", "--stdio"]
    }
  }
}
```

### For clients that only support local stdio servers

Use the [mcp-remote](https://www.npmjs.com/package/mcp-remote) bridge:

```json
{
  "mcpServers": {
    "excalidraw": {
      "command": "npx",
      "args": ["-y", "mcp-remote", "https://mcp.excalidraw.com"]
    }
  }
}
```

## Verification

After configuring and restarting your AI client, try:

```
Draw an architecture diagram showing a user connecting to an API server which talks to a database using excalidraw
```

## Updating

For the remote server: no update steps needed — the latest version is always available at `https://mcp.excalidraw.com`.

For local builds: pull the latest changes and rebuild.

## Uninstalling

Remove the `excalidraw` MCP entry from AI client configuration files.

## Common Issues

- **Diagram not rendering**: ensure your AI client supports MCP Apps (interactive HTML interfaces). Some older clients may not support this format.
- **Remote server unreachable**: check that `https://mcp.excalidraw.com` is accessible from your network. Try the local build as a fallback.
- **Local build fails**: ensure pnpm is installed (`npm install -g pnpm`) and Node.js LTS is available.
