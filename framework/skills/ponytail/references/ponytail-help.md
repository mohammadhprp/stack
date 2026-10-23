# Ponytail help

Use for `/ponytail-help`, “ponytail help”, or a request for the quick
reference. Display it once without changing mode, writing files, or persisting
anything.

| Operation | Trigger | Purpose |
|---|---|---|
| Mode | `/ponytail [lite\|full\|ultra\|off]` | Make the smallest complete change; name skipped work and risks. |
| Review | `/ponytail-review` | Review a change for bugs, safety, expected load, tests, speed, and unnecessary code. |
| Audit | `/ponytail-audit` | Apply the same quality review across a repository or named scope. |
| Debt | `/ponytail-debt [marker]` | Harvest `shortcut:` and older `ponytail:` comments into a ledger. |
| Gain | `/ponytail-gain` | Show upstream's published benchmark, not repo savings. |
| Help | `/ponytail-help` | Show this card. |

Say `stop ponytail` or `normal mode` to deactivate. Resume with `/ponytail`.
The default is `full`; `lite` builds the request and names a smaller option,
while `ultra` questions unjustified scope before building.

These are operation aliases inside one installed skill, not separately installed
commands. If the host does not expose slash commands, load `ponytail` and ask for
the operation in plain language. One-shot operations do not change mode.

Stack installs no upstream plugin or hooks, so automatic session activation,
prompt bands, environment/config-file defaults, and plugin auto-updates are not
provided. Update through Stack, then reinstall the skill into the target project.
See `ponytail-mode.md` beside this file for mode behavior and `upstream.md` for
the source version. Full upstream documentation: https://github.com/DietrichGebert/ponytail
