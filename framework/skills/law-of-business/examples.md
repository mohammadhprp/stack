# Law of Business Examples

Five worked scenarios showing routing, evidence, cited laws, and honest verdicts. In each, the agent maps the request to a reference, reads it in full, names specific laws and real cases, and refuses generic advice.

## Evaluating whether an idea is worth building

User: "I want to build an app that helps people split restaurant bills with friends. Is this a good business?"

Good agent behavior:
- Routes to [`references/market-and-demand.md`](references/market-and-demand.md) first, because the question is whether there is a business at all, not how to run one.
- Runs the fast triage and Part A of [`references/diagnostic-checklist.md`](references/diagnostic-checklist.md) before answering.
- States the honest verdict under the laws: a real but low-value problem, weak willingness to pay, and a distribution question, rather than a verdict based on the idea sounding nice.
- Cites laws by ID and name with cases: **A1. The Law of a Real Problem** and **Pumodo** (founders were not users of their own product), **A4. The Law of Willingness to Pay** and **Delight** (customers pay for information, not raw data), **A2. The Law of Market Size** and **Besomebody** (needed tens of millions of repeat bookers).
- Names the warning signs (users say "neat" but never pay; the problem is occasional and small) and a counter-move (get three paying users before building, price above cost to serve).
- Refuses to say "sounds great, go for it" without evidence, and does not pad the answer with generic startup advice.

## Diagnosing why a SaaS is stalling

User: "Our B2B SaaS grew fast for a year and then flattened. We add features constantly but churn is climbing. What's wrong?"

Good agent behavior:
- Routes to [`references/market-and-demand.md`](references/market-and-demand.md) and [`references/execution-and-operations.md`](references/execution-and-operations.md): stalling growth with rising churn can be a wrong-customer problem (A7), a distribution problem (A5), or a focus problem (C1).
- Reads both files in full, then runs the relevant checks in [`references/diagnostic-checklist.md`](references/diagnostic-checklist.md).
- Separates hypotheses with evidence: interview churned accounts; measure which acquisition channel actually retains; check whether features were built without customer input.
- Cites laws with cases: **A7. The Law of the Right Customer** and **VoterTide** (shipped features without client input) and **HelloParking** (rarely got out of the building); **C1. The Law of Focus** and **Intellibank** (features on features, unfocused).
- Does not accept "we need more features" as the diagnosis; treats it as the symptom the law is designed to expose.
- Ends with warning signs to watch (rising CAC, a single channel, no "no" list) and counter-moves (one ideal customer and channel per quarter; interview churned accounts this week) plus a re-check date.

## Raise a bridge or cut burn

User: "We have about five months of cash. Should we raise a bridge at a flat round or cut burn and extend runway?"

Good agent behavior:
- Routes to [`references/capital-and-economics.md`](references/capital-and-economics.md); this is B1 Runway, B7 Downward Spiral, and B4 Mismatched Backer territory.
- Reads the file in full before advising, then applies the scoring rule: five months is a red on **B1. The Law of Runway**, which is an automatic stop.
- Frames the real choice with cases: **Wow Air** ("ran out of time"), **Chef'd** (financing fell through and it closed immediately), **Patterbuzz** ("we ran out of money"), **Bridj** (bet on a deal instead of a round and the deal never closed).
- Separates the two decisions: whether the model deserves more capital at all (B2 Unit Economics, B5 Revenue Before Fundraising) before choosing the instrument.
- Notes the backer-mismatch risk of a flat bridge (B4: **Beepi** raised too much too soon; **Backplane**'s terms scared off new money) and refuses to treat a bridge as a plan.
- Gives a concrete recommendation with a decision rule (cut to a survivable base now, set the milestone the bridge would fund, and only raise against it) and a date to re-check.

## Should we pivot

User: "Growth is flat and the team keeps debating whether to pivot to a different customer segment. Should we pivot?"

Good agent behavior:
- Routes to [`references/execution-and-operations.md`](references/execution-and-operations.md) for **C7. The Law of Adaptation** and **C1. The Law of Focus**, and to [`references/market-and-demand.md`](references/market-and-demand.md) for A6 and A7.
- Reads the mapped files, then runs the checklist; distinguishes adaptation (responding to evidence, C7) from serial pivoting (C1).
- Asks for the evidence: what does the data say the current segment won't do, and what does the new segment actually demand? Treats a debate as a symptom of missing measurement.
- Cites laws with cases: **C7. The Law of Adaptation** and **Imercive** (stuck with the wrong strategy too long and pivoted only when capital was nearly gone), **Backchat** (could not adapt fast enough); **A6. The Law of Differentiation** and **Kato** ("Slack ate the world").
- Checks the pivot is fundable: does the company have the runway to execute a new plan (B1), or is this a decision made from the bottom of the runway?
- Ends with kill/iterate criteria and a review cadence, rather than a one-word "yes".

## An investor's diligence risk read

User: "We're considering investing in a quick-commerce startup doing 20-minute grocery delivery. What are the risks?"

Good agent behavior:
- Routes to [`references/capital-and-economics.md`](references/capital-and-economics.md) for **B2. The Law of Unit Economics** and [`references/execution-and-operations.md`](references/execution-and-operations.md) for **C4. The Law of Operational Reality**.
- Reads both files in full, then runs Part B and Part C of [`references/diagnostic-checklist.md`](references/diagnostic-checklist.md) and reports the reds by law.
- Names the known failure pattern with cases: **Fridge No More** (investors worried about bad order economics), **Miaoshenghou** (fresh-food margins of 10–20% couldn't cover rents), **Sprig** (owning production through delivery at scale was too complex), **Juno** (lost about $1M a day).
- Separates the company from the sector: the same category killed several companies, so the burden is on this team to show a different unit-economic structure, not just better execution.
- Uses [`references/case-index.md`](references/case-index.md) to trace any cited case back to its source title, so the diligence read is auditable.
- Closes with the specific reds found, the warning signs to monitor post-investment (CAC, contribution margin, ops cost per order), and the counter-moves the company must commit to.
