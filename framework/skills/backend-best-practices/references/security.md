# Security Reference

Apply practical security review across identity, access, secrets, data handling, and abuse resistance.

## When to Use

Use this when the task involves authentication, authorization, secrets, encryption, OWASP-style risks, injection, input validation, output encoding, rate limiting, abuse prevention, or least privilege. It is also useful when a request is vague, risky, touches production behavior, changes contracts, changes data, or needs a reviewable engineering plan. Do not use it for trivial text edits unless the edit changes engineering guidance.

## Principles

- Correctness and data integrity come before speed of implementation.
- Simplicity is a feature: fewer moving parts means fewer failure modes.
- Existing contracts must remain compatible unless a breaking change is approved.
- Every important decision should have a reason, an alternative considered, and an operational consequence.
- Work should be testable, observable, deployable, and reversible.
- Prefer explicit boundaries, clear names, and local reasoning over clever shared abstractions.
- Security and privacy are design inputs, not final review steps.

- Treat all input as untrusted, including internal calls, background jobs, files, and message payloads.
- Enforce least privilege for users, services, tokens, storage, and operational access.
- Secrets must never appear in source, logs, errors, metrics, traces, or examples.
- Use encryption deliberately for data in transit and sensitive data at rest.
- Prevent injection by separating commands from data and validating allowed shapes.
- Output encoding and safe error responses prevent data disclosure.
- Abuse controls such as rate limits should protect expensive and sensitive actions.

## Workflow

1. Identify data assets, entry points, and trust boundaries in the system.
2. Review authentication: is identity verified at every entry point?
3. Review authorization: is access checked against the resource, not assumed from the caller?
4. Review input handling: are all inputs validated, sanitized, or parameterized at every boundary?
5. Review secrets: are any credentials, tokens, or keys exposed in code, logs, configs, or errors?
6. Review data protection: is encryption applied for data in transit and sensitive data at rest?
7. Review abuse controls: rate limiting, payload limits, request throttling for expensive actions.
8. Test security boundaries: attempt injection, forced browsing, privilege escalation.
9. Document residual risks and recommended mitigations.

## Rules

- Never assume hidden requirements, traffic scale, compliance needs, or data retention rules.
- Do not introduce new infrastructure unless the current requirement cannot be met safely without it.
- Do not hide breaking changes in refactors.
- Do not weaken authorization, validation, transaction safety, or error handling to make implementation easier.
- Keep public contracts, migrations, and operational changes explicit in the deliverable.
- Reference related standards: references/standards/security.md.

## Deliverables

- Security review of authentication, authorization, and input handling.
- Secrets exposure analysis.
- Data protection review (in transit and at rest).
- Abuse control assessment.
- Residual risks and recommended mitigations.

## Common Mistakes

- Performing authorization checks only in the UI or frontend and not at the API boundary.
- Storing secrets in source code, environment files committed to git, or inline configs.
- Trusting input from internal services, background jobs, or message queues without validation.
- Using weak or outdated cryptographic algorithms.
- Ignoring rate limiting on authentication endpoints, allowing brute force.

## Failure Modes

- A vulnerability is introduced because input validation is missing on a new endpoint.
- A secret is exposed in logs or error responses, causing a credential leak.
- Authorization is checked on the first request but not re-verified for subsequent operations in a workflow.
- A security review misses an issue because it focused on code and ignored infrastructure config.

## Checklist

- [ ] Authentication is enforced at every external entry point.
- [ ] Authorization checks happen at the resource boundary, not just in the UI.
- [ ] All inputs are validated, sanitized, or parameterized.
- [ ] Secrets are never in source code, logs, errors, or example configs.
- [ ] Encryption is applied for data in transit (TLS) and sensitive data at rest.
- [ ] Rate limiting protects authentication and expensive endpoints.
- [ ] Output encoding prevents injection in responses to other systems.
- [ ] Dependencies are scanned for known vulnerabilities.
