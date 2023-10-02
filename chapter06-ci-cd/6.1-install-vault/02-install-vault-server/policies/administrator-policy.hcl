path "ssh-client-signer/roles/*" {
 capabilities = ["list"]
}
path "ssh-client-signer/sign/administrator" {
 capabilities = ["create","update"]
}
path "auth/token/lookup" {
  capabilities = ["read"]
}
path "kv/data/tokens/*" {
  capabilities = ["list","read","create","update","delete"]
}