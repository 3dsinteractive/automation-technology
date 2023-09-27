## Play Consumer-Producer

### Kafkacat URL
https://github.com/edenhill/kafkacat

1. Start all k8s object if not created
```bash
kubectl apply -f /root/automation-technology/chapter01-basic-knowledges/1.5-basic-kafka/01-run-kafka/
```

2. Check if kafka and zookeeper pods is running
```bash
kubectl get po -n basic-kafka
```
```bash
NAME                    READY   STATUS    RESTARTS   AGE
zk1-65559d76fc-t9swb    1/1     Running   0          14m
zk2-59494b946f-4rfwr    1/1     Running   0          14m
zk3-6c88b6f849-jrqz5    1/1     Running   0          14m
kfk1-d78b96bbc-ndlzf    1/1     Running   0          9m12s
kfk2-7bc9945669-77x6d   1/1     Running   0          9m12s
kfk3-fbbd79879-hlp25    1/1     Running   0          9m12s
client-util             1/1     Running   0          2m49s
```

3. Exec into client-util pod
```bash
kubectl exec -it client-util -n basic-kafka -- bash
root@client-util:/#
```

4. Start Producer using kafkacat, and send some messages
```bash
kafkacat -P -b "kfk1,kfk2,kfk3" -t "mytopic"
```
```bash
message1
message2
message3
```

5. Open another tab in terminal, and exec into client-util
```bash
kubectl exec -it client-util -n basic-kafka -- bash
root@client-util:/#
```

6. Start Consumer in the open terminal
```bash
kafkacat -C -b "kfk1,kfk2,kfk3" -t "mytopic"
```

7. Exit Consumer by (Ctrl + c)

8. Exit client-util 
```bash
exit
```

