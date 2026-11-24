provider "helm" {
  kubernetes = {
    config_path    = pathexpand(var.kubeconfig_path)
    config_context = var.kube_context
  }
}

locals {
  chart_dir = "${path.module}/../helm/incident-api"
  # Detect local template edits even when the chart path and version stay the same.
  chart_digest = sha256(jsonencode({
    for path in fileset(local.chart_dir, "**") : path => filesha256("${local.chart_dir}/${path}")
  }))
}

resource "helm_release" "api" {
  name             = "incident-api"
  chart            = local.chart_dir
  namespace        = var.namespace
  create_namespace = true
  atomic           = true
  description      = "Incident API chart ${local.chart_digest}"

  values = [yamlencode({
    image = {
      repository = var.image_repository
      tag        = var.image_tag
    }
  })]
}
