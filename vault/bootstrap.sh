#!/usr/bin/env bash
set -euo pipefail

: "${VAULT_TOKEN:?Set VAULT_TOKEN to the token printed by the local dev server}"
export VAULT_ADDR=http://127.0.0.1:18200
unset VAULT_AGENT_ADDR VAULT_PROXY_ADDR VAULT_NAMESPACE

if ! vault status -format=json | jq -e '.storage_type == "inmem" and .sealed == false' >/dev/null; then
  echo "Start the in-memory Vault dev server on 127.0.0.1:18200 first." >&2
  exit 1
fi

key=incident/practice/sqs-backlog
if ! vault kv get -mount=secret "$key" >/dev/null 2>&1; then
  vault kv put -cas=0 -mount=secret "$key" scenario=sqs-backlog
fi

vault policy write incident-reader "$(dirname "$0")/incident-reader.hcl"
