# Current handoff

The web app still renders the starter. The [v1 API](../../packages/README.md) serves the [AWS SQS backlog exercise](../../.aws/sqs-backlog.json). Fetch `/api/v1/scenarios/sqs-backlog` for its brief and evidence. Fetch `/api/v1/scenarios/sqs-backlog/debrief` when the learner requests the diagnosis.

Read the [OpenAPI contract](../../packages/api/v1/openapi.json) for response fields, or fetch it from the running API at `/api/v1/openapi.json`. Use a response’s `X-Request-ID` to locate its entry in the server log.

For the first canvas, start at `apps/web/app/page.tsx`, present the AWS evidence, and exercise both reads using the [feature workflow](../skills/implement-feature/SKILL.md). The API works directly from AWS fixtures. The optional [GORM model](../../packages/api/v1/persistence) provides a starting point for saving learner attempts when that feature is requested.

Start the API with `go run ./packages/api/v1/cmd/server` from the root. See the [testing rules](../rules/testing.md) for the root suites and the [infrastructure guide](../rules/infrastructure.md) for container build paths.
