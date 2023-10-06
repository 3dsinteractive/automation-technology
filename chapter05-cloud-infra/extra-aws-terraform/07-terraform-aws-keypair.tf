locals {
  key_export_path = path.cwd
  key_name = aws_key_pair.keypair.key_name
  key_id = aws_key_pair.keypair.key_pair_id
  public_key_value = aws_key_pair.keypair.public_key
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "keypair"{
  key_name   = "keypair-1"
  public_key = tls_private_key.private_key.public_key_openssh

  provisioner "local-exec" {
    command = "echo '${tls_private_key.private_key.private_key_pem}' > ${local.key_export_path}/${aws_key_pair.keypair.key_name}.pem"
  }
}