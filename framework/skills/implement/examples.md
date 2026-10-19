# Examples

## Implement a prepared task

**User:** “Implement this Notion task: <task-page-url>.”

Read the task properties, nested Spec brief, and required linked Jira details. Confirm the repository and a clean working tree. State the scope and verification plan, create a task branch following repository conventions, and make focused changes. Run checks and report the outcome, source URL, branch, and base. Leave service records unchanged and code uncommitted.

## Preserve an existing working tree

**User:** “Start the linked Notion login-validation task.”

Read the requirements and inspect the repository. If unrelated edits exist, stop before switching branches and ask whether to continue with them or use an isolated worktree. Do not stash, reset, or silently carry the edits to a new branch.

## Resume an existing branch

**User:** “Continue the Notion task on its existing branch.”

Verify the task identity, repository, branch contents, and current working tree before resuming. Do not recreate or reset the branch. Resolve material requirement conflicts before editing, then implement only the remaining scope and report checks actually run.
