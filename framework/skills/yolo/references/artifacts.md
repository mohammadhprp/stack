# Task records

1. Identify the target project's Git root and use its basename as the project slug. In a non-Git project use the confirmed project directory's basename. Ask if the target project is unclear.
2. Use the canonical issue key, or a short task title if there is no key, for the task slug. For both slugs, lowercase, replace runs of characters outside ASCII `a-z`, `0-9` with one hyphen, trim hyphens, and keep at most 64 characters without a trailing hyphen. Ask for a safe label if the result is empty. Never use raw URLs, absolute paths, or untrusted path components as slugs.
3. Expand `~` to the user's home directory. Require the records to stay beneath the resolved `~/.stack/yolo/` root. Quote paths. Refuse symlinked output components or files and paths that escape the root. Create the directory only when you have an intake record to write; turning on the mode alone creates nothing.
4. Record the canonical project path and task source or title in `01-task.md`. If the slug directory already exists, never assume it belongs to this run. Reuse it only when the user explicitly resumes the same task record and its project and task identities match. Otherwise choose a fresh task slug with a numeric suffix. Check destinations before writing; preserve all existing files and use numbered revisions rather than overwriting checkpoint records.

| Artifact | Contents |
| --- | --- |
| `01-task.md` | Source, project identity, scope, constraints, exact and inferred AC. |
| `02-plan.md` | Approved approach candidate, affected areas, test map, risks and baseline where relevant. |
| `03-implementation.md` | Changed paths, deviations, and focused checks. |
| `04-verification.md` | Reviewed revision or diff, exact checks and results, criterion-by-criterion evidence, omissions. |
| `05-submission.md` | Draft action and destination, scope, and copy-ready handoff. |
| `06-receipt.md` | What actually happened, final local or remote reference, and anything still open. |

Keep optional evidence in the same task directory. Save useful excerpts rather than dumping private tracker text, environment variables, credentials, or raw logs. Record approvals as history, never as authorization for a future session. Show absolute artifact paths at checkpoints. Do not commit or upload these records by default.
