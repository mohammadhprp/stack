# Figma MCP Go Installation

## Requirements

- Node.js and npm available on `PATH`.
- Figma desktop app.
- An MCP-compatible AI client.

## Installation

### 1. Install the Figma plugin

1. Download the latest `plugin.zip` from the [releases page](https://github.com/vkhanhqui/figma-mcp-go/releases).
2. In Figma Desktop: **Plugins → Development → Import plugin from manifest**.
3. Select `manifest.json` from `plugin.zip`.
4. Run the plugin inside any Figma file.

### 2. Configure your AI client

Add the MCP server to your client configuration.

## Configuration

| Field | Value |
| --- | --- |
| Server name | `figma-mcp-go` |
| Command | `npx` |
| Args | `-y`, `@vkhanhqui/figma-mcp-go` |

### Config file examples

Use the examples in [`configs/`](./configs/) as client-specific starting points:

| Client | Example | Typical location |
| --- | --- | --- |
| OpenCode | [`configs/opencode.json`](./configs/opencode.json) | OpenCode MCP configuration. |

### OpenCode

Copy [`configs/opencode.json`](./configs/opencode.json) into your OpenCode project root or user config directory and restart the OpenCode agent. The server will start automatically when the agent needs it.

## Verification

Run the MCP server command directly:

```bash
npx -y @vkhanhqui/figma-mcp-go
```

Then verify from your AI client:

- The `figma-mcp-go` MCP server is listed.
- The server starts without errors.
- Tools such as `get_document`, `get_selection`, or `create_frame` are visible to the agent.
- The Figma plugin shows a connected status.

## Updating

```bash
npx -y @vkhanhqui/figma-mcp-go@latest
```

Update the Figma plugin by re-importing from the latest `plugin.zip`.

## Uninstalling

1. Remove the `figma-mcp-go` MCP entry from AI client configuration files.
2. Remove the Figma plugin: **Plugins → Development → Manage plugins → remove figma-mcp-go**.

## Common Issues

- **Server starts but no tools are available**: confirm the Figma plugin is running in an open Figma file.
- **Client cannot find the server**: confirm the config is in the location used by that client and the server name is `figma-mcp-go`.
- **Plugin not listed**: use **Import plugin from manifest** rather than searching the plugin store.