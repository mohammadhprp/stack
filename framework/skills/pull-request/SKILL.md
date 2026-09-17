---
name: pull-request
description: Create or update a GitHub pull request (PR) for the current branch using the gh CLI. Use this skill whenever the user asks to open, create, update, or prepare a GitHub PR or pull request.
---

# `pull-request` skill instructions

Create or update a pull request for the current branch using `gh` CLI.

## Process

1. **Collect information**
   - Get the current branch name: `git branch --show-current`
   - Read the PR template from `.github/pull_request_template.md` if it exists.

2. **Format PR title**
   - Take the branch name, replace all `-` with spaces, and capitalize the first character.

3. **Collect commits and build summary**
   - List commits on the branch that are not on `develop`: `git log develop..HEAD --oneline`
   - Read each commit message and convert it to a bullet list summarizing user-facing changes.
   - Merge/squash related commits (for example, multiple commits for the same change).
   - Keep the summary concise, with one bullet per logical change.

4. **Fill template**
   - Set Summary to the bullet list from step 3.
   - Keep the Checklist section as-is.

5. **Present plan and confirm** - Show the source branch, target branch (`develop`), title, and filled description. Ask: "Shall I create this PR?" Push the changes if the user says yes.

6. **Create upon confirmation** - Use `gh pr create --title "<title>" --body "<body>" --base <target>` and show the resulting URL.
