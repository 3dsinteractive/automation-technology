## Install Docker

1. Install docker for ubuntu
```bash
apt update -y && apt install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt update -y && apt install -y docker-ce
```

2. Enable docker for systemctl
```bash
systemctl start docker
systemctl enable docker
```

3. Check docker version
```bash
docker version
```

4. Create docker group
```bash
groupadd docker
usermod -aG docker root
```