---
name: how
description: "Use for \"how does X work\", code walkthroughs before changing something, and placement / ownership / layering questions (\"where should this live\", \"which package owns this\", \"is this the right layer\"). Explains subsystem architecture, runtime flow, onboarding mental models. Use why for motivation."
disable-model-invocation: true
---

# How

Explore the codebase to answer "how does X work?" questions. Produce architectural explanations at the level of a senior engineer onboarding onto a subsystem, enough to build a working mental model, not so much that it reads like annotated source code.

Resolve every reference from the directory containing this loaded `SKILL.md`.
Explorers and explainers are read-only roles using the host's actual delegation
tool and supported parameters. Inherit the parent model unless the user explicitly
selects a model that the host lists as available. `auto` and `inherit-parent` mean
omit the override. Report rejected model choices and inherit rather than guessing
IDs. If the host has no delegation tool, run the roles sequentially yourself and
disclose that they were not independent. Do not edit source or publish findings.

## 1. Assess complexity

If the scope is ambiguous, state your interpretation and explore. The user can redirect.

- **Simple** (a single module, a small utility, a narrow question such as "how does function X work"): no explorers. Spawn one explainer that explores and explains in one pass, built from `references/explainer-prompt.md` without the explorer-findings section. Go to step 4.
- **Complex** (a subsystem spanning multiple files or services, a cross-cutting feature, a full architectural overview): go to step 2.

When in doubt, take the simple path.

**Other harnesses.** Use the tool exposed by the current host, not a guessed API.
Keep the role prompt, source scope, and read-only boundary; drop unsupported
parameters and disclose any loss of independence.

## 2. Explore (complex only)

Decompose the question into 2 to 4 angles, each a distinct slice of the subsystem. Spawn all explorers in a single message, each with `references/explorer-prompt.md` and its angle filled in.

## 3. Synthesize (complex only)

Once all explorers have returned, spawn one explainer with `references/explainer-prompt.md` and every explorer's findings filled in.

## 4. Present

Present the explainer's output. Light edits for clarity or context from the conversation are fine. Do not substantially rewrite it.

## Output format

The explanation uses the sections defined in `references/explainer-prompt.md`, dropping any that do not apply: Overview, Key Concepts, How It Works, Where Things Live, Gotchas.

## Optional critique

Keep Stack's critique mode when the user explicitly asks for architectural issues
or improvements. Explain first. Then give fresh read-only critics the explanation,
relevant source paths, `references/critic-prompt.md`, and
`references/critique-rubric.md`. Use supported user-selected models or parent
defaults. Judge findings as Act on, Consider, Noted, or Dismissed with source
evidence; present them after the explanation. Do not apply fixes.

## Source

Adapted from [backnotprop/pstack](https://github.com/backnotprop/pstack/blob/e44668c2aad92eacf4e84df47984121fefb11b82/skills/how/SKILL.md), MIT. Stack retains its optional critique references and uses host-native delegation.
