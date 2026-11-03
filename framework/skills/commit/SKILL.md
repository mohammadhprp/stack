---
name: commit
description: Create atomic Git commits with conventional messages. Use this skill whenever the user asks to commit changes, create commits, or prepare commit messages.
---

# `commit` skill instructions

Create Git commits only when the user requests commits. For a request to prepare commit messages, return the messages and stop before staging or committing.

## Process

1. **Analyze and plan** - Review conversation history, run `git status -s`, `git diff`, and `git diff --cached`, determine if changes should be one or multiple logical commits, group related files, and draft conventional commit messages (`type: description`) in imperative mood focusing on why.
2. **Present plan** - List files for each commit, show commit messages with type prefix, and proceed when the user already authorized those commits. Ask for confirmation only if the requested commit scope remains unclear.
3. **Execute authorized commits** - Use `git add` with specific files (never `-A` or `.`), check the staged diff contains only the authorized changes (preserve unrelated staging), create commits with planned messages, and show the result with `git log --oneline -n [N]`.

## Commit Message Format

Use conventional commit format: `type: description`

**Types:**
- `feat:` - New feature (user-facing)
- `fix:` - Bug fix (user-facing)
- `docs:` - Documentation only
- `chore:` - Maintenance, tooling, dependencies
- `refactor:` - Code restructuring without behavior change
- `test:` - Adding or updating tests
- `perf:` - Performance improvement
- `ci:` - CI/CD changes
