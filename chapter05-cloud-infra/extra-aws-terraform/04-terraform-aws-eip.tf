resource "aws_eip" "eip" {
    tags = {
      "Name" = "eip-1"
    }


    depends_on = [
      aws_internet_gateway.igw
    ]
}