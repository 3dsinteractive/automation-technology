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
echo "127.0.0.1    myvault.3dsinteractive.com" | sudo tee -a /etc/hosts
```

6. Setup vault endpoint
```bash
echo "export VAULT_ADDR='https://myvault.3dsinteractive.com:8200'" >> ~/.bashrc
```

7. Source bashrc
```bash
source ~/.bashrc
```


