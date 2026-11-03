---
name: pull-request
description: Create or update a GitHub pull request (PR) for the current branch using the gh CLI. Use this skill whenever the user asks to open, create, update, or prepare a GitHub PR or pull request.
---

# `pull-request` skill instructions

Create or update a pull request for the current branch using `gh` CLI.
For a request to prepare a PR, produce the title and body without pushing or publishing. Honor existing authorization to create or update; ask only when the requested scope is unclear.

## Process

1. **Collect information**
   - Get the current branch name: `git branch --show-current`
   - Resolve the target from the user's request, the app-managed target when applicable, or the existing PR's base. For a new PR, inspect repository defaults; do not assume `develop` exists.
   - Check for an existing PR with `gh pr view --json number,url,baseRefName,title,body`.
   - Read the PR template from `.github/pull_request_template.md` if it exists.

2. **Format PR title**
   - Take the branch name, replace all `-` with spaces, and capitalize the first character.

3. **Collect commits and build summary**
   - List commits and inspect the full diff against the resolved target: `git log <target>..HEAD --oneline` and `git diff <target>...HEAD`.
   - Read each commit message and convert it to a bullet list summarizing user-facing changes.
   - Group related commits in the summary (for example, multiple commits for the same change); do not rewrite Git history.
   - Keep the summary concise, with one bullet per logical change.

4. **Fill template**
   - Set Summary to the bullet list from step 3.
   - Keep the Checklist section as-is.

5. **Present the result** - Show the source branch, resolved target, title, and filled description. Stop here for preparation-only requests. Push the branch only within the user's authorized publishing scope.

6. **Create or update when authorized** - Write the exact description to a temporary file. Use `gh pr create --title "<title>" --body-file <file> --base <target>` for a new PR, or `gh pr edit <number> --title "<title>" --body-file <file>` for the existing PR. Preserve unrelated existing description content and show the resulting URL.
