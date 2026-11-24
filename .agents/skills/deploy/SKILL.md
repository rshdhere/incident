---
name: deploy
description: Prepare or carry out a requested release of the incident web app and verify the deployed learner route.
---

# Deploy the web app

1. Read `.agents/rules/infrastructure.md` and inspect `docs/runbooks/` and the target workspace's configuration. Identify the hosting project, intended environment and revision. At setup the checkout provided build scripts but no deployment target. A request to prepare a release can end with a runbook. A release request needs a known destination.
2. Use the repository lockfile and select `apps/web` for the practice site. Follow the root compute policy for heavy checks. In the web workspace, `bun run build` produces the Next.js production build. Establish that the tested source includes the intended local changes.
3. For an authorized deployment, use the configured delivery workflow and capture its deployment identifier. Keep the preceding release identifier or the provider's recovery procedure available.
4. Open the deployed route and walk the feature named in the release. Check asset loading, the learner interaction, and server errors where logs are available. Record the URL, release identifier, revision, checks and any outstanding verification. When delivery fails, report the failed step and last serving release.
