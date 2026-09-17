# Performance Reference

Diagnose and improve CPU, memory, network, storage, caching, and query behavior using evidence.

## When to Use

Use this when the task involves CPU, memory, network, caching, query optimization, parallelism, profiling, benchmarking, latency, throughput, or resource cost. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- Measure before optimizing and define the target metric: latency, throughput, cost, or capacity.
- CPU work should be bounded and avoid repeated expensive computation.
- Memory use should avoid unbounded aggregation, buffering, and cache growth.
- Network calls need timeouts, batching where appropriate, and payload discipline.
- Caching requires invalidation, freshness expectations, and failure behavior.
- Parallelism must preserve correctness and respect downstream limits.
- Benchmarks should reflect realistic data sizes and access patterns.

## Workflow

1. Define the performance target: latency p50/p95/p99, throughput RPS, resource budget.
2. Measure the current state before proposing any change. Profile CPU, memory, I/O, network.
3. Identify the bottleneck: slowest path, most allocated memory, most frequent call.
4. Formulate a targeted optimization based on evidence, not intuition.
5. Implement the smallest change that addresses the bottleneck.
6. Re-measure to confirm improvement and verify no regressions in other paths.
7. Add performance regression tests for critical paths.
8. Document the optimization: baseline, result, tradeoffs (complexity, readability, maintainability).

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/performance.md.

## Deliverables

- Baseline measurement and profiling results.
- Bottleneck identification with evidence.
- Targeted optimization with before/after comparison.
- Performance regression tests for critical paths.
- Tradeoff documentation: complexity, readability, maintainability.

## Common Mistakes

- Optimizing without measuring first, leading to wasted effort on non-bottlenecks.
- Adding caching without defining invalidation strategy, freshness requirements, or failure behavior.
- Premature optimization that adds complexity for unproven gains.
- Ignoring the performance impact of N+1 queries, unbounded data loading, and missing indexes.
- Optimizing for throughput at the expense of p99 latency or vice versa without understanding the requirement.

## Failure Modes

- An optimization improves p50 but makes p99 worse (tail at scale).
- A cache introduces stale data because invalidation was not designed.
- A query optimization breaks on a different database version or data distribution.
- A performance improvement reduces readability without sufficient documentation.

## Checklist

- [ ] A baseline measurement exists before the change.
- [ ] The bottleneck is identified through profiling, not guessing.
- [ ] Caching has defined invalidation, freshness, and failure behavior.
- [ ] N+1 queries and unbounded data loading are resolved.
- [ ] Performance regression tests cover critical paths.
- [ ] Tradeoffs against readability, complexity, and maintainability are documented.
- [ ] P50, P95, P99 latency targets are defined and measured.
