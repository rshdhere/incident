---
name: database-migration
description: Plan, implement or review a persistence change for incident scenarios and learner attempts when the task involves a database.
---

# Change a database schema

1. Read `.agents/rules/database.md` and inspect the GORM models and migration entry point under `packages/api/v1/persistence`. Identify the target driver, owner and persistence behavior. Use `tests/integration/gorm_test.go` as the SQLite regression check.
2. Describe the old and new record shape and how existing readers handle each. For example, if an attempt holds a scenario version, preserve enough information to interpret it after the scenario changes.
3. Use the chosen tool's versioned migrations. Where a rollout spans versions, stage compatible schema changes, backfill existing records in restartable batches as needed, and retire old fields after their readers have moved.
4. Verify against both a fresh disposable database and one with representative existing records. Check upgrade, retry behavior, application reads and recovery. Record the actual command, target environment, row-level expectations and recovery procedure before any requested rollout.
