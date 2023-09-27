Run Docker compose

1. Open directory
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.2-basic-docker/05-run-docker-compose
```

2. Run command to change owner of dockers directory
```bash
chown -Rf 1001:1001 dockers
```

3. Run command
```bash
docker compose up -d
```

```bash
docker ps
```

4. Run command
```bash
curl -X GET "localhost:8080/index.html"
```

5. Explain section in docker compose.yml

6. Run command
```bash
docker compose down
```