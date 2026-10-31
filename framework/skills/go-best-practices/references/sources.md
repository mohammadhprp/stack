# Sources and judgment

The skill is an independently written synthesis, researched 2026-10-31. Consult the
current documentation for exact APIs and behavior for the project's Go version.

## Requested sources

- [golang-base-practices-skills](https://github.com/cexll/golang-base-practices-skills/tree/26426d2b6ab9cb21d5cf5973a57a4406503573ee)
  covers error handling, concurrency, idioms, testing, services, databases and tooling.
  Its mandatory Gin/Kratos, GORM/Goose, DDD layout and 99% coverage are project choices,
  not universal requirements. Neither source repository displayed a root license; no
  source files or examples from them are bundled here.
- [practice-go](https://github.com/plutov/practice-go/tree/420c5bcce3067f90a5a906cc9896d408202800c4)
  is a programming challenge collection, not an authoritative style guide. Its emphasis
  on specifications, edge cases, tests, and comparable benchmarks informs the testing
  workflow; exercise-specific optimizations are not universal rules.

## Primary Go references

- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) and
  [Go Test Comments](https://go.dev/wiki/TestComments): idioms and test diagnostics.
- [Google Go Style](https://google.github.io/styleguide/go/guide),
  [style decisions](https://google.github.io/styleguide/go/decisions), and
  [best practices](https://google.github.io/styleguide/go/best-practices): detailed
  judgment and tradeoffs, not mandatory policy for other projects.
- [Effective Go](https://go.dev/doc/effective_go): useful core idioms, but its own
  introduction notes it is not actively maintained and does not cover modern Go fully.
- [Organizing a Go module](https://go.dev/doc/modules/layout): layout depends on the
  project; `cmd` and `internal` are useful conventions, not a required template.
- [Working with errors](https://go.dev/blog/go1.13-errors): wrapping exposes a cause
  to callers and is an API decision; use `errors.Is` and `errors.As` for chains.
- [Pipelines and cancellation](https://go.dev/blog/pipelines),
  [the memory model](https://go.dev/ref/mem), and
  [the race detector](https://go.dev/doc/articles/race_detector): concurrent design
  and its limits; older examples may predate modern context and loop semantics.
- [Security best practices](https://go.dev/doc/security/best-practices),
  [fuzzing](https://go.dev/doc/tutorial/fuzz), and
  [transactions](https://go.dev/doc/database/execute-transactions): concrete guidance
  for security, correctness and test strategy.

Repository instructions and established contracts win over optional style rules.
Verify third-party framework, ORM, and migration-tool APIs against their current docs.
