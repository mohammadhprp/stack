# Jira CLI Troubleshooting

Common `jira` CLI error scenarios and solutions.

## Table of Contents

- [Command not found](#command-not-found)
- [Authentication errors (401/403)](#authentication-errors-401403)
- [Configuration problems](#configuration-problems)
- [Command hangs in scripts](#command-hangs-in-scripts)
- [Transition errors](#transition-errors)
- [Empty or wrong issue lists](#empty-or-wrong-issue-lists)
- [Epic or issue creation fails on-premises](#epic-or-issue-creation-fails-on-premises)
- [Custom fields not applied](#custom-fields-not-applied)
- [Multiple Jira instances or projects](#multiple-jira-instances-or-projects)
- [View command pager issues](#view-command-pager-issues)

## Command not found

**Symptom:** `command not found: jira`

**Fix:**
1. Install jira-cli: download a binary from the [releases page](https://github.com/ankitpokhrel/jira-cli/releases), or use Homebrew (`brew install ankitpokhrel/jira-cli/jira-cli`)
2. Verify: `jira --version`

## Authentication errors (401/403)

**Symptom:** API errors like `GET .../myself: 401 Unauthorized` or 403 on every command.

**Fix:**
1. Verify the token is exported: `echo "${JIRA_API_TOKEN:?not set}"`
2. Jira Cloud: use an API token from https://id.atlassian.com/manage-profile/security/api-tokens (your email is the login; basic auth is the default)
3. On-premises with a Personal Access Token: set `JIRA_AUTH_TYPE=bearer`
4. On-premises with login credentials: export the password as `JIRA_API_TOKEN` (basic auth, the default)
5. Re-check with `jira me`
6. Tokens can also be supplied via `.netrc` or `keychain` instead of the environment variable

## Configuration problems

**Symptom:** `Please select installation type`, or prompts for server/project on every command.

**Fix:**
1. Run `jira init` and complete the wizard (installation type Cloud/Local, auth type, login, project, board)
2. To start over, re-run `jira init` and answer the regeneration prompt
3. Confirm the config path with `jira --help`; override it per command with `-c /path/to/config.yaml` or `JIRA_CONFIG_FILE`

## Command hangs in scripts

**Symptom:** `jira issue list` (and other list commands) never returns in CI or non-interactive shells.

**Cause:** list and create commands open an interactive TUI by default.

**Fix:**
1. Add `--plain` (or `--table`, `--raw`, `--csv`) to list commands
2. Add `--no-input` and all required flags to create/edit/move/assign commands
3. Pipe bodies via stdin instead of letting them prompt: `echo "Body" | jira issue create -s"Summary" -tTask --no-input`

## Transition errors

**Symptom:** `jira issue move ISSUE-1 "Done"` fails or reports an unknown transition.

**Fix:**
1. Status names are case-sensitive and must match the workflow exactly (e.g. `"In Progress"`, not `in progress`)
2. Run `jira issue move` with no arguments to list valid transitions for the issue
3. If a workflow field (comment, resolution, assignee) is rejected during a move, the Jira workflow does not allow it on that transition — drop the flag or ask an admin to enable it

## Empty or wrong issue lists

**Symptom:** `jira issue list` returns nothing, or issues from the wrong project.

**Fix:**
1. The CLI is scoped to the project/board chosen during `jira init` — re-run `jira init` to change it
2. Target another project for a single command: `jira issue list -pXYZ --plain`
3. Reset filters one at a time to isolate the culprit (`-a`, `-s`, `--created`, ...)
4. Sprint views only show the 25 most recent sprints; older sprints are not listed

## Epic or issue creation fails on-premises

**Symptom:** Creation fails or misbehaves on a non-English Jira Server/Data Center installation.

**Cause:** older Jira APIs do not return untranslated issue type names.

**Fix:** manually fill `epic.name`, `epic.link`, and `issue.types.*.handle` fields in the generated config file.

## Custom fields not applied

**Symptom:** `--custom "key=value"` appears to be ignored or errors.

**Fix:**
1. Custom fields must be configured in the generated config file to map names to field IDs
2. See the jira-cli discussion on custom fields for the exact config format
3. Validate the value type matches the field (text, option, user, ...)

## Multiple Jira instances or projects

**Symptom:** Need to work against two servers or project contexts.

**Fix:**
1. Generate a config per context with `jira init`
2. Switch with `jira issue list -c ./other_config.yaml` or `export JIRA_CONFIG_FILE=./other_config.yaml`

## View command pager issues

**Symptom:** `jira issue view` output is truncated or opens an unwanted pager.

**Fix:** `view` uses `less` by default; pipe output (`jira issue view ISSUE-1 | cat`) or configure a custom pager when embedding in scripts.

## mtls authentication issues

**Symptom:** Certificate errors with auth type `mtls`.

**Fix:**
1. Re-run `jira init`, select installation type `Local`, auth type `mtls`, and provide the CA cert, client key, and client cert
2. If `JIRA_API_TOKEN` is set, it is used together with `mtls`
