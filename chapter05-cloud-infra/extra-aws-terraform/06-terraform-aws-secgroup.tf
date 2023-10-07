resource "aws_security_group" "initial-sg" {
  name        = "init-sg"
  description = "init-sg"
  vpc_id      = aws_vpc.vpc.id

  ingress {
    description      = "allow ssh from public"
    from_port        = 22
    to_port          = 22
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  ingress {
    description      = "allow HTTP from public"
    from_port        = 80
    to_port          = 80
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  ingress {
    description      = "allow HTTP from public"
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "init-sg"
    Zone = "public"
  }
}



data "aws_network_interface" "ec2_eni" {
  id = aws_instance.ec2.primary_network_interface_id

  depends_on = [ 
    aws_instance.ec2
  ]
}
