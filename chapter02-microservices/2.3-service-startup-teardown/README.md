## Service Startup and Teardown

1. Open directory
```bash
cd /root/automation-technology/chapter02-microservices/2.3-service-startup-teardown
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
CONTAINER ID   IMAGE                             COMMAND                  CREATED         STATUS         PORTS                                                           NAMES
77543ea5c2cf   3dsinteractive/kafka:2.0-custom   "/app-entrypoint.sh …"   4 minutes ago   Up 4 minutes   9092/tcp, 0.0.0.0:9094->9094/tcp, :::9094->9094/tcp             05-asynctask-service-kafka-1
151990488a46   3dsinteractive/zookeeper:3.0      "/app-entrypoint.sh …"   4 minutes ago   Up 4 minutes   2888/tcp, 0.0.0.0:2181->2181/tcp, :::2181->2181/tcp, 3888/tcp   05-asynctask-service-zookeeper-1
ffb869159a3b   3dsinteractive/redis:4.0          "/app-entrypoint.sh …"   4 minutes ago   Up 4 minutes   0.0.0.0:6379->6379/tcp, :::6379->6379/tcp                       05-asynctask-service-redis-1
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
Scheduler: main.go 131 Tick at 22:18:21
Scheduler: main.go 131 Tick at 22:18:22
Scheduler: main.go 131 Tick at 22:18:23
Scheduler: main.go 131 Tick at 22:18:24
Scheduler: main.go 131 Tick at 22:18:25
Scheduler: main.go 131 Tick at 22:18:26
Scheduler: main.go 131 Tick at 22:18:27
Batch Consumer: main.go 106 Begin Batch
Batch Consumer: main.go 108 {"message_id":0}
Batch Consumer: main.go 108 {"message_id":1}
Batch Consumer: main.go 108 {"message_id":2}
Batch Consumer: main.go 110 End Batch
Batch Consumer: main.go 106 Begin Batch
Batch Consumer: main.go 108 {"message_id":3}
Batch Consumer: main.go 108 {"message_id":4}
Batch Consumer: main.go 108 {"message_id":5}
Batch Consumer: main.go 110 End Batch
Consumer: main.go 79 {"message_id":0}
Consumer: main.go 79 {"message_id":1}
Consumer: main.go 79 {"message_id":2}
Consumer: main.go 79 {"message_id":3}
Consumer: main.go 79 {"message_id":4}
Consumer: main.go 79 {"message_id":5}
Consumer: main.go 79 {"message_id":6}
Scheduler: main.go 131 Tick at 22:18:28
Consumer: main.go 79 {"message_id":7}
Scheduler: main.go 131 Tick at 22:18:29
Consumer: main.go 79 {"message_id":8}
Batch Consumer: main.go 106 Begin Batch
Batch Consumer: main.go 108 {"message_id":6}
Batch Consumer: main.go 108 {"message_id":7}
Batch Consumer: main.go 108 {"message_id":8}
Batch Consumer: main.go 110 End Batch
Scheduler: main.go 131 Tick at 22:18:30
Consumer: main.go 79 {"message_id":9}
Scheduler: main.go 131 Tick at 22:18:31
MS: microservice.go 662 Start cleanup
PROD: producer.go 83 Close successfullyScheduler: main.go 131 Tick at 22:18:21
Scheduler: main.go 131 Tick at 22:18:22
Scheduler: main.go 131 Tick at 22:18:23
Scheduler: main.go 131 Tick at 22:18:24
Scheduler: main.go 131 Tick at 22:18:25
Scheduler: main.go 131 Tick at 22:18:26
Scheduler: main.go 131 Tick at 22:18:27
Batch Consumer: main.go 106 Begin Batch
Batch Consumer: main.go 108 {"message_id":0}
Batch Consumer: main.go 108 {"message_id":1}
Batch Consumer: main.go 108 {"message_id":2}
Batch Consumer: main.go 110 End Batch
Batch Consumer: main.go 106 Begin Batch
Batch Consumer: main.go 108 {"message_id":3}
Batch Consumer: main.go 108 {"message_id":4}
Batch Consumer: main.go 108 {"message_id":5}
Batch Consumer: main.go 110 End Batch
Consumer: main.go 79 {"message_id":0}
Consumer: main.go 79 {"message_id":1}
Consumer: main.go 79 {"message_id":2}
Consumer: main.go 79 {"message_id":3}
Consumer: main.go 79 {"message_id":4}
Consumer: main.go 79 {"message_id":5}
Consumer: main.go 79 {"message_id":6}
Scheduler: main.go 131 Tick at 22:18:28
Consumer: main.go 79 {"message_id":7}
Scheduler: main.go 131 Tick at 22:18:29
Consumer: main.go 79 {"message_id":8}
Batch Consumer: main.go 106 Begin Batch
Batch Consumer: main.go 108 {"message_id":6}
Batch Consumer: main.go 108 {"message_id":7}
Batch Consumer: main.go 108 {"message_id":8}
Batch Consumer: main.go 110 End Batch
Scheduler: main.go 131 Tick at 22:18:30
Consumer: main.go 79 {"message_id":9}
Scheduler: main.go 131 Tick at 22:18:31
MS: microservice.go 662 Start cleanup
PROD: producer.go 83 Close successfully
```

7. Run command to cleanup
```bash
docker compose down
```