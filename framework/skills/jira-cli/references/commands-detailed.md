# Jira CLI Detailed Command Reference

Complete flag and option reference for the `jira` CLI. All list flags are POSIX-compliant and combine freely.

## Table of Contents

- [Global Options](#global-options)
- [Issue List](#issue-list)
- [Issue Create](#issue-create)
- [Issue Edit](#issue-edit)
- [Issue Assign](#issue-assign)
- [Issue Move (Transition)](#issue-move-transition)
- [Issue View](#issue-view)
- [Issue Link / Unlink](#issue-link--unlink)
- [Issue Clone](#issue-clone)
- [Issue Delete](#issue-delete)
- [Issue Comment](#issue-comment)
- [Issue Worklog](#issue-worklog)
- [Epic](#epic)
- [Sprint](#sprint)
- [Releases, Projects, Boards, Other](#releases-projects-boards-other)

## Global Options

```bash
jira --version              # Print version
jira init                   # Interactive config wizard (run once)
jira me                     # Print the current logged-in user
jira open [KEY-1]           # Open project or issue in browser
jira completion --help      # Shell completion setup (Bash/Zsh)
```

- `-c, --config <file>` — use an alternative config file (or `JIRA_CONFIG_FILE` env var)
- `-p, --project <KEY>` — override the configured project on list/create commands

## Issue List

```bash
jira issue list [flags]
```

| Flag | Description |
| --- | --- |
| `-a, --assignee <user>` | Filter by assignee; `x` = unassigned; use `$(jira me)` for yourself |
| `-r, --reporter <user>` | Filter by reporter |
| `-s, --status <status>` | Filter by status; prefix with `~` for "not" (e.g. `-s~Done`) |
| `-y, --priority <priority>` | Filter by priority (e.g. `High`) |
| `-t, --type <type>` | Filter by issue type (e.g. `Bug`) |
| `-l, --label <label>` | Filter by label; repeatable |
| `-w, --watching` | Issues you are watching |
| `-R, --resolution <value>` | Filter by resolution (e.g. `"Won't do"`, `Fixed`) |
| `--history` | Issues you recently interacted with |
| `--created <value>` | Relative window: `-7d`, `-1h`, `week`, `month` |
| `--updated <value>` | Same value format as `--created` |
| `--created-before <value>` | e.g. `-24w` |
| `--updated-before <value>` | Same format |
| `-q, --jql <query>` | Raw JQL within the project context |
| `--order-by <field>` | Order by field (e.g. `rank`, `created`) |
| `--reverse` | Reverse sort order |
| `--plain` | Non-interactive table output (use in scripts) |
| `--table` | Table view for explorer commands (epic/sprint) |
| `--raw` | Raw JSON output |
| `--csv` | CSV output |
| `--columns <list>` | Comma-separated columns to display |
| `--no-headers` | Omit the header row (for `awk`/parsing) |

Examples:

```bash
# Assigned to me, high priority, open, created this month
jira issue list -a$(jira me) -yHigh -sopen --created month

# Unassigned issues created this week
jira issue list -ax --created week

# Not Done, created before 24 weeks ago, assigned to someone
jira issue list -s~Done --created-before -24w -a~x

# Issues I reported this week, oldest first
jira issue list -r$(jira me) --created week --reverse
```

## Issue Create

```bash
jira issue create [flags]
```

| Flag | Description |
| --- | --- |
| `-t, --type <type>` | Issue type (required) |
| `-s, --summary <text>` | Summary/title (required) |
| `-b, --body <text>` | Description; supports GitHub- and Jira-flavored Markdown |
| `-y, --priority <value>` | Priority |
| `-l, --label <label>` | Label; repeatable |
| `-P, --parent <KEY>` | Attach to an epic (or parent in next-gen projects) |
| `--fix-version <value>` | Fix version |
| `--custom "key=value"` | Set custom fields; repeatable |
| `--template <file\|->` | Load description from file or stdin |
| `--no-input` | Skip all interactive prompts |

```bash
# Minimal non-interactive creation
jira issue create -tTask -s"Summary" -b"Description" --no-input

# Description piped from stdin
echo "Description" | jira issue create -s"Summary" -tTask --no-input

# Story attached to EPIC-42 with labels and fix version
jira issue create -tStory -s"Summary" -PEPIC-42 -lbackend --fix-version v2.0 --no-input
```

## Issue Edit

```bash
jira issue edit ISSUE-1 [flags]
```

Accepts the same field flags as create (`-s`, `-b`, `-y`, `-l`, `-C` component, `--fix-version`). Prefix a value with `-` to remove it:

```bash
# Remove label p2, add p1; swap fix version v1.0 for v2.0
jira issue edit ISSUE-1 --label -p2 --label p1 --fix-version -v1.0 --fix-version v2.0 --no-input
```

## Issue Assign

```bash
jira issue assign ISSUE-1 <user>
jira issue assign ISSUE-1 $(jira me)   # assign to self
jira issue assign ISSUE-1 default      # default assignee
jira issue assign ISSUE-1 x            # unassign
```

A partial name ("suffix") may prompt for selection if it matches multiple users.

## Issue Move (Transition)

```bash
jira issue move ISSUE-1 <status> [flags]
```

| Flag | Description |
| --- | --- |
| `--comment <text>` | Comment added during the transition (workflow must allow it) |
| `-R, --resolution <value>` | Set resolution while moving (e.g. `Fixed`) |
| `-a, --assignee <user>` | Reassign while moving |

```bash
jira issue move ISSUE-1 "In Progress"
jira issue move ISSUE-1 Done -RFixed -a$(jira me) --comment "Completed"
```

The status name must match the workflow exactly (case-sensitive). Run `jira issue move` with no arguments to pick from valid transitions interactively.

## Issue View

```bash
jira issue view ISSUE-1 [--comments <n>]
```

Renders the description as Markdown in the terminal (pager: `less` by default). `--comments 5` shows the 5 most recent comments.

## Issue Link / Unlink

```bash
jira issue link ISSUE-1 ISSUE-2 <link-type>   # e.g. Blocks, Relates, Duplicates
jira issue unlink ISSUE-1 ISSUE-2
jira issue link remote ISSUE-1 <url> "<text>" # add a remote web link
```

## Issue Clone

```bash
jira issue clone ISSUE-1 [flags]
```

| Flag | Description |
| --- | --- |
| `-s, -y, -a, -l, -C` | Override summary, priority, assignee, labels, components on the clone |
| `-H, --replace "find:replace"` | Case-sensitive text replacement in summary and description |

```bash
jira issue clone ISSUE-1 -s"Modified summary" -yHigh -a$(jira me)
jira issue clone ISSUE-1 -H"find me:replace with me"
```

## Issue Delete

```bash
jira issue delete ISSUE-1 [--cascade]
```

`--cascade` also deletes all subtasks.

## Issue Comment

```bash
jira issue comment add ISSUE-1 "Comment body" [flags]
```

| Flag | Description |
| --- | --- |
| `--internal` | Mark as an internal (restricted) comment |
| `--template <file\|->` | Load comment from file or stdin |

```bash
echo "Comment from stdin" | jira issue comment add ISSUE-1
```

The positional body takes precedence over `--template` when both are given.

## Issue Worklog

```bash
jira issue worklog add ISSUE-1 "<time>" [flags]
```

Time is positional and accepts values like `"10m"`, `"2h"`, `"2d 3h 30m"`.

```bash
jira issue worklog add ISSUE-1 "2d 3h 30m" --comment "Implementation" --no-input
```

## Epic

```bash
jira epic list [EPIC-KEY] [issue-list-flags]   # list epics, or issues in an epic
jira epic list EPIC-1 -ax -yHigh --table --plain
jira epic create -n"Epic name" -s"Summary" -yHigh -lbug -b"Description" --no-input
jira epic add EPIC-1 ISSUE-1 ISSUE-2           # up to 50 issues at once
jira epic remove ISSUE-1 ISSUE-2               # up to 50 issues at once
```

`epic list` supports all `issue list` filter flags except issue type. `epic create` requires `-n/--name`.

## Sprint

```bash
jira sprint list [SPRINT_ID] [flags]
```

| Flag | Description |
| --- | --- |
| `--current` | Issues in the current active sprint |
| `--prev` | Issues in the previous sprint |
| `--next` | Issues in the next planned sprint |
| `--state <list>` | e.g. `future,active` |
| `--table` | Table view (default is explorer) |

```bash
jira sprint list --current -a$(jira me) --table --plain
jira sprint list --state future,active --table --plain
jira sprint list SPRINT_ID -yHigh --table --plain   # all issue filters supported
jira sprint add SPRINT_ID ISSUE-1 ISSUE-2           # up to 50 issues at once
```

Only the 25 most recent sprints are shown. Get sprint IDs from `jira sprint list`.

## Releases, Projects, Boards, Other

```bash
jira release list [--project KEY]   # project versions; feature must be enabled on the instance
jira project list                   # all accessible projects
jira board list                     # boards in the configured project
jira open                           # open project in browser
jira open ISSUE-1                   # open issue in browser
```
