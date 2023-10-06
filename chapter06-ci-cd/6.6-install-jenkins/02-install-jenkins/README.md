## Install Jenkins

1. Go to VM 1
**Make sure you are on VM 1**

2. Open directory
```bash
cd /root/automation-technology/chapter06-ci-cd/6.6-install-jenkins/02-install-jenkins
```

2. Create logs directory
```bash
mkdir -p dockers/lb/logs
```

3. Change owner of lb directory
```bash
chown -Rf 1001 dockers/lb
```

4. Start jenkins
```bash
docker compose up -d
```

5. Get password from logs
```bash
docker logs 02-install-jenkins-jenkins-1
```
**Copy password to use later**

```bash
user=admin
password=xxxxxxx
```

6. Update jenkins
Click at Manage Jenkins -> Update Jenkins

7. Install Git plugin
Click at Manage Jenkins -> Plugins -> Git -> Install

8. Install SSH Credential plugin
Click at Manage Jenkins -> Plugins -> SSH Credential -> Install

9. Install SSH plugin
Click at Manage Jenkins -> Plugins -> SSH -> Install

9. Restart Jenkins

