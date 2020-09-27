package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("SET", "name", "hk")
	if err != nil {
		fmt.Println("redis set error:", err)
	}

	_, err = conn.Do("EXPIRE", "name", "5")
	if err != nil {
		fmt.Println("redis set EXPIRE:", err)
	}

	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}

	time.Sleep(time.Duration(6) * time.Second)
	name2, err2 := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err2)
	} else {
		fmt.Printf("Got name: %s \n", name2)
	}
}
