# API Design Reference

Design stable resource contracts, validation, compatibility, and operational behavior.

## When to Use

Use this when the task involves creating or changing service contracts, resources, endpoints, pagination, filtering, sorting, versioning, authentication, authorization, validation, error handling, idempotency, rate limiting, or backward compatibility. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- Model contracts around resources and state transitions, not storage tables.
- Pagination, filtering, and sorting must be bounded, deterministic, and documented.
- Versioning and compatibility decisions must protect existing clients.
- Authentication proves identity; authorization decides allowed actions per resource.
- Validation errors should be precise enough for clients to fix requests without leaking sensitive internals.
- Idempotency is required for retries that create orders, payments, inventory reservations, or notifications.
- Rate limiting and abuse controls should fail predictably with useful retry guidance.

## Workflow

1. Identify the resource, action, and state transitions the endpoint models.
2. Define the request contract: path params, query params, headers, body shape, auth requirements.
3. Define the response contract: status codes, body shape, error format, pagination/filtering shape.
4. Check backward compatibility with existing clients and contracts.
5. Add validation at the boundary: type checks, constraint checks, auth checks.
6. Add idempotency handling for mutating endpoints where clients may retry.
7. Add observability: request-level metrics, structured logs for errors, trace context.
8. Verify with contract tests and integration tests for the boundary behavior.
9. Document the endpoint: examples, limits, errors, auth, rate limits.

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/api.md, references/standards/security.md.

## Deliverables

- Request and response contract spec.
- Validation rules with error examples.
- Backward compatibility analysis.
- Idempotency and rate limiting design.
- Contract and integration test plan.

## Common Mistakes

- Returning internal error details in production error responses.
- Designing endpoints that mirror database tables instead of business operations.
- Forgetting idempotency on payment, order, and inventory endpoints.
- Making pagination unstable by sorting on non-deterministic fields.
- Ignoring rate limit headers and retry-after guidance.

## Failure Modes

- A client cannot parse the error response because the shape is inconsistent.
- A breaking change ships without versioning, breaking existing integrations.
- An unbounded filter or sort causes a database or memory outage.
- Rate limits are enforced without informative headers, causing opaque failures.

## Checklist

- [ ] Request validation covers type, format, required fields, and allowed values.
- [ ] Error responses are structurally consistent across all endpoints.
- [ ] Idempotency key is accepted and enforced on mutating endpoints.
- [ ] Pagination is bounded, deterministic, and has a default and max limit.
- [ ] Rate limit headers (X-RateLimit-*) are present where applicable.
- [ ] Backward compatibility is verified or a versioning strategy is documented.
- [ ] Auth and authz are enforced at the boundary, not in the client.
