# Deploy the incident API with Terraform

This root module installs the [local Helm chart](../helm/incident-api) on an existing Kubernetes cluster. The kubeconfig context is a required input. Terraform tracks the Helm release. Helm creates its target namespace if missing and leaves the namespace in place when the release is removed.

## Validate offline after initialization

From the repository root, download the provider pinned in the lockfile, then check the configuration.

```sh
terraform -chdir=infra init
terraform -chdir=infra fmt -check
terraform -chdir=infra validate
```

## Plan a release

Follow the [image preparation steps](../helm/README.md#prepare-the-image) first. The example uses the default `incident-api:dev` image in a kind cluster named `practice`. Set `TF_VAR_image_repository` and `TF_VAR_image_tag` for a registry-hosted image. Override `TF_VAR_kubeconfig_path` when your context is in a file other than `~/.kube/config`.

```sh
export TF_VAR_kube_context=kind-practice
export TF_VAR_namespace=incident
terraform -chdir=infra plan -out=release.tfplan
```

Review the context and plan before running `terraform -chdir=infra apply release.tfplan`. Check the release using the rollout and port-forward commands in the [chart guide](../helm/README.md#install-and-check). A failed install or upgrade uses Helm's atomic recovery. To restore an earlier successful release, plan and apply its chart revision and image tag.

The default state file is local to this module. Keep a separate state for each target cluster. For a shared environment, configure its remote state backend before the first apply.
