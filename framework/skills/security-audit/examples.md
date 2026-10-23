# Security Audit Examples

## Focused guidance

**Request:** "Does this handler let one tenant read another tenant's records?"

Read the handler, identity middleware, and storage query using the access-control
and data-isolation references. Trace the caller, tenant identity, ownership check,
and returned resource. Explain the source evidence and any missing runtime fact.
Do not launch a full audit or create report files for this focused question.

## Full codebase audit

**Request:** "Run a full security audit of this repository."

Resolve the target, source revision, scope, standard profile, and next unused
`~/.stack/<project-name>/security-audit/run-<N>/` directory. Follow all six phases,
track coverage, and independently verify candidates. Write `run-metadata.json`,
`architecture.md`, `coverage-ledger.json`, `findings.json`, `REPORT.md`,
`FINDINGS-DETAIL.md`, and `NEEDS-VALIDATION.md`. Run both bundled validators before
reporting completion. Return absolute artifact paths and remaining coverage gaps;
zero confirmed findings does not mean the target is secure.

## Missing execution or verification capability

**Request:** "Confirm this suspected parser issue and include it in the audit."

Trace the current source first. If the required network isolation, resource
limits, or other execution controls are unavailable, do not run the parser.
An independent verifier may retain a source-grounded `needs_validation` record
with the exact blocker and bounded validation plan; do not assign severity.
If independent verification itself is unavailable, leave the candidate unresolved
in the ledger, mark the run incomplete, and disclose that gap rather than
claiming confirmation or substituting self-review.
