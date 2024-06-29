package main

import (
	"fmt"
	"kehuguanli/model"
	"kehuguanli/service"
)

type aaa struct {
	key         string
	loop        bool
	Kehuservice *service.Kehuservice
}

func (a *aaa) list() {
	cc := a.Kehuservice.List()
	fmt.Println("------------客户列表-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(cc); i++ {
		fmt.Println(cc[i].Info())
	}

	fmt.Println("--------------完成---------------")
}

func (a *aaa) add() {
	fmt.Println("------------添加客户-------------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年纪")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("手机")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	vv := model.NewKehu2(name, gender, age, phone, email)
	if a.Kehuservice.Add(vv) {
		fmt.Println("------------添加成功-------------")
	}

}
func (a *aaa) del() {
	fmt.Println("------------删除客户-------------")
	fmt.Println("输入编号(000退出)")
	n := 0
	fmt.Scanln(&n)
	fmt.Println("确认删除? y/n")
	m := ""
	fmt.Scanln(&m)
	if m == "y" || m == "Y" {
		if a.Kehuservice.Del(n) {
			fmt.Println("------------删除成功-------------")
		}
	}

}
func (a *aaa) mainMenu() {
	for {
		fmt.Println("------------客户管理系统-------------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退    出")
		fmt.Println("            请选择(1-5)")
		fmt.Scanln(&a.key)
		switch a.key {
		case "1":
			a.add()
		case "2":
			fmt.Println("            请选择(1-5)")
		case "3":
			a.del()
		case "4":
			a.list()
		case "5":
			a.loop = false

		}
		if !a.loop {
			break
		}
	}
	fmt.Println("你已退出")
}
func main() {
	var a aaa = aaa{
		key:         "",
		loop:        true,
		Kehuservice: service.NewKehuservice(),
	}
	a.mainMenu()
}
