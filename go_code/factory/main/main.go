package main

import (
	"factory/model"
	"fmt"
)

func main() {
	// var stu = model.stu{
	// 	Name : "tom",
	// 	Score : 87.5,
	// }
	var stu = model.Say("tom", 12.4)
	fmt.Println(stu)
	fmt.Println(*stu)
}
