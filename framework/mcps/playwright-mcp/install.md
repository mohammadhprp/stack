# Playwright MCP Installation

## Requirements

- Node.js 18 or newer available on `PATH`.
- An MCP-compatible AI client.
- Browser binaries are downloaded automatically on first run (Chromium by default).

## Installation

The server is used directly with `npx` — no build step required.

```bash
npx @playwright/mcp@latest
```

## Configuration

| Field | Value |
| --- | --- |
| Server name | `playwright` |
| Command | `npx` |
| Args | `@playwright/mcp@latest` |

### Config file examples

Use the examples in [`configs/`](./configs/) as client-specific starting points:

| Client | Example | Typical location |
| --- | --- | --- |
| OpenCode | [`configs/opencode.json`](./configs/opencode.json) | OpenCode MCP configuration. |

### OpenCode

Copy [`configs/opencode.json`](./configs/opencode.json) into your OpenCode project root or user config directory and restart the OpenCode agent.

### Common flags

Pass these via the `args` array:

| Flag | Description |
| --- | --- |
| `--headless` | Run browser in headless (no UI) mode |
| `--browser=firefox` | Use Firefox instead of Chromium (also: `webkit`, `chrome`, `msedge`) |
| `--caps=vision` | Enable vision capability for coordinate-based interactions |
| `--caps=pdf` | Enable PDF generation capability |
| `--caps=devtools` | Enable DevTools capability |
| `--caps=network` | Enable network mocking/routing capability |
| `--caps=storage` | Enable cookie/localStorage management capability |
| `--isolated` | Keep browser profile in memory, do not save to disk |
| `--port=8931` | Listen on a port for SSE transport (standalone server) |
| `--user-data-dir=/path` | Custom user data directory for persistent profile |
| `--device="iPhone 15"` | Emulate a mobile device |
| `--no-sandbox` | Disable sandbox (needed in some Docker/CI environments) |
| `--viewport-size=1280x720` | Set browser viewport size |

Headless Chromium config example:

```json
{
  "mcp": {
    "playwright": {
      "type": "local",
      "enabled": true,
      "command": ["npx", "@playwright/mcp@latest", "--headless"]
    }
  }
}
```

### Docker

Playwright MCP supports Docker for headless Chromium:

```json
{
  "mcp": {
    "playwright": {
      "type": "local",
      "enabled": true,
      "command": ["docker", "run", "-i", "--rm", "--init", "--pull=always", "mcr.microsoft.com/playwright/mcp"]
    }
  }
}
```

Or as a long-lived service:

```bash
docker run -d -i --rm --init --pull=always \
  --entrypoint node \
  --name playwright \
  -p 8931:8931 \
  mcr.microsoft.com/playwright/mcp \
  /app/cli.js --headless --browser chromium --no-sandbox --port 8931 --host 0.0.0.0
```

## Verification

Run the MCP server command directly:

```bash
npx @playwright/mcp@latest
```

Then verify from your AI client:

- The `playwright` MCP server is listed.
- The server starts without errors.
- Tools are visible to the agent.

Recommended first prompt:

```
Navigate to https://playwright.dev and take a snapshot of the page
```

## Updating

The `@latest` tag ensures your client always uses the newest version. To force an update:

```bash
npx @playwright/mcp@latest
```

## Uninstalling

Remove the `playwright` MCP entry from AI client configuration files.

## Common Issues

- **Browser does not start**: ensure browser binaries are available or use `--browser` to specify an installed browser.
- **Sandbox errors on Linux**: add `--no-sandbox` to the args when running in Docker or CI.
- **Server not found**: confirm the config file is in the correct location and the JSON is valid.
- **Concurrent session conflicts**: persistent profiles cannot be shared. Use `--isolated` or a distinct `--user-data-dir` for parallel clients.
- **Port already in use**: specify a different port with `--port`.
