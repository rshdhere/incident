#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."
if ! (cd apps/web && bun -e 'require.resolve("next/package.json")') 2>/dev/null; then
  bun install
fi

# Give each service a process group so shutdown reaches its child servers too.
set -m
groups=()
stop_services() {
  for pid in "${groups[@]}"; do
    kill -TERM -- "-$pid" 2>/dev/null || true
  done
  wait || true
}
trap stop_services EXIT
trap 'exit 130' INT
trap 'exit 143' TERM

go run -mod=readonly ./packages/api/v1/cmd/server &
groups+=("$!")
(cd apps/web && exec bun run dev) &
groups+=("$!")

wait -n
