# YOLO examples

## Regression with a clear check

**User:** `/yolo Fix API-123. Empty searches should return an empty list, not an error.`

**Intake checkpoint:** “`AC-1`: empty search returns `[]`; `AC-2` (inferred): nonempty search is unchanged. No project files edited. Intake: `~/.stack/yolo/catalog-api/api-123/01-task.md`. Do you approve these criteria so I can design the solution?”

If the user asks why `AC-2` is inferred, answer and present the checkpoint again. Do not start design until the user approves. Once the solution is approved, show the diff and failing-before/passing-after regression result at the implementation checkpoint. Ask before starting final verification.

## Feature with multiple criteria

**User:** `/yolo Add CSV export for filtered results. Keep the existing order and include a header even when no rows match.`

After the user approves intake and design, implement and present the diff. Ask before final verification. In verification, separately prove the filter, order, and empty-results header. If all are met, ask to prepare submission. Then show the exact operation and destination and wait again before committing or opening a PR.

## Blocked verification and off

An acceptance criterion needs a service that is unavailable. At verification, mark it **Not verified**, state the exact blocked check, and do not offer to submit as complete.

**User:** `/yolo off`

**Assistant:** “YOLO is off. The task is paused at verification. I didn’t change the project, the branch, or remote state.”

If the user turns the mode back on in the same session, show the pending checkpoint. Activation is not approval. On a new session, start off even if the task records are still on disk.
