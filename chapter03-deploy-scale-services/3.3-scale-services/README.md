## Scale Service

1. Slides scale service

2. Open directory
```bash
cd /root/automation-technology/chapter03-deploy-scale-services/3.3-scale-services
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

10. Run command to check if every pods is running
```bash
kubectl get po -n scale-service
```
**Wait until every services is 1/1 Running**

11. Run command to start all application services
```bash
kubectl apply -f 01-application/.
```

12. Notice register-api and mail-consumer is started in different pod (different deployment)
```bash
kubectl get po -n scale-service
```

```bash
NAME                             READY   STATUS    RESTARTS   AGE
register-api-57f5b7cc87-rh4f6    1/1     Running   0          30s
register-api-57f5b7cc87-vmvdc    1/1     Running   0          30s
mail-consumer-6f648fb4b5-b2nfc   1/1     Running   0          30s
```

13. Look into sourcecode (main.go) to see that we read SERVICE_ID from env in deployment files
    Also look how register-api send message to mail-consumer

14. Also we read CACHE_SERVER and MQ_SERVERS from env too, to not hardcoding configuration in sourcecode

15. Run command
```bash
curl -X POST "http://kubernetes.docker.internal/citizen"
```

```javascript
{"citizen_id":"xxxxx","status":"success"}
```

16. Run command to see the logs of consumer
Copy pod name then run 
```bash
kubectl logs [pod-name] -n scale-service
Consumer: main.go 68 Mail has sent to {"citizen_id":"xxxxx"}
```

17. Run command to scale consumers
```bash
kubectl scale -f 01-application/06-mail-consumer.yml --replicas=3
deployment.apps/mail-consumer scaled
```

18. Run command to see consumer has scaled, while api has stay the same
```bash
kubectl get po -n scale-service
watch "microk8s kubectl get po -n scale-service"
```

19. Run command to play with replicas down
```bash
kubectl scale -f 01-application/06-mail-consumer.yml --replicas=1
kubectl get po -n scale-service
```

**See the mail-consumer scale down to 1**

20. Run command to cleanup
```bash
kubectl delete ns scale-service
```