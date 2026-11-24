---
name: orchestrate
description: Coordinate a multi-area incident project task across architecture, implementation, review, security and deployment roles.
---

# Coordinate a task

1. Name the requested result, affected workspace, and acceptance check. For multi-step changes, use `.agents/templates/implementation-plan.md`.
2. Assign independent questions to the available roles when delegation helps. Use `architect` for constraints and interfaces, `developer` for the owning implementation, `reviewer` for correctness, `security` for trust boundaries, and `devops` for build or delivery. Give each role the relevant paths and expected output. Assign one writer per file.
3. Sequence work that shares interfaces. For example, settle how an exercise identifies its evidence before connecting both the evidence list and canvas to it. Integrate the changes and resolve disagreements against the learner outcome.
4. Verify the integrated revision following `.agents/rules/testing.md`. Record completed work and checks in `.agents/state/progress.md`. When work spans sessions, use `.agents/templates/handoff.md` to refresh `.agents/state/handoff.md` with the next executable step.
