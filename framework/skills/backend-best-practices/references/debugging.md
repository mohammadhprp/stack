# Debugging Reference

Find root causes through reproduction, isolation, logs, metrics, traces, and controlled experiments.

## When to Use

Use this when the task involves root cause analysis, log analysis, reproduction, isolation, binary search debugging, monitoring, metrics, incidents, or unexplained production behavior. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- Reproduce before changing code when possible; if not possible, collect enough production evidence to form a falsifiable hypothesis.
- Isolate variables using binary search across versions, inputs, dependencies, and configuration.
- Logs, metrics, and traces should be correlated by request, job, user, order, or payment identifier.
- Fix root causes, not only symptoms, and add regression tests or alerts for recurrence.
- Preserve evidence during incidents; do not destroy state needed for diagnosis.

## Workflow

1. Reproduce the issue in a controlled environment before making any changes.
2. Gather evidence: logs, metrics, traces, stack traces, error rates, request samples.
3. Isolate variables: binary search through commits, configs, inputs, or environments.
4. Formulate a hypothesis that explains all observed symptoms, not just some.
5. Test the hypothesis with the smallest possible experiment.
6. Identify the root cause: code bug, config error, race condition, data corruption, dependency failure.
7. Fix the root cause, not the symptom. Add tests that would have caught it.
8. Add monitoring or alerting to detect recurrence.
9. Document the incident timeline, root cause, fix, and detection gap.

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/logging.md, references/standards/observability.md.

## Deliverables

- Reproduced issue and diagnostic evidence.
- Root cause analysis with supporting evidence.
- Fix with regression test and detection gap analysis.
- Monitoring or alerting additions for recurrence detection.
- Incident timeline and documentation.

## Common Mistakes

- Changing code before reproducing the issue.
- Fixing symptoms instead of finding and eliminating the root cause.
- Stopping at the first plausible cause without ruling out alternatives.
- Ignoring evidence that does not fit the current hypothesis.
- Failing to preserve diagnostic evidence before restarting or rolling back.

## Failure Modes

- The fix addresses the symptom but the root cause remains latent.
- Diagnostic evidence is lost because logs rotated or services restarted.
- A hypothesis is confirmed without disproving alternative explanations.
- The debugging process disturbs production traffic or state.

## Checklist

- [ ] The issue is reproduced in a non-production environment.
- [ ] Logs, metrics, and traces are correlated by request ID or trace ID.
- [ ] Variables are isolated one at a time (binary search approach).
- [ ] The root cause is identified and distinguished from symptoms.
- [ ] A regression test covers the fix and the detection gap.
- [ ] Monitoring or alerting can detect recurrence.
