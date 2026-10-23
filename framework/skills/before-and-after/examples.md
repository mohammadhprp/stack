# Before-and-after examples

## Local comparison

**Request:** "Compare these two screenshots."

Reuse the images, verify matching viewport and state, and format them locally.
Return the files or Markdown; do not upload or edit a PR without authorization.

## GitHub pull request

**Request:** "Add desktop before-and-after evidence to PR 123."

Confirm the repository and baseline, inspect images for secrets, format with the
installed `scripts/format.mjs`, and preserve the current PR body around the marked
block. Use native attachments only after checking installed `gh` support. Verify
the rendered PR; inspect partial success before retrying an upload.

## GitLab merge request

**Request:** "Add these images to our GitLab MR."

Follow `references/gitlab.md`: use the confirmed GitLab project and native upload
flow, preserve the current description, and verify the rendered comparison. Do
not use GitHub-only URLs or a public upload service as a fallback.
