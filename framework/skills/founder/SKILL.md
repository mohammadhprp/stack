---
name: founder
description: "Use this skill when a founder or business owner needs business strategy and next moves, go-to-market or launch planning (including a Product Hunt product launch), product/market or competitor intelligence, a PRD, pricing strategy, an SOP, marketing ideas, conversion-rate optimization (CRO), brand, ad, or email copy, cold outreach or email sequences, LinkedIn or X (Twitter) posts, viral hooks, or a lead-magnet give-away post. Plain-language triggers: grow my business, what should I do next, growth plan, launch plan, market entry, competitor analysis, product requirements document, tier structure and price points, process documentation, campaign ideas, landing-page audit, ads, sales emails, LinkedIn post, tweet, hook, and giveaway post."
license: MIT
---

# Founder

A consolidated skill for founder and business work. Every concern is covered by one reference file, and this file routes you to the right one. Each reference is a self-contained procedure: read it, follow its workflow, then run its own quality checklist.

## How to Apply

1. Map the request to the concerns in the index below. One request can map to several references (a launch, for example, needs `go-to-market-plan` plus copy and outreach).
2. Read each mapped `references/<topic>.md` before acting. It names any additional files it needs, such as its own templates or examples.
3. Several references require business context. When the mapped reference asks for it, read `FOUNDER_CONTEXT.md` from the project root. If it does not exist or is empty, ask the user for the missing pieces (company, industry, audience, value proposition, stage, competitors, brand voice) before producing output. A few references, such as `references/prd-generator.md`, deliberately do not use it.
4. Keep the output specific to the user's business. These references forbid generic advice; hold to that.
5. Finish by running the reference's own self-verification checklist before presenting.

## When to use this skill

- **Strategy and planning** — the user asks what to do next, how to grow, where a bottleneck is, how to enter a market, how to plan a launch, or what competitors are doing.
- **Product** — the user needs a PRD or product requirements, a pricing model or tier structure, or a standard operating procedure for a business process.
- **Copywriting and content** — the user needs ads, landing-page copy, sales-page copy, email copy, brand copy, marketing ideas, viral hooks, or a lead-magnet give-away post.
- **Social** — the user needs LinkedIn or X (Twitter) posts in a specific format or creator voice.
- **Outreach and launch** — the user needs a cold outreach sequence, LinkedIn DMs, follow-up emails, or a Product Hunt launch plan.
- **Conversion** — the user wants a landing-page audit or conversion-rate optimization recommendations.

## Do not use this skill when

- The task is general software engineering: writing, reviewing, refactoring, or debugging code, designing system architecture, or fixing a build. Use an engineering skill for that.
- The user needs infrastructure, security, or framework-specific guidance unrelated to the business, marketing, or product work above.

## References Index

Cross-cutting requests often need more than one reference file.

| Concern | Read |
| --- | --- |
| **Strategy & planning** | |
| Next moves, growth strategy, bottlenecks, strategic guidance | [`references/strategic-planning.md`](references/strategic-planning.md) |
| Go-to-market strategy, launch planning, market entry | [`references/go-to-market-plan.md`](references/go-to-market-plan.md) |
| Competitor analysis, market positioning, competitive intelligence | [`references/competitor-intel.md`](references/competitor-intel.md) |
| **Product** | |
| PRDs, product requirements, feature specs | [`references/prd-generator.md`](references/prd-generator.md) |
| Pricing strategy, tier structures, price points | [`references/pricing-strategist.md`](references/pricing-strategist.md) |
| SOPs, process documentation, operational guides | [`references/sop-creator.md`](references/sop-creator.md) |
| **Copywriting & content** | |
| Ads, landing-page and sales-page copy, email copy, brand voice | [`references/brand-copywriter.md`](references/brand-copywriter.md) |
| Viral hooks, opening lines, scroll-stoppers | [`references/viral-hook-creator.md`](references/viral-hook-creator.md) |
| Lead-magnet give-away posts that drive comments and DMs | [`references/lead-magnet-generator.md`](references/lead-magnet-generator.md) |
| Marketing ideas, growth tactics, campaign concepts | [`references/marketing-ideas.md`](references/marketing-ideas.md) |
| **Social** | |
| X (Twitter) posts and threads | [`references/x-writer.md`](references/x-writer.md) |
| LinkedIn posts | [`references/linkedin-writer.md`](references/linkedin-writer.md) |
| **Outreach & launch** | |
| Cold outreach, LinkedIn DMs, email sequences | [`references/outreach-specialist.md`](references/outreach-specialist.md) |
| Product Hunt launch plan | [`references/product-hunt-launch-plan.md`](references/product-hunt-launch-plan.md) |
| **Conversion** | |
| CRO, landing-page audits, conversion optimization | [`references/cro-optimization.md`](references/cro-optimization.md) |

## Business Context

Most references personalize their output from `FOUNDER_CONTEXT.md`, a file in the user's project root that describes their company, audience, value proposition, brand voice, goals, products, and competitors. Read it when a mapped reference calls for it. Do not create it for the user and do not require it where a reference says it is optional: when it is missing or empty, ask for the specific context the reference needs instead.
