# Local Vault exercise

Practice granting a responder read access to one KV v2 record for the AWS SQS backlog scenario. The Go catalog serves its public fixtures independently of Vault. This exercise uses Bash, the Vault CLI and jq.

## Start a disposable server

In a terminal, run the following command. Vault prints a temporary root token for this session. Stopping the process clears its in-memory data. [Vault dev server documentation](https://developer.hashicorp.com/vault/docs/concepts/dev-server)

```sh
vault server -dev -dev-listen-address=127.0.0.1:18200 -dev-no-store-token
```

In a second Bash terminal at the repository root, read the temporary token at the hidden prompt. The bootstrap uses the dev server's default `secret/` KV v2 mount, seeds a fictional record and installs the reader policy. Running it again preserves the existing version.

```bash
export VAULT_ADDR=http://127.0.0.1:18200
unset VAULT_AGENT_ADDR VAULT_PROXY_ADDR VAULT_NAMESPACE
read -rsp 'Local dev root token: ' VAULT_TOKEN
echo
export VAULT_TOKEN
./vault/bootstrap.sh
```

## Verify the reader

Mint a five-minute token and read the one allowed field. The first read should print `sqs-backlog`. The second read should return HTTP 403 because the policy grants access only to the SQS backlog record.

```bash
reader_token=$(vault token create -policy=incident-reader -no-default-policy -ttl=5m -field=token)
VAULT_TOKEN="$reader_token" vault kv get -mount=secret -field=scenario incident/practice/sqs-backlog
VAULT_TOKEN="$reader_token" vault read secret/data/incident/practice/another-scenario
unset reader_token VAULT_TOKEN
```

Stop the dev server with Ctrl-C. The policy targets the KV v2 API's data path. [KV v2 setup and policies](https://developer.hashicorp.com/vault/docs/secrets/kv/kv-v2/setup)
