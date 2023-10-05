## Alerting

1. Open directory 
```bash
cd /root/automation-technology/chapter08-monitoring-alerting/8.3-alerting
```

2. Get pod
```bash
kubectl get po -n tcir-app
```

3. Jump into client-util
```bash
kubectl get po -n tcir-app | grep client-util
kubectl exec -n tcir-app -it <CLIENT-UTIL-POD> -- bash
```

4. Query last 20 metrics
```bash
curl -X GET "http://els1:9200/metrics/_search" -H 'content-type: application/json' -d '{"size":20,"sort":[{"created_at":{"order":"desc"}}]}' | jq
```

5. Build
```bash
./deploy.sh
```

6. Deploy
```bash
/root/automation-technology/devopsctl-cli/devopsctl setup -d mon
```

7. Get pod
```bash
kubectl get po -n tcir-app
```