# Infrastructure

The Go catalog has a [container build](../../docker/api.Dockerfile), a [Helm chart](../../helm/README.md) and a [Terraform root](../../infra/README.md) targeting an existing Kubernetes cluster. Confirm the kubeconfig context, release owner and image tag for deployment work. Keep Dockerfiles under `docker/`, use the repository root as the build context and maintain only the root `.dockerignore`. Explicitly allow authored AWS fixture paths and the source files needed by each image. The Vault exercise uses a disposable loopback server.

For build work, use the owning app's manifest, `bun.lock` and repository resource guidance. For deployment work, establish the destination, the exact revision, required settings, a health check and a recovery path. Record the resulting project-specific procedure in `docs/runbooks/` once confirmed.

Use the [deployment workflow](../skills/deploy/SKILL.md) for an authorized release or deployment plan.

The [CI workflow](../../.github/workflows/ci.yml) verifies pull requests and pushes to main. The `.aws` directory holds authored practice fixtures. See the [exercise catalog](../../README.md#incident-exercises).
