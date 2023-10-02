## Install MicroK8S

1. Run command to install
```bash
snap install microk8s --classic
```

2. Add user to group microk8s
```bash
usermod -a -G microk8s $USER
```
```bash
chown -f -R $USER ~/.kube
```

3. Alias kubectl
```bash
echo "alias kubectl='microk8s kubectl'" >> ~/.bashrc
source ~/.bashrc
```

4. Start microk8s
```bash
microk8s start
```

5. Run command
```bash
kubectl get node
```
```
NAME           STATUS   ROLES    AGE     VERSION
pam-training   Ready    <none>   9m17s   v1.27.5
```

6. Run command to add domain in etc hosts
```bash
echo "127.0.0.1 kubernetes.docker.internal" | sudo tee -a /etc/hosts
```

7. Enable ingress
```bash
microk8s enable ingress
```
