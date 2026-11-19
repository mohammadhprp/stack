# Verify and review

Identify the exact diff being reviewed, including relevant uncommitted changes. Separate YOLO's changes from unrelated user work. Review on two axes: **task fit** (missing, incorrect, or extra behavior) and **repo standards** (conventions, correctness, and regression risk). Trace beyond the edited files where the change depends on callers, data contracts, or configuration. For material risk, find a test or real usage path that checks the critical assumption, not just a plausible explanation.

For each claim, identify the command or observation that proves it. Run the relevant checks on the final tree, read their output and exit status, and distinguish failed, skipped, and unavailable checks. Run the repository's required checks and exercise the real user-facing path where applicable. Earlier results do not prove a later diff. Do not use a green suite as a substitute for acceptance review.

| Criterion | Verdict | Evidence |
| --- | --- | --- |
| `AC-1`: exact requirement | Met / Not met / Not verified | Test or direct observation and result, or the precise gap |

Use **Met** only for a criterion supported by evidence, **Not met** for a demonstrated gap, and **Not verified** when evidence is missing, blocked, or inconclusive. If no acceptance criterion needs edits but an agreed check is blocked, report the blocker and wait; do not present the task as ready to submit. If fixes or a scope change are needed, ask to reopen the relevant phase. Refresh verification on the resulting diff before delivery.
