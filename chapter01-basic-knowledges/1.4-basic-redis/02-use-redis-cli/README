## Use Redis cli

1. Exec into client-util pod
```bash
kubectl exec -it client-util -n basic-redis -- bash
```

2. Start redis-cli
```bash
redis-cli -h redis
redis:6379>
```

3. Use SET command to set mykey = myvalue
```bash
SET "mykey" "myvalue"
OK
```

4. Use GET to get value from key
```bash
GET "mykey"
"myvalue"
```

5. Use EXPIRE command to expire mykey in 10 seconds
```bash
EXPIRE "mykey" 10
(integer) 1
```

**Wait 10 seconds**

6. Test get mykey when it is expired
```bash
GET "mykey"
(nil)
```

7. SET mykey2
```bash
SET "mykey2" "myvalue2"
OK
```

8. GET mykey2
```bash
GET "mykey2"
"myvalue2"
```

9. Use DEL command to delete mykey2
```bash
DEL "mykey2"
(integer) 1
```

10. GET mykey2 to see how it is deleted
```bash
GET "mykey2"
(nil)
```

11. SET mykey1 and mykey2
```bash
SET "mykey1" "value1"
SET "mykey2" "value2"
```

12. Use KEYS to list all keys using wildcard
```bash
KEYS "mykey*"
1) "mykey2"
2) "mykey1"
```

13. Exit from redis
```bash
exit
```

14. Exit from client-util
```bash
exit
```

15. Cleanup workshop
```bash
kubectl delete ns basic-redis
```