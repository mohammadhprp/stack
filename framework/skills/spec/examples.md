# Examples

## Plan assigned sprint work

**User:** “Spec my active sprint tasks and copy them to Notion.”

Read the configured environment and CLI help, retrieve all assigned active-sprint issues, and verify Notion mappings. Order ready tasks by dependencies and priority. Create task briefs and a sprint-plan page, then read them back. Report verified task counts and links, and confirm Jira was unchanged.

## Missing scheduling information

**User:** “Plan my sprint; I haven't set working days or capacity yet.”

Produce an ordered backlog without inventing hours or dates. Synchronize tasks with verified destination mappings, preserving existing dates. Ask one consolidated question if exact scheduling is needed. Report projects without mappings separately rather than borrowing another project's settings.

## Repeat a sync after a manual edit

**User:** “Sync my sprint again; I changed one Notion task's date.”

Find existing tasks by canonical Jira URL and compare current values with the last-written baseline. Preserve the edited date and report its conflict. Update only safe changed fields and managed blocks; leave notes untouched and avoid duplicate tasks or appended briefs. Count the conflicted task as blocked, not unchanged.
