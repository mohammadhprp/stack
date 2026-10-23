# Changelog

Notable changes to `stack` are recorded here.

## [Unreleased]

## [v0.1.6] - 2026-10-23

### Added

- Added GitLab CLI (`glab`) support alongside GitHub CLI (`gh`) in the `before-and-after` and `new-feature` skills.
- Added `security-audit` with source-first investigation, independent validation, structured findings, and coverage tracking.
- Added Diagram Design 2.6 resources for 45 visual types, semantic patterns, named profiles, diagram imports, SVG export, optional PNG export, and offline validation.

### Changed

- Refreshed Ponytail and 13 additional skills from pinned upstream revisions, preserving Stack's installed-relative paths, local additions, and safety rules.
- Updated Humanizer to 3.1.0 and expanded documentation, frontend design, architecture, and investigation guidance.
- Migrated `before-and-after` to agent-browser capture, local media formatting, and authorized native GitHub/GitLab attachments.
- Switched Sentry's default workflow to the read-only `sentry` CLI while retaining the legacy API fallback.

### Fixed

- Preserved surrounding PR description bytes when replacing a before-and-after media block.
- Replaced hard-coded model aliases and harness-specific delegation assumptions with host-native capabilities and verified model selections.

### Removed

- Removed `security-best-practices` in favor of `security-audit`.
- Removed the obsolete public third-party screenshot upload adapters.

### Licensing

- Retained upstream license notices. The refreshed `before-and-after` component uses PolyForm Shield 1.0.0, including a non-compete restriction; Stack's MIT license does not replace these terms.

## [v0.1.5] - 2026-10-19

### Added

- Added `spec` for planning assigned Jira sprint tasks and synchronizing briefs to Notion, and `implement` for building linked Notion tasks on dedicated Git branches with verification.

### Removed

- Removed the `taste`, `tlc-spec-driven`, `improve`, and `effective-html` skills from the embedded catalog.

## [v0.1.4] - 2026-10-13

### Added

- Add a walkthrough skill that generates grouped Markdown and HTML reports for PRs, MRs, and local changes.
