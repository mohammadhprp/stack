# Sentry Examples

## List recent production errors

User: "Check Sentry for what's broken in prod right now."

Good agent behavior:

- Confirm `SENTRY_AUTH_TOKEN` is set, and if missing, guide the user to create a read-only token without asking them to paste it in chat.
- Call the Sentry API using the bundled `sentry_api.py` script with the production environment and a 24h window.
- Return an ordered list with title, short ID, status, counts, and last-seen time.
- State explicitly when there are no results rather than padding the output.

## Summarize a specific issue

User: "What's going on with error ABC-123?"

Good agent behavior:

- Resolve the short ID to a numeric issue ID via the list endpoint, then fetch issue detail and events.
- Summarize the culprit, timestamp, environment, release, and URL without dumping raw stack traces.
- Redact PII such as emails and IPs from the output.
- Note that the API is called directly each time rather than cached.

## Production error report

User: "Give me a production error report for the last 24 hours."

Good agent behavior:

- Pull the top unresolved issues with default prod environment and 24h time range, limiting to a reasonable count.
- Order by most recent activity and group the highest-count errors first.
- Include the error title, short ID, event count, and affected environments for each.
- Flag anything needing immediate attention and offer to pull event details for the worst offenders.
