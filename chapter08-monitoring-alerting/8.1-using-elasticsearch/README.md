## Using Elasticsearch

1. Jump into client-util
```bash
kubectl get po -n tcir-app | grep client-util
kubectl exec -n tcir-app -it <CLIENT-UTIL-POD> -- bash
```

2. Check access to elasticsearch
```bash
curl -X GET http://els1:9200/_cluster/health | jq
```

3. Create mapping for metrics
```bash
curl -X PUT "http://els1:9200/metrics" -H "Content-Type: application/json" -d '{"mappings": {"properties": {"id": {"type": "keyword"}, "node_name": {"type": "keyword"}, "created_at": {"type": "date", "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"}, "metrics": {"type": "object", "dynamic": true}}}}'
```

4. Show index
```bash
curl -X GET "http://els1:9200/metrics"
```
