# Testing, performance, and review

Read this when designing tests, tuning performance, or reviewing Go changes.

## Tests

- Enumerate behavior before implementation: normal inputs, empty values, invalid inputs,
  boundaries, ordering and ties, and failure paths. Add regression cases for bugs. For
  algorithm work, consider input size, complexity, bytes versus Unicode, and overflow.
- Use table-driven subtests only when their checks are genuinely shared; separate tests
  when cases need different logic. `t.Helper()` makes helper failures point to callers;
  `t.TempDir()` and `t.Cleanup()` isolate resources. Respect repository test layout.
- Compare stable observable behavior, not incidental internal implementation. Use
  `errors.Is`/`errors.As` where error identity is contractual. Assert exact output bytes
  only when they are part of the contract, such as a file or wire format.
- Integration tests should isolate state and avoid real external services by default;
  when a real database or service is required, explicitly arrange an authorized test
  environment. Use `httptest` for local HTTP cases when suitable. Avoid changing process
  environment concurrently in parallel tests.
- For untrusted parsing, add useful seed inputs and properties to a fuzz test. A bounded
  example is `go test ./path/to/pkg -run=^$ -fuzz=^FuzzParse$ -fuzztime=30s`; the exact
  package and fuzz target must exist. Preserve minimized failing inputs as regressions.

## Measurement

- Measure a representative baseline and compare the candidate using the same inputs,
  machine, Go version, and settings; run more than once. For example,
  `go test ./path/to/pkg -run=^$ -bench=BenchmarkName -benchmem -count=5`.
- Profile CPU or heap only when the workload warrants it. Track realistic input sizes,
  throughput and tail latency where applicable, plus allocations; optimize the measured
  bottleneck. Use `benchstat` if already available, not as a mandatory dependency.
- Algorithmic improvements usually matter more than replacing formatting calls or
  preallocating blindly. Consider readability, contention, memory retention, and the
  cost of synchronization as well as benchmark numbers.

## Review checklist

Prioritize concrete defects: data loss, authorization and injection, wrong error or
transaction handling, goroutine leaks, races, broken contracts, and missing regression
tests. Trace callers and affected paths, including cancellation and cleanup. Then flag
API and maintainability issues only when they have a real cost. Include precise file and
line references, impact, and a minimal corrective direction. State checks not run and do
not pad the review with style preferences.

Sources: [Go test comments](https://go.dev/wiki/TestComments), [testing package](https://pkg.go.dev/testing),
[fuzzing tutorial](https://go.dev/doc/tutorial/fuzz), [profiling](https://go.dev/doc/diagnostics),
[practice-go](https://github.com/plutov/practice-go/tree/420c5bcce3067f90a5a906cc9896d408202800c4).
