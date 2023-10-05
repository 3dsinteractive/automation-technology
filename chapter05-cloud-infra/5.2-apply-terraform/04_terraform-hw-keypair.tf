resource "tls_private_key" "hw_private_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "huaweicloud_compute_keypair" "hw_keypair"{
  name   = "${local.resource_client}-keypair"
  public_key = tls_private_key.hw_private_key.public_key_openssh

  provisioner "local-exec" {
    command = "echo '${tls_private_key.hw_private_key.private_key_pem}' > ${local.key_export_path}/${huaweicloud_compute_keypair.hw_keypair.name}.pem"
  }
}

locals {
  key_export_path = "./keypair"
  key_name = huaweicloud_compute_keypair.hw_keypair.name
  key_id = huaweicloud_compute_keypair.hw_keypair.id
}

