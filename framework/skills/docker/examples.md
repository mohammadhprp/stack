# Docker Examples

## First-time Dockerization

User: "I just cloned this Node repo. Set it up so I can run it locally without installing Postgres on my Mac."

Good agent behavior:

- Read [`references/project-foundations.md`](references/project-foundations.md) and [`references/compose-patterns.md`](references/compose-patterns.md) before writing anything.
- Create all three files: a `.dockerignore` first (so the build context is small and no `.npmrc` leaks), a multi-stage `Dockerfile`, and a `compose.yaml` that defines Postgres as a service rather than suggesting `brew install postgres`.
- Read [`references/build-strategies.md`](references/build-strategies.md) for the Dockerfile: `.dockerignore`, pinned base tags, dependency-manifest-first layer order, a non-root `USER`, and `# syntax=docker/dockerfile:1`.
- Keep the first scaffold simple, bind published app ports to loopback, and add a `healthcheck` to Postgres so the app waits for readiness.
- Hand back the single command to run, and point to `checks/project-foundations.md` for the verification checklist.

## Dockerfile optimization review

User: "My Go image is 1.2 GB. Can you shrink it?"

Good agent behavior:

- Load [`references/build-strategies.md`](references/build-strategies.md), then read `references/build-strategies/multi-stage-builds.md` and `references/build-strategies/layer-caching.md`.
- Identify the missing multi-stage build and the build toolchain still present in the runtime image; propose a distroless or scratch runtime stage that copies only the compiled binary.
- Check credentials handling: confirm nothing is passed via `ARG`/`ENV` and that private module fetches use `RUN --mount=type=ssh` or `--mount=type=secret`.
- Confirm the runtime stage runs as non-root and that the base tag is pinned, not `latest`.
- Suggest the smallest diff and run `bash scripts/build-strategies/verify-build.sh` to report the new size and configured user before declaring success.

## Compose multi-service debugging

User: "`docker compose up` starts my API before Postgres is ready and it crashes. Also the DB port is exposed to my whole network."

Good agent behavior:

- Load [`references/compose-patterns.md`](references/compose-patterns.md) and `references/compose-patterns/service-dependencies.md`.
- Fix startup ordering with `depends_on` + `condition: service_healthy`, and add a `healthcheck` using `pg_isready` on the database (a health condition requires one).
- Take the published datastore port off `0.0.0.0` and bind it to loopback, since only local host tools need it.
- Validate without printing secrets: run `bash scripts/compose-patterns/verify-compose.sh` (or `docker compose config --quiet`), never a plain `docker compose config` that would render interpolated credentials into the transcript.
- Read `checks/compose-patterns.md` and walk its checklist for health checks, dependency ordering, and secret handling.

## Destructive command that must be confirmed first

User: "Docker is eating my disk. Just clean everything up."

Good agent behavior:

- Load [`references/destructive-guardrails.md`](references/destructive-guardrails.md) before proposing any command.
- Inspect first with `docker system df`, then lay out the options by risk: `docker container prune` and `docker image prune` are low risk; `docker image prune -a` removes tagged images; `docker system prune -a --volumes` also deletes unused anonymous volumes.
- State exactly what each option would delete and ask the user to pick — do not run anything yet.
- Never reach for `docker system prune -a --volumes` or `docker volume rm` as a reflex; named volumes such as database data need their own confirmation.
- If the request is Compose-scoped (`docker compose down -v`), route to [`references/compose-patterns.md`](references/compose-patterns.md) instead, which owns that guidance.

## Docker Agent configuration and run

User: "Make me a coding agent with Docker that can read files and run shell commands, and run it safely in CI."

Good agent behavior:

- Load [`references/agent-config.md`](references/agent-config.md) and write a valid `agent.yaml` with `model`, `description`, `instruction`, and `toolsets` (`filesystem`, `shell`).
- Load [`references/agent-run.md`](references/agent-run.md) for the run flags; for unattended CI use `--safety restricted` so unreviewed tool calls fail closed rather than auto-approving.
- Never hardcode a provider API key in the config; point at the provider's environment variable and suggest `docker agent setup`.
- Verify with `docker agent doctor ./agent.yaml` and `docker agent debug toolsets ./agent.yaml` before the real run, per `checks/agent-config.md` and `checks/agent-run.md`.

## Docker Sandboxes environment

User: "I want a check-in-able sandbox config that clones my repo before the agent starts, and I don't want the agent touching my source tree."

Good agent behavior:

- Load [`references/sandboxes-env.md`](references/sandboxes-env.md) and author an `sbxenv.yaml` with `schemaVersion: "1"`, an `agent:`, and a `lifecycle.initialize` command that clones the repo idempotently (it reruns on every attach).
- Point at [`references/sandboxes-lifecycle.md`](references/sandboxes-lifecycle.md) for the `--clone`/bind-mount isolation model and its preconditions.
- Keep secrets out of the checked-in file: show the `secrets:`/`registries:` fragments only if needed, using `ref:`/`command:` rather than a literal `value:`.
- Hand back `sbx env plan` so the user can review every host command and mount before approving `sbx env create`.
- For network access or credentials beyond the file, route to [`references/sandboxes-network-credentials.md`](references/sandboxes-network-credentials.md), and never mount a credentials file read-only into a sandbox.
