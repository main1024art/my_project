package main

import (
	"chanroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	n, err := conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("conn.Read err:", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err = conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn.Read err:", err)
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	return
}

// 处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	//循环读取客户端发送的信息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务端进行退出")
				return
			} else {
				fmt.Println("readPkg error:", err)
				return
			}

		}
		fmt.Println("readPkg:", mes)
	}
}

func main() {
	//提示
	fmt.Println("服务器8889")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Println("等待链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go process(conn)
	}

}
