# Verification

Use the scripts in the changed workspace's `package.json`. Both apps provide `bun run lint`, `bun run check-types` and `bun run build`. The UI package provides lint and type checks. An app type check runs Next route type generation first. Keep project test suites in the root `tests/` directory. Integration tests exercise the AWS v1 handlers and GORM migration in `tests/integration`. Run `go test ./packages/api/... ./tests/...` from the root module. Start the API for `tests/e2e/catalog.sh`. Use `go test ./tests/bench -run '^$' -bench .` for benchmarks. The [CI workflow](../../.github/workflows/ci.yml) also builds and probes the packaged catalog. The Helm and Terraform guides provide render and configuration checks.

For a learner-facing change, verify the affected route from its normal entry point. Pick one scenario and exercise the changed action, the expected next state and a relevant edge case. Record the action and observed outcome separately from lint and type-check results.

Follow the resource guidance in `AGENTS.md`. Check available RAM and swap before and during demanding work. Use Codex Cloud for builds, browser runs and heavy tests. Match the cloud checkout to the current commit and any working-tree changes involved in the check. If the cloud run cannot access those changes, report which checks remain unverified and continue lightweight work.
