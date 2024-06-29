package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func goods(a [10]int) {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	for _, v := range a {

		_, err = c.Do("lpush", "goo", v)
		if err != nil {
			fmt.Println(err)
		}
	}
	str, err := redis.Strings(c.Do("lrange", "goo", "0", "-1"))
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range str {
		fmt.Println(v)
	}
}
func main() {
	var aaa [10]int
	for i := 0; i < 10; i++ {
		aaa[i] = i + 5
	}
	fmt.Println(aaa)
	goods(aaa)

}
