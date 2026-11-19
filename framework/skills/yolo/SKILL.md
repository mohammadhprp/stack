---
name: yolo
description: "Explicit session mode for taking a coding task from intake to delivery with user approval between phases. Use /yolo, /yolo <task>, /yolo status, /yolo off, or 'stop yolo'."
argument-hint: "[<task>|status|off]"
disable-model-invocation: true
license: MIT
---

# YOLO

Take a coding task from intake to delivery. Within an approved phase, do the work without asking for permission at every routine step. Between phases, stop so the user can verify the result.

## Session mode

- `/yolo` turns the mode on. Start intake if the task is already clear; otherwise ask for a ticket, spec, or description. If a task is paused at a checkpoint, show it without advancing.
- `/yolo <task>` turns the mode on and starts intake for that task. If a different task is unfinished, ask whether to pause it and switch; never discard its work.
- `/yolo status` reports whether the mode is on, the task, the current phase, the record path, and any pending approval. It does not activate the mode, advance a phase, or write files.
- `/yolo off` or “stop yolo” turns it off and pauses the task. Leave files, branches, and remote state untouched. Turning it back on in the same session restores the pending checkpoint, not permission to advance.
- After completion, stay on for the next task in the same session. Start new sessions off. Existing records are not activation, approval, or permission to resume. If the user explicitly resumes an older record, confirm the task and working tree and present the relevant checkpoint again before taking further action.

## Approval checkpoints

End each phase with its deliverable, record path, evidence or limitations, and one named next action. Ask for explicit approval of that next action and **end the turn**. Do not start the next phase in the same turn or infer approval from silence, feedback, a question, a new command, or advance permission given before the checkpoint was shown. A clear “yes” in direct response to the pending checkpoint is enough. If the user requests changes, revise and present the checkpoint again. Reopen the relevant gate if scope, approach, diff, evidence, or submission destination changes materially.

Read only the reference for the current phase, plus conditional references as needed. Resolve them relative to this `SKILL.md`. They are ordinary Markdown, not skills to invoke. Do not install plugins or change tool permissions.

| Phase | Work and checkpoint |
| --- | --- |
| 1. Understand the task | Read [task and plan](references/task-and-plan.md). Present the source, goal, scope, constraints, and acceptance criteria (AC). Ask to start solution design. No project edits. |
| 2. Choose a solution | Use [task and plan](references/task-and-plan.md), and [diagnose](references/diagnose.md) for hard bugs or uncertainty. Present the approach, affected areas, risks, and test plan. Ask to implement. No project edits. |
| 3. Implement | Use [implement and test](references/implement-and-test.md). Present the actual diff, focused results, and deviations. Ask to start final verification. Do not commit or publish. |
| 4. Test and verify | Use [verify and review](references/verify-and-review.md). Present fresh check results and evidence for every AC. Ask to prepare submission only if the task is verified. If a fix is needed, ask to return to implementation and verify the revised change again. |
| 5. Prepare and submit | Use [submit](references/submit.md). Present the exact proposed delivery action and destination, then wait for explicit approval before taking it. Report the actual result afterward. |

## Generated records

All YOLO workflow artifacts (intake, plan, implementation notes, verification evidence, submission draft and receipt, and any scratch reports or captures) live in `~/.stack/yolo/<project-slug>/<task-slug>/`. Actual implementation, tests, and requested project docs stay in the target project. Follow [the artifact rules](references/artifacts.md) before writing anything. If the record directory cannot be written safely, report the blocker rather than falling back to the repository.

See [examples](examples.md) for checkpoints and [upstream sources](references/upstream.md) for attribution.
