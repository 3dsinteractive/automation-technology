## Keys Management

1. Open directory
```bash
cd /root/automation-technology/chapter06-ci-cd/6.2-keys-management
```

2. Generate some 16 digits password
```bash
openssl rand -base64 24 | head -c 16; echo
```

3. Create admin1
```bash
vault policy write admin admin-policy.hcl
```
```bash
vault write auth/userpass/users/admin1 policies="admin" password="<YOUR-PASSWORD>"
```

4. Generate some 16 digits password
```bash
openssl rand -base64 24 | head -c 16; echo
```

5. Create customer 1
```bash
vault write auth/userpass/users/cluster-customer1 policies="cluster-customer1" password="<YOUR-PASSWORD>"
```
```bash
vault policy write cluster-customer1 cluster-customer1-policy.hcl
```

6. Generate some 16 digits password
```bash
openssl rand -base64 24 | head -c 16; echo
```

7. Create customer 2
```bash
vault write auth/userpass/users/cluster-customer2 policies="cluster-customer2" password="<YOUR-PASSWORD>"
```
```bash
vault policy write cluster-customer2 cluster-customer2-policy.hcl
```

8. Login as admin1
```bash
vault login -method=userpass username=admin1
```

**Enter admin1 password**

8. Generate some access token for customer1
```bash
openssl rand -base64 24 | head -c 24; echo
```

9. Set access token to customer1 using generated access token
```bash
vault kv put kv/data/customers/customer1/els access_token=<YOUR_ACCESS_TOKEN_VALUE>
```

10. List all values in path kv/data/customers/customer1
```bash
vault kv list kv/data/customers/customer1
```

11. List all values in path kv/data/customers/customer1/els
```bash
vault kv get kv/data/customers/customer1/els
```

12. Generate some access token for customer2
```bash
openssl rand -base64 24 | head -c 24; echo
```

13. Set access token to customer2 using generated access token
```bash
vault kv put kv/data/customers/customer2/els access_token=<YOUR_ACCESS_TOKEN_VALUE>
```

14. List all values in path kv/data/customers/customer1
```bash
vault kv list kv/data/customers/customer2
```

15. List all values in path kv/data/customers/customer1/els
```bash
vault kv get kv/data/customers/customer2/els
```