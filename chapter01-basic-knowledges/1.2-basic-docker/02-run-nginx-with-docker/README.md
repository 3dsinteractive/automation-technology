## Run Nginx with Docker

1. Slides 1.6 Basic Docker (Why we need Docker?)

2. Open directory
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.2-basic-docker/02-run-nginx-with-docker
```

3. Run command to create dockers directory
```bash
mkdir -p dockers/nginx/nginx
```
```bash
chown -Rf 1001:1001 dockers
```

4. Run command
```bash
docker compose up -d
```

5. Run command
```bash
docker ps
```

```bash
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS          PORTS                                         NAMES
ebece58ff823   3dsinteractive/nginx:1.12   "/app-entrypoint.sh …"   26 seconds ago   Up 25 seconds   0.0.0.0:8080->8080/tcp, 0.0.0.0:8443->8443/tcp   nginx
```

6. Run command to request http from nginx
```bash
curl -X GET "http://localhost:8080"
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

7. Run command to stop nginx
```bash
docker compose down
```

```bash
Stopping nginx ... done
Removing nginx ... done
Removing network 1202-run-nginx-with-docker_default
```

8. Run command to check nginx is stop and remove
```bash
docker ps
```

```bash
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```
