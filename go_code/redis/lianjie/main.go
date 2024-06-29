package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	var aa int
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connect redis success")
	defer c.Close()
	_, err = c.Do("set", "key1", 998)
	if err != nil {
		fmt.Println(err)
		return
	}
	aa, err = redis.Int(c.Do("get", "key1"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(aa)
	fmt.Println("get redis success")
}
