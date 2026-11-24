# Workspace map

| Location | Responsibility |
| --- | --- |
| `apps/web/app` | Product routes, layouts, styles and bundled fonts. The starter home is `page.tsx`. |
| `apps/web/public` | Static assets served at the web app's root. |
| `apps/docs` | Separate Next.js starter app with its own routes, assets and dev server. |
| `packages/ui/src` | Shared React components consumed as `@repo/ui/<component>`. |
| `tooling/typescript-config`, `tooling/eslint-config` | Shared compiler and lint settings used by workspace consumers. |
| `packages/api/v1`, `packages/api/v2` | Active Go catalog and its GORM persistence model under v1. The v2 directory is reserved. |
| `docker/`, `.dockerignore` | All service Dockerfiles and the single root build-context filter. |
| `tests/integration`, `tests/e2e`, `tests/bench` | Integration suites, end-to-end checks and benchmarks. |
| `.aws` | Authored AWS interview scenarios with synthetic evidence and worked responses. The web starter does not load these files yet. |
| `docs/` | Repository documentation, including decisions and runbooks. |

A request to `/` on the web app renders `apps/web/app/page.tsx` inside `apps/web/app/layout.tsx`. For a trace across workspace boundaries, the starter page imports `@repo/ui/button`. The exports map in `packages/ui/package.json` resolves it to `packages/ui/src/button.tsx`, whose click handler displays an alert in the browser. Both apps consume that package.

Next.js generates route types and build output beneath each app's `.next/` directory. The build task's cached outputs are specified in `turbo.json`. The [Go API](../../packages/README.md) loads registered JSON fixtures at startup. The catalog route `GET /api/v1/scenarios` lists summaries. The scenario route returns evidence and the separate debrief route returns the worked response. The Next.js page is still a starter. Serving catalog reads requires only fixture files. GORM migrations are explicit operations against a caller-provided database.

The [Helm chart](../../helm/README.md) and [Terraform module](../../infra/README.md) deploy the same API image to an existing Kubernetes cluster. The optional [Vault lab](../../vault/README.md) runs independently.
