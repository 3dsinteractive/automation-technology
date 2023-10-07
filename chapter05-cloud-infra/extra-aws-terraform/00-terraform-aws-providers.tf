
provider "aws" {
  region = var.region
  default_tags {
    tags = {
      Name        = "Provider Tag"
      Environment = "${var.env}"
      Region      = "${var.region}"
    }
  }
}