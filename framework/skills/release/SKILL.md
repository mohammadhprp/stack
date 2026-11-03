---
name: release
description: Create a software release by reviewing commits, determining the semantic version, updating the changelog and package versions, committing release metadata, and tagging the result. Use this skill whenever the user asks to cut, prepare, publish, or tag a release.
---

# `release` skill instructions

Create a release from the current branch state.
For preparation-only requests, produce the version and changelog plan and stop before committing, tagging, pushing, or publishing. A local release does not authorize a remote push or publication. Preserve unrelated working-tree and staged changes.

## Process

1. **Review recent commits** - Run `git fetch --tags`, `git describe --tags --abbrev=0` to find the latest tag. Read full commit messages with `git log <latest-tag>..HEAD --format="%B%x00"` so breaking-change footers are included. If no previous tag exists, inspect all commits through `HEAD`. Read any existing `CHANGELOG.md` and the project's release conventions.
2. **Categorize commits** - Group commits by conventional commit type:
   - Any conventional type with an optional scope and `!` (such as `fix(api)!:`), or a `BREAKING CHANGE:` / `BREAKING-CHANGE:` footer → breaking change
   - `feat:` → minor feature
   - `fix:` → patch fix
   - `perf:`, `refactor:`, `test:` → patch (if no features)
   - `chore:`, `docs:`, `ci:` → filtered from changelog
3. **Determine next version** - Follow the project's release policy; otherwise use this semantic-version convention:
   - Breaking changes → increment major version (e.g., `1.2.3` → `2.0.0`)
   - New features → increment minor version (e.g., `1.2.3` → `1.3.0`)
   - Only fixes/refactors → increment patch version (e.g., `1.2.3` → `1.2.4`)
   - If no previous tag exists, propose `0.1.0`
4. **Present release plan** - Show the current version, new version, categorized changelog entries, and exact files and tag to change. Stop for preparation-only requests; proceed within existing authorization or clarify any unresolved scope.
5. **Execute the authorized release**:
   - Update `CHANGELOG.md`: create a new `## [vX.Y.Z]` section under `## Unreleased`, move categorized entries (excluding chore/docs/ci) into it, add the release date, and keep the `## Unreleased` section empty for future work.
   - Update project version metadata such as `package.json` and `package-lock.json` when present.
   - Stage only the changed release files, inspect the staged diff, and commit with `chore: release vX.Y.Z`. Do not include unrelated pre-staged changes; resolve that scope before committing.
   - Run `git tag -a vX.Y.Z -m "vX.Y.Z"`.
6. **Verify** - Run `git log --oneline -n 3` and `git tag --list --sort=-v:refname -n5` to confirm the release tag and commit are in place.
