# Persistence and migrations

The current workspace manifests and sources define no database client, schema, migrations or configured datastore. Check them again before starting a persistence change. An authored scenario and a learner's attempt have different owners and lifetimes. Identify which data the feature needs to retain and where it will be read.

When a persistence mechanism is chosen, record its schema owner and migration command alongside the implementation. For a schema change, identify old records and active readers, plan the order of compatible changes, and verify the result with representative existing data. Include recovery steps for changes that rewrite data.

Use the [database migration workflow](../skills/database-migration/SKILL.md) when a task involves stored data.
