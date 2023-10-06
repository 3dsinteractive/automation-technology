variable "region" {
  default     = "ap-southeast-1"
  description = "AWS region"
}

variable "env" {
  description = "Production, Test, Develop:"
}

locals {
  vpc_cidr = "10.0.0.0/16"
}

locals {
  public_subnet1  = "10.0.65.0/24"
}
