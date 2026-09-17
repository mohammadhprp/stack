---
name: backend-best-practices
description: Consolidated best practices for backend engineering — API design, data, security, testing, observability, performance, debugging, architecture, and refactoring.
version: 0.1.0
---

# Backend Best Practices

Consolidated best practices for backend engineering, organized by domain. Each section links to a detailed reference file. For exact API or framework syntax, verify with search-docs.

## Consistency First

Before applying any rule, check what the project already does. There are multiple valid approaches — the best choice is the one the codebase already uses, even if another pattern would be theoretically better. Inconsistency is worse than a suboptimal pattern.

Check sibling files, related controllers, models, or tests for established patterns. If one exists, follow it — don't introduce a second way. These rules are defaults for when no pattern exists yet, not overrides.

## Quick Reference

### 1. API Design → `references/api-design.md`

- Model contracts around resources and state transitions, not storage tables
- Pagination, filtering, and sorting must be bounded, deterministic, and documented
- Versioning and compatibility decisions must protect existing clients
- Authentication proves identity; authorization decides allowed actions per resource
- Validation errors should be precise enough for clients to fix requests without leaking sensitive internals
- Idempotency is required for retries that create orders, payments, inventory reservations, or notifications
- Rate limiting and abuse controls should fail predictably with useful retry guidance

### 2. Architecture Review → `references/architecture-review.md`

- Component boundaries should align with business capabilities
- High cohesion means related behavior changes together; low coupling means unrelated areas can change independently
- Boundaries should align to business capabilities, data ownership, and operational responsibility
- Layering must prevent policy decisions from leaking into transport, storage, or presentation glue
- Reliability design must include timeout, retry, fallback, and partial-failure behavior
- Scalability claims require bottleneck analysis, not generic distribution

### 3. Database Design → `references/database-design.md`

- Data models must encode business invariants through constraints whenever possible
- Transactions should cover the full consistency boundary and no more
- Locking strategy must be explicit for concurrent updates to orders, inventory, balances, or quotas
- Indexes should match real access patterns and be reviewed for write cost
- Normalize for integrity first; denormalize only with ownership and refresh rules
- Migrations must define backfill, verification, compatibility window, and rollback or mitigation

### 4. Debugging → `references/debugging.md`

- Reproduce before changing code when possible; if not, collect enough production evidence to form a falsifiable hypothesis
- Isolate variables using binary search across versions, inputs, dependencies, and configuration
- Logs, metrics, and traces should be correlated by request, job, user, order, or payment identifier
- Fix root causes, not only symptoms, and add regression tests or alerts for recurrence
- Preserve evidence during incidents; do not destroy state needed for diagnosis

### 5. Observability → `references/observability.md`

- Structured logs should describe events with stable fields and no sensitive data
- Metrics should track rates, errors, durations, saturation, and business outcomes
- Traces should show cross-boundary latency and failure points
- Health checks should distinguish process liveness from dependency readiness
- Alerts should be actionable, tied to user impact, and include runbook guidance
- Dashboards should support both real-time triage and trend review

### 6. Performance → `references/performance.md`

- Measure before optimizing and define the target metric: latency, throughput, cost, or capacity
- CPU work should be bounded and avoid repeated expensive computation
- Memory use should avoid unbounded aggregation, buffering, and cache growth
- Network calls need timeouts, batching where appropriate, and payload discipline
- Caching requires invalidation, freshness expectations, and failure behavior
- Parallelism must preserve correctness and respect downstream limits
- Benchmarks should reflect realistic data sizes and access patterns

### 7. Refactoring → `references/refactoring.md`

- Preserve externally visible behavior unless a behavior change is explicitly requested
- Make one conceptual change at a time
- Remove duplication by extracting the stable idea, not by forcing unrelated code through one abstraction
- Improve names when they reduce cognitive load for future changes
- Characterization tests are useful before touching risky legacy behavior
- Stop refactoring when the code is simpler enough for the current goal

### 8. Security → `references/security.md`

- Treat all input as untrusted, including internal calls, background jobs, files, and message payloads
- Enforce least privilege for users, services, tokens, storage, and operational access
- Secrets must never appear in source, logs, errors, metrics, traces, or examples
- Use encryption deliberately for data in transit and sensitive data at rest
- Prevent injection by separating commands from data and validating allowed shapes
- Output encoding and safe error responses prevent data disclosure
- Abuse controls such as rate limits should protect expensive and sensitive actions

### 9. Testing → `references/testing.md`

- Unit tests prove local rules and edge cases quickly
- Integration tests prove persistence, transactions, serialization, and external boundaries
- Contract tests protect clients and providers from incompatible changes
- End-to-end tests should cover critical user journeys, not every branch
- Deterministic tests avoid time, order, randomness, and shared-state flakiness
- Mocks should model behavior and failure, not implementation trivia

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

## How to Apply

1. Identify the domain your task falls into (API, database, security, etc.) and open the corresponding reference file.
2. Check sibling files for existing patterns — follow those first per Consistency First.
3. When a task spans multiple domains (e.g., a new endpoint touches API design, database design, security, and testing), consult all relevant reference files.
4. Verify syntax with search-docs where applicable.

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.

## Checklist

- [ ] Domain-specific reference files are consulted for the task at hand.
- [ ] Simpler alternatives were considered and documented.
- [ ] Data integrity and backward compatibility are protected.
- [ ] Security and authorization impact is reviewed.
- [ ] Tests cover normal paths, edge cases, and failure paths.
- [ ] Logs, metrics, traces, or health signals are included when operationally relevant.
- [ ] Deployment and rollback are understood and tested.
- [ ] The deliverable explains reasoning and evidence, not just code.
