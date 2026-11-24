# GORM persistence

The v1 `Attempt` model stores the learner ID, scenario ID and revision, reasoning, and creation time for one attempt. A caller with an open GORM database runs `persistence.Migrate(db)` to create or update the table explicitly. The read-only catalog can start from its AWS fixture without opening a database.

Run the migration check from the repository root:

```sh
go test ./tests/integration -run '^TestAttemptMigration$'
```

The test uses the SQLite GORM driver and a temporary database. It creates the initial schema, saves an attempt, runs the migration again, and reads back the original record. SQLite needs CGO and a C compiler for this check. The deployment database driver and connection settings can be chosen when a service starts storing attempts.
