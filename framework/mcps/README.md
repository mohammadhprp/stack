# MCP Catalog

A curated catalog of Model Context Protocol (MCP) servers for AI coding agents.

This catalog is framework-agnostic: each entry documents what the server does, when to use it, how to install it, how to configure it across common AI clients, and how to troubleshoot it.

## Available MCPs

| MCP | Purpose | Best fit |
| --- | --- | --- |
| [Excalidraw MCP](./excalidraw/README.md) | Stream hand-drawn Excalidraw diagrams with interactive fullscreen editing — architecture diagrams, flowcharts, wireframes. | Any project that benefits from visual diagrams, architecture visualization, or design sketching. |
| [Figma MCP Go](./figma-mcp-go/README.md) | Full read/write access to Figma designs via plugin bridge — no API token or rate limits. | Any project that works with Figma designs. |
| [Laravel Boost](./laravel-boost/README.md) | Gives AI agents Laravel application context, Laravel ecosystem documentation search, and safe local inspection tools. | Laravel applications. |
| [Playwright MCP](./playwright-mcp/README.md) | Cross-browser automation via accessibility snapshots — navigate, click, type, screenshot, evaluate, network control, storage management. | Any project needing browser automation or end-to-end testing. |
| [Supabase MCP](./supabase-mcp/README.md) | Manage Supabase projects, query databases, design schema, run migrations, deploy Edge Functions, manage Auth/Storage/Realtime, and search docs. | Any project using Supabase as backend. |

## Why MCPs Are Useful

MCP servers give AI clients a standard way to connect to project-aware tools and data sources. For coding agents, this can improve:

- **Context accuracy**: agents can inspect real application state instead of relying only on static files.
- **Tool consistency**: one server can work across many clients that support MCP.
- **Operational safety**: capabilities and permissions are explicit in configuration.
- **Repeatability**: teams can document the same install and verification steps for every developer.

## How to Choose an MCP

Use an MCP when it provides concrete project context or a controlled tool that the AI client cannot reliably infer from source code alone.

Before installing one, check:

- **Scope**: Does it match the framework, platform, or workflow in the project?
- **Capabilities**: Does it expose tools that materially improve the agent's work?
- **Permissions**: Can it read or modify sensitive data, run commands, or query databases?
- **Client support**: Does your AI client support local stdio MCP servers or the server's required transport?
- **Maintenance**: Is the server actively maintained and documented by an authoritative source?

## Catalog Entry Structure

Every MCP entry must use this structure:

```text
mcps/<mcp-slug>/
├── README.md
├── install.md
├── capabilities.md
├── troubleshooting.md
└── configs/
    ├── opencode.json
    └── .env.example
```

## Naming Conventions

- Use lowercase kebab-case directory names, for example `laravel-boost`.
- Use the upstream MCP name in page titles, for example `# Laravel Boost`.
- Use stable MCP server names in configuration examples, for example `laravel-boost`.
- Keep file names exactly as shown in the catalog structure.

## How to Contribute a New MCP

1. Create `mcps/<mcp-slug>/` using the standard structure.
2. Prefer official documentation for install commands, requirements, capabilities, and configuration.
3. Include client-specific config examples even when multiple clients share the same JSON schema.
4. Document where each config file belongs in the MCP's `install.md`.
5. Add limitations and permission notes in `capabilities.md`.
6. Add common startup, authentication, permission, and version issues in `troubleshooting.md`.
7. Add the MCP to the table in this README.

## Related Repository Areas

- [`skills/`](../skills/) contains task-specific procedures for AI coding agents.
- [`references/standards/`](../references/standards/) contains canonical engineering rules.
- [`references/templates/`](../references/templates/) contains reusable engineering deliverables.
