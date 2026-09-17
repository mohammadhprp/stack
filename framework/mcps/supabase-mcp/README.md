# Supabase MCP

## Overview

Supabase MCP is a hosted, officially maintained MCP server that lets AI agents manage Supabase projects — query databases, design schema, run migrations, deploy Edge Functions, manage Auth users and Storage, and more. It uses OAuth authentication and communicates over Streamable HTTP.

Official source:

- [Supabase MCP Server docs](https://supabase.com/docs/guides/ai-tools/mcp)
- [Tool reference](https://supabase.com/docs/guides/ai-tools/mcp#available-tools)
- [Configuration options](https://supabase.com/docs/guides/ai-tools/mcp#configuration-options)
- [GitHub: supabase/mcp](https://github.com/supabase/mcp)

## Features

- Execute SQL queries against your database.
- List tables, extensions, and migrations.
- Apply database migrations.
- Retrieve service logs (API, Postgres, Edge Functions, Auth, Storage, Realtime).
- Get security and performance advisors.
- Get project API URLs and publishable keys.
- Generate TypeScript types from database schema.
- List, get, and deploy Edge Functions.
- Manage Supabase projects and organizations.
- Get cost information and confirm costs.
- Search Supabase documentation.
- Create, merge, list, delete, reset, and rebase database branches (paid plans, experimental).
- List storage buckets and get/update storage configuration.
- OAuth authentication — no manual token management.
- Feature groups to restrict which tools are exposed.
- Read-only mode and project-scoped mode for security.

## Supported AI Clients

Supabase MCP officially documents setup for Claude Code, Cursor, VS Code (GitHub Copilot), Windsurf, ChatGPT, Claude.ai, Goose, Factory, Codex, Gemini CLI, Kiro, Kimi Code, Antigravity, and OpenCode.

- OpenCode

## When to Use

Use Supabase MCP when an AI coding agent needs to interact with Supabase projects — managing schema, querying data, deploying functions, or inspecting configuration.

Good fits:

- Designing tables and generating migrations from scratch.
- Querying data for reports or debugging.
- Deploying and listing Edge Functions.
- Generating TypeScript types from an existing schema.
- Managing Supabase projects (create, pause, restore).
- Searching Supabase documentation for API guidance.
- Creating and managing development branches.

Avoid using it with production data. Supabase MCP is designed for development and testing. Use read-only mode, project scoping, and feature groups to limit blast radius.

## Requirements

- A Supabase account with a project.
- An MCP-compatible AI client that supports remote HTTP (Streamable HTTP) MCP servers.
- For branching features: a paid Supabase plan.

## Related Skills

Relevant skills in this repository:

- [`backend-best-practices`](../../skills/backend-best-practices/SKILL.md): backend best practices for API design, testing, and debugging.