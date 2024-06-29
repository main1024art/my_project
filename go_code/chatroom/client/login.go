package main

import (
	"chanroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func login(userid int, userpwd string) (err error) {
	//fmt.Printf("userid = %d userwpd = %s\n", userid, userpwd)
	//return nil
	// 链接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	//延时关闭
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userid
	loginMes.UserPwd = userpwd
	//序列化loginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	//fmt.Printf("conn.Write 消息长度%d 内容%s", len(data), string(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	time.Sleep(time.Second * 20)
	fmt.Println("sleep 20 second")
	return
}
