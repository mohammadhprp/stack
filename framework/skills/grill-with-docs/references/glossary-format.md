# Glossary format

```markdown
# Ordering

Receives and tracks customer orders.

## Language

**Order**:
A customer's request for one or more items.
_Avoid_: Purchase, transaction

**Customer**:
A person or organization that places orders.
_Avoid_: Client, buyer, account
```

Pick one canonical term per concept and list rejected synonyms under `_Avoid_` when useful. Define what the thing *is* in one or two sentences. Include project-specific vocabulary, not general programming concepts. Group related terms under subheadings only when that makes the glossary easier to scan.

Most repositories use one root `GLOSSARY.md`. In a multi-context repository, the root `GLOSSARY-MAP.md` points to each context's glossary and can explain their relationships, for example:

```markdown
# Glossary Map

## Contexts

- [Ordering](./src/ordering/GLOSSARY.md): receives and tracks customer orders
- [Billing](./src/billing/GLOSSARY.md): generates invoices and processes payments

## Relationships

- **Ordering → Billing**: Ordering emits `OrderPlaced` to start billing.
```
