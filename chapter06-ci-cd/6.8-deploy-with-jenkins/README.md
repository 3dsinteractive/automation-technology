## Deploy with Jenkins


1. Copy id_rsa.pub of jenkins to k8s host
```bash
echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDW9NxDDpv4O3nBdZtxR8+1Ia4TRC7G09TMmHhqzVQyI118akE98ikW+E9V/wq1fIwCKgWHrCE0yEaEZ2v9fm6c1ucoJJ9s+FzOpUIcs3pKrsekXDfj2pYXl6/zJao4U0wz0v3ab/WHtljVNjguwTs5xXQSbpPF+5slLZt81J1UmHOQ3B4KgQ6dfLoyCAm9Eah+svB8jFtBf30NkIoDY79Ki6QwQhFkvpsOt8+u2yeWaM/deK3uVCnPaC/3tGAQuxir2k0L3r/8DxPMm8nzHtxPW4hLKfbTGFIu2crFbEpKOvCfvoPjR7oHB9+I5/p2pHwauI5DZ3JOTgbZsUFS6cZYQKti/gielt3q5ueu5wcfgM2Z6O8pZn440g8ZMGyj6dJVtMPRlTaZqQ73hyL+1TX/LaVFRoYolskXe4vsz93SfkJVO/LRBLNInqdyoxOrhqeUxjZEv6s5lDrltimLaAGentVsievOmWZE1bUx3j0WsSnvigahhxHFhUmija6Wzur2TRyTlrQDxfAU3zh53niaJn33kWlK3GFCtS4IVawsAMv2HXynGE1rG0WHY3L7DTXw9ibKsV32QQWNoUHhHIgRgTmJzIhTiLnjGARMN7irhjPIaerHhMummddlI+zK8DeaGbKhiYlNVFA2e+vSEFs6+FFq/iQ06rSB8IEVq9NEhw== paul3ds@Chaiyapongs-MacBook-Air-2.local" >> ~/.ssh/authorized_keys
```

2. Create Server in jenkins
Click Menu Manage Jenkins -> System -> SSH Remote Hosts -> Add
Hostname: Public IP or Private IP if K8S Host
Port: 22
Credential: root

**Click Check Connection** to test connection then **Click Save**


3. Create project
Name: deploy-tcir
Build Step: Execute shell script on remote host using ssh
SSH site: root@YOUR-IP-ADDRESS
Command:
cd /root/automation-technology/chapter06-ci-cd/6.5-devopsctl/05-build-devopsctl
./main setup -d tcir

4. Start Build