# Supabase MCP Capabilities

## What It Can Do

25+ tools organized into 8 feature groups. All groups except Storage are enabled by default.

### Database

| Tool | Description |
| --- | --- |
| `list_tables` | List all tables in the database |
| `list_extensions` | List available/installed Postgres extensions |
| `list_migrations` | List database migrations |
| `apply_migration` | Apply a database migration |
| `execute_sql` | Execute SQL queries |

### Debugging

| Tool | Description |
| --- | --- |
| `get_logs` | Retrieve service logs (API, Postgres, Edge Functions, Auth, Storage, Realtime) |
| `get_advisors` | Get security and performance advisors |

### Development

| Tool | Description |
| --- | --- |
| `get_project_url` | Get the API URL for a project |
| `get_publishable_keys` | Get publishable and legacy anon API keys for a project |
| `generate_typescript_types` | Generate TypeScript types from database schema |

### Edge Functions

| Tool | Description |
| --- | --- |
| `list_edge_functions` | List all Edge Functions |
| `get_edge_function` | Get a specific Edge Function |
| `deploy_edge_function` | Deploy an Edge Function |

### Account Management

Disabled when using `project_ref` scoping.

| Tool | Description |
| --- | --- |
| `list_projects` | List all Supabase projects |
| `get_project` | Get project details |
| `create_project` | Create a new Supabase project |
| `pause_project` | Pause a project |
| `restore_project` | Restore a paused project |
| `list_organizations` | List organizations |
| `get_organization` | Get organization details |
| `get_cost` | Get cost information for a project |
| `confirm_cost` | Confirm cost for a project |

### Docs

| Tool | Description |
| --- | --- |
| `search_docs` | Search Supabase documentation via the Content API (hybrid semantic + keyword search) |

### Branching (experimental, paid plans)

| Tool | Description |
| --- | --- |
| `create_branch` | Create a database branch |
| `list_branches` | List all branches |
| `delete_branch` | Delete a branch |
| `merge_branch` | Merge a branch into parent |
| `reset_branch` | Reset a branch to a previous state |
| `rebase_branch` | Rebase a branch onto parent |

### Storage (disabled by default)

| Tool | Description |
| --- | --- |
| `list_storage_buckets` | List storage buckets |
| `get_storage_config` | Get storage configuration |
| `update_storage_config` | Update storage configuration |

## What It Cannot Do

- It cannot work without OAuth authentication — no PAT required by default, but PAT is an option for CI.
- It cannot access Supabase resources outside the authenticated user's permissions.
- It cannot read/write files inside storage buckets (listing and config only).
- It cannot create or manage Auth users directly (only view logs).
- Branching requires a paid Supabase plan.
- Storage tools are disabled by default — must be explicitly enabled via `features` parameter.
- It cannot operate fully headless in its default OAuth mode — PAT auth required for CI.

## Configuration Options

URL query parameters to restrict scope and reduce blast radius:

- `read_only=true` — executes all queries as a read-only Postgres user (no inserts/updates/deletes).
- `project_ref=<id>` — scopes to a single project, disables account-level tools.
- `features=<groups>` — comma-separated list of feature groups to enable (e.g. `database,docs`).

## Security Best Practices

- **Never connect to production.** Use a development project with non-production data.
- **Use read-only mode** when connecting to any real data.
- **Scope to one project** with `project_ref` to limit blast radius.
- **Trim feature groups** to only what the task needs (e.g. `features=database,docs`).
- **Keep manual tool approval on** in the MCP client — review each tool call before executing.

## Common Workflows

### Explore database schema

1. Call `list_tables` to see available tables.
2. Call `execute_sql` with a `SELECT` query to inspect table structure or data.

### Generate TypeScript types

1. Call `list_tables` to see the schema.
2. Call `generate_typescript_types` to get typed definitions.

### Deploy an Edge Function

1. Call `list_edge_functions` to see existing functions.
2. Call `deploy_edge_function` with the function name and code.

### Search and apply from docs

1. Call `search_docs` to find Supabase documentation on a topic.
2. Call `execute_sql` to apply the documented approach.