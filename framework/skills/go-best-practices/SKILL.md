---
name: go-best-practices
description: >-
  Write, review, debug, refactor, explain, or port code to idiomatic Go. Use whenever
  working with Go or Golang code, .go files, go.mod, Go APIs, errors, contexts,
  goroutines, channels, testing, benchmarks, tooling, folder structure, HTTP services,
  or databases.
---

# Go best practices

Use this skill for Go implementation, review, debugging, and migrations. Adapt the
guidance to the task: a CLI, a library, and a server have different requirements.
Explicit user instructions and repository conventions take precedence over style
preferences; correctness, security, and public behavior are not optional.

## Workflow

1. Read the affected code, callers, tests, `go.mod`, and project instructions. Note the
   supported Go version, public contracts, build tags, existing tools, and integrations.
2. Identify observable behavior and failure modes before changing code. For a bug, add a
   focused failing test when feasible. Preserve existing APIs, error identity, ordering,
   output, and data unless the request authorizes a change.
3. Prefer the smallest direct solution and existing project patterns. Use the standard
   library when it fits; do not introduce frameworks, dependencies, layers, interfaces,
   or abstractions without a concrete need.
4. Run relevant checks below, report what ran and any limitations. For reviews, report
   actionable findings with file and line, consequence, and a concrete fix; do not edit
   code unless asked. For ports, characterize and test the original behavior first, and
   report intentional differences.

## Language and API design

- Run `gofmt`; use short, descriptive package and identifier names with consistent
  initialisms (`HTTP`, `ID`). Avoid repeating the package name in exported names.
- Design around the zero value when useful. Prefer concrete types and small interfaces
  defined by consumers when a real boundary needs them, not interfaces added solely for
  mocking. Use generics only when they simplify genuinely shared type-safe logic.
- Use pointer receivers when mutation, synchronization fields, or shared identity
  require them. Never copy a mutex after first use. Slices and maps share backing data;
  document ownership or copy when callers need independence.
- Handle errors deliberately. Add useful operation context; use `%w` when callers should
  inspect the cause with `errors.Is` or `errors.As`. Wrapping exposes a dependency as
  part of the API, so translate or avoid wrapping private implementation details.
  Return errors for expected failures; do not `panic` or exit from library code.
- Prefer `var s []T` for an empty slice unless the contract requires otherwise. In
  particular, JSON distinguishes `null` from `[]`; do not change wire behavior casually.
  Never assume map iteration order is stable.
- Keep public APIs small and document non-obvious contracts, ownership, concurrency,
  cancellation, and errors in the project's documentation style. Name return values only
  where they clarify the signature. Use `internal` for non-public packages when useful.

## Concurrency, I/O, and security

- Prefer synchronous operations unless concurrency solves a real problem. For every
  goroutine, identify its owner, how it exits, how it is joined, and how errors propagate.
  Bound concurrency and resource use. A buffer alone is not a cancellation strategy.
- Pass `context.Context` explicitly as the first argument for cancellable work; propagate
  caller cancellation, call returned cancel functions, and avoid storing request contexts
  in long-lived structs. Do not use context values as optional parameters.
- Protect shared mutable state with synchronization; coordinate sends and channel close
  so no sender can send after closure. Test cancellation, failure, and shutdown. The race
  detector sees only paths exercised at runtime.
- Close resources and handle errors that can affect correctness, especially writes,
  flushes, commits, and close of output files. Avoid process-wide exit before cleanup.
- Validate untrusted input, bound request and response sizes, parameterize SQL, and never
  log secrets. Use `crypto/rand` for security-sensitive randomness. Do not assume a path
  is safe merely because it was cleaned; account for traversal and symlinks.

## Tests and verification

- Exercise observable success, failures, boundaries, and relevant concurrency. Use
  table-driven tests when cases share the same logic, meaningful subtest names, temporary
  directories, deterministic synchronization, and failure messages showing input, got,
  and want. Compare error semantics rather than fragile strings where appropriate.
- Follow existing test layout and test libraries. Coverage locates untested code, not
  proof of correctness; do not impose an arbitrary percentage. Use fuzzing for parsers
  and other untrusted-input boundaries when appropriate.
- For new projects, put Go tests in a separate `tests/` tree mirroring the source
  packages, using external test packages and public APIs. Follow an existing repository's
  explicit test layout instead; see `references/project-structure.md` for limitations.
- For performance changes, define a representative workload, verify correctness first,
  benchmark time and allocations, profile the bottleneck, and compare under the same
  conditions. Do not apply preallocation, pooling, or micro-optimizations by default.
- Baseline checks, adjusted to the project and scope:

  ```sh
  gofmt -l <changed-go-files>
  go test ./...
  go vet ./...
  ```

  Run `go test -race ./...` for relevant concurrency where supported; use project CI
  checks, bounded fuzzing, benchmarks, or `govulncheck ./...` when appropriate and
  available. Tests and tools can execute project code or access external services; inspect
  first and respect the environment. Report skipped checks rather than claiming success.

## Focused references

- `references/project-structure.md`: package boundaries and folder layouts for Go projects.
- `references/concurrency.md`: lifecycle, cancellation, channels, locks, and shutdown.
- `references/services.md`: HTTP, databases, security, and operational boundaries.
- `references/testing-performance.md`: tests, fuzzing, benchmarks, and review checklist.
- `references/sources.md`: source selection, project-specific tradeoffs, and links.
