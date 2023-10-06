resource "aws_vpc" "vpc" {
  cidr_block = local.vpc_cidr
  enable_dns_hostnames = true

  tags = {
    Name = "vpc-1"
  }
}