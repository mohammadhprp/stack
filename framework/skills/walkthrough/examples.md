# Walkthrough examples

These excerpts illustrate grouping and attention labels. Actual output includes
the complete report structure in both Markdown and HTML, without emoji. Write
`walkthrough.json` under `~/.stack/<project-name>/walkthrough/`, then validate and
render it with the bundled scripts. See `references/schema.md` for a complete
input example and the shared presentation standard.

## 1. Feature spanning implementation, configuration, tests, and docs

Request: "Walk me through PR 42. It adds configurable session expiration."

Group the configuration contract, session logic, regression tests, and deployment
documentation under `Feature — Configurable session expiration`, rather than
creating four unrelated file sections.

Reading order:
1. `config/session.go`, expiration setting — defines the supported duration.
2. `internal/session/store.go`, expiry calculation — consumes that setting.
3. `tests/session_test.go`, expiration cases — describes expected boundaries.
4. `docs/deployment.md`, session settings — explains operator configuration.

Notice: The diff introduces a new expiration setting. Compare its default with
the existing deployment configuration before rollout.

Human decision: Confirm whether the configured expiration matches the product's
session policy. The code and tests alone cannot establish that policy.

## 2. MR containing an unrelated fix and refactor

Request: "Explain MR 18 for me; the payment handler has several changes."

Create separate groups for `Bug fix — Reject duplicate payment requests` and
`Refactor — Extract request parsing`, even if both modify the same handler.
Reference each group's verified symbols or line ranges separately. Put the
duplicate-request regression test in the bug-fix group.

Warning: The new regression test covers sequential retries, but does not exercise
concurrent requests. This is a verification gap, not proof of a race. Ask the
reviewer to check whether the persistence constraint also prevents concurrent
duplicates.

For the parsing refactor, explain the structural change and the evidence for
unchanged behavior; do not assume all edits in a refactor are behavior-preserving.

## 3. Local changes with incomplete verification

Request: "Give me a walkthrough of my staged and unstaged changes."

State that the source is the local working tree, not a published PR. Group a
dependency update with its lockfile under `Maintenance — Update HTTP client`;
place an unrelated README correction under `Documentation — Correct setup steps`.
Mention relevant untracked text files if included in the inspected scope.

Verification: Existing HTTP-client tests were inspected but not executed. No CI
results are available for the working tree.

Limitations: A changed binary asset cannot be explained from a text diff; record
its path and ask the reviewer to inspect it with an appropriate viewer.

Do not manufacture warnings or an approval verdict when no concrete issue was
identified. Deliver all groups together in `walkthrough.md` and `walkthrough.html`
under `~/.stack/<project-name>/walkthrough/`.
