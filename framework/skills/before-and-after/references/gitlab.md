# GitLab evidence

Load the installed `glab` skill when available and confirm the GitLab host,
project, and merge request before any remote action. Follow the main skill's
capture, output-path, baseline, privacy, and authorization rules. A local visual
comparison alone never authorizes an upload or description edit.

## Upload and format

Read `glab mr update --help` for the installed version. GitLab's experimental
`--attach` can append uploaded media while keeping the current description:

```sh
glab mr update "$MR" --repo "$REPO" --attach "$OUTPUT_DIR/captures/after.png"
```

Use it only when experimental attachment support is available and accepted. It
does not promise the GitHub formatter's local-reference rewriting behavior.

For a comparison table, use the project's native Markdown upload API after
authorization: `POST /projects/:id/uploads` with multipart field `file`. Use the
returned `markdown` or `url` values to build the table; do not substitute GitHub
attachment URLs or send files to a public third-party host. Check the CLI's API
help for multipart support; otherwise offer the GitLab UI upload flow. Never put
tokens in command arguments, logs, or Markdown.

Read and preserve the current MR description, replace only its
`<!-- before-and-after:start -->` / `<!-- before-and-after:end -->` block, and
write the merged description to `$OUTPUT_DIR/mr-description.md`. Keep before and
after order, captions, and revision facts accurate. Apply the authorized edit:

```sh
glab mr update "$MR" --repo "$REPO" --description-file "$OUTPUT_DIR/mr-description.md"
```

Re-read the MR and check rendered media and unchanged surrounding prose. After
partial failure, inspect what succeeded before retrying. Do not delete uploads
or temporary records without separate authorization.

## Sources

- https://docs.gitlab.com/cli/mr/update/
- https://docs.gitlab.com/api/project_markdown_uploads/
