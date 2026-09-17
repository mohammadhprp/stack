# Excalidraw MCP Troubleshooting

## Diagram not rendering

**Problem**

The `draw` tool runs but no diagram appears in the chat.

**Cause**

The AI client may not support MCP Apps (interactive HTML interfaces). Standard text-only MCP clients cannot render interactive diagrams.

**Fix**

- Check if your client supports MCP Apps. Claude Desktop and Codex CLI support it; some other clients may not.
- Try a simpler diagram description to rule out parsing issues.
- As a fallback, ask the agent to describe the diagram in text or ASCII art.

## Remote server unreachable

**Problem**

The client reports that `https://mcp.excalidraw.com` is unreachable.

**Cause**

Network connectivity issue or the server is temporarily down.

**Fix**

- Run `curl https://mcp.excalidraw.com` to test connectivity.
- Check for proxy or firewall rules blocking the connection.
- Fall back to the local build (see [install.md](./install.md)).

## Local build fails

**Problem**

`pnpm install` or `pnpm run build` fails.

**Cause**

Missing dependencies or Node.js version mismatch.

**Fix**

- Ensure pnpm is installed: `npm install -g pnpm`.
- Ensure Node.js LTS is installed (v18+).
- Run `pnpm install --frozen-lockfile` for a clean install.
- Check the `dist/index.js` file exists after build — this is the server entry point.

## Diagram content is wrong

**Problem**

The diagram doesn't match what you asked for.

**Cause**

The description was ambiguous or too vague for the rendering engine.

**Fix**

- Be more specific: include component names, positions, relationships, and directional flow.
- Break complex diagrams into multiple smaller requests.
- Use the interactive viewer to adjust elements after rendering.
- Example: instead of "draw a system diagram", say "draw a system with a React frontend, a Node.js API server, and a PostgreSQL database connected with arrows from frontend to API and API to database".
