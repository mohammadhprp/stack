# Docker Sandboxes kits (v3)

Consult the [upstream v3 kit workflow](https://github.com/docker/skills/tree/caeee2024f7527d890a29dd1732b62a112f56d0e/skills/docker-sandboxes-kits) for the current grammar and examples. Verify the installed `sbx` supports v3 before using it. The upstream workflow describes v3 as experimental; do not assume v2 was removed or every environment can run v3.

- A v3 kit is an OCI image with a descriptor in a manifest annotation. Author it in `<kit>.yaml` starting with `# syntax=docker/sandbox-kit:3` and `schemaVersion: "3"`, usually alongside `<kit>.dockerfile`.
- Use `workload` for a runnable kit, `mixin` for a composable extension, and `set` to group kits. Declare what the kit provides, what it requires, and the specific host capabilities it requests. Request only what it needs.
- Keep credentials out of descriptors and layers. Use capability requests and the runtime's credential store. Check network and mount requirements explicitly instead of porting v2 permissions.
- Build with Docker Buildx. Check conformance with `kit-tck validate`, inspect the kit with `sbx kit inspect` against an authorized daemon, and verify composition with a disposable `sbx run`. `sbx kit validate` does not replace conformance or a composition test. Report checks you could not run.
- For `sbxenv.yaml`, verify the current `agent:` and `kits:` schema and pin published kit references where reproducibility matters. Get explicit approval before publishing, signing, or mutating shared resources.

Do not mix v2 keys or commands into a v3 descriptor. Keep existing v2 files unchanged unless the user asks for a migration; write a separate descriptor and confirm parity before switching.
