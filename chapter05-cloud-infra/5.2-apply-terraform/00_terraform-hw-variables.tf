variable "region" {
  default     = "ap-southeast-2"
  description = "Huawei region"
}

variable "client" {
  description = "Client name (client-name)"
}

locals {
  resource_client = var.client
}

variable "hw_access_key" {
  description = "Huawei Account access key "
}

locals {
  resource_accesskey = var.hw_access_key
}

variable "hw_secret_key" {
  description = "Huawei Account secret key "
}

locals {
  resource_secretkey = var.hw_secret_key
}


