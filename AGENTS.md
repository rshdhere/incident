# Incident engineering practice

Build the learner-facing incident library and guided investigation canvas in `apps/web`. The current practice catalog focuses on AWS incidents for DevOps and backend interviews. The repo currently contains a starter UI. Read the [project brief](.agents/context/project.md) and [handoff](.agents/state/handoff.md) before continuing a product task.

## Backend layout

Implement active endpoints under `packages/api/v1` and expose them at `/api/v1`. Keep v2 work in `packages/api/v2`. Use GORM for Go persistence. Keep shared React components in `packages/ui` and authored AWS exercises under `.aws`.

Build containers from the repository root. Keep every Dockerfile in `docker/` and maintain one root `.dockerignore`. Put integration tests in `tests/integration/`, end-to-end tests in `tests/e2e/` and benchmarks in `tests/bench/`. Run Go commands from the root module. See the [API guide](packages/README.md).

Use the root [Taskfile](taskfile.yaml) for Go development. CI runs `task check` for formatting, golangci-lint and tests. Use `task --list` for focused commands and follow the compute policy below when choosing where to run them.

Name all project-owned YAML files with the lowercase `.yaml` extension, including dotfiles and GitHub Actions workflows. Preserve upstream filenames in installed skills. Keep skills and agent configuration excluded from GitHub language statistics through [.gitattributes](.gitattributes).

## Find the right guidance

- For workspace ownership, entry points and tools, see the [architecture](.agents/context/architecture.md), [tech stack](.agents/context/tech-stack.md) and [conventions](.agents/context/conventions.md).
- For scenarios, evidence and learner decisions, use the [domain guide](.agents/context/domain.md).
- Apply the rules relevant to the change: [coding](.agents/rules/coding.md), [testing](.agents/rules/testing.md), [security](.agents/rules/security.md), [database](.agents/rules/database.md), [infrastructure](.agents/rules/infrastructure.md) and [Git](.agents/rules/git.md).
- Use [implement-feature](.agents/skills/implement-feature/SKILL.md) or [fix-bug](.agents/skills/fix-bug/SKILL.md) for product changes. For reviews, use [code-review](.agents/skills/code-review/SKILL.md) or [security-audit](.agents/skills/security-audit/SKILL.md).
- Use [database-migration](.agents/skills/database-migration/SKILL.md) for schema changes, [deploy](.agents/skills/deploy/SKILL.md) for releases and [orchestrate](.agents/skills/orchestrate/SKILL.md) for coordinated work. Role definitions live in [.codex/agents](.codex/agents).
- For an upstream issue search or contribution, use [oss-contribution](.agents/skills/oss-contribution/SKILL.md). Check `git remote -v` and the intended upstream's contribution guide before preparing a PR.
- [Templates](.agents/templates) cover plans, PRs, bug reports and handoffs. Record completed milestones in [progress](.agents/state/progress.md) and the next actionable step in the [handoff](.agents/state/handoff.md).

## Resource usage and heavy compute

Monitor available RAM and swap with `free -h` before and during tests, builds, browser runs, video conversion and model inference. Run heavy workloads in Codex Cloud by default. Keep lightweight edits, searches and focused checks local when memory permits. If available RAM falls, swap rises, or a job risks exhausting memory, stop only the task's own heavy processes and move the work to Cloud.

If Cloud is unavailable, report the access or setup blocker and continue lightweight work. Verify the cloud checkout includes the exact current changes, including uncommitted inputs, before reporting a cloud test result.

<!-- BEGIN:turborepo-agent-rules -->

# This is NOT the Turborepo you know

Turborepo configuration, task behavior, and CLI commands can vary between installed versions and may differ from your training data. Resolve the `turbo` package from this file's directory or relevant workspace; in monorepos, it may not be visible from the repository root. For example, run `node -p "require.resolve('turbo/package.json')"` from a workspace that depends on `turbo`.

Read `docs/README.md` inside that installed package first, then read the relevant pages from its `docs/` directory before changing Turborepo configuration or commands. Heed deprecation notices. These bundled docs match the installed package version and are available without network access.

This block is written and re-added by `turbo` before repository-scoped commands when an AI agent is detected. In the Turborepo source repository, its template is defined in `crates/turborepo-cli/src/cli/agent_guidance.rs`. Removing the managed block while updates are enabled means a later qualifying invocation will add it again. Set `"agentGuidance": true` in the root `turbo.json` or `turbo.jsonc` to opt in; this does not remove an existing block. Keep the block committed with your work to avoid an uncommitted change on the next agent invocation.
<!-- END:turborepo-agent-rules -->
