## Apply Terraform

1. Open directory
```bash
cd /root/automation-technology/chapter05-cloud-infra/5.2-apply-terraform
```

2. Run command
```bash
terraform init
```

3. Run command
```bash
terraform plan
```

4. Input value name
```bash
Client name (client-name)

  Enter a value: [your-name]
```

5. Input value access key
```bash
Huawei Account access key 

  Enter a value: [access-key-from-file]
```

6. Input value secret key
```bash
Huawei Account secret key 

  Enter a value: [secret-key-from-file]
```

7. Run command
```bash
terraform apply
```

8. Input value name of each service deployed ex. [your-name]-vpc, [your-name]-ecs 
```bash
Client name (client-name)

  Enter a value: [your-name]
```

9. Input value access key
```bash
Huawei Account access key 

  Enter a value: [access-key]
```

10. Input value secret key
```bash
Huawei Account secret key 

  Enter a value: [secret-key]
```

10. Prompt will ask confirm action
```bash
Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: [type 'yes']
```

12. Output terraform process done
```
Outputs:

ecs1_piblic_ip = <sensitive>
ecs2_piblic_ip = <sensitive>

```

13. Output terraform process done
```
Outputs:

ecs1_piblic_ip = <sensitive>
ecs2_piblic_ip = <sensitive>

```

14. Run command and save the result to use next chapter
```bash
terraform output ecs1_piblic_ip
```

15. Run command and save the result to use next chapter
```bash
terraform output ecs2_piblic_ip
```

16. Run command
```bash
chmod 400 ./keypair/[your-name]-keypair.pem
```

