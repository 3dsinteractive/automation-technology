## Run Docker compose

1. See what is Docker compose file in Slides

2. Open directory
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.2-basic-docker/05-run-docker-compose
```

3. Run command to change owner of dockers directory
```bash
chown -Rf 1001:1001 dockers
```

4. Run command
```bash
docker compose up -d
```

```bash
docker ps
docker ps --format '{{.Image}}\t{{.Status}}\t{{.Ports}}'
```

5. Run command
```bash
curl -X GET "localhost:8080/index.html"
```

6. Explain section in docker-compose.yml

7. Run command
```bash
docker compose down
```