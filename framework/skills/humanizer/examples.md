# Humanizer Examples

## Humanize a pasted paragraph

User: "Make this sound less like AI: 'The annual festival stands as a testament to the town's vibrant cultural heritage, showcasing a rich tapestry of music, art, and culinary traditions that symbolize its enduring legacy.'"

Good agent behavior:

- Run the draft-to-audit-to-final loop internally.
- Preserve every fact, cut the significance framing and the rule of three, and avoid copula-substituting verbs like "stands as" and "showcases."
- Deliver the draft, the "still-AI" bullets, and the final rewrite without any em dashes.

```text
The festival runs every July. It includes a music stage, an art market, and food stalls run by local restaurants.
```

## Edit a file in place

User: "Humanize the prose in docs/onboarding.md but leave the code blocks and frontmatter alone."

Good agent behavior:

- Read the file, then run the loop internally rather than pasting the whole rewrite back.
- Touch only the prose: skip code blocks, frontmatter, data, and link targets.
- Rewrite the file in place so it contains only the final rewrite.
- Report a short summary of what changed, with the before-and-after for the worst offenders.

## Match the author's voice

User: "I wrote this post, sample included. It uses em dashes and short fragments a lot. Keep my voice, just cut the AI tells."

Good agent behavior:

- Read the sample first and note its sentence lengths, vocabulary, punctuation, and transitions.
- Match the author's habits instead of only deleting AI patterns: keep em dashes at roughly the sample's frequency rather than banning them.
- Do not upgrade casual words or regularize deliberate quirks.

## Avoid fabricating facts

User: "Make this sound natural, but I only have these details: 'The museum opened in 1972 and holds roughly 4,000 objects.'"

Good agent behavior:

- Rewrite freely but never add a fact, name, number, date, or citation not in the source.
- If a sentence needs real-world detail to work, write the plain version without it or ask for the detail.
- Deliver a rewrite that is more natural without inventing claims about the museum's architect or visitors.
