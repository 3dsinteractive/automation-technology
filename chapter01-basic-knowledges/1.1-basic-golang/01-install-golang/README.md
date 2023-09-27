## Install Golang

1. Open home directory
```bash
cd /root
```

2. Install Golang on ubuntu
```bash
apt update -y 
wget https://dl.google.com/go/go1.20.1.linux-amd64.tar.gz
tar -xvf go1.20.1.linux-amd64.tar.gz
mv go /usr/local

echo "export GOPATH=/usr/local/go" >> ~/.bashrc
echo "export PATH=/usr/local/go/bin:$PATH" >> ~/.bashrc
source ~/.bashrc
```

2. After install open terminal and run command
```bash
go version
```