# Refactoring Reference

Improve code structure safely while preserving behavior through small verified steps.

## When to Use

Use this when the task involves removing duplication, simplifying code, reducing complexity, improving naming, preserving behavior, or making small incremental changes. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- Preserve externally visible behavior unless a behavior change is explicitly requested.
- Make one conceptual change at a time.
- Remove duplication by extracting the stable idea, not by forcing unrelated code through one abstraction.
- Improve names when they reduce cognitive load for future changes.
- Characterization tests are useful before touching risky legacy behavior.
- Stop refactoring when the code is simpler enough for the current goal.

## Workflow

1. Identify the improvement goal: reduce duplication, simplify logic, improve naming, extract module.
2. Verify existing test coverage. If coverage is weak, add characterization tests first.
3. Make one conceptual change at a time. Commit or stage after each safe step.
4. Preserve external behavior: same inputs produce same outputs, same errors, same side effects.
5. Run the full test suite after each change step.
6. Compare before and after: verify the improvement did not introduce subtle behavioral differences.
7. Remove dead code, commented code, and unused dependencies found during the refactor.
8. Stop when the code is clear enough for the current maintenance needs - avoid over-refactoring.

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/naming.md, references/standards/testing.md.

## Deliverables

- Improved code with one conceptual change per step.
- Passing test suite before and after each step.
- Characterization tests for legacy code with weak coverage.
- Before/after comparison confirming behavior preservation.
- Removed dead code, commented code, and unused dependencies.

## Common Mistakes

- Refactoring and changing behavior in the same step.
- Removing tests or code that seems unused but is actually relied upon implicitly.
- Over-abstracting: creating interfaces, base classes, or indirection before a clear need exists.
- Refactoring for performance without measuring the baseline.
- Leaving the code in a half-refactored state.

## Failure Modes

- A refactor introduces a subtle behavior change that tests do not catch.
- Over-refactoring produces a more complex design than the original.
- A rename misses a reference, causing a production issue.
- A shared abstraction created during refactoring becomes a coupling point that is hard to undo.

## Checklist

- [ ] Existing tests pass before and after each refactoring step.
- [ ] One conceptual change is made at a time.
- [ ] External behavior is preserved (same inputs, outputs, errors, side effects).
- [ ] Dead code, commented code, and unused dependencies are removed.
- [ ] Naming reflects business meaning after the refactor.
- [ ] Refactored code is simpler, not more complex.
- [ ] Characterization tests exist for legacy code with weak coverage.
