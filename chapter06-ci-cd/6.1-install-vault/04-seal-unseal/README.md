## Seal - Unseal

1. Init vault
```bash
vault operator init
```

2. Write down all unseal keys and initial root access token
```bash
Unseal Key 1: xxxxx
Unseal Key 2: xxxxx
Unseal Key 3: xxxxx
Unseal Key 4: xxxxx
Unseal Key 5: xxxxx

Initial Root Token: xxxxx
```

2. Unseal Vault
```bash
vault operator unseal
```

**Useal with vault unseal key 3 times**

```bash
Key             Value
---             -----
Seal Type       shamir
Initialized     true
Sealed          false
Total Shares    5
Threshold       3
Version         1.9.4
Storage Type    file
Cluster Name    vault-cluster-1c7ca582
Cluster ID      8635e343-e39d-87c9-cccf-f90e5f2f2f1a
HA Enabled      false
```

3. Login
```bash
vault login
```

**Use root token**