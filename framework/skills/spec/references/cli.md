# CLI reference

These command forms were checked against installed CLI help. Recheck help at runtime because versions and Notion API schemas change. Placeholders such as `ISSUE-KEY`, `DATA_SOURCE_ID`, and `PAGE_ID` must be replaced with verified values. Use JSON serialization and body files rather than interpolating issue text into shell commands.

## Jira: ankitpokhrel/jira-cli

```sh
jira version
jira me
jira issue list --help
jira issue view --help
jira sprint list --help
jira board list --help
```

The CLI defaults to a configured project. Its documented cross-project query uses `project IS NOT EMPTY`; include that explicit project clause to prevent implicit default-project scoping:

```sh
jira issue list --jql 'project IS NOT EMPTY AND assignee = currentUser() AND sprint in openSprints()' --raw --paginate 0:100
jira issue list --jql 'project IS NOT EMPTY AND assignee = currentUser() AND sprint in openSprints()' --raw --paginate 100:100
jira issue view ISSUE-KEY --raw
jira issue view ISSUE-KEY --plain --comments 20
jira sprint list --state active --table --plain --columns id,name,start,end,state
```

Continue issue offsets in batches of at most 100 until the response metadata or an empty page establishes completion; do not stop after the first 100 issues. Inspect the actual JSON shape rather than assuming a bare array. Use a stable ordering when supported, deduplicate keys, and report if the source changed during retrieval.

`jira sprint list` help describes a top-50 sprint list; do not assume it enumerates every board or sprint, or that its issue-pagination flags paginate sprint metadata. Resolve sprint identity from issue fields and the relevant configured boards. If metadata remains unavailable, mark dates/board identity unknown and avoid exact scheduling or inventing a plan identity.

Use `--project PROJECT_KEY` for project-specific reads and `--config CONFIG_PATH` for an existing confirmed site configuration. Verify board-selection capabilities from help/config; do not invent a `--board` flag. Read only the needed non-secret configuration fields; do not dump configuration files or debug output that might contain credentials.

Construct browser URLs from the verified Jira site plus `/browse/ISSUE-KEY`, respecting any configured site context path. The issue's REST `self` URL is not a browser link. Raw fields supply status categories, relationships, and custom field data; plain views can help inspect comments. Increasing `--comments` is not proof that every relevant comment was retrieved: check totals and disclose CLI limits.

Do not use issue edit/create/move/delete/assign, sprint membership mutations, or any other Jira write command in this workflow.

## Notion: ntn

```sh
ntn --version
ntn whoami
ntn api --help
ntn api ls
ntn pages get --help
ntn pages create --help
ntn pages edit --help
```

Inspect endpoint help/specs before preparing bodies:

```sh
ntn api v1/databases/DATABASE_ID --help
ntn api v1/data_sources/DATA_SOURCE_ID --spec
ntn api v1/data_sources/DATA_SOURCE_ID/query --spec
ntn api v1/pages/PAGE_ID --spec
ntn api v1/blocks/BLOCK_ID/children --spec
ntn api v1/blocks/BLOCK_ID --spec
```

Read schema, properties, and page content:

```sh
ntn api v1/databases/DATABASE_ID
ntn api v1/data_sources/DATA_SOURCE_ID
ntn api v1/pages/PAGE_ID
ntn pages get PAGE_ID --json
ntn api v1/blocks/PAGE_ID/children page_size==100
```

Paginate every list/query with `has_more` and `next_cursor`, using `start_cursor` in the appropriate query parameters or JSON body. Recursively read child blocks when inspecting managed subtrees or legacy identities. Never send a `collection://` URI as an API ID.

For exact Jira URL lookup, serialize this body to a temporary file and invoke `ntn api v1/data_sources/DATA_SOURCE_ID/query -X POST --data @query.json`:

```json
{
  "filter": {
    "property": "Link",
    "url": {"equals": "https://example.com/browse/PROJ-123"}
  },
  "page_size": 100
}
```

Use this filter only after verifying `Link` is a URL property; adapt to the actual schema otherwise. Search does not replace a database property query for exact identity matching.

For task creation, inspect the Pages API schema and prepare a body with a `data_source_id` parent, verified typed properties, and the managed block subtree. Create through `ntn api v1/pages -X POST --data @create.json` so properties and identity content are included together where supported. For updates:

```sh
ntn api v1/pages/PAGE_ID -X PATCH --data @properties.json
ntn api v1/blocks/BLOCK_ID -X PATCH --data @block.json
ntn api v1/blocks/PARENT_BLOCK_ID/children -X PATCH --data @children.json
```

Respect current request-size and block limits from endpoint specs. When replacing a managed subtree, prefer updating existing blocks; if removal is needed, use only verified skill-owned block IDs and the endpoint's documented archive/delete operation. Never remove the whole page or user-owned blocks. After a partial write, read current state before retrying.

For new sprint-plan pages beneath a confirmed parent, Markdown creation is available:

```sh
ntn pages create --parent page:PARENT_PAGE_ID --content '# Sprint plan' --json
```

Add and track the managed subtree using the block API. `ntn pages get` includes property frontmatter by default, but `ntn pages create`/`edit` ignore most frontmatter properties. Set task properties through `ntn api`, not Markdown frontmatter.

**Preservation warning:** `ntn pages edit` (also exposed as `update` in some versions) replaces page content. Do not use it to refresh managed sections on existing pages, and do not use `--allow-deleting-content`. Truncated Markdown or `unknown_block_ids` require block-level inspection, not whole-page replacement.

Use existing authentication; `NOTION_API_TOKEN` overrides stored credentials. Never put tokens in body files, tracked configuration, or logs. Keep temporary payloads private and remove only files created by the current run when finished. On rate limits, respect server retry guidance; on permission/schema failures, stop the affected write and report it rather than retrying blindly.
