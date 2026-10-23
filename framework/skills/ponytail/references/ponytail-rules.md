# Ponytail rules

Lazy means efficient, not careless. The best code is the code never written.

Before writing code, stop at the first rung that holds:

1. Does this need to be built at all? Skip unrequested flexibility.
2. Does it already exist in this codebase? Reuse it the way the project does.
3. Does the standard library or platform do it? Prefer the project's own component when one exists.
4. Does an already-installed dependency solve it? Use it.
5. Can this be one line a reader understands at a glance? Make it one line.
6. Only then: write the minimum code that fully works.

The ladder runs after understanding the problem, not instead of it. Read the
task and touched code, trace the real flow end to end, then choose the rung.
Include callers, tests, fixtures, configuration, and exports in scope. Preserve
the project's structure and any validation or error handling in code you move.

For bug fixes, find every caller of the function being changed and fix the
shared root cause once. Do not scatter symptom guards across callers.

Do not simplify away input validation at trust boundaries, error handling that
prevents data loss, security, accessibility, hardware calibration, or anything
explicitly requested. Non-trivial lazy code leaves one runnable check behind;
trivial one-liners need no test.

End replies with skipped work, unrun checks, and material risks. Requested reports
still receive their full output, not a coding-response length limit.

Mark deliberate shortcuts with a `shortcut:` comment naming the ceiling and
upgrade path, for example:

```text
# shortcut: global lock, per-account locks if throughput matters
```
