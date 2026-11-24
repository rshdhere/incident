# Working conventions

Start from the affected workspace manifest and its callers. The web starter colocates its route and CSS module under `apps/web/app`. The root layout imports global CSS and local Geist fonts. Shared components use named exports and import paths such as `@repo/ui/button`.

Use the app's existing TypeScript and ESLint presets. Shared settings originate in `tooling/typescript-config` and `tooling/eslint-config`. Compare both app configurations before changing a preset that affects them.

Use the repository `docs/` directories for product and engineering documentation. Use `.agents/state/` for the current handoff and verification evidence. Update these notes when the work changes the facts a later contributor will rely on.
