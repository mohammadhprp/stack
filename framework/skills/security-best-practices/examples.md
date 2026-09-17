# Security Best Practices Examples

## Write secure-by-default code in a new project

User: "Start a new FastAPI project. Make it secure from the start."

Good agent behavior:

- Identify the language and framework (Python + FastAPI) and load the matching reference file.
- Read every relevant reference, including the general backend security file, before writing code.
- Apply secure defaults from the start: parameterized queries, input validation, safe secrets handling.
- Flag critical violations to the user as they appear rather than burying them in a report.

## Passive vulnerability detection while working

User: "Add a profile endpoint that returns users by their auto-incrementing ID."

Good agent behavior:

- Notice the use of incrementing public IDs and flag it as a high-impact issue: it leaks resource counts and enables ID guessing.
- Suggest a UUID4 or random hex string instead, with a one-sentence impact statement.
- Notify the user and ask whether to fix the finding before proceeding.
- Fix one finding at a time with a concise comment tied to the specific best practice.

## Full security report

User: "Produce a security report for the codebase."

Good agent behavior:

- Identify all languages and frameworks, frontend and backend, and load every matching reference.
- Write the report to `security_best_practices_report.md` (or where the user asks), starting with an executive summary.
- Group findings by severity with numeric IDs and include file:line references for each.
- Focus on the most critical findings and give each a one-sentence impact statement.
- Offer to fix findings one at a time and check for regressions against the project's tests.
