## Batch Consumer

1. Slide Batch Consumer Service

2. Open directory
```bash
cd /root/automation-technology/chapter02-microservices/2.2-microservices-types/03-batch-consumer-service
```

2. Start kafka and zookeeper
```bash
docker compose up -d
```

3. Wait until container is ready
```bash
docker ps
```
```bash
CONTAINER ID   IMAGE                             COMMAND                  CREATED          STATUS         PORTS                                                           NAMES
11fd6a7ff959   3dsinteractive/kafka:2.0-custom   "/app-entrypoint.sh …"   9 seconds ago    Up 8 seconds   9092/tcp, 0.0.0.0:9094->9094/tcp, :::9094->9094/tcp             02-consumer-service-kafka-1
09f00c0618e8   3dsinteractive/zookeeper:3.0      "/app-entrypoint.sh …"   10 seconds ago   Up 9 seconds   2888/tcp, 0.0.0.0:2181->2181/tcp, :::2181->2181/tcp, 3888/tcp   02-consumer-service-zookeeper-1
```

4. Run command to init project
```bash
go mod init automationworkshop/main
go mod tidy
```

5. Run command to build project
```bash
go build .
```

6. Run program
```bash
./main
```

```bash
Batch Consumer:  Begin Batch
Batch Consumer:  {"message_id":0}
Batch Consumer:  {"message_id":1}
Batch Consumer:  {"message_id":2}
Batch Consumer:  End Batch
Batch Consumer:  Begin Batch
Batch Consumer:  {"message_id":3}
Batch Consumer:  End Batch
Batch Consumer:  Begin Batch
Batch Consumer:  {"message_id":4}
Batch Consumer:  {"message_id":5}
Batch Consumer:  {"message_id":6}
Batch Consumer:  End Batch
Batch Consumer:  Begin Batch
Batch Consumer:  {"message_id":7}
Batch Consumer:  {"message_id":8}
Batch Consumer:  End Batch
Batch Consumer:  Begin Batch
Batch Consumer:  {"message_id":9}
Batch Consumer:  End Batch
```

7. Run command to cleanup
```bash
docker compose down
```