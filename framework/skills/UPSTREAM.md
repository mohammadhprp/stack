# Upstream review — 2026-11-19

The review compares the files Stack actually ships, not just repository HEADs. Stack keeps its skill IDs, installed-relative paths, local references, and approval rules. See the individual skills for source pins and licenses.

| Skill | Source | Result |
| --- | --- | --- |
| `adhd` | [ayghri/i-have-adhd](https://github.com/ayghri/i-have-adhd) | Updated explicit activation and examples; keeps `/adhd` and completeness rules. |
| `agent-browser` | [vercel-labs/agent-browser](https://github.com/vercel-labs/agent-browser) | Updated the CLI guide index; available guides depend on the installed CLI version. |
| `diagram-design` | [cathrynlavery/diagram-design](https://github.com/cathrynlavery/diagram-design) | Updated to 2.6.168 with PlantUML import, geometry checks, and more examples. |
| `docker` | [docker/skills](https://github.com/docker/skills) | Added v3 kit guidance; kept the existing v2 references for legacy projects. |
| `find-docs` | [upstash/context7](https://github.com/upstash/context7) | Synced wording and recorded provenance. |
| `find-skills` | [vercel-labs/skills](https://github.com/vercel-labs/skills) | Recorded provenance; retained source inspection and project-local installation instead of adopting global unattended installs. |
| `humanizer` | [blader/humanizer](https://github.com/blader/humanizer) | Updated to 4.0.0 with later wording and informal-text fixes; retained source-grounded editing. |
| `laravel-best-practices` | [laravel/boost](https://github.com/laravel/boost) | Refreshed architecture, routing, caching, and test isolation guidance; retained Stack’s standalone references rather than importing Blade templates. |
| `merge-risk`, `tdd` | [mattpocock/skills](https://github.com/mattpocock/skills) | Adopted relevant guidance without replacing Stack’s risk scoring or bug-fix workflow. |
| `perf-web-optimization` | [tech-leads-club/agent-skills](https://github.com/tech-leads-club/agent-skills) | Updated audit routing; bundled reference files were unchanged. |
| `ponytail` | [DietrichGebert/ponytail](https://github.com/DietrichGebert/ponytail) | Recorded 5.3.3; bundled workflows did not change. |
| `resolve-merge-conflicts` | [warpdotdev/common-skills](https://github.com/warpdotdev/common-skills) | Recorded provenance; upstream helper matched; retained portable paths and overwrite safeguards. |
| `show-me` | [humanlayer/skills](https://github.com/humanlayer/skills) | Matched explicit activation while retaining portable preview instructions. |
| `stack-diagram` | [caru-ini/stack-diagram](https://github.com/caru-ini/stack-diagram) | Updated vendor icon support; did not bundle public gallery uploads or deletion. |
| `unslop` | [backnotprop/pstack](https://github.com/backnotprop/pstack), [petergyang/no-ai-slop](https://github.com/petergyang/no-ai-slop) | Reviewed the condensed rules and kept Stack’s numbered supplements. |

No newer applicable content was found for `before-and-after`, `docs-writer`, `frontend-design`, `go-best-practices`, `grill-with-docs`, `rust-best-practices`, `sentry`, `skill-creator`, or `untangle-codebase`. Changes to unrelated upstream skills, plugins, and services were not imported.

No matching external skill source was verified for `backend-best-practices`, `changelog`, `commit`, `founder`, `gh`, `glab`, `jira-cli`, `law-of-business`, `merge-request`, `new-feature`, `pull-request`, `release`, `review`, `security-audit`, `spec`, or `walkthrough`. These were left as-is; a tool's own repository is not necessarily the source of Stack's skill for that tool.
