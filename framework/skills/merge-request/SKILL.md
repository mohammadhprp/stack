---
name: merge-request
description: Create a GitLab merge request (MR) for the current branch using the glab CLI. Use this skill whenever the user asks to open, create, update, or prepare a GitLab MR or merge request.
---

# `merge-request` skill instructions

Create or update a merge request for the current branch. A request to prepare an MR produces a draft description only; it does not authorize pushing or publishing.

## Process

1. **Collect information**
   - Get the current branch name: `git branch --show-current`
   - Read the MR template from `.gitlab/merge_request_templates/default.md` if it exists.
   - Check for an existing open MR for this branch. For updates, read its target, description, and checked items; preserve unrelated content rather than creating a duplicate.
   - Resolve the target from the user's request or the host's documented target metadata, then repository guidance or the GitLab project's default branch. Do not assume `develop`; ask if the target remains ambiguous. Use a verified local or remote ref for it as `TARGET_REF` below.

2. **Format MR title**
   - Take the branch name, replace all `-` with spaces, and capitalize the first character.

3. **Collect commits and build summary**
   - List commits on the branch that are not on the resolved target: `git log "${TARGET_REF}..HEAD" --oneline`.
   - Read `git diff "${TARGET_REF}...HEAD"` to verify what the commits actually change. Do not include uncommitted changes in a summary of the published branch.
   - Read each commit message and convert it to a bullet list summarizing user-facing changes.
   - Merge/squash related commits (for example, multiple commits for the same change).
   - Keep the summary concise, with one bullet per logical change.

4. **Fill template**
   - Set Summary to the bullet list from step 3.
   - Keep the Checklist section as-is.

5. **Present the result** - Show the source branch, resolved target, title, and filled description. Stop here for preparation requests. If creation/update and any needed push are already authorized, proceed; otherwise ask for the missing authorization. Do not commit working-tree changes automatically.

6. **Publish when authorized** - Check the installed `glab mr create --help` or `glab mr update --help` for supported flags. Update an existing MR by its ID with `glab mr update`; otherwise use `glab mr create` with:
   - `--source-branch`: Current branch
   - `--target-branch`: Resolved target branch
   - `--title`: Prepend `Draft: ` to the formatted branch name
   - `--description-file`: A file containing the exact filled description, preserving newlines and literal shell characters
   - Set assignees only when requested, using usernames. Keep project defaults for squash and source-branch removal unless the user or repository specifies otherwise.

7. **Show the resulting URL.**
