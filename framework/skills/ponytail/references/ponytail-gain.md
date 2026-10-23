# Ponytail gain

Use for `/ponytail-gain`, “what does ponytail save”, or a request for the
measured-impact scoreboard. This is a one-shot display: do not change mode,
write files, or persist anything.

Display upstream's published Ponytail 5 benchmark, not measurements of this
repository: headless Claude Code (Opus 5.5, default effort), 39 tasks, five runs
each. Eighteen tasks have hidden correctness and safety checks. Figures are
geometric means of per-task median ratios against the same agent without the
skill. The benchmark disables shell access and runs generated checks afterwards;
it is not evidence for Stack's adaptation or every harness.

Source: https://github.com/DietrichGebert/ponytail/blob/09c434ab73ca43761c5d8b0c5d341926c2fba710/benchmarks/results/2026-10-07-agentic.md

```text
  ponytail gain           upstream benchmark · 39 tasks x 5 runs · Opus 5.5

  no-skill        ####################  100%
  Lines of code   #########...........   47%   down 53%
  Output tokens   ###########.........   55%   down 45%
  Cost            ###############.....   74%   down 26%
  Time            ############........   59%   down 41%
  Hidden checks passed   97%  (no-skill 96%)
  Tests where the logic needs one   98%  (no-skill 68%)

  This repo:  /ponytail-debt  (shortcuts you deferred)
              /ponytail-audit (what's still cuttable)
```

These are benchmark figures, not this repository. Never invent a per-repo
savings number: an unbuilt alternative is not a measurable baseline. The debt
ledger can count actual markers; an audit proposes cuts, not realized savings.
