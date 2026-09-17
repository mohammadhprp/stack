# Supabase MCP Installation

## Requirements

- A Supabase account with a project.
- An MCP-compatible AI client that supports remote HTTP MCP servers (Streamable HTTP).
- No local infrastructure — the server is hosted by Supabase at `https://mcp.supabase.com/mcp`.

## Configuration

Supabase MCP is a remote Streamable HTTP server — no local command or install needed. Authentication uses OAuth, completed interactively on first use.

| Field | Value |
| --- | --- |
| Server name | `supabase` |
| Transport type | Streamable HTTP |
| URL | `https://mcp.supabase.com/mcp` |

### Config file examples

Use the examples in [`configs/`](./configs/) as client-specific starting points:

| Client | Example | Typical location |
| --- | --- | --- |
| OpenCode | [`configs/opencode.json`](./configs/opencode.json) | OpenCode MCP configuration. |

### OpenCode

Copy [`configs/opencode.json`](./configs/opencode.json) into your OpenCode project root or user config directory. Then authenticate:

```bash
opencode mcp auth supabase
```

This opens your browser to complete the OAuth flow.

### For clients that only support local stdio servers

Some MCP clients only support local stdio servers. Use the [mcp-remote](https://www.npmjs.com/package/mcp-remote) bridge to connect:

```json
{
  "mcpServers": {
    "supabase": {
      "command": "npx",
      "args": ["-y", "mcp-remote", "https://mcp.supabase.com/mcp"]
    }
  }
}
```

## Configuration options

Append URL query parameters to restrict server scope:

| Parameter | Description | Example |
| --- | --- | --- |
| `read_only=true` | Execute all queries as a read-only Postgres user | `?read_only=true` |
| `project_ref=<id>` | Scope to a specific project (disables account tools) | `?project_ref=abc123` |
| `features=<groups>` | Enable only specific tool groups (comma-separated) | `?features=database,docs` |

Parameters can be combined: `https://mcp.supabase.com/mcp?project_ref=abc123&read_only=true`

## Verification

After configuring and authenticating, ask your AI agent:

```
List all tables in the database using MCP tools.
```

## Updating

Supabase MCP is hosted by Supabase — no update steps needed. The latest version is always available at `https://mcp.supabase.com/mcp`.

## Uninstalling

Remove the `supabase` MCP entry from AI client configuration files. Optionally, revoke the OAuth connection in your Supabase organization settings.

## Common Issues

- **Authentication fails**: make sure you complete the OAuth flow when prompted. Run the auth command again to trigger a fresh flow.
- **Client does not support remote servers**: use the `mcp-remote` bridge as a stdio wrapper.
- **Permission errors**: check that your Supabase user has the correct permissions for the project or organization.
- **Tool not found**: verify which feature groups are enabled — some tools are disabled by default (Storage, Branching).