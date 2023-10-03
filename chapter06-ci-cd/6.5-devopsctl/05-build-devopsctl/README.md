## DevOpsCtrl

1. Increase vm.maxmapcount (for elasticsearch)
```bash
echo "vm.max_map_count=262144" | sudo tee -a /etc/sysctl.conf
sysctl -p
```

2. Open directory
```bash
cd /root/automation-technology/chapter06-ci-cd/6.5-devopsctl/04-build-devopsctl
```

3. Run command to init project
```bash
go mod init automationworkshop/main
go mod tidy
```

4. Run command to build project
```bash
go build .
```

5. Run setup for database
```bash
./main setup -d db
```

6. Run setup for application
```bash
./main setup -d tcir
```

7. Get pods
```bash
kubectl get po -n tcir-app
```