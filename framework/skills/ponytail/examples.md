# Ponytail examples

## Smallest complete change

**Request:** "Use ponytail full to add a date input."

Inspect the surrounding UI first. Reuse the project's date component if one
exists; otherwise prefer a native input over a new dependency. Update the form's
data flow and relevant checks, not just the visible control. End with what was
skipped, what was not checked, and any material risk.

## Review a change

**Request:** "Ponytail review the last commit."

Read that commit, its connected callers, and tests. Check correctness, security,
expected load, tests, performance, and unnecessary code. Report verified cases
under Must fix, Should fix, and Nice to have. Each numbered finding explains what
the code does, the problem, the smallest fix, and the consequence of leaving it.
Do not apply fixes or publish a remote review. Keep the current coding mode.

## Debt ledger and old markers

**Request:** "List the ponytail shortcuts we deferred."

Search actual source comments for both `shortcut:` and legacy `ponytail:` while
excluding ignored files. A comment such as `// shortcut: global lock, use
per-account locks when throughput matters` becomes a file-and-line ledger row
with its ceiling and upgrade trigger. Tag missing triggers rather than inventing
them. Do not write a ledger file unless asked.

## Skill-only help

**Request:** "Ponytail help. Can I enable the prompt band?"

Show the operation card once without activating or changing mode. Explain that
Stack installs the skill instructions only; the prompt band needs upstream's
separate plugin and is not part of this package. Do not write configuration or
claim that plugin commands are installed.
