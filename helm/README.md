# Incident API chart

The chart runs the Go catalog service and exposes port 8080 inside the cluster. Run these commands from the repository root to check the manifests offline.

```sh
helm lint --strict helm/incident-api
helm template incident-api helm/incident-api --namespace incident
```

## Prepare the image

Build the service and the AWS exercise fixture into a local image.

```sh
docker build -t incident-api:dev -f docker/api.Dockerfile .
```

For a kind cluster named `practice`, load it with `kind load docker-image incident-api:dev --name practice`. For a remote cluster, push an image to its registry and set `image.repository` and `image.tag` to match. Use a distinct tag for each revision to trigger rollout and avoid reusing a cached image.

## Install and check

For direct Helm management, the following example targets the `kind-practice` kubeconfig context. Use the [Terraform configuration](../infra/README.md) if Terraform owns the release. Choose one release owner per namespace.

```sh
helm upgrade --install incident-api helm/incident-api \
  --kube-context kind-practice --namespace incident --create-namespace --wait
kubectl --context kind-practice -n incident rollout status deployment/incident-api
kubectl --context kind-practice -n incident port-forward service/incident-api 8080:8080
```

In another terminal, request the health endpoint and catalog.

```sh
curl -fsS http://localhost:8080/healthz
curl -fsS http://localhost:8080/api/v1/scenarios
```

For a Helm-managed release, run `helm history incident-api --kube-context kind-practice -n incident` to find the previous healthy revision. Restore it with `helm rollback incident-api REVISION --kube-context kind-practice -n incident --wait`.
