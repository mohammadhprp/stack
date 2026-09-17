---
name: effective-html
description: Create or redesign self-contained single-file HTML artifacts, including reports, explainers, landing pages, presentations, tools, wireframes, prototypes, plans, and diagrams. Route the request to the narrowest reference playbook, load only the guidance that applies, and verify the result at wide and narrow viewports. Do not use for ordinary application implementation when standalone HTML is not the deliverable.
---

# Effective HTML

Build one self-contained HTML file that makes the subject clearer, easier to use, or easier to understand. This skill combines creative direction with specialized workflows for structure, fidelity, planning, and visual explanation. The standard is consistent care, not a consistent look.

## Route the request first

Read the brief and choose the narrowest playbook that owns the main review question:

| Request | Reference |
|---|---|
| Palette, typography, composition, theming, or overall visual direction | [`design-artifact`](references/design-artifact.md) |
| Unsettled information hierarchy, navigation, task flow, or responsive structure | [`html-wireframe`](references/html-wireframe.md) |
| Polished mockup or working interactive flow | [`html-prototype`](references/html-prototype.md) |
| Plan, roadmap, implementation sequence, or rollout document | [`html-plan`](references/html-plan.md) |
| Relationships, sequence, topology, state, hierarchy, or quantitative structure | [`html-diagram`](references/html-diagram.md) |
| Report, explainer, presentation, landing page, data story, or broad tool | Continue with this skill and load the applicable references below |

The selected reference owns its specialist decisions. When a task spans domains, apply the relevant references in sequence and reconcile conflicts by user priority, project conventions, and the artifact's purpose. Do not load every reference by default.

## Read the room before designing

Inspect the user's request and supplied material. In a repository, look for `AGENTS.md`, `CLAUDE.md`, design-system documentation, tokens, existing components, and nearby artifacts.

Authority runs in this order:

1. The user's explicit visual and functional instructions.
2. The project's established design system and conventions.
3. The subject matter, audience, and purpose of the artifact.
4. Your own design judgment.

Before coding, settle the audience and job, form, visual register, fidelity, and useful interaction. Use real content and keep assumptions visible.

## Load only the guidance the artifact needs

- Reports, briefs, plans, explainers, and decks: [`documents-and-presentations`](references/documents-and-presentations.md)
- Interfaces, calculators, and tools: [`interfaces`](references/interfaces.md)
- Architecture, process, sequence, state, hierarchy, and concept diagrams: [`diagrams`](references/diagrams.md)
- Quantitative charts, tables, metrics, and data stories: [`charts-and-data`](references/charts-and-data.md)
- Visual direction when no project design system exists: [`creative-direction`](references/creative-direction.md)

Read every reference that materially applies, then give the artifact one coherent direction.

## Build contract

- Produce one `.html` file with essential CSS and JavaScript inline. It must work when opened directly without a build step or external service unless the user permits one.
- Use real content. Do not fill prominent space with placeholder copy, decorative statistics, or controls that do nothing.
- Let content determine structure. Make sequences ordered, comparisons scannable, interfaces stateful, and diagrams legible.
- Use semantic HTML, responsive layout, accessible contrast, visible keyboard focus, and reduced-motion handling.
- Keep accidental horizontal overflow out of the page body. Put intentionally broad content in a contained scrolling or pannable region.
- Define a small set of CSS tokens for the chosen direction and use them consistently.
- Treat motion as explanation or feedback. Remove animation that adds no meaning or useful feedback.
- Follow the user's or project's theme policy. When none exists, consider durable light and dark themes where they improve use.

## Verify and hand off

Write the file to the requested location, or choose a clear filename in the current workspace. When browser tooling is available, inspect wide and narrow viewports, exercise controls, check the console, and fix clipping, overlap, illegible text, broken states, and accidental overflow.

Before delivery, check that the visual direction belongs to the subject rather than a generic neighboring topic. Return the absolute path and a short description of the artifact's visual and interaction choices, including important assumptions and deliberately omitted behavior.
