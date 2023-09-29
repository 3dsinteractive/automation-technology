## ParallelTask Service

1. Slide ParallelTask Service

2. Open directory
```bash
cd /root/automation-technology/chapter02-microservices/2.2-microservices-types/06-paralleltask-service
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

8. Run command
```bash
curl -X POST "localhost:8080/citizen/batch?task_id=email_a&worker_count=5" -d '{"input1":"value1"}'
```

```javascript
{"task_id":"email_a"}
```

9. Run command
```bash
curl -X GET "localhost:8080/citizen/batch?task_id=email_a"
```

**The response can be formatted with jq**

```bash
apt install -y jq
```

```bash
curl -X GET "localhost:8080/citizen/batch?task_id=email_a" | jq
```

```javascript
{
  "status": "complete",
  "workers": [
    {
      "code": 200,
      "error": "",
      "response": {
        "result": "123"
      },
      "status": "complete",
      "worker_id": "ptask-email_a-4751997750760398084"
    },
    {
      "code": 200,
      "error": "",
      "response": {
        "result": "123"
      },
      "status": "complete",
      "worker_id": "ptask-email_a-7504504064263669287"
    },
    {
      "code": 200,
      "error": "",
      "response": {
        "result": "123"
      },
      "status": "complete",
      "worker_id": "ptask-email_a-1976235410884491574"
    },
    {
      "code": 200,
      "error": "",
      "response": {
        "result": "123"
      },
      "status": "complete",
      "worker_id": "ptask-email_a-3510942875414458836"
    },
    {
      "code": 200,
      "error": "",
      "response": {
        "result": "123"
      },
      "status": "complete",
      "worker_id": "ptask-email_a-2933568871211445515"
    }
  ]
}
```

10. Run command to cleanup
```bash
docker compose down
```