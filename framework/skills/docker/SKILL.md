---
name: docker
description: "Use this skill for any Docker work — writing, reviewing, or optimizing Dockerfiles (multi-stage builds, layer caching, BuildKit secrets, non-root users, image size), Dockerizing or scaffolding a project for the first time (.dockerignore, Dockerfile, compose.yaml), and wiring or debugging Docker Compose services (health checks, dependency ordering, volumes, networks, development overrides). Use it before any destructive Docker command — docker rm, docker system prune, docker image prune, docker volume rm, sbx rm and the like — even when the user just says 'clean up', 'wipe everything', 'start fresh', or 'force remove'. It also covers Docker Agent (cagent): agent.yaml configuration of agents, models/providers, toolsets and sub_agents; running agents with docker agent run (--safety, --sandbox, aliases, worktrees); and serving, sharing, or evaluating them (serve, share, eval). And it covers standalone Docker Sandboxes (sbx): the sandbox lifecycle, declarative sbxenv.yaml environments, spec.yaml kits, network policy, and credentials."
license: Apache-2.0
---

# Docker

Docker guidance for this repository, organized as an index of reference files. Each reference file explains what to do and why. It spans first-time Dockerization, image builds, Compose, destructive-command safety, Docker Agent, and Docker Sandboxes.

## How to apply

1. Identify which area(s) the task touches: project foundations, image builds, Compose, destructive-command safety, Docker Agent, or Docker Sandboxes.
2. Map each concern to the reference index below and read every mapped reference file before editing. Skip the unrelated ones.
3. Read the matching file under `checks/` for the verification steps, and run the project's own build, test, and lint commands.
4. Make the smallest coherent change that satisfies the task. Match the layout and naming already in the project instead of introducing a second pattern.
5. Never bake credentials into an image, a Compose file, an `agent.yaml`, or a checked-in `sbxenv.yaml`/kit — use BuildKit secret mounts, `env_file`, provider environment variables, or the sandbox credential store.
6. Before running anything destructive, state exactly what will be lost and get explicit confirmation. See [`references/destructive-guardrails.md`](references/destructive-guardrails.md).

## When to use this skill

Activate this skill when any of the following applies.

**Foundations & builds**

- A user asks to set up, initialize, or Dockerize a project, or the project needs an initial `Dockerfile`, `compose.yaml`, or `.dockerignore`.
- A user wants to add a service dependency (database, cache, message queue) to a project, or to run/develop it locally without host-level installs.
- A user writes, reviews, or optimizes a Dockerfile, or says the image is too large, the build is slow, or the container needs hardening for production. Covers multi-stage builds, layer caching, `.dockerignore`, non-root users, and image size.

**Compose**

- A user creates or edits `compose.yaml`, adds or changes services, sets up a development override, or debugs service startup ordering or connectivity. Plain-language triggers: "wire the services together", "add a database to my stack", "set up local development with multiple containers".

**Safety**

- A user asks to "clean up", "clear the cache", "start fresh", "wipe everything", "nuke it", "reset", "force remove", or "tear down" Docker resources without naming a command.
- The agent is considering `docker rm`, `docker rm -f`, `docker container prune`, `docker kill`, `docker system prune`, `docker rmi`, `docker image prune -a`, `docker network rm`, `docker network prune`, `docker builder prune`, `docker buildx rm`, `docker context rm`, `docker volume rm`/`docker volume prune`, or `sbx rm`/`sbx prune` as a fix for an unrelated problem.
- A cross-product overview of destructive Docker commands is needed.

**Docker Agent**

- A user creates, edits, or reviews an `agent.yaml`/`agent.yml`/`agent.hcl`: agents, models/providers, built-in or MCP toolsets, and multi-agent teams with `sub_agents`. Plain-language triggers: "build an AI agent with Docker", "make a coding agent config", "add a tool to my agent", "set up a team of agents".
- A user runs an agent (`docker agent run`), chooses or debugs a `--safety`/`--yolo`/approval mode, isolates it with `--sandbox`, sets up an alias or worktree, or debugs credentials with `docker agent doctor`. Plain-language triggers: "run my agent", "make my agent auto-approve tool calls", "run this agent safely", "why can't my agent see my API key".
- A user exposes an agent as a server (`serve mcp/api/a2a/acp/chat`), distributes it via an OCI registry (`share push/pull`), or measures it with `docker agent eval`. Plain-language triggers: "turn my agent into an MCP server", "let Claude Desktop use my agent", "publish my agent to Docker Hub", "test my agent in CI".

**Docker Sandboxes**

- A user starts, reattaches to, lists, stops, or removes a local `sbx` sandbox; wants a `--clone` workspace or extra read-only workspaces; copies files, publishes a port, or runs `sbx exec`; or cleans up stopped sandboxes. Plain-language triggers: "run claude in a sandbox", "isolate an agent from my repo", "give an agent its own git clone", "clean up old sandboxes".
- A user authors, plans, or runs a declarative `sbxenv.yaml` environment (`sbx env create/run/plan/exec/rm`), needs host lifecycle commands tied to a sandbox, or parameterizes a shared environment with `args`. Plain-language triggers: "check in a sandbox config", "make onboarding reproducible", "run a setup script before the agent starts", "define arguments for a shared sandbox environment".
- A user authors, validates, packages, signs, or composes a kit `spec.yaml` (`sbx kit add/inspect/pack/pull/push/sign/validate/verify`). Plain-language triggers: "add a tool to a sandbox agent", "build a reusable sandbox extension", "publish a kit to a registry", "give a mixin its own credentials and network access".
- A user configures what a sandbox can reach on the network or which credentials it authenticates with (`sbx policy`, `sbx secret`). Plain-language triggers: "let the agent call an internal API", "block all network access", "give the agent a GitHub token", "use a private registry image for a sandbox".

## Do not use this skill when

- The operation is non-Docker (git, filesystem, cloud resources) — this skill covers Docker work only.
- The user explicitly wants to avoid Docker and no container, image, Compose, Docker Agent, or sandbox work is involved.
- The task is Docker Cloud Sandboxes (`sbx --cloud`) — out of scope; this skill covers the local sandbox daemon only.

## References Index

Every reference file below starts from the container/agent/sandbox semantics its area owns. Cross-cutting tasks often need more than one.

| Concern | Read |
| --- | --- |
| **Foundations & builds** | |
| Scaffolding a project, first-time Dockerization, `.dockerignore`/`Dockerfile`/`compose.yaml` layout, Dockerized dependencies | [`references/project-foundations.md`](references/project-foundations.md) |
| Multi-stage builds, layer caching, BuildKit secrets and SSH mounts, non-root users, image size | [`references/build-strategies.md`](references/build-strategies.md) |
| **Compose** | |
| Service definitions, health checks, dependency ordering, volumes, networks, environment variables, development overrides, Compose debugging | [`references/compose-patterns.md`](references/compose-patterns.md) |
| **Safety** | |
| Destructive `docker`/`sbx` commands, confirmation rules, and the cross-reference index of which file owns each command | [`references/destructive-guardrails.md`](references/destructive-guardrails.md) |
| **Docker Agent** | |
| `agent.yaml` — agents, models/providers, toolsets, `sub_agents`, commands | [`references/agent-config.md`](references/agent-config.md) |
| Running agents — `docker agent run`, `--safety`, `--sandbox`, aliases, worktrees, `doctor` | [`references/agent-run.md`](references/agent-run.md) |
| Serving, sharing, and evaluating — `serve`, `share`, `eval` | [`references/agent-deploy.md`](references/agent-deploy.md) |
| **Docker Sandboxes** | |
| Sandbox lifecycle — `sbx run/create/ls/stop/rm/prune`, workspace bind mount vs `--clone`, `exec`/`cp`/`ports` | [`references/sandboxes-lifecycle.md`](references/sandboxes-lifecycle.md) |
| `sbxenv.yaml` — declarative environments, host lifecycle hooks, args, merge layers, `secrets`/`registries`/`bindings` | [`references/sandboxes-env.md`](references/sandboxes-env.md) |
| Kits (`spec.yaml`, schemaVersion 2) — authoring, validation, composition, and distribution | [`references/sandboxes-kits.md`](references/sandboxes-kits.md) |
| Network policy and credentials — `sbx policy`, `sbx secret`, registry pull credentials, deny-over-allow precedence | [`references/sandboxes-network-credentials.md`](references/sandboxes-network-credentials.md) |

## Assets

Copy-ready examples, grouped by the reference that uses them.

- Foundations: `assets/project-foundations/Dockerfile.simple`, `assets/project-foundations/dockerignore-example`, `assets/project-foundations/compose-dev.yaml`.
- Builds: `assets/build-strategies/Dockerfile.golang`, `assets/build-strategies/Dockerfile.nodejs`, `assets/build-strategies/Dockerfile.python`, `assets/build-strategies/dockerignore-example`.
- Compose: `assets/compose-patterns/compose-web-app.yaml`, `assets/compose-patterns/compose-dev-override.yaml`, `assets/compose-patterns/bad-vs-good.md`.
- Docker Agent: `assets/agent-config/team-agent.yaml`, `assets/agent-deploy/eval-session-example.json`.
- Docker Sandboxes: `assets/sandboxes-env/sbxenv.yaml`, `assets/sandboxes-kits/spec-sandbox.yaml`, `assets/sandboxes-kits/spec-mixin.yaml`.

## Scripts

Bundled verification scripts, run from the skill root (or by the path shown) with `--help` for usage and exit-status details:

- `bash scripts/project-foundations/verify-setup.sh [--help]` — checks `.dockerignore`, `Dockerfile`, and `compose.yaml`, then validates the Compose configuration.
- `bash scripts/build-strategies/verify-build.sh [--help] [IMAGE_NAME]` — builds the image and reports its size and configured user.
- `bash scripts/compose-patterns/verify-compose.sh [--help]` — validates `compose.yaml` with `docker compose config --quiet`.

## Checks

Manual verification runbooks and checklists, one per reference:

- [`checks/project-foundations.md`](checks/project-foundations.md)
- [`checks/build-strategies.md`](checks/build-strategies.md)
- [`checks/compose-patterns.md`](checks/compose-patterns.md)
- [`checks/destructive-guardrails.md`](checks/destructive-guardrails.md)
- [`checks/agent-config.md`](checks/agent-config.md)
- [`checks/agent-run.md`](checks/agent-run.md)
- [`checks/agent-deploy.md`](checks/agent-deploy.md)
- [`checks/sandboxes-lifecycle.md`](checks/sandboxes-lifecycle.md)
- [`checks/sandboxes-env.md`](checks/sandboxes-env.md)
- [`checks/sandboxes-kits.md`](checks/sandboxes-kits.md)
- [`checks/sandboxes-network-credentials.md`](checks/sandboxes-network-credentials.md)
