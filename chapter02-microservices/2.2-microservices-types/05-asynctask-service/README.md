## AsyncTask Service

1. Slide AsyncTask Service

2. Open directory
```bash
cd /root/automation-technology/chapter02-microservices/2.2-microservices-types/05-asynctask-service
```

3. Start kafka and zookeeper
```bash
docker compose up -d
```

4. Wait until container is ready
```bash
docker ps
docker ps --format '{{.Image}}\t{{.Status}}\t{{.Ports}}'
```
```bash
CONTAINER ID   IMAGE                             COMMAND                  CREATED         STATUS         PORTS                                                           NAMES
77543ea5c2cf   3dsinteractive/kafka:2.0-custom   "/app-entrypoint.sh …"   4 minutes ago   Up 4 minutes   9092/tcp, 0.0.0.0:9094->9094/tcp, :::9094->9094/tcp             05-asynctask-service-kafka-1
151990488a46   3dsinteractive/zookeeper:3.0      "/app-entrypoint.sh …"   4 minutes ago   Up 4 minutes   2888/tcp, 0.0.0.0:2181->2181/tcp, :::2181->2181/tcp, 3888/tcp   05-asynctask-service-zookeeper-1
ffb869159a3b   3dsinteractive/redis:4.0          "/app-entrypoint.sh …"   4 minutes ago   Up 4 minutes   0.0.0.0:6379->6379/tcp, :::6379->6379/tcp                       05-asynctask-service-redis-1
```

5. Run command to init project
```bash
go mod init automationworkshop/main
go mod tidy
```

6. Run command to build project
```bash
go build .
```

7. Run program
```bash
./main
```

8. Open another terminal and run command 
```bash
curl -X POST "localhost:8080/citizen/register" -d '{"firstname":"chaiyapong"}'
```
```javascript
{"ref":"atask-xxxxxxxxxxx"}
```

6. Run command 
```bash
curl -X GET "localhost:8080/citizen/register?ref=atask-xxxxxxxxxxx"
```
```javascript
{"code":200,"data":{"id":"123"},"status":"success"}
```

7. Run command for cleanup
```bash
docker compose down
```