---
name: fix-bug
description: Diagnose and fix a defect in the incident app, shared UI, or scenario content using the affected learner journey.
---

# Fix a bug

1. Capture the route, starting state, learner action, expected behavior and observed result. Use `.agents/templates/bug-report.md` when collecting a report. Replace real service logs with a minimal synthetic example.
2. Start at the owning workspace. For a web interaction, trace the route under `apps/web/app` through component state and any imported `@repo/ui` component. For incorrect incident guidance, compare the hint and explanation to the evidence visible at that stage.
3. Use the existing `diagnosing-bugs` skill when the cause is unclear. Find the smallest check that demonstrates the defect, make the correction, then repeat that check with the same initial state. Exercise a neighboring case that should still work.
4. Apply `.agents/rules/testing.md` and report the cause, changed path and verification evidence. If the issue remains, record the smallest known reproduction and the observation needed to distinguish the remaining causes.
