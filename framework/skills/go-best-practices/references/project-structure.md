# Go project and folder structure

Read this when starting, organizing, or reviewing a Go project. Start with the existing
repository and its instructions; these examples are options, not required templates.
Do not move files or change public import paths just to match a diagram.

## Choose the smallest layout that fits

A small library can keep its package at the module root:

```text
go.mod
parser.go
tests/
  parser_test.go
```

A small command can likewise use a root `main.go` and other `package main` files. Add
packages when they represent meaningful responsibilities, not for every type or layer.

A project with multiple commands or both a library and commands might grow into:

```text
go.mod
client.go                 # importable root package, if needed
tests/
  client_test.go
cmd/
  tool/
    main.go               # command entry point
internal/
  config/
    config.go             # implementation private to the parent tree
```

For a server, a common choice is `cmd/server/main.go` for startup and `internal/` for
application packages. Group code by coherent feature or responsibility, such as
`internal/orders/`, rather than creating empty `handlers/`, `services/`, `repositories/`,
or DDD layers by default. Place migrations, configuration examples, and deployment files
where the project and its tools expect them; only add directories with an actual use.

## Boundaries and naming

- A Go package is normally the Go files in one directory; a module contains packages
  rooted at its `go.mod`. Keep related code together and split when a separate package
  improves ownership, reuse, or dependency direction. Avoid import cycles.
- `internal` is enforced by the Go toolchain: only code in the tree rooted at the
  directory immediately above `internal` may import packages under it. It is not a
  substitute for access control or a reason to expose unstable public packages.
- Keep executable entry points focused on configuration, dependency wiring, lifecycle,
  and error reporting. Keep reusable logic in appropriately scoped packages when useful.
- Name packages for what they provide, not vague categories such as `util` or `common`.
  Do not introduce `pkg/` merely to denote public code: any non-`internal` package may
  be importable and its supported API should be intentional.
- For new projects, put `*_test.go` files under a separate `tests/` tree mirroring source
  package paths, using external test packages that import and exercise the real packages'
  exported APIs. Place fixtures under that tree's `testdata/` as appropriate. Go treats
  each test directory as a separate package, so tests there cannot access unexported
  identifiers; use public behavior rather than exporting internals just for tests.
  Preserve an existing repository's explicit test layout instead of moving its tests.
  Use `go test ./...` to discover the separate test packages and `-coverpkg` when measuring
  coverage of the source packages rather than only the test packages.
- Before reorganizing a published library, check imports, generated code, build scripts,
  `go:embed` paths, and consumers. Moving packages or changing a module path can break
  callers; preserve compatibility or get approval for a breaking change.

Sources: [Organizing a Go module](https://go.dev/doc/modules/layout),
[Go package names](https://go.dev/blog/package-names),
[internal directory rules](https://pkg.go.dev/cmd/go#hdr-Internal_Directories).
