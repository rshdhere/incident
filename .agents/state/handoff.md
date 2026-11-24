# Current handoff

The web library lists live scenarios and links to each guided investigation. The [v1 API](../../packages/README.md) serves the [AWS SQS backlog exercise](../../.aws/sqs-backlog.json). Fetch `/api/v1/scenarios/sqs-backlog` for its brief and evidence. Fetch `/api/v1/scenarios/sqs-backlog/debrief` when the learner requests the diagnosis.

Read the [OpenAPI contract](../../packages/api/v1/openapi.json) for response fields, or fetch it from the running API at `/api/v1/openapi.json`. Use a response’s `X-Request-ID` to locate its entry in the server log.

The canvas under `apps/web/app/practice/[id]` presents evidence, session-only notes and an on-demand debrief. Verify the learner path against the running Go API in a browser. The API works directly from AWS fixtures. The optional [GORM model](../../packages/api/v1/persistence) provides a starting point for saving learner attempts when that feature is requested.

Use `task --list` for the backend commands. The catalog CI job runs `task check` for formatting, lint and tests.

Start the API with `go run ./packages/api/v1/cmd/server` from the root. See the [testing rules](../rules/testing.md) for the root suites and the [infrastructure guide](../rules/infrastructure.md) for container build paths.
