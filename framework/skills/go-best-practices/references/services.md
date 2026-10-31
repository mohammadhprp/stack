# Services, storage, and security

Read this for HTTP handlers, clients, persistence, migrations, or security-sensitive work.
Use an existing framework and its current documentation when the project already chose
one; neither a web framework, ORM, migration tool, nor DDD layout is mandatory.

## HTTP and service lifecycle

- Propagate `r.Context()` into downstream operations. Limit request bodies before
  decoding, validate syntax and domain constraints, and return appropriate status codes
  without disclosing internal errors. Define authentication and authorization at the
  trust boundary, not merely in the client.
- Configure suitable server read-header, read, write, and idle timeouts for the workload;
  streaming endpoints may need different settings. Configure client deadlines, reuse
  clients/transports, close response bodies, and bound response reads. Check HTTP status
  codes; successful transport does not imply a successful response.
- Graceful shutdown stops accepting traffic and waits for in-flight requests with a
  bounded fresh context. Separately stop and join background workers. Account for
  long-lived connections and deployment termination deadlines.
- Log actionable, structured context at the boundary where an error is handled. Avoid
  logging and returning the same error repeatedly, and redact credentials and personal
  data. Retrying requires bounded attempts, cancellation, backoff, and idempotent effects.

## Database operations

- Reuse `sql.DB` as a concurrent connection pool; tune limits and lifetimes for the
  database and deployment instead of copying fixed numbers. Parameterize input using
  the driver's placeholder syntax, and close rows and check `rows.Err()` after iteration.
- Use `BeginTx`, defer rollback after a successful begin, perform every transactional
  operation through `tx`, and check `Commit`. Results are not confirmed if commit fails;
  in distributed systems its outcome may be uncertain, so do not blindly retry writes.
  Choose isolation and locking to preserve the actual invariant under concurrent access.
- Make schema changes compatible with running application versions and existing data;
  review backfills, locks, rollback and data loss before execution. Never apply a
  migration to a real environment without the user's authorization.

## Security boundaries

- Validate at trust boundaries, limit work and input sizes, and use prepared or
  parameterized queries rather than constructing SQL from input. For user-selected paths,
  handle both traversal and symlink races using suitable confined filesystem APIs for the
  supported Go version; lexical cleaning alone is not confinement.
- Use `html/template` for HTML, appropriate output encoding for other contexts, and
  `crypto/rand` for secrets. Guard HTTP clients that fetch user-provided URLs against
  SSRF, including redirects and DNS/IP changes. Do not leak secrets via errors or logs.
- Scan with `govulncheck` when available; review dependency updates and toolchain
  compatibility instead of changing versions blindly. Follow existing security policy.

Sources: [net/http](https://pkg.go.dev/net/http), [database transactions](https://go.dev/doc/database/execute-transactions),
[SQL injection](https://go.dev/doc/database/sql-injection), [Go security](https://go.dev/doc/security/best-practices),
[filesystem confinement](https://go.dev/blog/osroot).
