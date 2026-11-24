# Tools and commands

The root `package.json` declares Node.js 24 or newer, Bun 1.2.3 and the `apps/*`, `packages/*` and `tooling/*` workspaces. Use `bun.lock` for reproducible installs. Both app manifests currently pin Next.js 16.3.8, React 19.2.8 and TypeScript 7.0.2. Check the manifest and lockfile for the versions relevant to a change.

| From | Command | Purpose |
| --- | --- | --- |
| Root | `go run ./packages/api/v1/cmd/server` | Start the v1 catalog on port 8080. |
| Root | `go test ./packages/api/... ./tests/...` | Compile Go packages and run tests, including the GORM SQLite migration check. |
| Root | `bun install --frozen-lockfile` | Install dependencies from the recorded lockfile. |
| `apps/web` | `bun run dev` | Start the product app on port 3000. |
| `apps/docs` | `bun run dev` | Start the docs app on port 3001. |
| Either app | `bun run lint`, `bun run check-types` | Run ESLint, then separately generate Next route types and type-check. |
| `packages/ui` | `bun run lint`, `bun run check-types` | Check shared components. |
| Either app | `bun run build` | Produce the app's production build. |

Root scripts use Turborepo for workspace-wide development, lint, type checking and builds. Read the installed Turbo docs as directed by `AGENTS.md` before changing those commands or their configuration. Run heavy checks in Codex Cloud, following [the testing rules](../rules/testing.md).
