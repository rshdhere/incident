# Infrastructure

The app manifests and `turbo.json` define local development and build tasks. The starter contains deployment links but supplies no project deployment pipeline, AWS stack or Kubernetes manifests. The exercise topics describe the learning material. Hosting remains an open project decision.

For build work, use the owning app's manifest, `bun.lock` and repository resource guidance. For deployment work, establish the destination, the exact revision, required settings, a health check and a recovery path. Record the resulting project-specific procedure in `docs/runbooks/` once confirmed.

Use the [deployment workflow](../skills/deploy/SKILL.md) for an authorized release or deployment plan.

The [CI workflow](../../.github/workflows/ci.yml) verifies pull requests and pushes to main. The cloud-provider directories hold practice fixtures. See the [exercise catalog](../../README.md#incident-exercises).
