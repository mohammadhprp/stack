# Playwright MCP

## Overview

Playwright MCP lets coding agents automate and interact with web pages using [Playwright](https://playwright.dev). It uses structured accessibility snapshots instead of pixel-based input, so no vision model is needed. Published and maintained by Microsoft.

Official source:

- [GitHub: microsoft/playwright-mcp](https://github.com/microsoft/playwright-mcp)

## Features

- **Fast and lightweight**: Uses Playwright's accessibility tree, not pixel-based input or screenshots.
- **LLM-friendly**: Operates purely on structured data — no vision models required.
- **Cross-browser**: Supports Chromium, Firefox, and WebKit.
- **Deterministic tool application**: Avoids ambiguity common with screenshot-based approaches.
- **Persistent profiles**: Session data saved between runs with automatic workspace-scoped isolation.
- **Headless support**: Run headed or headless.
- **Capability system**: Opt-in tool categories (network, storage, devtools, vision, PDF).

## Supported AI Clients

Playwright MCP officially documents setup for Amp, Antigravity, Claude Code, Claude Desktop, Cline, Codex, Copilot, Cursor, Factory, Gemini CLI, Goose, Junie, Kiro, LM Studio, OpenCode, Qodo Gen, VS Code, Warp, and Windsurf.

## When to Use

Use Playwright MCP when an AI coding agent needs to navigate, interact with, or inspect web pages.

Good fits:

- Automating browser interactions (click, fill forms, navigate, type).
- Taking accessibility snapshots for structured page understanding.
- E2E test debugging and inspection.
- Capturing screenshots and console logs.
- Evaluating JavaScript in page context.
- Simulating network conditions and mocking requests.
- Managing cookies, localStorage, and session state.

Avoid using it when no browser interaction is needed or when you need pixel-level visual inspection (use `--caps=vision` if coordinates are required).

## Requirements

- Node.js 18 or newer.
- An MCP-compatible AI client.

## Related Skills

Relevant skills in this repository:

- [`backend-best-practices`](../../skills/backend-best-practices/SKILL.md): backend best practices for testing, debugging, and performance analysis.
