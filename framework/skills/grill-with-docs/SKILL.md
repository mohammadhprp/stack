---
name: grill-with-docs
description: "Grill a plan or design against the codebase and its docs, and record settled domain terms and significant decisions as you go. Invoke explicitly with /grill-with-docs or by asking to grill with docs."
disable-model-invocation: true
---

# Grill with Docs

Sharpen a plan or design in one session. Interview the user and build the project's shared language at the same time.

1. Read [the interview workflow](references/grilling.md) and [the domain-modeling workflow](references/domain-modeling.md) before asking the first round. These are bundled files, not skills to invoke or install.
2. Identify the target repository. Read its applicable agent instructions, existing glossaries or glossary map, relevant ADRs, and the code that can answer factual questions. If there is no writable repository, tell the user and run the interview without writing files.
3. Ask decision questions in rounds, with a recommendation for each. Challenge ambiguous terms and discrepancies between the plan, docs, and code. Check what is actually settled after each round before asking the next one.
4. Write each newly settled domain term to the right `GLOSSARY.md` immediately. Offer an ADR only when a settled decision meets all three gates in the domain-modeling workflow, and write it only if the user agrees. Create docs lazily and leave other decisions in the conversation.
5. When the interview is done, summarize the agreed design, identify the docs changed or say why none were needed, and ask the user to confirm the shared understanding. Do not implement the plan as part of this skill.

See [examples](examples.md) for interview rounds and a documentation update.

## Source

Adapted from Matt Pocock's [grill-with-docs](https://github.com/mattpocock/skills/blob/5863806dd4e15c63174bee4e7515a4256792ecfa/skills/engineering/grill-with-docs/SKILL.md), [grilling](https://github.com/mattpocock/skills/blob/5863806dd4e15c63174bee4e7515a4256792ecfa/skills/productivity/grilling/SKILL.md), and [domain-modeling](https://github.com/mattpocock/skills/tree/5863806dd4e15c63174bee4e7515a4256792ecfa/skills/engineering/domain-modeling). See [LICENSE](LICENSE).
