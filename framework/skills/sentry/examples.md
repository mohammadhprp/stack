# Sentry Examples

## List recent production errors

User: "Check Sentry for what's broken in prod right now."

Good agent behavior:

- Confirm `SENTRY_AUTH_TOKEN` is set, and if missing, guide the user to create a read-only token without asking them to paste it in chat.
- Verify the installed `sentry` CLI, confirm the selected project, and use `sentry issue list` with the user's production environment, a 24h window, and JSON output.
- Return an ordered list with title, short ID, status, counts, and last-seen time.
- State explicitly when there are no results rather than padding the output.

## Summarize a specific issue

User: "What's going on with error ABC-123?"

Good agent behavior:

- Use the short ID directly with `sentry issue view ABC-123 --json`, then read its events when needed.
- Summarize the culprit, timestamp, environment, release, and URL without dumping raw stack traces.
- Redact PII such as emails and IPs from the output.
- Report the actual CLI results and any unavailable fields; do not run AI jobs or mutate issues.

## Production error report

User: "Give me a production error report for the last 24 hours."

Good agent behavior:

- Pull the top unresolved issues with default prod environment and 24h time range, limiting to a reasonable count.
- Order by most recent activity and group the highest-count errors first.
- Include the error title, short ID, event count, and affected environments for each.
- Flag anything needing immediate attention and offer to pull event details for the worst offenders.
