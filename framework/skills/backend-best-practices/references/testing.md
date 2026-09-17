# Testing Reference

Design deterministic unit, integration, contract, and end-to-end testing strategies.

## When to Use

Use this when the task involves unit, integration, contract, end-to-end, testability, mocking, deterministic behavior, or release confidence. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- Unit tests prove local rules and edge cases quickly.
- Integration tests prove persistence, transactions, serialization, and external boundaries.
- Contract tests protect clients and providers from incompatible changes.
- End-to-end tests should cover critical user journeys, not every branch.
- Deterministic tests avoid time, order, randomness, and shared-state flakiness.
- Mocks should model behavior and failure, not implementation trivia.

## Workflow

1. Identify what is being tested: a unit of logic, an integration boundary, a contract, or an end-to-end journey.
2. Choose the narrowest test type that provides sufficient confidence for the risk level.
3. Write the test starting with the expected behavior: given X, when Y, then Z.
4. Cover normal paths, edge cases (empty, null, boundary values), and failure paths (timeouts, errors, invalid input).
5. Keep tests deterministic: no dependence on time, random values, shared state, or external service availability.
6. Use mocks at integration boundaries to model real behavior and realistic failures, not implementation details.
7. Run the full test suite locally before pushing.
8. Review test coverage: are there untested branches, error paths, or concurrent scenarios?

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/testing.md.

## Deliverables

- Test plan covering normal paths, edge cases, and failure paths.
- Deterministic tests with given/when/then structure.
- Mock strategy modeling boundary behavior.
- Contract tests for public API boundaries.
- Coverage analysis of untested branches and error paths.

## Common Mistakes

- Writing tests that mirror implementation details, breaking when the code is refactored.
- Testing only the happy path and ignoring error handling and edge cases.
- Using mocks that model framework internals instead of boundary behavior.
- Creating flaky tests that depend on timing, ordering, or shared state.
- Writing overly broad integration tests that are slow and brittle.

## Failure Modes

- A test passes in CI but fails locally (or vice versa) because of environment differences.
- A refactor breaks many tests because they were coupled to implementation instead of behavior.
- An integration test becomes a maintenance burden because it exercises too many paths at once.
- False-positive tests give false confidence while real bugs exist in untested paths.

## Checklist

- [ ] Tests cover normal paths, edge cases, and failure paths.
- [ ] Tests are deterministic: no flakiness from time, randomness, or shared state.
- [ ] Mocks model boundary behavior, not implementation internals.
- [ ] The test suite runs in CI and is fast enough for quick feedback.
- [ ] Coverage includes error handling and concurrent scenarios where applicable.
- [ ] Tests are readable: given/when/then structure with descriptive names.
- [ ] Contract tests exist for public API boundaries.
