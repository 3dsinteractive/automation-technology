## Deploy Service

1. Slides Deploy Service

2. Open directory
```bash
cd /root/automation-technology/chapter03-deploy-scale-services/3.1-deploy-services
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

This will build the Dockerfile and push image to [your-docker-repository-name]/automation-technology

8. Explain deploy.sh (Each comments will explain itself)

9. Explain Dockerfile

10. Explain entrypoint.sh

11. Run command 
```bash
cd k8s
```

12. Run command
```bash
kubectl apply -f .
```

13. Run command
```bash
kubectl get po -n deploy-service
```

```bash
NAME                            READY   STATUS    RESTARTS   AGE
register-api-854c48c45c-nd8hh   1/1     Running   0          27s
register-api-854c48c45c-9fvkl   1/1     Running   0          27s
```

14. Run command
```bash
kubectl get svc -n deploy-service
```

```bash
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
register-api   ClusterIP   10.152.183.47   <none>        8080/TCP   43s
```

15. Run command
```bash
kubectl get ing -n deploy-service
```

```bash
NAME      CLASS    HOSTS                        ADDRESS     PORTS   AGE
ingress   public   kubernetes.docker.internal   127.0.0.1   80      65s
```

16. Run command
```bash
curl -X POST "http://kubernetes.docker.internal/citizen"
```

```javascript
{"status":"success"}
```

17. Run command
```bash
curl -X PUT "http://kubernetes.docker.internal/citizen/123"
```

```javascript
{"id":"123"}
```

18. Run command
```bash
curl -X GET "http://kubernetes.docker.internal/citizen/123?page=2"
```

```javascript
{"id":"123","page":"2"}
```

19. Run command
```bash
curl -X DELETE "http://kubernetes.docker.internal/citizen/123"
```

```javascript
{"status":"success"}
```

20. Run command
```bash
kubectl delete ns deploy-service
```