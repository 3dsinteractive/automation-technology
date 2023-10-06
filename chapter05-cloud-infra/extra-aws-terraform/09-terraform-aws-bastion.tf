locals {
  node_type = "t2.small"
}

resource "aws_instance" "ec2" {
  ami                    = data.aws_ami.ubuntu.image_id
  instance_type          = local.node_type
  key_name               = aws_key_pair.keypair.key_name
  vpc_security_group_ids = [aws_security_group.initial-sg.id]
  subnet_id              = aws_subnet.public1.id
  associate_public_ip_address = true

  ebs_block_device {
    device_name = "/dev/sda1"
    volume_size = 50
  }

  provisioner "file" {
    source = pathexpand("~/.aws")
    destination = "/home/ubuntu"

    connection {
      type        = "ssh"
      host        = self.public_ip
      user        = "ubuntu"
      private_key = file("${path.cwd}/${aws_instance.ec2.key_name}.pem")
    }
  }
}

resource "aws_eip_association" "eip_asso" {
  instance_id   = aws_instance.ec2.id
  allocation_id = aws_eip.eip.id
}