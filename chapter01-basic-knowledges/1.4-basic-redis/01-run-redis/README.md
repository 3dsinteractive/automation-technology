## Run Redis

1. See what is Redis in Slides

2. Enter directory
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.4-basic-redis/01-run-redis
```

3. Create k8s namespace
```bash
kubectl apply -f 00-namespace.yml
```

```bash
namespace/basic-redis created
```

4. Check if namespace has created
```bash
kubectl get ns
```

```bash
NAME              STATUS   AGE
kube-system       Active   36m
kube-public       Active   36m
kube-node-lease   Active   36m
default           Active   36m
ingress           Active   9m38s
basic-redis       Active   31s
```

5. Create redis deployment
```bash
kubectl apply -f 01-deployment.yml
```

```bash
deployment.apps/redis created
```

6. Check redis deployment has created
```bash
kubectl get po -n basic-redis
```

```bash
NAME                    READY   STATUS    RESTARTS   AGE
redis-599994f84-h5n6r   1/1     Running   0          2m30s
```

**Wait until the STATUS is Running**

7. Create redis service
```bash
kubectl apply -f 02-service.yml 
```

```bash
service/redis created
```

8. Check redis service has created
```bash
kubectl get svc -n basic-redis
```
```bash
NAME    TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
redis   ClusterIP   None         <none>        6379/TCP   22s
```

9. Create client-util pod
```bash
docker pull opcellent/util:2.0
kubectl apply -f 03-client-util.yml
```
```bash
pod/client-util created
```

10. Check if client-util pod has created
```bash
kubectl get po -n basic-redis
```

```bash
NAME                    READY   STATUS    RESTARTS   AGE
redis-599994f84-h5n6r   1/1     Running   0          11m
client-util             1/1     Running   0          7m33s
```

**Wait until the STATUS is Running**

11. Exec into client-util pod
```bash
kubectl exec -it client-util -n basic-redis -- bash
root@client-util:/#
```

12. Run redis-cli to connect to redis
```bash
redis-cli -h redis
redis:6379>
```

13. Exit from redis-cli
```bash
exit
root@client-util:/#
```

14. Exit from client-util
```bash
exit
```

15. Do not cleanup workshop, we will use it in next workshop