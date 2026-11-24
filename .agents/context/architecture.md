# Workspace map

| Location | Responsibility |
| --- | --- |
| `apps/web/app` | Product routes, layouts, styles and bundled fonts. The starter home is `page.tsx`. |
| `apps/web/public` | Static assets served at the web app's root. |
| `apps/docs` | Separate Next.js starter app with its own routes, assets and dev server. |
| `packages/ui/src` | Shared React components consumed as `@repo/ui/<component>`. |
| `tooling/typescript-config`, `tooling/eslint-config` | Shared compiler and lint settings used by workspace consumers. |
| `.aws`, `.azure`, `.gcp`, `.k8s` | Authored interview scenarios with synthetic evidence and worked responses. The web starter does not load these files yet. |
| `docs/` | Repository documentation, including decisions and runbooks. |

A request to `/` on the web app renders `apps/web/app/page.tsx` inside `apps/web/app/layout.tsx`. For a trace across workspace boundaries, the starter page imports `@repo/ui/button`. The exports map in `packages/ui/package.json` resolves it to `packages/ui/src/button.tsx`, whose click handler displays an alert in the browser. Both apps consume that package.

Next.js generates route types and build output beneath each app's `.next/` directory. The build task's cached outputs are specified in `turbo.json`. The starter has no incident data service, API routes, database, or background workers. Revisit this map as a first feature establishes those boundaries.
