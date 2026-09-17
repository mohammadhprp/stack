# Database Design Reference

Design data models, constraints, transactions, indexes, and migrations.

## When to Use

Use this when the task involves modeling entities such as users, products, orders, payments, inventory, notifications, and reports; changing relationships, constraints, transactions, locking, indexes, normalization, denormalization, migrations, or integrity rules. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- Data models must encode business invariants through constraints whenever possible.
- Transactions should cover the full consistency boundary and no more.
- Locking strategy must be explicit for concurrent updates to orders, inventory, balances, or quotas.
- Indexes should match real access patterns and be reviewed for write cost.
- Normalize for integrity first; denormalize only with ownership and refresh rules.
- Migrations must define backfill, verification, compatibility window, and rollback or mitigation.

## Workflow

1. Understand the entities, relationships, cardinalities, and access patterns.
2. Model the schema with constraints at the database level: PKs, FKs, unique, check, not null.
3. Define indexes based on real query patterns, not guessed columns.
4. Design transactions around consistency boundaries, not convenience.
5. Plan the migration: forward SQL, rollback SQL, data backfill, verification query.
6. Consider locking strategy: optimistic vs pessimistic, deadlock risk, lock duration.
7. Add observability: query timing, connection pool metrics, migration progress.
8. Verify with integration tests that exercise constraints and concurrent access.
9. Document the data model: entity relationships, invariants, migration plan.

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/database.md, references/standards/performance.md.

## Deliverables

- Schema with entity relationships, constraints, and indexes.
- Migration plan with forward, rollback, backfill, and verification steps.
- Transaction and locking strategy.
- Integration test plan covering constraints and concurrent access.
- Data model documentation.

## Common Mistakes

- Skipping database-level constraints and relying only on application validation.
- Adding indexes without analyzing query plans and access patterns.
- Running large migrations without backfill strategy or rollback plan.
- Using transactions that are too broad (hold locks too long) or too narrow (miss consistency boundaries).
- Normalizing past the point of practical query performance without measuring.

## Failure Modes

- A migration locks production tables for minutes because it runs ALTER TABLE on large data.
- Data becomes inconsistent because a constraint was enforced only in the application layer.
- A query that passed staging tests performs poorly at production scale due to missing index.
- A rollback is impossible because the migration is irreversible.

## Checklist

- [ ] Database-level constraints enforce entity invariants (PK, FK, unique, check, not null).
- [ ] Indexes match real access patterns, verified with query plans.
- [ ] Migrations are reversible with a tested rollback script.
- [ ] Transaction boundaries match the consistency boundary.
- [ ] Locking strategy is explicit and deadlock risk is evaluated.
- [ ] Migration has a data backfill and verification step for existing rows.
- [ ] Connection pool limits and query timeouts are configured.
