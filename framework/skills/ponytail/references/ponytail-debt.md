# Ponytail debt

Use for `/ponytail-debt`, “what did ponytail defer”, or a request for the
shortcut ledger. Read and report only unless the user explicitly asks to write
the ledger to a file.

Search the repository for current and legacy comment markers, skipping `.git`
and paths ignored by `.gitignore` (dependencies, build output, and caches):

```bash
git grep --no-index --exclude-standard -nE '(#|//|/[*]) ?(shortcut|ponytail):'
```

Add comment prefixes for the target language when needed. If the user supplies a
custom marker (for example `/ponytail-debt TODO`), search for that marker instead.
Exit status 1 means no matches, not a failed scan; other errors must be reported.
Skip text that is not a real deferral, such as a note about a keyboard shortcut.

Treat each genuine marker as one row, grouped by file:

`<file>:<line>, <what was simplified>. ceiling: <limit>. upgrade: <trigger>.`

Extract the ceiling and upgrade trigger from the comment. Mark entries with no
upgrade path or trigger as `no-trigger`. End with
`<N> markers, <M> with no trigger.` If there are no hits, say
`No shortcut debt. Clean ledger.` Optional ownership comes from
`git blame -L<line>,<line>`; do not invent owners or upgrade triggers.
