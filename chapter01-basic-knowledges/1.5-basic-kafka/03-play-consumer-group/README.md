## Play Consumer Group

### Kafkacat URL
https://github.com/edenhill/kafkacat

1. Start all k8s objects
```bash
kubectl apply -f /root/automation-technology/chapter01-basic-knowledges/1.5-basic-kafka/01-run-kafka/
```

2. Check if kafka and zookeeper pods is running
```bash
kubectl get po -n basic-kafka
```
```bash
NAME                    READY   STATUS    RESTARTS   AGE
zk1-65559d76fc-t9swb    1/1     Running   0          22m
zk2-59494b946f-4rfwr    1/1     Running   0          22m
zk3-6c88b6f849-jrqz5    1/1     Running   0          22m
kfk1-d78b96bbc-ndlzf    1/1     Running   0          16m
kfk2-7bc9945669-77x6d   1/1     Running   0          16m
kfk3-fbbd79879-hlp25    1/1     Running   0          16m
client-util             1/1     Running   0          10m
```

3. Exec into client-util pod
```bash
kubectl exec -it client-util -n basic-kafka -- bash
root@client-util:/#
```

4. Start Producer using kafkacat, and send some messages
```bash
kafkacat -P -b "kfk1,kfk2,kfk3" -t "mytopic1"
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

6. Start Consumer 1 in new terminal
```bash
kafkacat -C -b "kfk1,kfk2,kfk3" -G mygroup mytopic1
```

7. Open another tab in terminal, then exec into client-util shell
```bash
kubectl exec -it client-util -n basic-kafka -- bash
root@client-util:/#
```

8. Start Consumer 2 in new terminal
```bash
kafkacat -C -b "kfk1,kfk2,kfk3" -G mygroup mytopic1
```

9. Send some more messages in Producer
```bash
message4
message5
message6
message7
message8
```

10. Check each Consumers will receive new messages

12. Exit from Consumer using (Ctrl + c)

13. When Consumer exit, see the rebalance process happen

14. Exit from all consumer using (Ctrl + c)

15. Exit from all client-util
```bash
exit
```

16. Cleanup workshop
```bash
kubectl delete ns basic-kafka
```



