## Realife App TCIR (Thai Citizen ID Card Register)

1. Slides TCIR

2. Open directory
```bash
cd /root/automation-technology/chapter04-implement-reallife-app/4.2-tcir-application
```

3. Register for https://hub.docker.com

4. Create public repository call [your-docker-repository-name]/automation-technology

5. Make sure you are logged in with your docker account
```bash
docker login
username:
password:
```

6. Update file deploy.sh change
```bash
DOCKER_REPOSITORY=
to
DOCKER_REPOSITORY=[your-docker-repository-name]
```

7. Run command (deploy.sh will be the file used to build your project, especially when it integrate with ci/cd)
```bash
./deploy.sh
```

8. Run command
```bash
cd k8s
```

9. Run command to start databases
```bash
kubectl apply -f 00-databases/.
```

10. Run command to check if every pods is OK
```bash
kubectl get po -n tcir-app
```
**Wait until every services is 1/1 Running**

11. Run command to start all services
```bash
kubectl apply -f 01-application/.
```

12. Notice all services is started in different pod (different deployment)
```bash
kubectl get po -n tcir-app
```
```bash
NAME                                 READY   STATUS    RESTARTS   AGE
batch-ptask-api-55df49f66f-gvnnz     1/1     Running   0          2m14s
batch-ptask-api-55df49f66f-z8ggg     1/1     Running   0          2m14s
batch-ptask-worker-87c757d95-879rt   1/1     Running   0          2m14s
batch-ptask-worker-87c757d95-8whkr   1/1     Running   0          2m14s
batch-ptask-worker-87c757d95-pdsmd   1/1     Running   0          2m14s
batch-scheduler-c49fb8bbd-l7thf      1/1     Running   0          2m14s
client-util                          1/1     Running   1          33m
external-api-57668cf778-lmdgh        1/1     Running   0          2m14s
kfk1-86886b6b84-rm6s5                1/1     Running   4          36m
kfk2-5b69dfcdb4-gm6kg                1/1     Running   4          36m
kfk3-6d4c8874c6-4w42p                1/1     Running   4          36m
mail-consumer-dbf5fbdd6-9tf22        1/1     Running   0          2m14s
mail-consumer-dbf5fbdd6-ds48q        1/1     Running   0          2m14s
mail-consumer-dbf5fbdd6-gwdgs        1/1     Running   0          2m14s
redis-577d58dd6c-jthr6               1/1     Running   1          36m
register-api-79db7f8cd4-hnlzx        1/1     Running   0          2m14s
register-api-79db7f8cd4-t7fmv        1/1     Running   0          2m14s
zk1-76cc547698-g6cvp                 1/1     Running   1          36m
zk2-7bb59d6788-pswh7                 1/1     Running   1          36m
zk3-566db54d6b-j4w6m                 1/1     Running   1          36m
```

13. Look into sourcecode (main.go) to see that we read SERVICE_ID from env in deployment files
    And also look at how we connect each services together using HTTP, Consumer, AsyncTask, Scheduler, ParallelTask

14. Run command
```bash
curl -X POST "http://kubernetes.docker.internal/api/citizen"
```

```javascript
{"ref":"atask-5577006791947779410"}
```

15. Run command
```bash
curl -X GET "http://kubernetes.docker.internal/api/citizen?ref=atask-5577006791947779410"
```

```javascript
{"code":200,"data":{"citizen_id":"5577006791947779410","status":"success"},"status":"success"}
```

16. Run command to get pod name of [mail-consumer-xxxx]
    And use pod name to get logs
```bash
kubectl get po -n tcir-app | grep mail-consumer
```
```bash
kubectl logs [pod1] -n tcir-app
kubectl logs [pod2] -n tcir-app
kubectl logs [pod3] -n tcir-app
```
```bash
Consumer: main.go 112 Mail confirmation has sent to 5577006791947779410
```

17. Run command
```bash
kubectl exec -it client-util -n tcir-app -- bash
kafkacat -b "kfk1,kfk2,kfk3" -L | grep "confirm"
```

```bash
topic "when-citizen-has-confirmed" with 50 partitions:
```

18. Check message in partitions
```bash
kafkacat -b "kfk1,kfk2,kfk3" -C -e -t when-citizen-has-confirmed
Find {"citizen_id":"5577006791947779410"} in one of 50 partitions
exit
```

19. The scheduler will run at midnight and batch create & delivery 
    id card to all of citizen in topic when-citizen-has-confirmed

20. Run command to cleanup
```bash
kubectl delete ns tcir-app
```