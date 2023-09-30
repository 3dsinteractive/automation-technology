## Frequent commands

### Install zsh
```bash
apt install -y zsh
```

### Install ohmyzsh
```bash
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

### When Get Lost
```bash
cd ~/automation-technology
```

### Docker
```bash
docker ps -a

docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
```

### Image Pull
```bash
docker pull opcellent/util:2.0
docker pull 3dsinteractive/redis:5.0
docker pull 3dsinteractive/kafka:2.0-custom
docker pull 3dsinteractive/zookeeper:3.0
```

### Git
```bash
git pull
git status
git checkout HEAD file_path
```

### K8S
```bash
kubectl get node
kubectl get ns
kubectl get po -n namespace
kubectl get svc -n namespace
kubectl get rs -n namespace
kubectl get ing -n namespace
kubectl get deploy -n namespace

kubectl logs pod_name -n namespace
kubectl describe po pod_name -n namespace
```