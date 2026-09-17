# Observability Reference

Design logs, metrics, traces, health checks, dashboards, and alerts for production operations.

## When to Use

Use this when the task involves structured logging, metrics, distributed tracing, health checks, dashboards, alerting, service-level indicators, or incident readiness. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- Structured logs should describe events with stable fields and no sensitive data.
- Metrics should track rates, errors, durations, saturation, and business outcomes.
- Traces should show cross-boundary latency and failure points.
- Health checks should distinguish process liveness from dependency readiness.
- Alerts should be actionable, tied to user impact, and include runbook guidance.
- Dashboards should support both real-time triage and trend review.

## Workflow

1. Identify the operational signals needed: what questions will operators ask during normal operation and incidents?
2. Add structured logs with stable field names, correlation IDs, and severity levels.
3. Define metrics: rates (RPS), errors, durations (latency), saturation (CPU, memory, connections), and business outcomes.
4. Add distributed tracing for requests that cross service boundaries.
5. Implement health check endpoints that distinguish liveness (process alive) from readiness (able to serve traffic).
6. Set up dashboards for triage (real-time) and trends (daily/weekly).
7. Configure alerts that are actionable, documented with runbooks, and free of noise.
8. Test observability: verify logs appear, metrics have values, traces propagate, alerts fire.

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/observability.md, references/standards/logging.md.

## Deliverables

- Structured logging schema with field names and severity levels.
- Metrics definition: rates, errors, durations, saturation, business outcomes.
- Distributed tracing configuration and propagation plan.
- Health check implementation (liveness vs readiness).
- Dashboard layouts and alert rules with runbook links.
- Observability test plan.

## Common Mistakes

- Logging everything at the same level, making it impossible to distinguish critical from noise.
- Creating dashboards that look impressive but cannot be used to diagnose an incident.
- Setting alerts that page without a clear mitigation step.
- Adding tracing to only some services, breaking the end-to-end trace.
- Forgetting to log or metric business outcomes (orders placed, payments completed).

## Failure Modes

- An incident cannot be diagnosed because logs are too verbose or too sparse.
- An alert fires but no one knows what to do because there is no runbook.
- Metrics have high cardinality labels, causing monitoring infrastructure to fail.
- Traces are sampled away for the exact request that failed.

## Checklist

- [ ] Logs are structured with consistent field names and correlation IDs.
- [ ] Sensitive data is excluded from logs, metrics, and traces.
- [ ] Metrics track rates, errors, durations, saturation, and business outcomes.
- [ ] Health checks distinguish liveness from readiness.
- [ ] Alerts are actionable and linked to a runbook.
- [ ] Dashboards support both triage and trend analysis.
- [ ] Tracing covers all services in the request path.
