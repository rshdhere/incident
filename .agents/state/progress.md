# Progress

| Milestone | Status | Evidence |
| --- | --- | --- |
| Project agent setup | Prepared | [Routing guide](../../AGENTS.md), [Codex roles](../../.codex/agents), [workflows](../skills) |
| AWS incident practice | Authored | [SQS backlog investigation](../../.aws/sqs-backlog.json) |
| Contributor tooling | Configured | [Husky hook](../../.husky/pre-commit) passed accept/reject checks. [CI workflow](../../.github/workflows/ci.yml) passed static validation. [Claude Code](../../.claude/CLAUDE.md) discovered 26 skills and five roles. |
| Catalog and infrastructure bootstrap | Implemented locally | The [root test suites](../../tests) exercised v1 listing, evidence and debrief, a [GORM migration retry](../../tests/integration/gorm_test.go), a short benchmark and a live end-to-end run of the AWS-only binary. [Helm](../../helm/README.md) rendered, [Terraform](../../infra/README.md) validated, and the [local Vault lab](../../vault/README.md) passed its reader checks. Container startup is covered by the new CI job and awaits a run on this revision. |
| API contract and operations | Implemented locally | [OpenAPI contract](../../packages/api/v1/openapi.json), [request log checks](../../tests/integration/request_log_test.go), and [server lifecycle checks](../../tests/e2e/server_lifecycle.py). Live startup and clean exits on SIGINT and SIGTERM verified. |
| Go developer tooling | Configured | Root [Taskfile](../../Taskfile.yml) and [golangci-lint configuration](../../.golangci.yml) supply the same Go checks locally and in CI. |
| Incident library and guided canvas | Planned | [Product brief](../context/project.md) |

When a milestone changes, record the delivered user behavior and the check that exercised it. Link the specific change or artifact when available.
