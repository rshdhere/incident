# Persistence and migrations

GORM owns the v1 learner-attempt schema in [persistence](../../packages/api/v1/persistence). Run migrations explicitly on a caller-provided database. The catalog serves authored AWS fixtures independently of learner state.

Keep the scenario revision on each attempt so a later change to an exercise preserves the meaning of saved answers. Verify that a migration can be retried with existing records using the [integration test](../../tests/integration/gorm_test.go). SQLite is the test driver. Choose a driver and recovery procedure for the intended deployment when introducing persistent writes.

Use the [migration workflow](../skills/database-migration/SKILL.md) when a task changes stored data.
