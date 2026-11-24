#!/usr/bin/env bash
set -euo pipefail

base_url=${1:-http://127.0.0.1:8080}
base_url=${base_url%/}
curl -fsS --max-time 5 --retry 10 --retry-connrefused --retry-delay 1 "$base_url/healthz" |
  jq -e '.status == "ok"' >/dev/null

scenarios="$base_url/api/v1/scenarios"
curl -fsS --max-time 5 "$base_url/api/v1/openapi.json" |
  jq -e '.openapi and .paths["/api/v1/scenarios/{id}/debrief"].get' >/dev/null
curl -fsS --max-time 5 "$scenarios" |
  jq -e 'all(.[]; .platform == "aws") and any(.[]; .id == "sqs-backlog")' >/dev/null
curl -fsS --max-time 5 "$scenarios/sqs-backlog" |
  jq -e '.id == "sqs-backlog" and (.question | length > 0) and (.evidence | length > 0) and (has("workedResponse") | not)' >/dev/null
curl -fsS --max-time 5 "$scenarios/sqs-backlog/debrief" |
  jq -e '[.diagnosis, .mitigation, .recovery] | all(.[]; type == "string" and length > 0)' >/dev/null
echo "SQS catalog, evidence and debrief passed."
