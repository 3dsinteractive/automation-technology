## Run Kafka

1. Enter directory
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.5-basic-kafka/01-run-kafka
```

2. Create namespace
```bash
kubectl apply -f 00-namespace.yml 
```
```bash
namespace/basic-kafka created
```

3. Check namespace exists
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

4. Create zookeeper deployment
```bash
kubectl apply -f 01-deployment-zk.yml
```

```bash
deployment.apps/zk1 created
deployment.apps/zk2 created
deployment.apps/zk3 created
```

5. Check zookeeper deployment
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

6. Create zookeeper service
```bash
kubectl apply -f 02-service-zk.yml
```
```bash
service/zk1 created
service/zk2 created
service/zk3 created
```

7. Check zookeeper service
```bash
kubectl get svc -n basic-kafka
```
```bash
NAME   TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                      AGE
zk1    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   8s
zk2    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   8s
zk3    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   8s
```

8. Create kafka deployment
```bash
kubectl apply -f 03-deployment-kfk.yml
```
```bash
deployment.apps/kfk1 created
deployment.apps/kfk2 created
deployment.apps/kfk3 created
```

9. Check kafka deployment
```bash
kubectl get po -n basic-kafka
```
```bash
NAME                    READY   STATUS              RESTARTS   AGE
kfk1-86886b6b84-xkfh2   0/1     ContainerCreating   0          14s
kfk2-5b69dfcdb4-kwwk2   0/1     ContainerCreating   0          14s
kfk3-6d4c8874c6-ll7nh   0/1     ContainerCreating   0          14s
zk1-76cc547698-jhngx    1/1     Running             0          12m
zk2-7bb59d6788-rc8s5    1/1     Running             0          12m
zk3-566db54d6b-g579s    1/1     Running             0          12m
```

**Wait until STATUS = Running and READY = 1/1**

10. Create kafka service
```bash
kubectl apply -f 04-service-kfk.yml
```
```bash
service/kfk1 created
service/kfk2 created
service/kfk3 created
```

11. Check kafka service
```bash
kubectl get svc -n basic-kafka
```
```bash
NAME   TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                      AGE
kfk1   ClusterIP   None         <none>        9092/TCP                     28s
kfk2   ClusterIP   None         <none>        9092/TCP                     28s
kfk3   ClusterIP   None         <none>        9092/TCP                     28s
zk1    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   14m
zk2    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   14m
zk3    ClusterIP   None         <none>        2181/TCP,2888/TCP,3888/TCP   14m
```

12. Create client-util pod
```bash
kubectl apply -f 05-client-util.yml
```
```bash
pod/client-util created
```

13. Exec into client-util pod
```bash
kubectl exec -it client-util -n basic-kafka -- bash
```

14. Run kafkacat -L to list brokers and topics
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

15. Exit client-util
```bash
exit
```

16. Do not cleanup workshop, we will use it in next workshop


