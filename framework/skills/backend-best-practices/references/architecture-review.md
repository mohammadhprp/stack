# Architecture Review Reference

Evaluate system boundaries, coupling, cohesion, complexity, reliability, and scalability.

## When to Use

Use this when the task involves reviewing coupling, cohesion, layering, boundaries, complexity, maintainability, reliability, or scalability. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- High cohesion means related behavior changes together; low coupling means unrelated areas can change independently.
- Boundaries should align to business capabilities, data ownership, and operational responsibility.
- Layering must prevent policy decisions from leaking into transport, storage, or presentation glue.
- Reliability design must include timeout, retry, fallback, and partial-failure behavior.
- Scalability claims require bottleneck analysis, not generic distribution.

## Workflow

1. Understand the architecture's scope, boundaries, and stated goals.
2. Map the system: components, data flows, ownership boundaries, external dependencies.
3. Evaluate coupling and cohesion: can components change independently?
4. Check layering: does infrastructure leak into domain or policy into detail?
5. Assess reliability: timeouts, retries, circuit breakers, fallbacks, bulkheads.
6. Assess scalability: bottleneck analysis, load distribution, data partitioning.
7. Identify single points of failure and missing failure modes.
8. Compare against simpler alternatives that satisfy the same requirements.
9. Summarize key risks, tradeoffs, and recommended follow-ups.

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/architecture.md.

## Deliverables

- Architecture diagram or component map (text description).
- Coupling and cohesion analysis per boundary.
- Reliability and scalability assessment.
- Concrete risks with recommended mitigations.
- Simpler alternatives analysis.

## Common Mistakes

- Reviewing implementation details instead of architectural boundaries.
- Accepting complex designs without requiring evidence that simpler alternatives fail.
- Ignoring operational concerns like deployment, rollback, and observability.
- Treating architecture as fixed rather than evaluating changeability.
- Overlooking data ownership and consistency boundaries.

## Failure Modes

- The review focuses on code style and misses structural coupling issues.
- A component becomes a deployment bottleneck because ownership boundaries are unclear.
- The architecture passes review but cannot be operated because observability was not considered.
- Scaling requires a rewrite because data partitioning was not considered early.

## Checklist

- [ ] Component boundaries align with business capabilities.
- [ ] Dependency direction follows the stable-dependency principle.
- [ ] Each component has a clear data ownership boundary.
- [ ] Failure isolation (bulkheads, circuit breakers) is present where needed.
- [ ] The architecture supports incremental deployment and rollback.
- [ ] Observability is designed in, not added after.
- [ ] A simpler architecture was evaluated and ruled out with evidence.
