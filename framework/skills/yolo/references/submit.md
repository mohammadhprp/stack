# Prepare and submit

Start with the verified diff, approved criteria, and latest check results. If the tree changed materially after verification, reopen verification. Follow the target repo's contribution and submission rules. Use existing authorized tracker and forge tooling only if needed; an unavailable tool or ambiguous destination is a blocker, not a reason to guess.

Prepare a concise handoff describing what changed and why, the source task, evidence, and any residual risk. For a PR or MR include the title, draft description, base and head branches, relevant before/after evidence, and any rollback concern. Do not include secrets, private records, or unrelated changes.

Before acting, show the exact operation and destination:

- **Local handoff** (default): leave changes uncommitted for review.
- **Commit:** show the proposed message, branch, and included paths.
- **Push and open or update a PR/MR:** show which commit or changes will be pushed, the remote, head and base, whether a review request will be created or updated, and its draft content.

Ask for explicit approval covering the concrete action. A request during intake is not approval of the prepared submission. Verify scope before staging or pushing; do not include unrelated user edits or generated records. Never implicitly merge, deploy, force-push, delete work, or change a ticket's status. If a submission fails or only partly succeeds, report the actual state and identifiers. Record the result and leave YOLO on for the next task.
