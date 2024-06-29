package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	i := 1
	var a int

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	fmt.Println("输入一个数")

	for ; i <= 10; i++ {
		fmt.Scanln(&a)
		if a == randomNumber {
			if i == 1 {
				fmt.Println("天才")
			} else if i == 2 && i == 3 {
				fmt.Println("崇明")
			} else if i >= 4 && i <= 9 {
				fmt.Println("一般般")
			}
		}
	}
}
