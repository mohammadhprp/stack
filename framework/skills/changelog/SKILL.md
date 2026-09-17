---
name: changelog
description: Create, add, or update CHANGELOG.md entries for repository changes. Use this skill whenever the user asks to maintain a changelog, document release notes, or add an Unreleased entry.
---

# `changelog` skill instructions

Maintain `CHANGELOG.md` entries for the user's changes.

## Process

1. **Review and categorize** - Review conversation history, read the current `CHANGELOG.md`, determine if changes are `Added`, `Changed`, `Fixed`, `Removed`, `Deprecated`, or `Security`, and read the existing `## Unreleased` section.
2. **Group related changes** - Combine related changes into single bullet points. Use past tense ("Added...", "Fixed..."). Include file paths or component names in backticks when helpful. Match existing style and tone.
3. **Add entries** - Insert new bullet points under the correct heading within `## Unreleased`. Create the `## Unreleased` section with relevant headings if it does not exist. Preserve all existing entries.
4. **Verify** - Read the final `CHANGELOG.md` to confirm entries are in the right section, correctly formatted, and no existing entries were altered or removed.

## Entry Format

```markdown
## Unreleased

### Added
- New features, entries, additions.

### Changed
- Changes in existing functionality, refactors, renames.

### Fixed
- Bug fixes, corrections.

### Removed
- Removed features, files, entries.

### Deprecated
- Soon-to-be-removed features.

### Security
- Vulnerabilities, security fixes.
```

Group entries by section. Order sections: Added, Changed, Fixed, Removed, Deprecated, Security. Within each section, entries are reverse-chronological (newest first). Keep descriptions concise but informative; include the file path or component name when it adds clarity.
