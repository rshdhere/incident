---
name: implement-feature
description: Implement a learner-facing feature in the incident practice site, from an observable scenario outcome through focused verification.
---

# Implement a feature

1. Read `.agents/context/project.md` and `.agents/context/domain.md`. Turn the request into one learner action and one observable result. For example, a learner compares SQS queue depth with the Lambda mapping state to explain a backlog.
2. Trace the entry route under `apps/web/app`. Use the v1 API under `packages/api/v1` for the catalog and register its authored fixtures from `.aws`. Check existing exports in `packages/ui/src` before adding a shared primitive.
3. Implement the smallest complete path to the result. For a guided exercise, make the starting evidence sufficient to reason from and reveal the explanation at the point the task specifies. For canvas controls, verify keyboard use and access to the same evidence on a narrow screen.
4. Follow `.agents/rules/testing.md`. From `apps/web`, use `bun run lint` and `bun run check-types` as relevant. Exercise the actual learner path and record the result alongside the revision tested. Run heavy checks in Codex Cloud under the root resource policy.
