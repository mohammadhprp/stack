---
name: unslop
description: Cut AI tells from any writing. Must always apply.
disable-model-invocation: true
---

# Unslop

Edit text to remove AI patterns and add human voice.

## Process

1. Scan for the patterns below.
2. Rewrite. Preserve meaning, match intended tone.
3. Add soul (see next section).
4. Self-audit: "What makes this obviously AI generated?" Fix remaining tells.
5. Output the full edited draft plus a short **What changed** section listing the patterns fixed.
6. Run the checks in [eval.md](references/eval.md). If any check fails, fix the draft and re-check.

## Detect mode

When the user asks whether a piece is slop, or asks to audit, scan, or flag a draft without rewriting: name each pattern from this skill that appears, quote the line, and give the fix in a few words. Do not rewrite, score the draft, or guess whether AI wrote it. Offer to edit the draft after.

## Adding soul

Removing patterns is half the job. Sterile, voiceless writing is just as obvious.

- **Have opinions.** React to facts instead of neutrally listing pros and cons.
- **Vary rhythm.** Short sentences. Then longer ones that take their time. Mix it up.
- **Acknowledge complexity.** "Impressive but also kind of unsettling" beats "impressive."
- **Use "I" when it fits.** First person isn't unprofessional.
- **Let some mess in.** Perfect structure looks machine-made.
- **Be specific.** Not "this is concerning" but "there's something unsettling about agents churning away at 3am."

## Patterns to detect and fix

### Content

1. **Puffery.** "pivotal moment", "testament to", "evolving landscape", "setting the stage for", "indelible mark", "deeply rooted". Cut puffery, state what happened.
2. **Name-dropping.** Listing media outlets without context. Pick one, say what was said.
3. **Superficial -ing phrases.** "highlighting...", "ensuring...", "reflecting...", "showcasing...", "fostering...". Delete or expand with real sources.
4. **Promotional language.** "nestled", "vibrant", "breathtaking", "groundbreaking", "renowned", "stunning", "must-visit". Use neutral descriptions.
5. **Vague attributions.** "Experts believe", "Industry reports suggest", "Some critics argue". Name the source or delete.
6. **Formulaic challenges.** "Despite challenges... continues to thrive." Replace with specific facts.

### Language

7. **AI vocabulary.** Additionally, beacon, crucial, cutting-edge, delve, empower, embark, elevate, enduring, enhance, ever-evolving, fostering, game changer, garner, interplay, intricate, landscape (abstract), meticulous, multifaceted, paradigm shift, paramount, pivotal, realm, robust, showcase, streamline, supercharge, tapestry (abstract), testament, this changes everything, this is huge, transformative, underscore, vibrant. Replace with plain words.
8. **Fancy ways to say "is".** "serves as", "stands as", "boasts", "features". Just say "is" or "has".
9. **"Not just X, but Y."** State the point directly instead.
10. **Rule of three.** Forcing ideas into groups of three. Use the natural number.
11. **Synonym cycling.** Protagonist, main character, central figure, hero all in one paragraph. Pick one, repeat it.
12. **False ranges.** "from X to Y" where X and Y aren't on a meaningful scale. List topics directly.

### Style

13. **Em dash overuse.** Avoid em dashes entirely. Use periods or commas only (no parentheses, no en dashes, no hyphen-as-dash substitutes). Em dashes are an AI tell, and reaching for parentheses instead just trades one tell for another. If a thought needs separation, end the sentence or use a comma.
14. **Colon overuse.** Colons are fine before a list or example. Not as mid-sentence connectors. "If you're coming from traditional automation: instead of registering event handlers, you describe conditions" adds nothing with the colon. Rewrite to let the point stand on its own without comparison framing. "Describing when the scheduler should fire works best as plain English." Same meaning, no crutch punctuation.
15. **Boldface overuse.** Don't bold every proper noun or acronym.
16. **Inline-header lists.** The tell is a bold label and colon that restates the line: "**Performance:** Performance improved...". Convert those to prose. A bold lead-in that ends in a period, names the item, and is followed by genuinely new detail ("**Schema in TypeScript.** Tables live in one file.") is fine, not a tell.
17. **Title case headings.** Use sentence case.
18. **Decorative emojis.** Remove from headings and bullets.
19. **Curly quotes.** Replace with straight quotes.

### Communication artifacts

20. **Chatbot phrases.** "I hope this helps!", "Let me know if...", "Of course!", "Certainly!", "Found the smoking gun!" Remove.
21. **Cutoff disclaimers.** "While specific details are limited..." Find sources or remove.
22. **Sycophantic tone.** "Great question! You're absolutely right!" Respond directly.

### Filler

23. **Filler phrases.** "In order to" becomes "To". "Due to the fact that" becomes "Because". "It is important to note that" gets deleted. Often-empty phrases that delay the point: it's worth noting, at the end of the day, when it comes to, at its core, in today's world, in the age of, in the world of, the reality is, the truth is, in terms of, with regard to, going forward, in this article, let's dive in. Cut them unless part of the writer's recognizable voice.
24. **Excessive hedging.** "could potentially possibly be argued that it might" becomes "may". Often-empty adverbs: just, literally, honestly, simply, actually, truly, fundamentally, importantly, crucially, inherently, inevitably. Cut when they add nothing; keep when they carry emphasis, uncertainty, contrast, or spoken rhythm.
25. **Generic conclusions.** "The future looks bright." State specific plans or facts.

### Jargon

26. **Abstract metaphor nouns.** Substrate, wedge, vector, locus, vantage, nexus, primitive (as noun), harness (as metaphor), surface (as in "API surface"), bedrock, scaffolding (as metaphor), modality, paradigm, gold-plating, ratchet (as metaphor), evacuate (for moving code), endgame, north star, flywheel. These read as technical but usually have a plainer concrete word. "Substrate" becomes "base". "Wedge in" becomes "add". "Vector" becomes "way" or "method". "Gold-plating" becomes "more than the job needs". "Ratchet" becomes the mechanism's real name or "a limit that only tightens". "Evacuate" becomes "move out". "Endgame" becomes "the last phase". Pick the concrete word.

### Plain speech

27. **Say what it does, not how it feels.** "the database stays close at hand", "SQL you can read", "types that follow your schema" name a feeling. The fix names the mechanism or a number: "`.toSQL()` returns the exact string sent to the database", "a column rename fails the build". Ask what the sentence tells the reader to do or know, then write that. If you can't restate it as a concrete instruction, fact, or number, cut it. One more check: if the sentence could appear unchanged in another project's docs, it says nothing about this one. Cut it.
28. **Shorten or split dense sentences.** If the reader has to backtrack to parse a sentence, break it in two or drop clauses. One idea per sentence.
29. **Active voice.** Prefer it. Catch "is/are/was/were + past participle" and name the actor: "queries are validated" becomes "the compiler validates queries", "the file is parsed by the loader" becomes "the loader parses the file". Passive is fine only when the actor is unknown or genuinely doesn't matter.
30. **Cut adverbs, or use a stronger verb.** "runs quickly" becomes "is fast" or the number. "significantly improves" becomes the measured delta. An adverb propping up a weak verb means the verb is wrong.
31. **Prefer the plain word.** "utilize" becomes "use", "leverage" becomes "use", "facilitate" becomes "help", "numerous" becomes "many", "in the event that" becomes "if". The fancier synonym is rarely clearer.

### Voice preservation (from no-ai-slop)

32. **Minimum effective edit.** Fix AI patterns, errors, repetition, and unclear passages. Leave strong human sentences alone. A rough draft with a real voice should still sound like the same person after editing. Do not make every paragraph equally tidy.
33. **Show, don't label.** Cut commentary that labels a point important, surprising, subtle, or obvious instead of demonstrating why ("That last part matters more than it sounds", "The key point is", "As you can see", "This distinction matters", redundant "In other words"). If the prose already shows the point, delete the aside. Otherwise replace it with facts.
34. **Protect the specific fact.** Don't smooth a useful detail into generic importance. "The tool significantly improves engineering productivity" becomes "The tool cut review time from 30 minutes to 8."
35. **Portability test.** If a sentence could move unchanged to another person, company, country, or product, it is filler. Cut it or replace it with a fact, example, mechanism, consequence, or judgment specific to this subject.

### Dramatic setups and endings (from no-ai-slop)

36. **Throat-clearing openers.** "Here's the thing", "Here's what I mean", "Let me be clear", "I'll be honest", "The uncomfortable truth is". Cut and state the point. Keep a personal aside only when it creates context, tension, or character.
37. **Faux-insight setups.** "What nobody tells you", "What most people get wrong", "The part everyone misses", "This is the part most people skip". These flatter the writer as the lone expert. Cut the setup; make the claim stand on its own.
38. **Binary contrasts.** "It's not X. It's Y.", "The question isn't X, it's Y." State Y directly. "The question isn't the model. It's the eval." becomes "The eval matters more than the model." (See also 9.)
39. **Negative listing.** "Not a X. Not a Y. A Z." Just say Z.
40. **Rhetorical setups.** "What if I told you...", "Think about it:", "Plot twist:", self-answered "Question? Answer." pairs. Drop them and make the point.
41. **Colon reveals.** A noun phrase, a colon, then a lowercase dramatic reveal: "The best part: it learns." Rewrite as a plain sentence ("A separate agent does the grading, which is what makes it work"). Colons are for lists, labels, and quotes, not fake drama. (See also 14.)
42. **Dramatic fragmentation.** "That's it. That's the whole thing.", "X. And Y. And Z." Use complete sentences unless the fragment is clearly the writer's own cadence.
43. **Robotic rhythm.** Repeated sentence shapes, identical paragraph structures, stacked punchy fragments. Vary the shape only when it helps the point.
44. **Fake-profound kickers.** Cut the final "deep" line when it turns the point into a metaphor, aphorism, or mic-drop ("The future isn't coming. It's already here."). Do not rewrite it into a better metaphor. Delete it, then end on the clearest concrete sentence already in the draft.
45. **Summary-recap endings.** "In conclusion", "Ultimately", "Overall", or a final paragraph restating the piece. The reader was just there. End on the last concrete point, takeaway, or next action instead. (See also 25.)
46. **Formatting slop.** Bullet lists where two sentences of prose would read better, headers over two-sentence sections. Format follows the content, not decorates it. (See also 15, 18.)

## Source

Patterns 32–46, detect mode, and `references/eval.md` adapted from [no-ai-slop](https://github.com/petergyang/no-ai-slop) by Peter Yang, MIT License.
