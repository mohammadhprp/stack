---
name: arena
description: "Spawn N parallel candidates at the same task, pick a base, graft the strongest parts of the losers into it. Use for /arena, 'arena this', 'throw it in the arena', or when one attempt at a non-trivial artifact would lock in the wrong shape."
disable-model-invocation: true
---

# Arena

## Installed-skill boundaries

Resolve resources from the directory containing this loaded `SKILL.md`. Use the
host's delegation tool and supported parameters, not another harness's tool names.
Use parent-model defaults unless the user explicitly selects available models;
discover supported IDs before passing a model override. Never guess model names.
If delegation is unavailable, explain the limitation and offer sequential
alternatives; do not claim independent parallel judging occurred.

Named principles below describe the discipline in the text, not dependencies on
separately installed skills. Keep candidates isolated in task-specific worktrees
or scratch directories. Do not commit, publish, delete user work, or start
implementation outside the requested scope.

Open a todolist with one entry per phase (Frame, Fan out, Cross-judge, Pick, Graft, Verify) before launching anything.

## Phase A: Frame

Every candidate gets the same prompt, so the prompt is the contract.

1. State the artifact each candidate is producing.
2. Derive the rubric, 3-6 concrete gradeable criteria for what success looks like on *this* task. It is the picker's tool in Phase D. Candidates only see the task.
3. Pick the runners using the host's available capabilities. Honor explicitly configured user choices only when supported; `auto` or `inherit-parent` means omit the model override. If a chosen model is rejected, report it and use the parent default rather than guessing another ID. Use distinct model families when available and requested; otherwise disclose that independent candidates share a model. Select reasoning effort and speed only through supported host settings.
4. Give each candidate its own output location: a task worktree where available, otherwise an approved scratch directory per candidate. No shared writable output.

## Phase B: Fan out

Launch all N candidates concurrently when the host supports it, each with the task, the path to the shared grounding, its own output path, and instructions to produce the artifact and a short rationale naming the alternatives it considered and what it rejected. Use the host's actual background/delegation parameters.

If a candidate produces no output, proceed with N-1 and note the dropout.

## Phase C: Cross-judge

After all Phase B candidates complete, launch a fresh read-only judge. Honor an explicitly configured supported cross-judge model; otherwise inherit the parent model. Prefer a different family only when available and authorized. It sees the rubric and candidates by path label, scores each criterion, and recommends a base with rationale. It runs in parallel with your reading in Phase D. Do not spawn it while candidates are still writing.

## Phase D: Pick a base

Read every candidate end to end before picking. Score each against the rubric criterion by criterion, not on holistic feel, and compare with the cross-judge. Agreement on the base confirms the pick. Disagreement means one of you is biased or the rubric was ambiguous. Read both rationales before deciding.

Pick the base a future maintainer can extend most easily without breaking invariants. When two feel tied, prefer the cleaner boundary or smaller API, per the Laziness Protocol.

## Phase E: Graft

Walk each losing candidate once more for what is worth porting into the base, usually one or two things per candidate, not most of it. Fold each graft in by hand, per the **redesign-from-first-principles** principle skill, so the result stays coherent under one mental model. Don't paste mechanically.

If the candidates converge on one shape, that is strong agreement. Note it and ship the consensus shape. No graft is needed. If they wildly diverge, Phase A was under-specified. Reframe and re-run rather than averaging the divergence.

## Phase F: Verify

Verify the synthesized artifact like any other output, per the **prove-it-works** principle skill. If verification finds a problem the arena missed, either Phase A was wrong (re-frame and re-run) or a candidate caught it and you missed the graft (go back to Phase E). Don't paper over it.

## Outputs

One synthesized artifact, with a short synthesis note alongside naming the base and why, the cross-judge's verdict, each graft with its source candidate, what was rejected and why, any convergence or dropouts, and the verification result.

## Source

Adapted from [backnotprop/pstack](https://github.com/backnotprop/pstack/blob/e44668c2aad92eacf4e84df47984121fefb11b82/skills/arena/SKILL.md), MIT. Stack uses host-native delegation and discovered model IDs rather than pstack-specific settings and model aliases.
