## Push docker image

1. Slides 1.2 Basic Docker (Docker images and containers)

2. Register for https://hub.docker.com

3. Create public repository call [your-repository-name]/mynginx:1.0

4. Login docker 
```bash
docker login
[username] : [your-repository-name]
[password] : xxxxxxxxx
```

5. Run command (Change [your-repository-name] to your registered docker name)
```bash
docker push [your-repository-name]/mynginx:1.0
```