## Healthcheck Readiness

1. Slides service healthcheck

2. Open directory
```bash
cd /root/automation-technology/chapter03-deploy-scale-services/3.2-healthcheck-readiness
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

**Wait until redis is 1/1 Running**

10. Run command
```bash
kubectl get po -n healthcheck
```

```bash
NAME                     READY   STATUS    RESTARTS   AGE
redis-577d58dd6c-c6595   1/1     Running   0          3m55s
```

11. Run command to start application
```bash
kubectl apply -f 01-application/.
```

12. Run command to make sure all services is Running
```bash
kubectl get po -n healthcheck
```

```bash
register-api-758f494ccd-7xqmx   1/1     Running   0          109s
register-api-758f494ccd-gbtdg   1/1     Running   0          109s
```

13. Run command
```bash
curl -X POST "http://kubernetes.docker.internal/citizen"
```

```javascript
{"status":"success"}
```

14. Run command 
```bash
watch "microk8s kubectl get po -n healthcheck"
```

15. Open other terminal and Run command to kill redis
```bash
cd /root/automation-technology/chapter03-deploy-scale-services/3.2-healthcheck-readiness/k8s/00-databases
```
```bash
kubectl delete -f 01-redis.yml
```

16. Wait to see service healthcheck fail
**Notice RESTARTS will increase**

17. Run command to get redis back
```bash
kubectl apply -f 01-redis.yml
```

18. Wait to see service come back

19. Explain healthcheck endpoint in sourcecode

20. Run command
```bash
kubectl delete ns healthcheck
```