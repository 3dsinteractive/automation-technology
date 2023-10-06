## Install vault client

1. Go to home directory
```bash
cd ~/
```

2. Install required program
```bash
apt -y install unzip
```

3. Install vault
```bash
wget https://releases.hashicorp.com/vault/1.6.3/vault_1.6.3_linux_amd64.zip
```

```bash
unzip vault_1.6.3_linux_amd64.zip
```

```bash
sudo mv vault /usr/local/bin/
```

4. Validate Installable
```bash
vault --version
```

5. Setup host file
```bash
echo "<MACHINE-1-PRIVATE-IP>   myvault.3dsinteractive.com" | sudo tee -a /etc/hosts
```
**Replace MACHINE-1-PRIVATE-IP with machine 1 IP Address**

6. Setup vault endpoint
```bash
echo "export VAULT_ADDR='https://myvault.3dsinteractive.com:8200'" >> ~/.bashrc
source ~/.bashrc
```

7. Login as admin1
```bash
vault login -method=userpass username=admin1
```

8. IMPORTANT Write down vault token
**Don't forget to write down vault token**

9. List all values in path kv/data/customers/customer1
```bash
vault kv list kv/data/customers/customer1
```

10. List all values in path kv/data/customers/customer1/els
```bash
vault kv get kv/data/customers/customer1/els
```