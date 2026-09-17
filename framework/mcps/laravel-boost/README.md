# Laravel Boost

## Overview

Laravel Boost is Laravel's first-party MCP server and AI-assistance package for Laravel applications. It gives AI coding agents project-aware Laravel context, local inspection tools, documentation search, generated AI guidelines, and optional agent skills.

Official sources:

- [Laravel Boost product page](https://laravel.com/ai/boost)
- [Laravel Boost documentation](https://laravel.com/docs/boost)

## Features

- Exposes an MCP server from the Laravel application through `php artisan boost:mcp`.
- Reads PHP, Laravel, database, package, and Eloquent model information.
- Inspects routes, configuration, logs, browser errors, and database schema.
- Executes database queries and Tinker code through MCP tools.
- Searches Laravel's hosted documentation API with package-aware results.
- Generates AI guidelines and skills for supported Laravel ecosystem packages.

## Supported AI Clients

Laravel Boost officially documents setup for Cursor, Claude Code, Codex, Gemini CLI, GitHub Copilot in VS Code, and Junie. The MCP server can also be registered manually in other MCP-compatible clients using the documented command and arguments.

- OpenCode

## When to Use

Use Laravel Boost when an AI coding agent is working inside a Laravel application and needs current, project-specific Laravel context.

Good fits:

- Building or modifying Laravel features.
- Inspecting routes, models, configuration, logs, or database schema.
- Asking Laravel-version-specific implementation questions.
- Upgrading or maintaining Laravel ecosystem packages supported by Boost guidelines or skills.

Avoid using it as a generic MCP outside Laravel projects.

## Requirements

- A Laravel application.
- PHP and Composer available in the project environment.
- `laravel/boost` installed as a development dependency.
- An AI client that supports MCP, or one of the clients Boost can configure during `php artisan boost:install`.
- Local access to run `php artisan boost:mcp` from the Laravel application root.

## Related Skills

Relevant skills in this repository:

- [`backend-best-practices`](../../skills/backend-best-practices/SKILL.md): backend best practices for API design, database design, debugging, testing, and security.
