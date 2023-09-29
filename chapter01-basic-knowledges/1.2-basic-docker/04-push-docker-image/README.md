## Push docker image

1. See What is build and push Docker image in Slides

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
docker tag 3dsinteractive/mynginx:1.0 [your-repository-name]/mynginx:1.0
docker push [your-repository-name]/mynginx:1.0
```