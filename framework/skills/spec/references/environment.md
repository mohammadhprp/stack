# Environment configuration

Keep actual values in your private shell configuration, such as `~/.zshsecrets`, not in skill files. Export them so child processes can read them. This file documents placeholders only; the skill has no built-in personal settings.

## Single-project setup

Set these exports in the user's existing private shell configuration. Do not create or edit that configuration without permission:

```sh
export JIRA_API_TOKEN='your-jira-token'
export NOTION_API_TOKEN='your-notion-token'
export SPEC_JIRA_SITE='https://jira.example.com'
export SPEC_JIRA_PROJECT_KEY='PROJ'
export SPEC_NOTION_TASKS_DATABASE_ID='your-database-id'
export SPEC_NOTION_TASKS_DATA_SOURCE_ID='your-data-source-id'
export SPEC_NOTION_PROJECT_PAGE_ID='your-project-relation-page-id'
export SPEC_NOTION_SUB_PROJECT='your-sub-project'
export SPEC_NOTION_TAGS_JSON='["your-tag"]'
export SPEC_NOTION_SPRINT_PARENT_PAGE_ID='your-sprint-plan-parent-page-id'
export SPEC_TIMEZONE='your-IANA-timezone'
```

Use bare Notion IDs, not page URLs or `collection://` URIs. Confirm them through the current Notion schema before writing. If a sub-project or tags are deliberately unused, use an empty string or `[]` respectively; distinguish explicitly empty variables from missing variables.

Optional scheduling and board settings:

| Variable | Format |
| --- | --- |
| `SPEC_JIRA_BOARD_IDS_JSON` | JSON array of board IDs |
| `SPEC_WORKING_DAYS_JSON` | JSON array of weekday names |
| `SPEC_CAPACITY_JSON` | JSON object with explicit amount, unit, period, and known commitments |
| `SPEC_DATE_POLICY` | User-confirmed scheduling policy in plain text |

Absent scheduling settings mean ordered backlog only, not invented hours, dates, or workweek. Missing sprint parent allows task synchronization but blocks sprint-plan creation. Authentication may alternatively use existing CLI credentials; environment tokens override where supported. `JIRA_CONFIG_FILE` can select an existing native Jira CLI configuration.

## Multiple projects

Optionally export `SPEC_PROJECTS_JSON` as a JSON array of profile objects. Each object uses the same uppercase variable names as the single-project setup, with native JSON values for `_JSON` fields (arrays/objects, not encoded strings). Require `SPEC_JIRA_SITE` and `SPEC_JIRA_PROJECT_KEY` in each record; permit `JIRA_CONFIG_FILE` for an existing site-specific CLI configuration. Never include tokens in these records.

An exact matching record replaces the simple profile for that site/project. Do not inherit another project's relation, sub-project, tags, scheduling, or destination. Use the simple profile only for its exact site/project when no matching record exists. Reject duplicate records and invalid JSON. Do not save this variable into skill files or a profile cache.

Set the environment variables above before using the skill.
