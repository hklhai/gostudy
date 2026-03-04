## download image
docker pull redis:latest

## start container
docker run -itd --name redis-rdb -p 6379:6379 redis


## 进入容器设置kv
``` 
docker exec -it redis-rdb /bin/bash

redis-cli

set test 123

get test

// 设置过期时间
EXPIRE test 5

```

## 查看所有key
``` 
keys *
```

## 查看key的过期时间
``` 
127.0.0.1:6379> ttl OPObNsLpHVsngS0AVl5Phk/i1oz+VwlLocicWiRw7pQiQ2s4Hb66WMFrPNHOypycW3qDbvyQSi7zBcMPz6DrjuWXBS08jITK9d1HDIR59Z1snj4ve1s1grKlKt7sg3Nl
(integer) 2697
```

## 删除key
``` 
127.0.0.1:6379> del s
(integer) 1
```