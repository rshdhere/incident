variable "kubeconfig_path" {
  description = "Path to the kubeconfig file for the target cluster."
  type        = string
  default     = "~/.kube/config"
}

variable "kube_context" {
  description = "Exact kubeconfig context name of the target cluster."
  type        = string
  nullable    = false

  validation {
    condition     = trimspace(var.kube_context) != ""
    error_message = "Supply the target kubeconfig context name."
  }
}

variable "namespace" {
  description = "Namespace for the incident-api Helm release."
  type        = string
  default     = "incident"
}

variable "image_repository" {
  description = "Container image repository accessible to the cluster."
  type        = string
  default     = "incident-api"
}

variable "image_tag" {
  description = "Tag of the API image to deploy."
  type        = string
  default     = "dev"
}
