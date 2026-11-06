---
name: untangle-codebase
description: Refactor structural coupling in a Laravel backend or a selected area while preserving its observable behavior.
disable-model-invocation: true
argument-hint: "[optional area or concern to focus on]"
---

Untangle the Laravel backend within the user's requested scope. Identify what should change together, redesign the most consequential coupling, and implement and verify the changes. Respect existing project conventions and explicit limits.

## Optional cleanup

Unless the user already specified their choices, ask which of these additions to include:

| Addition | What it does |
| --- | --- |
| **Comments** | Remove comments that only restate the code. |
| **Tests** | Remove tests that wouldn't catch a real bug, and fix the ones that fail without one. |

Do only the selected cleanup. Keep tests as a regression safety net during refactoring. Preserve comments explaining external constraints or non-obvious decisions.

## Discover and plan

Before editing, inspect the application's Laravel and PHP versions, dependency and test configuration, project conventions, and the relevant routes, controllers, requests, models, services, jobs, events, listeners, policies, migrations and tests. Trace the actual execution paths and use Git history when useful. Record concrete problems and proposed fixes in a decision log with file references.

Look for hidden links, needless coupling, surplus wrappers or unused code, and special cases that indicate a missing domain concept. Check Laravel-specific seams:

- Duplicated validation, authorization, business rules or query constraints across HTTP endpoints, jobs and commands.
- Business workflows coupled to controllers, request state, facades, global state or a particular external provider, where that coupling makes real changes harder.
- Implicit behavior in Eloquent observers, model events, scopes, accessors and relationship loading; trace their callers and side effects before moving logic.
- Transaction boundaries, concurrency, idempotency, queue retries and event dispatch timing across database writes and external calls.
- N+1 queries, inconsistent tenant or authorization constraints, and APIs whose response or error behavior depends on hidden conventions.

Distinguish intentional framework conventions and necessary business exceptions from harmful coupling. Do not treat every conditional, retry, abstraction or framework dependency as a defect. Prefer cohesive domain behavior and Laravel's existing facilities over speculative interfaces, repositories or a wholesale folder reorganization.

## Implement and verify

Group related problems under the smallest coherent design change. Refactor incrementally, keeping public routes, response shapes, authorization, persisted data, transaction semantics and asynchronous side effects stable unless the user requested a change. Preserve migration history and avoid destructive database commands or production services; use an isolated test database and safe test configuration. Do not dispatch real external calls or jobs during verification.

Run the relevant existing tests before and after changes where practical, then the project's appropriate broader checks. Add or adjust behavior-focused tests when needed to protect affected workflows, especially authorization, persistence, transactions and queued side effects. If test cleanup was selected, perform it after the structural work and retain coverage of real regressions. Investigate failures rather than weakening tests to make them pass.

Revisit the decision log and address worthwhile issues within scope. Report changes, verification and any remaining risks or unverified behavior. Use subagents only when the user has authorized delegation and the host supports it; otherwise review the work directly.

## Decision log example

- Centralize order submission in the existing order workflow, keeping HTTP validation at the boundary and dispatching the receipt after the transaction commits.
  - `OrderController` and `SubmitOrderJob` duplicate total calculation and can diverge.
  - The controller dispatches a receipt before the order transaction finishes.
  - Verify both entry points, rollback behavior and the receipt dispatch with the project's test facilities.

Adapted from [caldoai/skills](https://github.com/caldoai/skills/tree/f3cf47774975a7d9fca45730babf4ff7ad765ada/skills/untangle-codebase), MIT.
