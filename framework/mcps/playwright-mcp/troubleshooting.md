# Playwright MCP Troubleshooting

## Server won't start

**Problem**

The AI client reports that `playwright` failed to start.

**Cause**

The client cannot run `npx`, Node.js is not installed, or browser binaries are missing.

**Solution**

- Run `npx @playwright/mcp@latest` manually to verify the server starts.
- Check Node.js is installed: `node --version` (requires 18+).
- On first run, Playwright downloads browser binaries automatically. Ensure a working internet connection.
- Check network access if the npm package has not been cached locally.

## Browser does not launch

**Problem**

The server starts but no browser window appears.

**Cause**

Headless mode is enabled, or the browser binary is not found.

**Solution**

- If using `--headless`, the browser runs without a UI — this is expected.
- Use `--browser` to specify an installed browser engine (chromium, firefox, webkit).
- Set `PLAYWRIGHT_MCP_EXECUTABLE_PATH` or pass `--executable-path` to point to a custom browser binary.

## Sandbox errors on Linux

**Problem**

The browser fails to start with sandbox-related errors on Linux.

**Cause**

Chrome/Chromium sandbox requires specific kernel settings not available in Docker or containerized environments.

**Solution**

- Pass `--no-sandbox` in the server args.
- Or set the `PLAYWRIGHT_MCP_NO_SANDBOX` environment variable.
- Note: disabling sandbox reduces security. Use only in trusted environments.

## "Cannot find module" errors

**Problem**

The server fails with module resolution errors.

**Cause**

The npm package is corrupted or incompletely cached.

**Solution**

- Clear npm cache: `npm cache clean --force`.
- Retry with `npx @playwright/mcp@latest`.

## Configuration errors

**Problem**

The MCP server is not listed, or the client ignores the config file.

**Cause**

The config file is in the wrong location, the JSON is invalid, or the client expects a different key name.

**Solution**

- Validate the JSON with `python -m json.tool <file>` or another JSON parser.
- Confirm the client-specific config location in the client's documentation.
- Keep the server name as `playwright`.

## Concurrent session conflicts

**Problem**

A second MCP client fails to start or shows stale data from another session.

**Cause**

The persistent profile (at `~/Library/Caches/ms-playwright/mcp-...` on macOS) can only be used by one browser instance at a time.

**Solution**

- Start additional clients with `--isolated` to keep profiles in memory only.
- Or set a distinct `--user-data-dir` for each concurrent client.
- Or change the workspace root so a different `workspace-hash` is derived.

## Port already in use

**Problem**

The standalone MCP server (with `--port`) fails to bind.

**Solution**

- Specify a different port: `--port 8932`.
- Or kill the process using the port: `lsof -ti:8931 | xargs kill`.

## Docker issues

**Problem**

The Docker container fails to run or the browser does not work.

**Cause**

Docker image only supports headless Chromium. Missing kernel capabilities for browser execution.

**Solution**

- Ensure `--headless` and `--no-sandbox` are passed as args.
- Use `--pull=always` to get the latest image.
- Verify the container has the `--init` flag for proper signal handling.
- For full control, build the Docker image locally with `docker build -t mcr.microsoft.com/playwright/mcp .`.

## Performance issues

**Problem**

The MCP server is slow or times out frequently.

**Solution**

- Adjust action timeout: `--timeout-action 10000` (default 5000ms).
- Adjust navigation timeout: `--timeout-navigation 120000` (default 60000ms).
- Reduce snapshot depth with `--snapshot-mode=none` if you don't need the full accessibility tree.
- Limit console output level: `--console-level=error`.

## Extension mode not working

**Problem**

`--extension` flag does not connect to the running browser.

**Cause**

The Playwright browser extension is not installed, or Chrome is not running with the extension enabled.

**Solution**

- Install the Playwright Extension from [microsoft/playwright/packages/extension](https://github.com/microsoft/playwright/tree/main/packages/extension).
- Ensure Chrome/Edge is running with the extension enabled before starting the MCP server.
- This mode requires Chrome or Edge only; Firefox and WebKit are not supported.
