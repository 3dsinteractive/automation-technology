## Run Kafka

1. See what is Kafka in Slides

2. Enter directory
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.5-basic-kafka/01-run-kafka
```

3. Create namespace
```bash
kubectl apply -f 00-namespace.yml 
```
```bash
namespace/basic-kafka created
```

4. Check namespace exists
```bash
kubectl get ns
```
```bash
NAME              STATUS   AGE
kube-system       Active   66m
kube-public       Active   66m
kube-node-lease   Active   66m
default           Active   66m
ingress           Active   39m
basic-kafka       Active   7s
```

5. Create zookeeper service
```bash
kubectl apply -f 02-service-zk.yml
```
```bash
service/zk1 created
service/zk2 created
service/zk3 created
```

6. Create zookeeper deployment
```bash
kubectl apply -f 01-deployment-zk.yml
```

```bash
deployment.apps/zk1 created
deployment.apps/zk2 created
deployment.apps/zk3 created
```

7. Check zookeeper deployment
```bash
kubectl get po -n basic-kafka
```
```bash
NAME                   READY   STATUS    RESTARTS   AGE
zk1-65559d76fc-t9swb   1/1     Running   0          4m36s
zk2-59494b946f-4rfwr   1/1     Running   0          4m36s
zk3-6c88b6f849-jrqz5   1/1     Running   0          4m36s
```

**Wait until STATUS = Running and READY = 1/1**

8. Check zookeeper service
```bash
kubectl get svc -n basic-kafka
```
```bash
NAME   TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                      AGE
zk1    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   8s
zk2    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   8s
zk3    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   8s
```
9. Create kafka service
```bash
kubectl apply -f 04-service-kfk.yml
```
```bash
service/kfk1 created
service/kfk2 created
service/kfk3 created
```

10. Create kafka deployment
```bash
kubectl apply -f 03-deployment-kfk.yml
```
```bash
deployment.apps/kfk1 created
deployment.apps/kfk2 created
deployment.apps/kfk3 created
```

11. Check kafka deployment
```bash
kubectl get po -n basic-kafka
```
```bash
NAME                    READY   STATUS    RESTARTS   AGE
zk1-65559d76fc-t9swb    1/1     Running   0          10m
zk2-59494b946f-4rfwr    1/1     Running   0          10m
zk3-6c88b6f849-jrqz5    1/1     Running   0          10m
kfk1-d78b96bbc-ndlzf    1/1     Running   0          4m53s
kfk2-7bc9945669-77x6d   1/1     Running   0          4m53s
kfk3-fbbd79879-hlp25    1/1     Running   0          4m53s
```

**Wait until STATUS = Running and READY = 1/1**

12. Check kafka service
```bash
kubectl get svc -n basic-kafka
```
```bash
NAME   TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                      AGE
zk1    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   5m46s
zk2    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   5m46s
zk3    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   5m46s
kfk1   ClusterIP   None         <none>        9092/TCP                     6s
kfk2   ClusterIP   None         <none>        9092/TCP                     6s
kfk3   ClusterIP   None         <none>        9092/TCP                     6s
```

13. Create client-util pod
```bash
kubectl apply -f 05-client-util.yml
```
```bash
pod/client-util created
```

14. Exec into client-util pod
```bash
kubectl exec -it client-util -n basic-kafka -- bash
```

15. Run kafkacat -L to list brokers and topics
```bash
kafkacat -b "kfk1,kfk2,kfk3" -L
```
```bash
Metadata for all topics (from broker -1: kfk1:9092/bootstrap):
 3 brokers:
  broker 2 at kfk2:9092
  broker 3 at kfk3:9092
  broker 1 at kfk1:9092
 0 topics:
```

16. Exit client-util
```bash
exit
```

17. Do not cleanup workshop, we will use it in next workshop


