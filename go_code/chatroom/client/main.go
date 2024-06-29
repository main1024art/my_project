package main

import "fmt"

// 定义两个变量  id pwd
var userId int
var userPwd string

func main() {
	//接受用户选择
	var key int
	var loop = true

	for loop {
		fmt.Println("---------------------登录--------------------")
		fmt.Println("\t\t\t\t\t1, 登陆聊天室")
		fmt.Println("\t\t\t\t\t2, 注册")
		fmt.Println("\t\t\t\t\t3, 退出")
		fmt.Println("\t\t\t\t\t选择 1 - 3")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录")
			loop = false
		case 2:
			fmt.Println("zhuce")
			loop = false
		case 3:
			fmt.Println("tuichu")
			loop = false
		default:
			fmt.Println("shuru you wu")
		}
	}
	if key == 1 {
		fmt.Println("shuru id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("shuru pwd")
		fmt.Scanf("%s\n", &userPwd)
		err := login(userId, userPwd)
		fmt.Println()
		if err != nil {
			fmt.Println("denglu shibai")
		} else {
			fmt.Println("denglu chenggong")
		}
	} else if key == 2 {
		fmt.Println("zhuce")
	}
}
