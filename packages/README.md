# Incident API

The Go module starts at the repository root. The [v1 server](api/v1/cmd/server) serves authored AWS incidents. [v2](api/v2) is reserved for the next API version.

Run from the repository root:

```sh
go run ./packages/api/v1/cmd/server
```

In a second terminal:

```sh
curl -fsS http://localhost:8080/api/v1/scenarios
curl -fsS http://localhost:8080/api/v1/scenarios/sqs-backlog
curl -fsS http://localhost:8080/api/v1/scenarios/sqs-backlog/debrief
```

The exercise response contains the brief, question and evidence. The debrief contains the worked diagnosis, mitigation and recovery check. The server defaults to port 8080. Set `PORT` to override it or `INCIDENT_ROOT` to point to a different fixture root. Probes use `GET /healthz`.

## API contract and request logs

Fetch `GET /api/v1/openapi.json` for the [OpenAPI contract](api/v1/openapi.json). It ships inside the binary, so clients can read the contract from the same release that serves the catalog.

Each response includes an `X-Request-ID`. Use it to find the matching JSON access log on stderr, including the HTTP method, route, status and duration. For example, issue a request with a recognizable ID and look up that ID in the server log:

```sh
curl -si -H 'X-Request-ID: practice-001' http://localhost:8080/api/v1/scenarios/sqs-backlog
```

## Server lifecycle

The server validates its environment before listening. Set timeout values as positive Go durations such as `250ms` or `10s`.

| Setting | Default | Applies to |
| --- | --- | --- |
| `HTTP_READ_TIMEOUT` | `10s` | Reading the request |
| `HTTP_WRITE_TIMEOUT` | `15s` | Writing the response |
| `HTTP_IDLE_TIMEOUT` | `1m` | Waiting for the next request on a keep-alive connection |
| `HTTP_SHUTDOWN_TIMEOUT` | `10s` | Draining active requests on SIGINT or SIGTERM |

When setting the shutdown timeout for Kubernetes, leave enough time within the pod's termination grace period for the process to exit.

These additions adapt the request correlation, server lifecycle and OpenAPI publishing patterns from [Sriniously’s Go boilerplate](https://github.com/sriniously/go-boilerplate/tree/ff3656aa83c74d037a10d52bd41b2b2fa171c310/apps/backend).

## Ownership

| Path | Purpose |
| --- | --- |
| [api/v1](api/v1) | HTTP handlers, catalog service, evidence types, content validation and [sample requests](api/v1/sandbox/catalog.http) |
| [api/v1/persistence](api/v1/persistence) | GORM model and explicit migration for learner attempts |
| [ui](ui) | Shared React components for the web and docs apps |
| [tests](../tests) | Integration tests, benchmarks and live HTTP end-to-end checks |
| [docker](../docker) | Service Dockerfiles built from the repository root |

To add an AWS exercise, register its authored JSON file in [the catalog](api/v1/services/catalog.go), include it in [the Dockerfile](../docker/api.Dockerfile) and [the root build-context allowlist](../.dockerignore), and extend the integration checks.

## Verification

From the repository root:

```sh
go test ./packages/api/... ./tests/...
go test ./tests/bench -run '^$' -bench BenchmarkScenarioDetail
tests/e2e/catalog.sh http://127.0.0.1:8080
```

Start the API before running the end-to-end check. The GORM integration test uses SQLite and requires CGO and a C compiler.

## Container

```sh
docker build -f docker/api.Dockerfile -t incident-api:dev .
docker run --rm --read-only -p 127.0.0.1:8080:8080 incident-api:dev
```

The image runs as UID 65532 and includes the AWS fixture. The [Helm chart](../helm/README.md) deploys the same image.
