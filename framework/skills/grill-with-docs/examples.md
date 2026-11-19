# Grilling Skill — Example Outputs

Examples of how `grill-with-docs` stress-tests a plan through design-tree rounds and records settled vocabulary.

## Example 1: New feature plan (frontier round, recommendations, wait)

**User:** "Grill me on my plan to add Redis caching to the API."

**What the agent does:**

Maps the design tree (invalidation strategy branches off what is cached; TTL branches off invalidation; cache key shape branches off endpoints covered), computes the frontier, and asks round 1:

```
❓ **Q1** - **What gets cached**: Which endpoints or queries go through Redis — everything, read-heavy endpoints only, or specific hot queries?

➡️ Read-heavy endpoints only. Caching everything multiplies invalidation bugs for little gain.

---

❓ **Q2** - **Where Redis lives**: Managed service (Upstash/ElastiCache), self-hosted container, or reuse existing infrastructure?

➡️ Managed service. One less stateful thing to operate for a first caching pass.
```

Then waits. Q3 (TTL values) stays out of this round — it depends on Q1's answer.

**Result:** Whole frontier asked at once, each with a recommendation, no guessing at unsettled answers.

## Example 2: Facts via sub-agent, decisions via user (don't block)

**User:** Answers Q1 with "specific hot queries" and Q2 with "self-hosted container."

**What the agent does:**

Settled decisions push the frontier outward. To prepare Q3 (TTL per query), the agent checks access logs and slow-query stats instead of asking the user which queries are hot. Asks the rest of the frontier now (eviction policy, invalidation trigger); only the TTL question waits for the evidence.

**Result:** Facts looked up, decisions put to the user, no blocking on exploration.

## Example 3: Session close (frontier empty, confirm before acting)

**User:** Answers the final round on invalidation triggers.

**What the agent does:**

Frontier is empty — every branch visited (scope, hosting, TTLs, eviction, invalidation). States the shared understanding back in one summary and asks for confirmation. Does not write any code or config until the user confirms.

**Result:** Nothing silently assumed, no implementation before confirmed understanding.

## Example 4: Resolve terminology and write it immediately

**User:** “The customer can cancel one item without cancelling the entire order.”

**What the agent does:** The glossary defines “Cancellation” for a whole Order, and the code currently agrees. The agent asks whether this change introduces a distinct “Line Cancellation,” recommending that the existing term retain its meaning. When the user agrees, the agent adds “Line Cancellation” to the applicable `GLOSSARY.md` before the next round.

**Result:** The glossary records the meaning of the term, not how cancellation is implemented. No ADR is written unless a later decision clears all three gates and the user agrees to record it.
