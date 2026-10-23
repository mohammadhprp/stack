---
name: "sentry"
description: "Use when the user asks to inspect Sentry issues or events, summarize recent production errors, or pull basic Sentry health data via the Sentry CLI; perform read-only queries using the `sentry` command."
---


# Sentry (Read-only Observability)

Inspect with the `sentry` CLI, not the unrelated `sentry-cli` release uploader.
Check `sentry --help` and command help for the installed version before use.
Confirm the organization, project, environment, and time range; auto-detection
is a hint, not authorization to inspect another project. Use read-only queries
and GET-only API fallbacks. Do not resolve, assign, delete, deploy, create plans,
or start AI jobs unless separately requested. Never print credentials or raw PII.

## Quick start

- If not already authenticated, ask the user to run `sentry auth login` or set `SENTRY_AUTH_TOKEN` as an env var.
- The CLI auto-detects org/project from DSNs in `.env` files, source code, config defaults, and directory names. Only specify `<org>/<project>` if auto-detection fails or picks the wrong target.
- Defaults: time range `24h`, environment `production`, limit 20.
- Always use `--json` when processing output programmatically. Use `--json --fields` to select specific fields and reduce output size.
- Use `sentry schema <resource>` to discover API endpoints quickly.

If the CLI is not installed, give the user these steps:
1. Ask the user to install the CLI using its official instructions at https://cli.sentry.dev/.
2. Authenticate: `sentry auth login`
3. Confirm authentication: `sentry auth status`
- Never ask the user to paste the full token in chat. Ask them to set it locally and confirm when ready.

## Core tasks (use Sentry CLI)

Use the `sentry` CLI for all queries. It handles authentication, org/project detection, pagination, and retries automatically. Use `--json` for machine-readable output.

### 1) List issues (ordered by most recent)

```bash
sentry issue list \
  --query "is:unresolved environment:production" \
  --period 24h \
  --limit 20 \
  --json --fields shortId,title,priority,level,status
```

If auto-detection doesn't resolve org/project, pass them explicitly:
```bash
sentry issue list {your-org}/{your-project} \
  --query "is:unresolved environment:production" \
  --period 24h \
  --limit 20 \
  --json
```

### 2) Resolve an issue short ID to issue detail

```bash
sentry issue view {ABC-123} --json
```

Use the short ID format (e.g., `ABC-123`), not the numeric ID.

### 3) Issue detail

```bash
sentry issue view {ABC-123}
```

### 4) Issue events

```bash
sentry issue events {ABC-123} --limit 20 --json
```

### 5) Event detail

```bash
sentry event view {your-org}/{your-project}/{event_id} --json
```

### Optional AI operations

Upstream also describes `sentry issue explain` and `sentry issue plan`. They are
not part of this read-only default: check installed help and obtain explicit
authorization for service-side jobs or quota use before considering them.

## Fallback: arbitrary API access

For endpoints not covered by dedicated CLI commands, use `sentry api`:
```bash
sentry api /api/0/organizations/{your-org}/ --method GET
```

Use `sentry schema` to discover available API endpoints:
```bash
sentry schema issues
```

## Inputs and defaults

- `org_slug`, `project_slug`: auto-detected by the CLI from DSNs, env vars, and directory names. Override with positional `{your-org}/{your-project}` if auto-detection fails.
- `time_range`: default `24h` (pass as `--period 24h`).
- `environment`: default `production`; use the actual environment selected by the user.
- `limit`: default 20 (pass as `--limit`).
- `search_query`: optional `--query` parameter, uses Sentry search syntax (e.g., `is:unresolved`, `assigned:me`).
- `issue_short_id`: use directly with `sentry issue view`.

## Output formatting rules

- Issue list: show title, short ID, status, counts, times, environments, and tags
  only when returned. Fetch details when needed; do not invent fields omitted by
  `--fields`. Order by most recent activity when that is the requested view.
- Event detail: include culprit, timestamp, environment, release, url.
- If no results, state explicitly.
- Redact PII in output (emails, IPs). Do not print raw stack traces.
- Never echo auth tokens.

## Golden test inputs

- Org: `{your-org}`
- Project: `{your-project}`
- Issue short ID: `{ABC-123}`

Example prompt: "List the top 10 open issues for prod in the last 24h."
Expected: ordered list with titles, short IDs, counts, last seen.

## Compatibility fallback

The existing `scripts/sentry_api.py` is retained for users who explicitly need
the older API workflow. Resolve its path from this loaded `SKILL.md` directory,
never `~/.agent-skills` or another fixed harness location. It is a read-only
fallback, not the default CLI implementation; report which path was used.

## Source

Adapted from [openai/skills](https://github.com/openai/skills/blob/49f948faa9258a0c61caceaf225e179651397431/skills/.curated/sentry/SKILL.md). Stack retains its token/PII safeguards and legacy fallback, verifies target context, and excludes AI service jobs from the read-only default.
