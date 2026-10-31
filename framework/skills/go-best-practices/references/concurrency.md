# Concurrency and cancellation

Read this when code starts goroutines, shares state, uses channels, or shuts down workers.

## Lifecycle

- Start with a synchronous function. If concurrency is necessary, identify who starts and
  stops each goroutine and how the caller waits for completion. Cancellation requests
  stopping; it does not itself join workers or force blocking operations to finish.
- Propagate the request context into HTTP, database, and other operations that accept it.
  For derived contexts call the cancel function even on success. Use a fresh bounded
  context for graceful shutdown when the request or server context is already canceled.
- Make blocking channel sends and receives interruptible when downstream can stop early.
  Check that underlying I/O can actually be canceled. Bound concurrent operations using
  the existing worker pool or a suitable limit; account for memory and open connections.
- If using `errgroup`, understand that its derived context is canceled on the first error
  or when `Wait` returns. Collect results without unsynchronized writes and wait for all
  workers before reading shared output.

## Coordination

- A channel's sender or coordinating owner closes it only after all sends complete;
  receivers normally do not close channels they do not own. Closing is not required for
  every channel. Do not use `len(ch)` or a speculative buffer size as synchronization.
- Prefer a mutex for straightforward shared state; use channels when communication and
  ownership transfer fit. Keep locks scoped, avoid copying synchronization values and
  holding locks during unbounded I/O or callbacks, and protect all accesses to an
  invariant consistently. Atomics do not automatically protect compound invariants.
- Check Go version and loop form when reasoning about closures: Go 1.22 introduced
  per-iteration variables for loops declaring them under the relevant language version;
  variables assigned from an outer scope can still be shared.
- Test with explicit coordination, not sleeps used as evidence of ordering. Exercise
  early cancellation, a worker failure, stopped consumers, and shutdown. Run `-race`
  where supported; a clean run only covers executed interleavings.

Sources: [Go pipelines](https://go.dev/blog/pipelines), [context](https://pkg.go.dev/context),
[memory model](https://go.dev/ref/mem), [race detector](https://go.dev/doc/articles/race_detector),
[loop variable change](https://go.dev/blog/loopvar-preview).
