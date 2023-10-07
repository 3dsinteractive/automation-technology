# Public Subnet
resource "aws_subnet" "public1" {
  vpc_id     = aws_vpc.vpc.id
  cidr_block = local.public_subnet1
  availability_zone = "${var.region}a"
  tags = {
    Name = "public-subnet-01"
    Zone = "public"
  }
}
