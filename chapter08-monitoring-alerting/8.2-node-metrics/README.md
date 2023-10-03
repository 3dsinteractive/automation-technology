## Node Metrics


1. Open directory 
```bash
cd /root/automation-technology/chapter08-monitoring-alerting/8.2-node-metrics
```

2. Run command
```bash
go mod init automationworkshop/main
go mod tidy
```

3. Deploy
```bash
/root/automation-technology/devopsctl-cli/devopsctl setup -d mon
```

4. Get pod
```bash
kubectl get po -n tcir-app
```

5. Jump into client-util
```bash
kubectl get po -n tcir-app | grep client-util
kubectl exec -n tcir-app -it <CLIENT-UTIL-POD> -- bash
```

6. Show metrics
```bash
curl -X GET "http://els1:9200/metrics/_search" | jq
```
