## Unit Test

1. Open directory 
```bash
cd /root/automation-technology/chapter07-automate-test/7.1-unit-test
```

2. Run command
```bash
go mod init automationworkshop/main
go mod tidy
```

3. Run Test
```bash
go test .
```

4. Check Method **onPostClient(ctx IContext, cfg IConfig, rnd IRandom)** in main.go

5. Check following files
- main_test.go
- context_mock.go
- producer_mock.go
- config_mock.go
- random_mock.go