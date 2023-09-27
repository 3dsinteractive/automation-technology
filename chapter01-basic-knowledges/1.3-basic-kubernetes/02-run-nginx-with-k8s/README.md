## Run nginx with k8s

1. Open directory
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.3-basic-kubernetes/02-run-nginx-with-k8s
```

2. Run command to create namespace
```bash
kubectl apply -f 00-namespace.yml
```

```bash
namespace/basic-k8s created
```

3. Check for namespace (basic-k8s)
```bash
kubectl get ns
```

```bash
NAME              STATUS   AGE
basic-k8s         Active   19s
default           Active   6h40m
kube-node-lease   Active   6h40m
kube-public       Active   6h40m
kube-system       Active   6h40m
```

4. Create deployment for nginx
```bash
kubectl apply -f 01-deployment.yml
```

```bash
deployment.apps/nginx created
```

5. Check for nginx deployment
```bash
kubectl get pod -n basic-k8s
```

```bash
NAME                    READY   STATUS    RESTARTS   AGE
nginx-7c7ddc7b7-cvb8z   1/1     Running   0          64s
```

6. Create service for nginx
```bash
kubectl apply -f 02-service.yml
```

```bash
service/nginx created
```

7. Check for service nginx
```bash
kubectl get svc -n basic-k8s
```

```bash
NAME    TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)             AGE
nginx   ClusterIP   10.152.183.221   <none>        8080/TCP,8443/TCP   23s
```

8. Create ingress rule to nginx service
```bash
kubectl apply -f 03-ingress.yml
```

```bash
ingress.networking.k8s.io/ingress created
```

9. Check for ingress rule
```bash
kubectl get ing -n basic-k8s
```

```bash
NAME      CLASS    HOSTS                        ADDRESS   PORTS   AGE
ingress   <none>   kubernetes.docker.internal             80      37s
```

10. Test access service via ingress
```bash
curl -X GET "http://kubernetes.docker.internal"
```

```html
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

11. Run command to cleanup
```bash
kubectl delete ns basic-k8s
```
