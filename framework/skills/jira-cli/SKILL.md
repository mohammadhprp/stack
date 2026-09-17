---
name: jira-cli
description: Expert guidance for using the Jira CLI (jira) to manage Jira issues, sprints, epics, comments, transitions, and worklogs from the command line. Use this skill when the user needs to interact with Jira resources or perform Jira workflows — listing, creating, editing, or transitioning issues, running JQL queries, or scripting against Jira data — even if they don't explicitly mention the jira CLI.
allowed-tools: Bash, Read, Grep, Glob
---

# Jira CLI (jira) Skill

Provides guidance for using `jira`, the feature-rich interactive Jira command line, to perform Jira operations from the terminal.

## When to Use This Skill

Invoke when the user needs to:
- Search, list, or view Jira issues
- Create, edit, assign, clone, or transition issues
- Add comments, links, or worklogs to issues
- Work with epics, sprints, releases, or project boards
- Run JQL queries or script against Jira data

## Prerequisites

Verify jira installation before executing commands:
```bash
jira --version
```

If not installed, inform the user and provide platform-specific installation guidance.

## Setup Quick Start

`jira` requires one-time initialization before first use:

```bash
# Interactive setup wizard: installation type (Cloud/Local), auth type, project, board
jira init
```

Authentication via environment variables:

```bash
# Jira Cloud: API token from https://id.atlassian.com/manage-profile/security/api-tokens
export JIRA_API_TOKEN=your-token

# On-premises with a Personal Access Token
export JIRA_API_TOKEN=your-token
export JIRA_AUTH_TYPE=bearer
```

- Use `jira issue list -c /path/to/config.yaml` or `export JIRA_CONFIG_FILE=/path/to/config.yaml` to work with multiple projects or instances
- Check authentication and current user: `jira me`

## Core Workflows

### Searching and Viewing Issues

```bash
# 1. List issues assigned to you (sorted by created, descending)
jira issue list -a$(jira me) --plain

# 2. Combine filters: high priority, open, created this month, label backend
jira issue list -yHigh -s"To Do" --created month -lbackend -a$(jira me) --plain

# 3. View issue details with recent comments
jira issue view ISSUE-1 --comments 5
```

Filter flags: `-a` assignee, `-r` reporter, `-s` status (prefix `~` for "not"), `-y` priority, `-l` label (repeatable), `-t` type, `-w` watching, `-R` resolution, `--created`/`--updated` with values like `-7d`, `week`, `month`.

> **IMPORTANT:** list commands open an interactive TUI by default. Always pass `--plain` (or `--table`, `--raw`, `--csv`) when running non-interactively, or the command will hang waiting for keyboard input.

### Creating an Issue

```bash
# 1. Non-interactive creation with all required fields
jira issue create -tBug -s"New Bug" -yHigh -lbug -b"Bug description" --no-input

# 2. Attach to an epic on creation
jira issue create -tStory -s"Summary" -PEPIC-42 --no-input

# 3. Body from stdin or template
echo "Description from stdin" | jira issue create -s"Summary" -tTask --no-input
jira issue create --template /path/to/template.tmpl -s"Summary" -tTask --no-input
```

Supports GitHub-flavored and Jira-flavored Markdown in descriptions and comments. Use `--custom "key=value"` for custom fields.

### Editing and Assigning

```bash
# Edit summary, priority, labels
jira issue edit ISSUE-1 -s"Updated summary" -yHigh -lbug --no-input

# Prefix minus (-) removes labels, components, or fix versions
jira issue edit ISSUE-1 --label -p2 --label p1 --fix-version -v1.0 --fix-version v2.0 --no-input

# Assign to a user, to self, to default, or unassign
jira issue assign ISSUE-1 "Jon Doe"
jira issue assign ISSUE-1 $(jira me)
jira issue assign ISSUE-1 default
jira issue assign ISSUE-1 x
```

### Transitioning Issues

```bash
# Move an issue to a new status (name must match the workflow exactly)
jira issue move ISSUE-1 "In Progress"

# Move with comment, resolution, and assignee
jira issue move ISSUE-1 Done -RFixed -a$(jira me) --comment "Completed"
```

### Comments, Links, and Worklogs

```bash
# Add a comment (positional body, template, or stdin)
jira issue comment add ISSUE-1 "My comment body"

# Internal (restricted) comment
jira issue comment add ISSUE-1 "Internal note" --internal

# Link two issues with a link type such as Blocks, Relates, Duplicates
jira issue link ISSUE-1 ISSUE-2 Blocks

# Add a remote web link
jira issue link remote ISSUE-1 https://example.com "Example text"

# Log work: time is positional, e.g. "2d 3h 30m"
jira issue worklog add ISSUE-1 "2d 3h 30m" --comment "Implementation" --no-input
```

### Sprints and Epics

```bash
# List issues in the current active sprint
jira sprint list --current -a$(jira me) --table --plain

# List issues in a specific sprint (get the ID from `jira sprint list`)
jira sprint list SPRINT_ID -yHigh --table --plain

# Add up to 50 issues to a sprint
jira sprint add SPRINT_ID ISSUE-1 ISSUE-2

# List issues in an epic
jira epic list EPIC-1 --table --plain

# Create an epic (epic name via -n is required)
jira epic create -n"Epic epic" -s"Everything" -b"Epic description" --no-input

# Add or remove issues in an epic (up to 50 at once)
jira epic add EPIC-1 ISSUE-1 ISSUE-2
jira epic remove ISSUE-1 ISSUE-2
```

## Common Patterns

### Scripting and Automation

The default view is an interactive UI. For shell pipelines, use machine-friendly output:

```bash
# Tab-separated plain output with selected columns, no headers
jira issue list -a$(jira me) --plain --columns key,summary,status --no-headers

# Raw JSON or CSV for parsing
jira issue list --raw
jira issue list --csv
```

### JQL Queries

Execute raw JQL within the configured project context using `-q/--jql`:

```bash
jira issue list -q "summary ~ cli AND status != Done"
```

### Working With a Different Project

Use `-p KEY` to target a project other than the one in the config:

```bash
jira issue list -pXYZ -a$(jira me) --plain
```

### Navigation Shortcuts (Interactive Mode Only)

In the default TUI: `v` view issue, `m` transition, `ENTER` open in browser, `c` copy URL, `q` quit. Not usable in scripts — use `--plain` instead.

## Best Practices

1. **Verify setup first**: `jira me` confirms auth and config are working
2. **Always pass `--no-input` with explicit flags** in scripts; interactive prompts will hang
3. **Use `$(jira me)`** for self-assignment and self-filtering
4. **Use `--plain --columns ... --no-headers`** for any output you intend to parse
5. **Match transition names exactly** to the Jira workflow; list interactively first if unsure

## Common Commands Quick Reference

**Issues:**
- `jira issue list -a$(jira me) --plain` - Your assigned issues
- `jira issue list -w` - Issues you are watching
- `jira issue list --history` - Issues you recently interacted with
- `jira issue create -tBug -s"Summary" -b"Body" --no-input` - Create issue
- `jira issue view ISSUE-1 --comments 5` - View details
- `jira issue move ISSUE-1 "In Progress"` - Transition issue
- `jira issue clone ISSUE-1 -s"New summary"` - Clone with modifications
- `jira issue delete ISSUE-1 --cascade` - Delete with subtasks

**Epics:**
- `jira epic list` - List epics
- `jira epic list EPIC-1` - List issues in an epic
- `jira epic create -n"Name" -s"Summary" --no-input` - Create epic

**Sprints:**
- `jira sprint list --current --table --plain` - Current sprint issues
- `jira sprint list --prev --table --plain` - Previous sprint issues
- `jira sprint add SPRINT_ID ISSUE-1 ISSUE-2` - Add issues to sprint

**Projects and Boards:**
- `jira project list` - List accessible projects
- `jira board list` - List boards in the configured project
- `jira open ISSUE-1` - Open issue in browser
- `jira me` - Show current user

## Progressive Disclosure

For detailed command documentation, refer to:
- **references/commands-detailed.md** - Comprehensive command reference with all flags and options
- **references/quick-reference.md** - Condensed command cheat sheet
- **references/troubleshooting.md** - Detailed error scenarios and solutions

Load these references when:
- User needs specific flag or option details
- Troubleshooting authentication, config, or permission issues
- Working with advanced features (JQL, custom fields, multiple configs)

## Common Issues Quick Fixes

**"command not found: jira"** - Install jira-cli or verify PATH

**Auth errors (401/403)** - Check `JIRA_API_TOKEN` is exported; for on-premises PAT auth set `JIRA_AUTH_TYPE=bearer`

**"Please select installation type" / config missing** - Run `jira init` to generate the config

**List command hangs in a script** - The default output is an interactive TUI; add `--plain`

**"Transition not found" or move fails** - Status name must match the workflow exactly; run `jira issue move` without arguments to see valid transitions

**Empty issue lists** - Verify the project/board chosen during `jira init`; try `-p PROJECT_KEY`

For detailed troubleshooting, load **references/troubleshooting.md**.

## Notes

- Works with both Jira Cloud and on-premises (Server/Data Center) installations
- Supports `basic`, `bearer` (PAT), and `mtls` (client certificates) auth types
- Sprint lists show only the 25 most recent sprints
- Descriptions and comments accept GitHub-flavored and Jira-flavored Markdown
- Shell completion is available via `jira completion --help`
