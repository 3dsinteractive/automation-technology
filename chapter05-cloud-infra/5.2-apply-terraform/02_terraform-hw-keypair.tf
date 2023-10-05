resource "tls_private_key" "hw_private_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "huaweicloud_compute_keypair" "hw_keypair"{
  name   = "${local.resource_client}-keypair"
  public_key = tls_private_key.hw_private_key.public_key_openssh
}

locals {
  key_name = huaweicloud_compute_keypair.hw_keypair.name
  key_id = huaweicloud_compute_keypair.hw_keypair.id
}

output "private_key_value" {
  value = tls_private_key.hw_private_key.private_key_pem
  sensitive = true
}
