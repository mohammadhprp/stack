# Domain modeling

Actively sharpen the project's vocabulary as the design takes shape:

- Challenge terms that conflict with the glossary. If it defines “cancellation” differently, ask which meaning applies.
- Propose a precise canonical term when a word is vague or overloaded. For example, distinguish a Customer from a User rather than calling both an “account.”
- Test relationships with concrete edge cases. If the user says partial cancellation is allowed but the code cancels only entire Orders, surface the contradiction instead of silently choosing a version.

## Glossary

Read `GLOSSARY-MAP.md` at the repository root if it exists. Use it to find the glossary for the context being discussed; ask if the context is unclear. Otherwise use the root `GLOSSARY.md`.

When the user settles a term, update the relevant glossary in that round, not at the end. Create the glossary only when there is a term to write. Use the [glossary format](glossary-format.md) and preserve unrelated entries. A glossary contains only domain vocabulary, not implementation details, specs, or scratch notes.

## Architecture decision records

Offer an ADR only if the decision is all three of:

1. **Hard to reverse:** changing course would cost meaningful time or effort.
2. **Surprising without context:** a future reader would reasonably ask why.
3. **A real trade-off:** there were genuine alternatives and reasons to prefer one.

If the user accepts, record the decision using the [ADR format](adr-format.md). Use root `docs/adr/` for system-wide decisions and the relevant context's `docs/adr/` for context-specific ones in a multi-context repo. Create the directory only when writing the first ADR. If any gate fails, leave the decision in the conversation.
