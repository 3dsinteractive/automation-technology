## Build docker image

1. Open directory
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.2-basic-docker/03-build-docker-image
```

2. Run command
```bash
docker build -t 3dsinteractive/mynginx:1.0 .
```

3. Run command
```bash
docker compose up -d
```

```bash
Creating network "03-build-docker-image_default" with the default driver
Creating 03-build-docker-image_nginx_1 ... done
```

4. Run command
```bash
curl -X GET localhost:8080/index.html
```

```html
<html>
<header><title>Hello World</title></header>
<body>
    <h1>Hello world</h1>
</body>
</html>
```

5. Explain Dockerfile

6. Run command to stop nginx
```bash
docker compose down
```

```bash
Stopping nginx ... done
Removing nginx ... done
Removing network 1202-run-nginx-with-docker_default
```

7. Run command to check nginx is stop and remove
```bash
docker ps
```

```bash
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```


