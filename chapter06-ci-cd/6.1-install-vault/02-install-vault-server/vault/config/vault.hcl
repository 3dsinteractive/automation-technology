disable_mlock = true
ui            = true

listener "tcp" {
  address       = "0.0.0.0:8200"
  tls_cert_file = "/vault/cert/myvault_3dsinteractive_com.pem"
  tls_key_file  = "/vault/cert/myvault_3dsinteractive_com.key"
}

storage "file" {
  path = "/vault/data"
}
