package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("connected to redis server")
	_, err = conn.Do("hmset", "m1", "name", "红孩儿", "age", "18", "skill", "火")
	if err != nil {
		fmt.Println(err)
	}
	str, err := redis.Strings(conn.Do("hgetall", "m1"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	for _, str := range str {
		fmt.Println(str)
	}

}
