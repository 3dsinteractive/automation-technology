## Install Docker

1. Go to VM 1
**Make sure you are on VM 1**

2. Install docker for ubuntu
```bash
apt update -y && apt install -y apt-transport-https ca-certificates curl software-properties-common
```
```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```
```bash
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```
```bash
apt update -y && apt install -y docker-ce
```

3. Enable docker for systemctl
```bash
systemctl start docker
```
```bash
systemctl enable docker
```

4. Check docker version
```bash
docker version
```

5. Create docker group
```bash
groupadd docker
```
```bash
usermod -aG docker root
```

6. Create user 1001
```bash
useradd -u 1001 --badname --no-create-home 1001
```
```bash
usermod -aG docker 1001
```