# Upstream source

- Repository: https://github.com/DietrichGebert/ponytail
- Version: 5.3.2 (`package.json`)
- Reviewed revision: `09c434ab73ca43761c5d8b0c5d341926c2fba710`
- Snapshot: https://github.com/DietrichGebert/ponytail/tree/09c434ab73ca43761c5d8b0c5d341926c2fba710
- License: MIT; the upstream notice is included in `../LICENSE`.

## Source mapping

| Bundled content | Upstream source |
|---|---|
| `../SKILL.md`, `ponytail-mode.md` | `skills/ponytail/SKILL.md` |
| `ponytail-rules.md` | `.agents/rules/ponytail.md` |
| `ponytail-review.md` | `skills/ponytail-review/SKILL.md` |
| `ponytail-audit.md` | `skills/ponytail-audit/SKILL.md` |
| `ponytail-debt.md` | `skills/ponytail-debt/SKILL.md` |
| `ponytail-gain.md` | `skills/ponytail-gain/SKILL.md`, `benchmarks/results/2026-10-07-agentic.md` |
| `ponytail-help.md` | `skills/ponytail-help/SKILL.md` |

## Stack adaptations

Stack distributes one skill with operation references, not upstream's six skills
or its plugins, commands, hooks, and status UI. Slash forms identify operations
when supported by the host; they do not guarantee installed commands. Mode starts
when the skill is activated, not automatically at session startup. Environment
and config-file mode resolution, band controls, and plugin update instructions
are deliberately omitted because no runtime implements them in this package.

Resource paths resolve from the installed `SKILL.md`. Use the target project's
test conventions. Review and audit remain read-only and leave remote state alone;
one-shot operations leave mode unchanged. Benchmark figures are attributed to
upstream, not presented as results for this adaptation. The gain card uses ASCII
bars and links to the pinned benchmark instead of a missing repository file.

The `principle-*.md` files predate this sync and are local supplements, not files
from this upstream revision. They are retained unchanged and are optional; their
project-specific examples do not impose architecture on another repository.
