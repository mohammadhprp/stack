# Figma MCP Go Troubleshooting

## Server won't start

**Problem**

The AI client reports that `figma-mcp-go` failed to start.

**Cause**

The client cannot run `npx`, Node.js is not installed, or the npm package is unavailable.

**Solution**

- Run `npx -y @vkhanhqui/figma-mcp-go` manually to verify the server starts.
- Check Node.js is installed: `node --version && npm --version`.
- Check network access if the package has not been cached locally.
- Ensure the config file contains the correct command and args.

## Figma plugin not connecting

**Problem**

The MCP server starts but reports no Figma connection.

**Cause**

The Figma plugin is not running, not imported, or not open in a file.

**Solution**

- Confirm the plugin was imported via **Plugins → Development → Import plugin from manifest**.
- Run the plugin inside the Figma file you want to edit.
- Verify the plugin shows a connected status after running.

## No tools available

**Problem**

The AI client shows the server is running but no tools are listed.

**Cause**

The server started without a Figma connection, or the plugin bridge failed to initialise.

**Solution**

- Restart the Figma plugin.
- Restart the MCP server.
- Ensure the Figma file is open and the plugin is running in that file.

## Configuration errors

**Problem**

The MCP server is not listed, or the client ignores the config file.

**Cause**

The config file is in the wrong location, the JSON is invalid, or the client expects a different configuration wrapper.

**Solution**

- Validate the JSON with `python -m json.tool <file>` or another JSON parser.
- Confirm the client-specific config location in the client's documentation.
- Keep the server name as `figma-mcp-go`.
- Use the official command `npx` and args `-y`, `@vkhanhqui/figma-mcp-go`.

## Version mismatch

**Problem**

Some tools are missing or behave differently than documented.

**Cause**

The MCP server version and the Figma plugin version are out of sync.

**Solution**

- Update both the server and plugin to the latest version.
- Re-download `plugin.zip` from the latest release and re-import into Figma.
- Clear the npm cache: `npm cache clean --force` and retry.

## Export tools fail

**Problem**

`save_screenshots` or `export_frames_to_pdf` fail with file system errors.

**Cause**

The server process does not have write permission to the target directory.

**Solution**

- Provide an absolute path to a writable directory.
- Check filesystem permissions on the target directory.
- Use `get_screenshot` (base64 output) instead when disk writes are not required.