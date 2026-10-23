# Ponytail mode

Use this operation for `/ponytail`, `/ponytail lite`, `/ponytail full`,
`/ponytail ultra`, or `/ponytail off`.

Ponytail is active every response until the user says `stop ponytail`, `normal
mode`, or `/ponytail off`. The default is **full**.

## Levels

| Level | Behavior |
|---|---|
| **lite** | Build what was asked, then name the lazier alternative in one line. |
| **full** | Make the smallest complete change, preserving project conventions, checks, and safety. Default. |
| **ultra** | Also question unjustified parts of the request before building. |

Use the main skill's ladder and rules for the actual coding task. A mode switch
does not itself change files or create configuration.

## Skill-only behavior

Stack installs instructions, not the upstream plugin or session hooks. Do not
claim automatic startup activation, prompt bands, environment/config-file mode
resolution, or plugin auto-updates. Select the level in the conversation. A host
may expose `/ponytail`; otherwise load the skill and request the operation in
plain language. `/ponytail off`, `stop ponytail`, and `normal mode` deactivate it.
