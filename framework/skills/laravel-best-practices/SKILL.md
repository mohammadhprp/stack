---
name: laravel-best-practices
description: "Apply this skill whenever writing, reviewing, or refactoring Laravel PHP code. This includes creating or modifying controllers, models, migrations, form requests, policies, jobs, scheduled commands, service classes, and Eloquent queries. Triggers for N+1 and query performance issues, caching strategies, authorization and security patterns, validation, error handling, queue and job configuration, route definitions, and architectural decisions. Also use for Laravel code reviews and refactoring existing Laravel code to follow best practices. Covers any task involving Laravel backend PHP code patterns."
license: MIT
metadata:
  author: laravel
---

# Laravel Best Practices

Best practices for Laravel, organized as an index of reference files. Each reference file teaches what to do and why. For exact API syntax, verify with `search-docs`.

## Consistency First

Before applying any reference, check what the application already does. Laravel offers multiple valid approaches, and the best choice is the one the codebase already uses, even if another pattern would be theoretically better. Inconsistency is worse than a suboptimal pattern.

Check sibling files, related controllers, models, or tests for established patterns. If one exists, follow it. Don't introduce a second way. These references are defaults for when no pattern exists yet, not overrides.

## How to Apply

1. Check the changed files, nearby code, project configuration, and relevant tests for established patterns. Deviate only for a correctness or security defect, and call the deviation out.
2. Map every affected concern to the reference index below. Read each mapped reference file before editing. Skip unrelated reference files.
3. Make the smallest coherent change. Keep the application's architecture and naming instead of introducing a second pattern for the same job.
4. Verify version-sensitive Laravel APIs for the installed version with `search-docs`, or inspect the installed framework when it is unavailable.
5. Run the narrowest relevant tests first, then the project's formatting and static-analysis checks when the change warrants them.
6. Re-read the diff against every mapped reference before finishing.

## References Index

Cross-cutting changes often need more than one reference file.

| Concern | Read |
| --- | --- |
| Query count, eager loading, indexes, large datasets | [`references/db-performance.md`](references/db-performance.md) |
| Production monitoring, slow queries, exception dashboards | [`references/pulse.md`](references/pulse.md) |
| Subqueries, aggregates, complex ordering and query plans | [`references/advanced-queries.md`](references/advanced-queries.md) |
| Models, relationships, scopes, casts | [`references/eloquent.md`](references/eloquent.md) |
| Full-text search, indexing, searchable models | [`references/scout.md`](references/scout.md) |
| Authentication, authorization, input safety, secrets, uploads | [`references/security.md`](references/security.md) |
| OAuth2 servers, clients, token authentication | [`references/passport.md`](references/passport.md) |
| Roles and permissions with spatie/laravel-permission | [`references/spatie-laravel-permission.md`](references/spatie-laravel-permission.md) |
| Model media uploads, conversions, responsive images | [`references/spatie-medialibrary.md`](references/spatie-medialibrary.md) |
| Form Requests and validation references | [`references/validation.md`](references/validation.md) |
| Controllers, route binding, resources, middleware | [`references/routing.md`](references/routing.md) |
| Schema changes, columns, foreign keys, indexes | [`references/migrations.md`](references/migrations.md) |
| Jobs, retries, uniqueness, batches, Horizon | [`references/queue-jobs.md`](references/queue-jobs.md) |
| Cache lifetime, invalidation, locks, memoization | [`references/caching.md`](references/caching.md) |
| Outbound requests, retries, timeouts, fakes | [`references/http-client.md`](references/http-client.md) |
| Exceptions, reporting, rendering, log context | [`references/error-handling.md`](references/error-handling.md) |
| Events and notifications | [`references/events-notifications.md`](references/events-notifications.md) |
| Mailables and mail assertions | [`references/mail.md`](references/mail.md) |
| Scheduled tasks and overlap protection | [`references/scheduling.md`](references/scheduling.md) |
| Database backups, cleanup, monitoring, notifications | [`references/spatie-laravel-backup.md`](references/spatie-laravel-backup.md) |
| Collections, lazy iteration, bulk operations | [`references/collections.md`](references/collections.md) |
| Blade components, attributes, composers | [`references/blade-views.md`](references/blade-views.md) |
| Tailwind CSS classes, v4 configuration, extraction | [`references/tailwindcss.md`](references/tailwindcss.md) |
| Environment values and application configuration | [`references/config.md`](references/config.md) |
| Feature flags and per-user feature checks | [`references/pennant.md`](references/pennant.md) |
| Pest/PHPUnit patterns, factories, fakes | [`references/testing.md`](references/testing.md) |
| Pest 4 syntax, datasets, browser and architecture tests | [`references/pest-testing.md`](references/pest-testing.md) |
| Naming, helpers, file boundaries, PHP style | [`references/style.md`](references/style.md) |
| Actions, services, dependencies, application structure | [`references/architecture.md`](references/architecture.md) |
| Detecting and recording this app's conventions as rules | [`references/infer-conventions.md`](references/infer-conventions.md) |

## Decision rule

- Prefer framework features and existing application abstractions over new helpers or dependencies.
- Avoid speculative abstractions. Extract code when it creates a clear domain boundary, removes meaningful duplication, or makes behavior independently testable.
- Keep database access out of Blade views and prevent hidden N+1 queries across controllers, resources, jobs, and serialization.
