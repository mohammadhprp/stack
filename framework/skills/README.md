# Skill Catalog

A curated catalog of task-specific procedures for AI coding agents.

This catalog is framework-agnostic: each skill defines when to activate, a step-by-step workflow, and the engineering judgment expected of the agent.

## Available Skills

| Skill | Purpose | Best fit |
| --- | --- | --- |
| [adhd](./adhd/SKILL.md) | Shape output for ADHD readers: action first, numbered steps, restated state, no tangents. | Any task where the response must be immediately actionable without preamble or recap. |
| [agent-browser](./agent-browser/SKILL.md) | Automate browser and Electron workflows for navigation, testing, screenshots, and data extraction. | Website interaction, browser automation, exploratory testing, QA, and Electron desktop app workflows. |
| [architect](./architect/SKILL.md) | Sketch architecture and module boundaries before implementation. | Non-trivial design and implementation work. |
| [arena](./arena/SKILL.md) | Compare parallel candidate solutions and synthesize the strongest result. | Non-trivial artifacts and design alternatives. |
| [backend-best-practices](./backend-best-practices/SKILL.md) | Consolidated backend engineering practices for API design, architecture, data, security, testing, observability, performance, debugging, and refactoring. | Any backend task that benefits from structured domain guidance across multiple backend disciplines. |
| [bro](./bro/SKILL.md) | Restate the last message in plain human language without jargon. | Simplifying technical explanations. |
| [changelog](./changelog/SKILL.md) | Create or update `CHANGELOG.md` entries. | Release notes and unreleased changes. |
| [commit](./commit/SKILL.md) | Create atomic Git commits with conventional messages. | Committing reviewed changes. |
| [diagram-design](./diagram-design/SKILL.md) | Create technical and product diagrams as standalone HTML files with inline SVG and an opinionated editorial design system. | Architecture, data, process, and other diagrams that communicate more clearly than prose or tables. |
| [docs-writer](./docs-writer/SKILL.md) | Write, review, and edit documentation files with consistent structure, tone, and technical accuracy. | Creating docs, reviewing markdown files, writing READMEs, or updating `/docs` directories. |
| [effective-html](./effective-html/SKILL.md) | Create self-contained HTML artifacts with routed guidance for design, wireframes, prototypes, plans, and diagrams. | Standalone HTML reports, explainers, interfaces, wireframes, prototypes, plans, and diagrams. |
| [find-skills](./find-skills/SKILL.md) | Discover, evaluate, and install skills from the open agent skills ecosystem. | Finding an installable skill for a specialized task or extending an agent's capabilities. |
| [frontend-design](./frontend-design/SKILL.md) | Distinctive, intentional visual design for new UI or reshaping existing UI — aesthetic direction, typography, and choices that don't read as templated defaults. | Building new interfaces, reshaping existing UI, or escaping generic AI-generated design looks. |
| [grilling](./grilling/SKILL.md) | Grill plans and decisions relentlessly through design-tree rounds until shared understanding is reached. | Stress-testing ideas, surfacing hidden assumptions, and scoping plans before implementation. |
| [glab](./glab/SKILL.md) | Use the GitLab CLI (glab) to manage merge requests, issues, pipelines, and repositories from the command line. | Any project hosted on GitLab (SaaS or self-hosted). |
| [humanizer](./humanizer/SKILL.md) | Remove signs of AI-generated writing from text — inflated importance, promotional language, em dash overuse, rule of three, AI vocabulary, and filler phrases. | Editing or reviewing prose to make it sound more natural and human-written. |
| [improve](./improve/SKILL.md) | Audit repositories as a read-only senior advisor and produce prioritized implementation plans for another agent. | Repository-wide audits, improvement roadmaps, and implementation handoffs. |
| [how](./how/SKILL.md) | Explain subsystem architecture, runtime flow, ownership, and layering. | Code walkthroughs and architecture questions. |
| [jira-cli](./jira-cli/SKILL.md) | Use the Jira CLI (jira) to manage issues, sprints, epics, comments, transitions, and worklogs from the command line. | Any project using Jira (Cloud or self-hosted) for issue tracking. |
| [commit](./commit/SKILL.md) | Create atomic Git commits with conventional messages. | Committing reviewed changes. |
| [pull-request](./pull-request/SKILL.md) | Create or update GitHub pull requests (PRs). | GitHub contribution workflows. |
| [merge-request](./merge-request/SKILL.md) | Create GitLab merge requests (MRs). | GitLab contribution workflows. |
| [review](./review/SKILL.md) | Perform comprehensive code quality reviews. | Diff reviews and production-readiness checks. |
| [changelog](./changelog/SKILL.md) | Create or update `CHANGELOG.md` entries. | Release notes and unreleased changes. |
| [release](./release/SKILL.md) | Prepare and tag semantic-versioned releases. | Release management. |
| [laravel-best-practices](./laravel-best-practices/SKILL.md) | Apply Laravel PHP patterns for Eloquent, validation, security, testing, caching, and architecture. | Writing, reviewing, or refactoring Laravel PHP code. |
| [perf-web-optimization](./perf-web-optimization/SKILL.md) | Optimize web performance: bundle size, images, caching, lazy loading, and overall page speed. | A slow site, large bundles, layout shifts, poor Time to Interactive, or low Lighthouse scores. |
| [security-best-practices](./security-best-practices/SKILL.md) | Perform language- and framework-specific security best-practice reviews and suggest improvements, with support for Python, JavaScript/TypeScript, and Go. | Explicit security reviews or reports, secure-by-default coding, or auditing existing code for major vulnerabilities. |
| [sentry](./sentry/SKILL.md) | Inspect Sentry issues, summarize production errors, and pull health data via the Sentry API (read-only). | Checking production errors, summarizing recent crashes, or pulling a production error report. |
| [show-me](./show-me/SKILL.md) | Help the user understand the current topic visually with concise diagrams, code-shape sketches, and focused HTML artifacts. | Explaining logic, control flow, UI structure, file responsibility, or data flow with the smallest useful visual. |
| [skill-creator](./skill-creator/SKILL.md) | Create new skills, modify and improve existing skills, and measure skill performance. | Designing new skills, optimizing existing skills, or running evals. |
| [tlc-spec-driven](./tlc-spec-driven/SKILL.md) | Feature planning and implementation with adaptive phases, EARS testable requirements, atomic Conventional Commits, and independent verification. | Planning features, implementing with verification and atomic commits, or validating an implementation against a spec. |
| [taste](./taste/SKILL.md) | Unified design-quality skill spanning anti-slop frontend design, image generation and image-to-code, brand kits, minimalist and industrial UI, redesigns, and Google Stitch design systems. | Landing pages, portfolios, redesigns, website/mobile/brand imagery, or any UI that must not look AI-generated. |
| [tdd](./tdd/SKILL.md) | Drive clear bug fixes with focused failing regression tests. | Explicit TDD and regression-test workflows. |
| [unslop](./unslop/SKILL.md) | Remove AI writing patterns and restore a natural human voice. | Editing prose to remove generic or artificial phrasing. |
| [why](./why/SKILL.md) | Investigate the evidence and rationale behind code and design decisions. | Historical design-reasoning and motivation investigations. |

## How Skills Are Organized

Each skill entry follows this structure:

```text
skills/<skill-name>/
├── SKILL.md       # Full skill definition: triggers, workflow, standards references (Required)
├── examples.md    # Example output showing expected format (Required)
├── references/    # Docs loaded into context as needed (Optinal)
│   └── schemas.md      
└── scripts/      # Executable code (bash/node/python) for deterministic/repetitive tasks (Optinal)
    └── validate.sh
```

## How to Choose a Skill

Use the skill table above to find the entry closest to the task at hand. Skills are activated based on the `description` field in their frontmatter, which defines when an AI agent should load them.

Before activating a skill, check:

- **Fit**: does the skill's purpose match the current task?
- **Scope**: some skills are general, others are narrow (Laravel Best Practices).
- **Combination**: complex tasks may need multiple skills (e.g., API Design + Security + Testing).
- **Sequence**: some skills are best used in order (e.g., Architecture Review before Database Design).

## How to Contribute a New Skill

1. Create `skills/<skill-name>/` with `SKILL.md` and `examples.md`.
2. The `name` in the frontmatter must be **kebab-case** and match the directory name (e.g., `name: my-skill` for `skills/my-skill/`).
3. Write a clear `description` in the frontmatter — this determines when agents activate the skill.
4. Define when to activate, a step-by-step workflow, standards references, and expected outcomes.
5. Add 3+ realistic examples in `examples.md` showing good agent behavior.
6. Add the skill to the table in this README.

## Related Repository Areas

- [`mcps/`](../mcps/) contains the MCP (Model Context Protocol) server catalog.
- [`plugins/`](../plugins/) contains the OpenCode plugin catalog.
- [`references/standards/`](../references/standards/) contains canonical engineering rules.
- [`references/templates/`](../references/templates/) contains reusable engineering deliverables.
