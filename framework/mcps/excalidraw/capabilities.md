# Excalidraw MCP Capabilities

## What It Can Do

Excalidraw MCP renders hand-drawn style diagrams from text descriptions. It has a single tool — `draw` — that accepts natural language descriptions of what to draw.

### The `draw` Tool

| Aspect | Description |
| --- | --- |
| Input | Plain-text description of the diagram (e.g., "a user connecting to a server which talks to a database") |
| Output | Interactive Excalidraw HTML interface in the chat |
| Style | Hand-drawn (sketch-like) rendering |

### What You Can Draw

- **Architecture diagrams**: system components, services, databases, external integrations
- **Flowcharts**: decision trees, process flows, state transitions
- **Entity relationship diagrams**: tables, relationships, cardinality
- **Wireframes**: UI layouts, screen mockups, component placement
- **Network topology**: servers, clients, load balancers, network paths
- **Concept maps**: relationships between ideas, hierarchies
- **Sequence diagrams**: interaction flows between components

## What It Cannot Do

- It cannot modify existing diagrams (each `draw` call creates a new diagram).
- It cannot export diagrams to image formats programmatically — use the interactive viewer's built-in export.
- It cannot save diagrams to a file system — diagrams exist in the chat session.
- It cannot render without an MCP Apps-compatible client.
- It does not support authentication or access control — the remote server is public.
- It cannot read or edit existing Excalidraw files (.excalidraw format).

## Best Practices

- Be specific in descriptions: include component names, relationships, and directional flow.
- Request diagrams that clarify architectural concepts rather than trivial visual decoration.
- Use the interactive viewer to zoom, pan, and adjust elements after rendering.
- The diagram is best for conceptual understanding — use architecture diagrams, not pixel-perfect UI mockups.

## Common Workflows

### Architecture documentation

1. Describe the system architecture (services, databases, clients, external APIs).
2. Ask Excalidraw to render it as an architecture diagram.
3. Use the interactive viewer to explore and adjust.
4. Reference the diagram in design docs or ADRs.

### Explaining design decisions

1. Describe the current and proposed architecture.
2. Ask Excalidraw to draw both for visual comparison.
3. Annotate differences and tradeoffs on the rendered diagrams.
