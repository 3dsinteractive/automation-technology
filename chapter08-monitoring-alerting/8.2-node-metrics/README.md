## Node Metrics


1. Open directory 
```bash
cd /root/automation-technology/chapter08-monitoring-alerting/8.2-node-metrics
```

2. docker login
```bash
docker login
```

3. Build
```bash
./deploy.sh
```

4. Deploy
```bash
/root/automation-technology/devopsctl-cli/devopsctl setup -d mon
```

5. Get pod
```bash
kubectl get po -n tcir-app
```

6. Jump into client-util
```bash
kubectl get po -n tcir-app | grep client-util
kubectl exec -n tcir-app -it <CLIENT-UTIL-POD> -- bash
```

7. Show metrics
```bash
curl -X GET "http://els1:9200/metrics/_search" -H 'content-type: application/json' -d '{"size":20,"sort":[{"created_at":{"order":"desc"}}]}' | jq
```
