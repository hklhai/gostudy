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
```
