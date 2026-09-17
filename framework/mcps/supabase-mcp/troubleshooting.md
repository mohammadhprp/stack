# Supabase MCP Troubleshooting

## Connection fails

**Problem**

The AI client reports that `supabase` cannot connect.

**Cause**

The client does not support remote HTTP MCP servers, or the URL is misconfigured.

**Solution**

- Verify the URL is `https://mcp.supabase.com/mcp`.
- If your client only supports local stdio servers, use the `mcp-remote` bridge: `npx -y mcp-remote https://mcp.supabase.com/mcp`.
- Check the client's documentation for how to add a remote MCP server.

## Authentication fails

**Problem**

The connection succeeds but tools return authentication errors.

**Cause**

The OAuth flow was not completed or the token has expired.

**Solution**

- Run the auth command again: `opencode mcp auth supabase`.
- Make sure you select the correct Supabase organization during OAuth.
- Try disconnecting and reconnecting. Look for a "Clear authentication" or "Disconnect" option in the client's MCP settings.

## Permission errors

**Problem**

Tools return 403 or project not found errors.

**Cause**

The authenticated user does not have access to the requested project or organization, or the project ref is incorrect.

**Solution**

- Verify your Supabase user has the correct permissions for the project.
- Check that the `project_ref` parameter matches an actual project ID.
- Use account-level tools (no `project_ref`) to list available projects and organizations first.

## Read-only mode blocks writes

**Problem**

INSERT, UPDATE, DELETE queries fail silently or return errors.

**Cause**

The server is configured with `read_only=true`, which executes all queries as a read-only Postgres user.

**Solution**

- Remove `read_only=true` from the URL if you need write access.
- Use branching to make changes in an isolated environment.

## Tool not available

**Problem**

A specific tool is missing when the AI agent lists available tools.

**Cause**

The tool belongs to a feature group that is not enabled.

**Solution**

- Verify which feature groups are enabled via the `features` URL parameter.
- Storage tools are disabled by default — add `features=storage` or include `storage` in the comma-separated list.
- Branching tools are disabled by default and require a paid plan.
- Account tools are disabled when `project_ref` is set.

## Rate limited

**Problem**

Tools return rate limit errors.

**Cause**

Supabase API rate limits have been exceeded.

**Solution**

- Reduce the number of parallel tool calls.
- Space out requests between agent steps.
- Wait before retrying.

## CI/CD authentication

**Problem**

Cannot complete the OAuth flow in a CI environment.

**Cause**

Browser-based OAuth is not available in CI.

**Solution**

- Create a personal access token (PAT) in your Supabase dashboard.
- Pass it as an `Authorization: Bearer <token>` header in the MCP server configuration.
- Always scope the server to a project and use read-only mode in CI.