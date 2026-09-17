# Figma MCP Go

## Overview

Figma MCP Go is an open-source Figma MCP server with full read/write access via a Figma plugin bridge. It does not use the Figma REST API, so there are no rate limits and no API token required. It exposes 73 tools for reading and writing Figma documents, styles, variables, components, prototypes, and more.

Official source:

- [GitHub: vkhanhqui/figma-mcp-go](https://github.com/vkhanhqui/figma-mcp-go)

## Features

- No Figma API token required.
- No rate limits — free plan friendly.
- Read and write live Figma data via plugin bridge — 73 tools total.
- Full design automation: styles, variables, components, prototypes, and content.
- Built-in design prompts: `read_design_strategy`, `design_strategy`, and more.
- Export tools: screenshots, PDFs, and design tokens.

## Supported AI Clients

Figma MCP Go works with any MCP-compatible AI client. The server is started via `npx` and communicates over stdio.

- OpenCode

## When to Use

Use Figma MCP Go when an AI coding agent needs to read or modify Figma designs directly.

Good fits:

- Converting designs to code.
- Updating text, colors, or layouts across Figma frames programmatically.
- Extracting design tokens, styles, or variable definitions.
- Automating component creation, renaming, or restructuring.
- Generating screenshots or PDF exports of frames.

Avoid using it as a generic MCP outside Figma-related workflows.

## Requirements

- Node.js and npm available on `PATH` (for `npx`).
- A Figma desktop app to install and run the companion plugin.
- An AI client that supports local stdio MCP servers.

## Related Skills

Relevant skills in this repository:

- [`backend-best-practices`](../../skills/backend-best-practices/SKILL.md): backend best practices for API design, testing, and debugging.