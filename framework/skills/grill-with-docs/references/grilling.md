# Interview workflow

Map the design as a tree: every decision branches into decisions that depend on it. Work in rounds. The **frontier** is the set of open decisions whose prerequisites are already settled. Ask the whole frontier in one round, number each question, and recommend an answer with its rationale. Then wait.

```text
❓ **Q1** - **<question title>**: <question and relevant options>

➡️ <recommended answer and why>

---

❓ **Q2** - **<question title>**: <question and relevant options>

➡️ <recommended answer and why>
```

Where possible, word the question so “yes” accepts the recommendation. Recompute the frontier after each reply. Save questions that depend on an unsettled answer for a later round.

Look up facts instead of asking the user for information you can find. Use the host's native subagent tool only if available and authorized; otherwise investigate directly. Do not substitute another provider or managed session without authorization. While a fact is still being checked, defer its dependent questions and ask the rest of the frontier. Put decisions, not lookups, to the user.

Finish when there are no unresolved branches in scope, and confirm the shared understanding before any implementation.
