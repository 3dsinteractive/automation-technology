## Install Vault

1. Open directory
```bash
cd /root/automation-technology/chapter06-ci-cd/6.1-install-vault/02-install-vault
```

2. Change owner of vault directory
```bash
chown -Rf 100 vault
```

3. Install Make
```bash
apt -y install make
```

4. Run vault
```bash
make start
```