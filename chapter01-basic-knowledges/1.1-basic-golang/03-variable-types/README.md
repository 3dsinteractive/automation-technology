## Variables and Types

1. Open directory 
```bash
cd /root/automation-technology/chapter01-basic-knowledges/1.1-basic-golang/03-variable-types
```

2. Run command
```bash
go mod init automationworkshop/main
go mod tidy
```

3. Run command
```bash
go run main.go
```

```bash
string =  This variable type string
int =  305
bool =  true
---
interface =  This variable type string
interface =  305
interface =  true
---
map =  map[citizen_id:1234 firstname:Chaiyapong lastname:Lapliengtrakul]
map JSON =  {"citizen_id":"1234","firstname":"Chaiyapong","lastname":"Lapliengtrakul"}
No Gender specify
---
slice =  [item 1 item 2 item 3]
slice JSON =  ["item 1","item 2","item 3"]
---
Citizen =  &{Chaiyapong Lapliengtrakul 1234}
Citizen JSON =  {"firstname":"Chaiyapong","lastname":"Lapliengtrakul","citizen_id":"1234"}
---
theMap is nil
theCitizen is nil
---
Gender is Unspecify
Gender =  UNSPECIFY
---
```