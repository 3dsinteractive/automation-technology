## DevOpsCtrl

1. Increase vm.maxmapcount (for elasticsearch)
```bash
echo "vm.max_map_count=262144" | sudo tee -a /etc/sysctl.conf
sysctl -p
```

2. Open directory
```bash
cd /root/automation-technology/chapter06-ci-cd/6.5-devopsctl/05-build-devopsctl
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

6. Get database pods
```bash
kubectl get po -n tcir-app
```

7. Run setup for application
```bash
./main setup -d tcir -t <vault-token>
```

8. Get application pods
```bash
kubectl get po -n tcir-app
```

9. Run setup util
```bash
./main setup -d util
```

10. Get all pods
```bash
kubectl get po -n tcir-app
```

11. Copy main to devopsctl
```bash
cp /root/automation-technology/chapter06-ci-cd/6.5-devopsctl/05-build-devopsctl/main /root/automation-technology/devopsctl-cli/devopsctl
```