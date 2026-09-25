# Founder Examples

## Go-to-market plan

User: "We just launched our MVP and have 12 beta users. How do we go to market?"

Good agent behavior:

- Map the request to `references/go-to-market-plan.md` and read it before answering.
- Read `FOUNDER_CONTEXT.md` from the project root for product, ICP, stage, and competitor context; ask only for genuinely missing pieces instead of guessing.
- Diagnose readiness first, then deliver exactly 3 ranked strategies, each with a concrete playbook, metrics to track, milestones, and a first action doable in 30-60 minutes.
- Adapt tactics to the stated stage (post-MVP, validating) instead of generic advice like "do content marketing".

## LinkedIn and X posts with voice matching

User: "Write a LinkedIn post and a couple of tweets about why I switched from an agency to building my own SaaS."

Good agent behavior:

- Map the social request to `references/linkedin-writer.md` and `references/x-writer.md` and read both.
- For X, auto-select a creator voice that fits the topic, confirm it with the user, then write 3 posts in 3 different formats matched to that voice's structure and rhythm.
- For LinkedIn, auto-select 2 different formats, match the voice DNA in the reference posts, and make each post ready to paste with a first two lines that survive the "see more" cutoff.
- Reuse `references/viral-hook-creator.md` when stronger opening hooks are needed, and blend the brand voice when `FOUNDER_CONTEXT.md` exists.

## Landing-page CRO audit

User: "Here's my landing page URL, review my landing page and tell me what to fix."

Good agent behavior:

- Map the request to `references/cro-optimization.md` and read it before analyzing.
- Read `references/cro-optimization/cro_principles.md`, `references/cro-optimization/landing_page_patterns.md`, and `references/cro-optimization/element_audit_framework.md` as that reference requires.
- Fetch the page, extract the elements, then audit against the 13 principles and compare to the high-converting patterns.
- Return findings prioritized by conversion impact with before/after copy examples, a testing roadmap, and a "what's working well" section, not generic advice.

## Cold outreach sequence

User: "Write a cold email sequence to book demos with marketing directors at mid-size SaaS companies."

Good agent behavior:

- Map the request to `references/outreach-specialist.md` and read it, along with its `references/outreach-specialist/outreach-templates.md` and `references/outreach-specialist/sequence-strategy.md`.
- Ask only the missing diagnostic questions (platform, offer result, proof, warm or cold, desired length), then build a default 3-message sequence.
- Keep the first touch under the platform limit, lead with the prospect, use one CTA per message, and make each follow-up add new value instead of "just bumping this".
- Verify against the reference's checklist; no em dashes, no AI slang, no "just following up".

## PRD request

User: "Turn my idea for a churn-prevention tool into a PRD I can hand to a coding agent."

Good agent behavior:

- Map the request to `references/prd-generator.md` and read `references/prd-generator/prd_template.md` first.
- Do not read `FOUNDER_CONTEXT.md`: the reference deliberately treats the PRD as standalone.
- Ask targeted clarifying questions from the question bank (fewer is better), then fill every applicable template section.
- Give every feature testable acceptance criteria, keep P0 at roughly 30-40%, include field-level data models, and always include the "Implementation Notes for AI" section before saving and converting the document.
