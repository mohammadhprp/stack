# Jira CLI Quick Reference

Condensed cheat sheet. Use `--plain` (or `--table`, `--raw`, `--csv`) and `--no-input` when running non-interactively.

## Setup

```bash
jira init                  # one-time config wizard (Cloud/Local, auth, project, board)
export JIRA_API_TOKEN=...  # Jira Cloud API token
export JIRA_AUTH_TYPE=bearer  # only for on-premises PAT auth
jira me                    # verify auth and current user
```

## List Issues

```bash
jira issue list --plain
jira issue list -a$(jira me) --plain                          # assigned to me
jira issue list -r$(jira me) --created week --plain           # reported by me this week
jira issue list -w --plain                                    # watching
jira issue list --history                                     # recently interacted
jira issue list -yHigh -s"To Do" --created month -lbackend --plain
jira issue list -ax --created week --plain                    # unassigned
jira issue list -s~Done --plain                               # status is not Done
jira issue list -q "summary ~ cli" --plain                    # raw JQL
```

Filter flags: `-a` assignee (`x` = unassigned), `-r` reporter, `-s` status (`~` = not), `-y` priority, `-l` label (repeatable), `-t` type, `-w` watching, `-R` resolution, `--created`/`--updated` (`-7d`, `-1h`, `week`, `month`), `--created-before`/`--updated-before`, `--order-by <field> --reverse`.

## Output Modes

```bash
jira issue list --plain --columns key,summary,status --no-headers   # TSV, no header
jira issue list --raw        # JSON
jira issue list --csv        # CSV
```

## Create / Edit / Assign

```bash
jira issue create -tBug -s"Summary" -b"Body" -yHigh -lbug --no-input
jira issue create -tStory -s"Summary" -PEPIC-42 --no-input              # attach to epic
echo "Body" | jira issue create -s"Summary" -tTask --no-input           # body via stdin
jira issue edit ISSUE-1 -s"New summary" --no-input
jira issue edit ISSUE-1 --label -p2 --label p1 --no-input               # "-" removes
jira issue assign ISSUE-1 "Jon Doe"
jira issue assign ISSUE-1 $(jira me)                                    # self
jira issue assign ISSUE-1 default                                       # default assignee
jira issue assign ISSUE-1 x                                             # unassign
```

## Transition

```bash
jira issue move ISSUE-1 "In Progress"
jira issue move ISSUE-1 Done -RFixed -a$(jira me) --comment "Done"
```

Status name must match the workflow exactly.

## Comments / Links / Worklog

```bash
jira issue comment add ISSUE-1 "Comment body"
jira issue comment add ISSUE-1 "Internal note" --internal
jira issue link ISSUE-1 ISSUE-2 Blocks
jira issue unlink ISSUE-1 ISSUE-2
jira issue link remote ISSUE-1 https://example.com "Example text"
jira issue worklog add ISSUE-1 "2d 3h 30m" --comment "Note" --no-input
```

## Clone / Delete

```bash
jira issue clone ISSUE-1 -s"New summary" -a$(jira me)
jira issue clone ISSUE-1 -H"find:replace"          # replace text in summary/description
jira issue delete ISSUE-1 --cascade                # also deletes subtasks
```

## Epics

```bash
jira epic list --table --plain
jira epic list EPIC-1 --table --plain              # issues in an epic
jira epic create -n"Name" -s"Summary" -b"Body" --no-input
jira epic add EPIC-1 ISSUE-1 ISSUE-2               # up to 50
jira epic remove ISSUE-1 ISSUE-2                   # up to 50
```

## Sprints

```bash
jira sprint list --current --table --plain
jira sprint list --prev --table --plain
jira sprint list --next --table --plain
jira sprint list --state future,active --table --plain
jira sprint list SPRINT_ID -yHigh --table --plain  # issue filters work here too
jira sprint add SPRINT_ID ISSUE-1 ISSUE-2          # up to 50
```

## Projects, Boards, Misc

```bash
jira project list
jira board list
jira release list --project KEY
jira open ISSUE-1          # open in browser
jira me                    # current user
jira issue list -pXYZ ...  # override configured project
jira issue list -c ./config.yaml ...   # alternative config file
```
