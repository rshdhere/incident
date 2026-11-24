# Verification

Use the scripts in the changed workspace's `package.json`. Both apps provide `bun run lint`, `bun run check-types` and `bun run build`. The UI package provides lint and type checks. An app type check runs Next route type generation first. The starter defines no automated test script.

For a learner-facing change, verify the affected route from its normal entry point. Pick one scenario and exercise the changed action, the expected next state and a relevant edge case. Record the action and observed outcome separately from lint and type-check results.

Follow the resource guidance in `AGENTS.md`. Check available RAM and swap before and during demanding work. Use Codex Cloud for builds, browser runs and heavy tests. Match the cloud checkout to the current commit and any working-tree changes involved in the check. If the cloud run cannot access those changes, report which checks remain unverified and continue lightweight work.
