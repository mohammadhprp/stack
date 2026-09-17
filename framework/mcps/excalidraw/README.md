# Excalidraw MCP

## Overview

Excalidraw MCP lets coding agents stream hand-drawn Excalidraw diagrams with smooth viewport camera control and interactive fullscreen editing. It renders diagrams as interactive HTML interfaces directly in the chat (MCP Apps).

Official source:

- [GitHub: excalidraw/excalidraw-mcp](https://github.com/excalidraw/excalidraw-mcp)
- [Hosted server](https://mcp.excalidraw.com)

## Features

- **Hand-drawn diagrams**: Render architecture diagrams, flowcharts, wireframes, and illustrations in the signature Excalidraw hand-drawn style.
- **Interactive editing**: Fullscreen editing with zoom, pan, and element manipulation.
- **Smooth camera control**: Animated viewport transitions for diagram exploration.
- **Streaming output**: Diagrams stream incrementally as they're drawn.
- **No API key required**: The remote server requires no authentication.

## When to Use

Use Excalidraw MCP when you need to visualize:

- System architecture diagrams and component relationships
- Data flow diagrams showing how data moves through a system
- Wireframes and UI mockups
- Flowcharts and decision trees
- Entity relationship diagrams
- Any concept easier to understand visually than as text

## Tools

The server exposes a single tool — `draw` — which accepts a plain-text description of what to draw.

| Tool | Description |
| --- | --- |
| `draw` | Draw an Excalidraw diagram from a text description |
