# Current handoff

The learner app in `apps/web` still renders the Turborepo starter. The project setup establishes shared agent guidance for the first product task.

## Suggested first slice

When implementation is requested, use the [feature workflow](../skills/implement-feature/SKILL.md) to take one authored incident through the canvas. Use the existing [Kubernetes worker restart exercise](../../.k8s/worker-restarts.json). Present its brief, let the learner inspect the evidence and explain a diagnosis, then reveal the worked response.

Start at `apps/web/app/page.tsx`. Agree on the steps and evidence for the first scenario before generalizing the UI. Persistence, scoring and deployment are open design choices. See the [testing rules](../rules/testing.md) when selecting checks.
