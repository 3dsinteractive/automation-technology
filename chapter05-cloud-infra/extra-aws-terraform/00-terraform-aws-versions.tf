
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.67.0"
    }

    random = {
      source  = "hashicorp/random"
      version = ">= 3.5.1"
    }

    local = {
      source  = "hashicorp/local"
      version = "~> 2.4.0"
    }

    null = {
      source  = "hashicorp/null"
      version = "~> 3.2.1"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.20.0"
    }
  }

  required_version = ">= v1.4.6"
}
