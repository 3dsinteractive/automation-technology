## Consumer Service

1. Slides Consumer service

2. Open directory
```bash
cd /root/automation-technology/chapter02-microservices/2.2-microservices-types/02-consumer-service
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

3. Run command to init project
```bash
go mod init automationworkshop/main
go mod tidy
```

4. Run command to build project
```bash
go build .
```

5. Run program
```bash
./main
```

```bash
Consumer:  {"message_id":0}
Consumer:  {"message_id":1}
Consumer:  {"message_id":2}
Consumer:  {"message_id":3}
Consumer:  {"message_id":4}
Consumer:  {"message_id":5}
Consumer:  {"message_id":6}
Consumer:  {"message_id":7}
Consumer:  {"message_id":8}
Consumer:  {"message_id":9}
```

6. Run command to cleanup project
```bash
docker compose down
```
