## Install Docker

1. Install docker for ubuntu
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

2. Enable docker for systemctl
```bash
systemctl start docker
```
```bash
systemctl enable docker
```

3. Check docker version
```bash
docker version
```

4. Create docker group
```bash
groupadd docker
```
```bash
usermod -aG docker root
```

5. Create user 1001
```bash
useradd -u 1001 --badname --no-create-home 1001
```
```bash
usermod -aG docker 1001
```

6. Pull required images
```bash
docker pull 3dsinteractive/elasticsearch:6.8.23-custom
docker pull 3dsinteractive/kafka:2.0-custom
docker pull 3dsinteractive/zookeeper:3.0
docker pull 3dsinteractive/redis:5.0
docker pull opcellent/util:2.0
docker pull 3dsinteractive/kibana:6.8
```