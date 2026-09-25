# Cross-Reference Destructive Command Index

A single-page index of destructive or irreversible Docker commands documented across this skill's reference files, so agents and reviewers can see the full picture without hunting through every reference. Each row's detail lives in the owning reference — this table only tracks what exists and where. Every command below requires explicit user confirmation before running, except the narrow Tier 1 exception in [`destructive-guardrails.md`](../destructive-guardrails.md) — see that reference's Core guidance for exactly which cases qualify.

| Command | What's lost | Owning reference |
|---|---|---|
| `docker rm` | Stopped container and its writable layer; low-risk removal (Tier 1 in the guardrail model) | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker rm -f` | Container removed without graceful shutdown; un-persisted in-container state | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker container prune` | All stopped containers on the host at once | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker kill` | Container killed via SIGKILL with no grace period; always Tier 2 | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker system prune` (esp. `-a`/`--volumes`) | Stopped containers, unused networks, dangling/all unused images, build cache, and (with `--volumes`) unused *anonymous* volume data | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker rmi` / `docker image rm` | A specific image | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker image prune -a` | All images not used by an existing container, including tagged ones | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker network rm` | A specifically named network's configuration | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker network prune` | All unused user-defined networks and their configuration | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker builder prune` (esp. `-a`) | Build cache; with `-a`, also internal helper/frontend images and cache shared with other build outputs | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker buildx rm` | A builder instance's configuration/state (not its build cache) | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker context rm` | Local context configuration (endpoint, TLS references) for a Docker host | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker volume rm` / `docker volume prune` (standalone, no Compose project in play) | Volume data, directly | [`destructive-guardrails.md`](../destructive-guardrails.md) |
| `docker compose down -v` / `docker compose down --volumes` | Named volumes and their data (e.g. database state) | [`compose-patterns.md`](../compose-patterns.md) |
| `docker volume rm` / `docker volume prune` (a Compose project's volumes) | Volume data, directly | [`compose-patterns.md`](../compose-patterns.md) |
| `docker compose rm -v` | Anonymous volumes attached to removed containers | [`compose-patterns.md`](../compose-patterns.md) |
| `sbx rm` / `sbx prune` | Sandbox containers, Git worktrees, state, and sandbox-scoped secrets; for a clone-mode sandbox, any unfetched commits too | [`sandboxes-lifecycle.md`](../sandboxes-lifecycle.md) |
| Docker Desktop destructive commands | Pending — see PR #14, not yet merged. Do not assume content until it ships. | *pending* |

## Notes

- This table is a routing aid, not a replacement for the owning reference's detail. Read the owning reference (`docker-cli-destructive-commands.md` here, or [`compose-patterns.md`](../compose-patterns.md)) before advising on or running any of these commands.
