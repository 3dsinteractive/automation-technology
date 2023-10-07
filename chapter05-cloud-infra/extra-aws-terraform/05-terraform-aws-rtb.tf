# Public Route Table
## Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
resource "aws_route_table" "public-rtb" {
  vpc_id = aws_vpc.vpc.id

  route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "public-rtb-1"
  }


  depends_on = [
    aws_internet_gateway.igw
  ]
}


# Subnets & Route Tables Assocation
# Subnets Associate with Public Route Table
resource "aws_route_table_association" "pub1" {
  subnet_id      = aws_subnet.public1.id
  route_table_id = aws_route_table.public-rtb.id


  depends_on = [
    aws_route_table.public-rtb
  ]
}
