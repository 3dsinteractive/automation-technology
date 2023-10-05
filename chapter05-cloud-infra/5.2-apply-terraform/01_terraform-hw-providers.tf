terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.38.1"
    }
    
    kubernetes = {
      source = "hashicorp/kubernetes"
      version = ">= 2.11.0"
    }

    tls = {
      source = "hashicorp/tls"
      version = ">= 3.4.0"
    }

    random = {
      source = "hashicorp/random"
      version = ">= 3.3.1"
      }
  }
}

provider "huaweicloud" {
  region     = "ap-southeast-2"
  access_key = "${local.resource_accesskey}"
  secret_key = "${local.resource_secretkey}"
}
