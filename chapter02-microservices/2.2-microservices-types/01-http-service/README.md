## HTTP service

1. Open directory
```bash
cd /root/automation-technology/chapter02-microservices/2.2-microservices-types/01-http-service
```

2. Run command to init project
```bash
go mod init automationworkshop/main
go mod tidy
```

3. Run command to build project
```bash
go build .
```

4. Run program
```bash
./main
```

5. Run command
```bash
curl -X POST "localhost:8080/citizen"
{"status":"success"}
```

6. Run command
```bash
curl -X PUT "localhost:8080/citizen/123"
{"id":"123"}
```

7. Run command
```bash
curl -X GET "localhost:8080/citizen/123?page=2"
{"id":"123","page":"2"}
```

8. Run command
```bash
curl -X DELETE "localhost:8080/citizen/123"
{"status":"success"}
```

9. Explain the service in source code